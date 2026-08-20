package handler

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/content-services/content-sources-backend/pkg/tasks"
	"github.com/content-services/content-sources-backend/pkg/tasks/client"
	"github.com/content-services/content-sources-backend/pkg/tasks/payloads"
	"github.com/content-services/content-sources-backend/pkg/tasks/queue"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/dao"
	ce "github.com/content-services/content-sources-backend/pkg/errors"
	"github.com/content-services/content-sources-backend/pkg/rbac"
	"github.com/content-services/content-sources-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

const maxCoverageUploadSizeBytes = 50 * 1024 * 1024 // 50 MiB

type CoverageReportHandler struct {
	DaoRegistry dao.DaoRegistry
	TaskClient  client.TaskClient
}

func checkLightwellBeaconAndLensAccessible(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := CheckLightwellBeaconAndLensAccessible(c.Request().Context()); err != nil {
			return err
		}
		return next(c)
	}
}

func RegisterCoverageReportRoutes(engine *echo.Group, daoReg *dao.DaoRegistry, taskClient *client.TaskClient) {
	ch := CoverageReportHandler{
		DaoRegistry: *daoReg,
		TaskClient:  *taskClient,
	}
	addRepoRoute(engine, http.MethodPost, "/coverage_reports/", ch.createCoverageReport, rbac.RbacVerbWrite, checkLightwellBeaconAndLensAccessible)
	addRepoRoute(engine, http.MethodGet, "/coverage_reports/:uuid", ch.getCoverageReport, rbac.RbacVerbRead, checkLightwellBeaconAndLensAccessible)
	addRepoRoute(engine, http.MethodGet, "/coverage_reports/:uuid/packages", ch.listCoverageReportPackages, rbac.RbacVerbRead, checkLightwellBeaconAndLensAccessible)
}

// CreateCoverageReport godoc
// @Summary      Create coverage report
// @ID           createCoverageReport
// @Description  Upload a manifest file and start coverage analysis.
// @Tags         coverage_reports
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "Manifest file (CycloneDX, SPDX, etc.)"
// @Success      201 {object} api.CoverageReportResponse
// @Failure      400 {object} ce.ErrorResponse
// @Failure      401 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /coverage_reports/ [post]
func (ch *CoverageReportHandler) createCoverageReport(c echo.Context) error {
	accountID, orgID := getAccountIdOrgId(c)

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return ce.NewErrorResponse(http.StatusBadRequest, "Error binding parameters", "File is required")
	}
	if fileHeader.Size <= 0 {
		return ce.NewErrorResponse(http.StatusBadRequest, "Error reading upload", "Size must be greater than 0")
	}
	if fileHeader.Size > maxCoverageUploadSizeBytes {
		return ce.NewErrorResponse(http.StatusRequestEntityTooLarge, "Error reading upload", "File exceeds maximum upload size")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return ce.NewErrorResponse(http.StatusBadRequest, "Error opening upload", err.Error())
	}
	defer file.Close()

	uploadUUID := uuid.NewString()
	storageKey := "coverage-uploads/" + uploadUUID

	hash := sha256.New()
	fileBytes, err := io.ReadAll(io.TeeReader(file, hash))
	if err != nil {
		return ce.NewErrorResponse(http.StatusBadRequest, "Error reading upload", err.Error())
	}

	if err := uploadCoverageManifestToS3(c.Request().Context(), storageKey, fileBytes); err != nil {
		return ce.NewErrorResponse(http.StatusInternalServerError, "Error uploading coverage report", err.Error())
	}

	sha256Hex := hex.EncodeToString(hash.Sum(nil))
	sizeBytes := int64(len(fileBytes))

	reportParams := dao.CreateCoverageReportParams{OrgID: orgID}
	if accountID != "" {
		reportParams.AccountID = utils.Ptr(accountID)
	}
	uploadParams := dao.CreateCoverageUploadParams{
		UUID:       uploadUUID,
		StorageKey: storageKey,
		Sha256:     sha256Hex,
		SizeBytes:  sizeBytes,
	}

	report, err := ch.DaoRegistry.CoverageReport.Create(c.Request().Context(), reportParams, uploadParams)
	if err != nil {
		return ce.NewErrorResponse(ce.HttpCodeForDaoError(err), "Error creating coverage report", err.Error())
	}

	ch.enqueueCoverageAnalysisEvent(c, report, uploadUUID)

	return c.JSON(http.StatusCreated, report)
}

