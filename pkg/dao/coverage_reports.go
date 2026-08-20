package dao

import (
	"context"
	"time"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/config"
	ce "github.com/content-services/content-sources-backend/pkg/errors"
	"github.com/content-services/content-sources-backend/pkg/models"
	"gorm.io/gorm"
)

type CreateCoverageReportParams struct {
	OrgID     string
	AccountID *string
}

type CreateCoverageUploadParams struct {
	UUID       string
	StorageKey string
	Sha256     string
	SizeBytes  int64
}

type coverageReportDaoImpl struct {
	db *gorm.DB
}

func (d coverageReportDaoImpl) Create(ctx context.Context, report CreateCoverageReportParams, upload CreateCoverageUploadParams) (api.CoverageReportResponse, error) {
	var modelReport models.CoverageReport
	var modelUpload models.CoverageUpload

	coverageReportCreateParamsToModels(report, upload, &modelReport, &modelUpload)
	modelReport.Status = config.TaskStatusPending

	err := d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&modelUpload).Error; err != nil {
			return err
		}
		return tx.Create(&modelReport).Error
	})
	if err != nil {
		return api.CoverageReportResponse{}, coverageReportDBErrorToApi(err)
	}

	var resp api.CoverageReportResponse
	coverageReportModelToApi(modelReport, &resp)
	return resp, nil
}

func (d coverageReportDaoImpl) Fetch(ctx context.Context, orgID string, reportUUID string) (api.CoverageReportResponse, error) {
	var modelReport models.CoverageReport
	result := d.db.WithContext(ctx).Where("uuid = ? AND org_id = ?", reportUUID, orgID).First(&modelReport)
	if result.Error != nil {
		return api.CoverageReportResponse{}, coverageReportDBErrorToApi(result.Error)
	}

	var resp api.CoverageReportResponse
	coverageReportModelToApi(modelReport, &resp)
	return resp, nil
}

func (d coverageReportDaoImpl) InternalOnlyFetchCoverageUpload(ctx context.Context, uploadUUID string) (models.CoverageUpload, error) {
	var upload models.CoverageUpload
	result := d.db.WithContext(ctx).Where("uuid = ?", uploadUUID).First(&upload)
	if result.Error != nil {
		return models.CoverageUpload{}, coverageReportDBErrorToApi(result.Error)
	}
	return upload, nil
}

func (d coverageReportDaoImpl) SetAnalysisTaskUUID(ctx context.Context, reportUUID string, taskUUID string) error {
	result := d.db.WithContext(ctx).
		Model(&models.CoverageReport{}).
		Where("uuid = ?", reportUUID).
		Update("analysis_task_uuid", taskUUID)
	if result.Error != nil {
		return coverageReportDBErrorToApi(result.Error)
	}
	if result.RowsAffected == 0 {
		return coverageReportDBErrorToApi(gorm.ErrRecordNotFound)
	}
	return nil
}

func (d coverageReportDaoImpl) UpdateCoverageReportStatus(ctx context.Context, reportUUID string, status string, errMsg *string) error {
	switch status {
	case config.TaskStatusPending, config.TaskStatusRunning, config.TaskStatusCompleted, config.TaskStatusFailed:
	default:
		return &ce.DaoError{BadValidation: true, Message: "Invalid coverage report status."}
	}

	updates := map[string]interface{}{
		"status": status,
	}
	if status == config.TaskStatusFailed && errMsg != nil {
		updates["analysis_task_error"] = *errMsg
	}
	if status == config.TaskStatusCompleted {
		now := time.Now().UTC()
		updates["completed_at"] = now
	}

	result := d.db.WithContext(ctx).
		Model(&models.CoverageReport{}).
		Where("uuid = ?", reportUUID).
		Updates(updates)
	if result.Error != nil {
		return coverageReportDBErrorToApi(result.Error)
	}
	if result.RowsAffected == 0 {
		return coverageReportDBErrorToApi(gorm.ErrRecordNotFound)
	}
	return nil
}

func coverageReportDBErrorToApi(e error) *ce.DaoError {
	if dbError, ok := e.(models.Error); ok && dbError.Validation {
		return &ce.DaoError{BadValidation: true, Message: dbError.Message}
	}
	daoErr := ce.DaoError{Message: e.Error()}
	daoErr.Wrap(e)
	return &daoErr
}

func coverageReportCreateParamsToModels(report CreateCoverageReportParams, upload CreateCoverageUploadParams, modelReport *models.CoverageReport, modelUpload *models.CoverageUpload) {
	modelReport.OrgID = report.OrgID
	modelReport.AccountID = report.AccountID

	modelUpload.UUID = upload.UUID
	modelUpload.StorageKey = upload.StorageKey
	modelUpload.Sha256 = upload.Sha256
	modelUpload.SizeBytes = upload.SizeBytes
}

func coverageReportModelToApi(model models.CoverageReport, resp *api.CoverageReportResponse) {
	resp.UUID = model.UUID
	resp.Status = model.Status
	resp.CreatedAt = model.CreatedAt.UTC()
	if model.InputFormat != nil {
		resp.InputFormat = *model.InputFormat
	}
	if model.CompletedAt != nil {
		completedAt := model.CompletedAt.UTC()
		resp.CompletedAt = &completedAt
	}
	if model.Total != nil {
		resp.Total = *model.Total
	}
	if model.ExactMatches != nil {
		resp.ExactMatches = *model.ExactMatches
	}
	if model.PartialMatches != nil {
		resp.PartialMatches = *model.PartialMatches
	}
	if model.Unmatched != nil {
		resp.Unmatched = *model.Unmatched
	}
	if model.AnalysisTaskError != nil {
		resp.AnalysisTaskError = *model.AnalysisTaskError
	}
	if model.AnalysisTaskUUID != nil {
		resp.AnalysisTaskUUID = *model.AnalysisTaskUUID
	}
	if model.EcosystemCoverageSummary != nil {
		resp.EcosystemCoverageSummary = make([]api.EcosystemCoverageSummary, len(*model.EcosystemCoverageSummary))
		for i, entry := range *model.EcosystemCoverageSummary {
			resp.EcosystemCoverageSummary[i] = api.EcosystemCoverageSummary{
				Ecosystem:      entry.Ecosystem,
				Total:          entry.Total,
				ExactMatches:   entry.ExactMatches,
				PartialMatches: entry.PartialMatches,
				Unmatched:      entry.Unmatched,
			}
		}
	}
}
