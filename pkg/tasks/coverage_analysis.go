package tasks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/dao"
	"github.com/content-services/content-sources-backend/pkg/db"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/tasks/payloads"
	"github.com/content-services/content-sources-backend/pkg/tasks/queue"
	"github.com/rs/zerolog/log"
)

const maxCoverageManifestSizeBytes = 50 * 1024 * 1024 // 50 MiB

const coverageAnalysisNotImplementedMsg = "coverage analysis not implemented"

type CoverageAnalysis struct {
	ctx     context.Context
	payload *payloads.CoverageAnalysisPayload
	daoReg  *dao.DaoRegistry
}

func CoverageAnalysisHandler(ctx context.Context, task *models.TaskInfo, _ *queue.Queue) error {
	opts := payloads.CoverageAnalysisPayload{}
	if err := json.Unmarshal(task.Payload, &opts); err != nil {
		return fmt.Errorf("payload incorrect type for %s", config.CoverageAnalysisTask)
	}

	logger := LogForTask(task.Id.String(), task.Typename, task.RequestID)
	ctxWithLogger := logger.WithContext(ctx)

	daoReg := dao.GetDaoRegistry(db.DB)
	ca := CoverageAnalysis{
		ctx:     ctxWithLogger,
		payload: &opts,
		daoReg:  daoReg,
	}
	err := ca.Run()
	if err == nil {
		return nil
	}
	if errors.Is(err, context.Canceled) {
		return nil
	}
	msg := err.Error()
	if updateErr := daoReg.CoverageReport.UpdateCoverageReportStatus(context.Background(), opts.CoverageReportUUID, config.TaskStatusFailed, &msg); updateErr != nil {
		log.Error().Errs("errors", []error{err, updateErr}).Str("coverage_report_uuid", opts.CoverageReportUUID).Msg("failed to update coverage report status")
	}
	return err
}

func (c *CoverageAnalysis) Run() error {
	upload, err := c.daoReg.CoverageReport.InternalOnlyFetchCoverageUpload(c.ctx, c.payload.CoverageUploadUUID)
	if err != nil {
		return fmt.Errorf("fetch coverage upload: %w", err)
	}

	manifestBytes, err := downloadCoverageManifestFromS3(c.ctx, upload.StorageKey)
	if err != nil {
		return fmt.Errorf("download coverage manifest: %w", err)
	}

	if err := verifyCoverageManifest(manifestBytes, upload); err != nil {
		return err
	}

	if err := c.daoReg.CoverageReport.UpdateCoverageReportStatus(c.ctx, c.payload.CoverageReportUUID, config.TaskStatusRunning, nil); err != nil {
		return err
	}

	// Plumbing-only: parser/matcher not implemented yet. Analysis failed, task succeeded.
	msg := coverageAnalysisNotImplementedMsg
	if err := c.daoReg.CoverageReport.UpdateCoverageReportStatus(c.ctx, c.payload.CoverageReportUUID, config.TaskStatusFailed, &msg); err != nil {
		return err
	}
	return nil
}

func verifyCoverageManifest(manifestBytes []byte, upload models.CoverageUpload) error {
	if int64(len(manifestBytes)) != upload.SizeBytes {
		return fmt.Errorf("manifest size mismatch")
	}

	hash := sha256.Sum256(manifestBytes)
	if hex.EncodeToString(hash[:]) != upload.Sha256 {
		return fmt.Errorf("manifest sha256 mismatch")
	}
	return nil
}

func downloadCoverageManifestFromS3(ctx context.Context, storageKey string) ([]byte, error) {
	cfg := config.Get().Clients.Lightwell.CoverageUploads
	if cfg.Name == "" {
		return nil, fmt.Errorf("coverage uploads bucket is not configured")
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %w", err)
	}
	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
		if cfg.URL != "" {
			o.BaseEndpoint = aws.String(cfg.URL)
		}
	})

	out, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &cfg.Name,
		Key:    &storageKey,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download report from S3: %w", err)
	}
	defer out.Body.Close()

	body, err := io.ReadAll(io.LimitReader(out.Body, maxCoverageManifestSizeBytes+1))
	if err != nil {
		return nil, fmt.Errorf("failed to read report from S3: %w", err)
	}
	if int64(len(body)) > maxCoverageManifestSizeBytes {
		return nil, fmt.Errorf("manifest exceeds maximum size")
	}
	return body, nil
}