func uploadCoverageManifestToS3(ctx context.Context, storageKey string, fileBytes []byte) error {
	cfg := config.Get().Clients.Lightwell.CoverageUploads
	if cfg.Name == "" {
		log.Warn().Msg("Not configured to upload to S3")
		return nil
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithRegion(cfg.Region),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKey, cfg.SecretKey, "")))
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %w", err)
	}
	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = true
		if cfg.URL != "" {
			o.BaseEndpoint = aws.String(cfg.URL)
		}
	})

	_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &cfg.Name,
		Key:    &storageKey,
		Body:   bytes.NewReader(fileBytes),
	})
	if err != nil {
		return fmt.Errorf("failed to upload report to S3: %w", err)
	}
	log.Info().Msgf("Uploaded %s to s3", storageKey)
	return nil
}

// GetCoverageReport godoc
// @Summary      Get coverage report
// @ID           getCoverageReport
// @Description  Return a coverage report by UUID.
// @Tags         coverage_reports
// @Produce      json
// @Param        uuid path string true "Coverage report UUID"
// @Success      200 {object} api.CoverageReportResponse
// @Failure      404 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /coverage_reports/{uuid} [get]
func (ch *CoverageReportHandler) getCoverageReport(c echo.Context) error {
	_, orgID := getAccountIdOrgId(c)

	report, err := ch.DaoRegistry.CoverageReport.Fetch(c.Request().Context(), orgID, c.Param("uuid"))
	if err != nil {
		return ce.NewErrorResponse(ce.HttpCodeForDaoError(err), "Error fetching coverage report", err.Error())
	}

	return c.JSON(http.StatusOK, report)
}

// ListCoverageReportPackages godoc
// @Summary      List coverage report packages
// @ID           listCoverageReportPackages
// @Description  Return paginated packages for a completed coverage report.
// @Tags         coverage_reports
// @Produce      json
// @Param        uuid path string true "Coverage report UUID"
// @Param        search query string false "Filter by package name"
// @Param        ecosystem query string false "Filter by ecosystem"
// @Param        status query string false "Filter by package match status (possible values: in_network, not_in_network)"
// @Param        offset query int false "Starting point for pagination. Default: 0"
// @Param        limit query int false "Number of items per page. Default: 100"
// @Success      200 {object} api.CoverageReportPackagesResponse
// @Failure      400 {object} ce.ErrorResponse
// @Failure      404 {object} ce.ErrorResponse
// @Failure      500 {object} ce.ErrorResponse
// @Router       /coverage_reports/{uuid}/packages [get]
func (ch *CoverageReportHandler) listCoverageReportPackages(c echo.Context) error {
	req := api.ListCoverageReportPackagesRequest{}
	if err := c.Bind(&req); err != nil {
		return ce.NewErrorResponse(http.StatusBadRequest, "Error binding parameters", err.Error())
	}

	response, err := stubListCoverageReportPackages(c.Param("uuid"), req, ParsePagination(c))
	if errors.Is(err, errStubCoverageReportNotFound) {
		return ce.NewErrorResponse(http.StatusNotFound, "Coverage report not found", "Report is not available or analysis is incomplete")
	}
	if err != nil {
		return ce.NewErrorResponse(http.StatusInternalServerError, "Error loading fixture", err.Error())
	}

	return c.JSON(http.StatusOK, response)
}

func (ch *CoverageReportHandler) enqueueCoverageAnalysisEvent(c echo.Context, report api.CoverageReportResponse, uploadUUID string) uuid.UUID {
	accountID, orgID := getAccountIdOrgId(c)
	payload := payloads.CoverageAnalysisPayload{CoverageReportUUID: report.UUID, CoverageUploadUUID: uploadUUID}
	task := queue.Task{
		Typename:  config.CoverageAnalysisTask,
		Payload:   payload,
		OrgId:     orgID,
		AccountId: accountID,
		RequestID: c.Response().Header().Get(config.HeaderRequestId),
	}
	taskID, err := ch.TaskClient.Enqueue(task)
	logger := tasks.LogForTask(taskID.String(), task.Typename, task.RequestID)
	if err != nil {
		logger.Error().Msg("error enqueuing task")
	}
	if err == nil {
		if err = ch.DaoRegistry.CoverageReport.SetAnalysisTaskUUID(c.Request().Context(), report.UUID, taskID.String()); err != nil {
			logger.Error().Msg("error updating analysis task UUID")
		} else {
			report.AnalysisTaskUUID = taskID.String()
		}
	}

	return taskID
}
