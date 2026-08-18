package dao

import (
	"context"

	"github.com/content-services/content-sources-backend/pkg/api"
	"github.com/content-services/content-sources-backend/pkg/config"
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

func (d coverageReportDaoImpl) CreateCoverageReport(ctx context.Context, report CreateCoverageReportParams, upload CreateCoverageUploadParams) (api.CoverageReportResponse, error) {
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
		return api.CoverageReportResponse{}, err
	}

	var resp api.CoverageReportResponse
	coverageReportModelToApi(modelReport, &resp)
	return resp, nil
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
