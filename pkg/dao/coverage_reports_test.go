package dao

import (
	"context"
	"testing"

	"github.com/content-services/content-sources-backend/pkg/config"
	ce "github.com/content-services/content-sources-backend/pkg/errors"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CoverageReportDaoSuite struct {
	*DaoSuite
}

func TestCoverageReportDaoSuite(t *testing.T) {
	m := DaoSuite{}
	suite.Run(t, &CoverageReportDaoSuite{DaoSuite: &m})
}

func (s *CoverageReportDaoSuite) coverageReportDao() coverageReportDaoImpl {
	return coverageReportDaoImpl{db: s.tx}
}

func (s *CoverageReportDaoSuite) createCoverageReport() (string, string) {
	uploadUUID := uuid.NewString()
	report, err := s.coverageReportDao().Create(context.Background(),
		CreateCoverageReportParams{
			OrgID:     orgIDTest,
			AccountID: utils.Ptr("account-1"),
		},
		CreateCoverageUploadParams{
			UUID:       uploadUUID,
			StorageKey: "coverage-uploads/" + uploadUUID,
			Sha256:     "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
			SizeBytes:  1024,
		},
	)
	require.NoError(s.T(), err)
	return report.UUID, uploadUUID
}

func (s *CoverageReportDaoSuite) TestCreateCoverageReport() {
	reportUUID, uploadUUID := s.createCoverageReport()

	var readReport models.CoverageReport
	err := s.tx.Where("uuid = ?", reportUUID).First(&readReport).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), config.TaskStatusPending, readReport.Status)
	assert.Equal(s.T(), orgIDTest, readReport.OrgID)
	assert.Equal(s.T(), "account-1", *readReport.AccountID)

	var readUpload models.CoverageUpload
	err = s.tx.Where("uuid = ?", uploadUUID).First(&readUpload).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), "coverage-uploads/"+uploadUUID, readUpload.StorageKey)
	assert.Equal(s.T(), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", readUpload.Sha256)
	assert.Equal(s.T(), int64(1024), readUpload.SizeBytes)
}

func (s *CoverageReportDaoSuite) TestFetchCoverageReport() {
	reportUUID, _ := s.createCoverageReport()

	report, err := s.coverageReportDao().Fetch(context.Background(), orgIDTest, reportUUID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), reportUUID, report.UUID)
	assert.Equal(s.T(), config.TaskStatusPending, report.Status)
}

func (s *CoverageReportDaoSuite) TestFetchCoverageReportWrongOrg() {
	reportUUID, _ := s.createCoverageReport()

	_, err := s.coverageReportDao().Fetch(context.Background(), "other-org", reportUUID)
	require.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
}

func (s *CoverageReportDaoSuite) TestInternalOnlyFetchCoverageUpload() {
	_, uploadUUID := s.createCoverageReport()

	upload, err := s.coverageReportDao().InternalOnlyFetchCoverageUpload(context.Background(), uploadUUID)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), uploadUUID, upload.UUID)
	assert.Equal(s.T(), "coverage-uploads/"+uploadUUID, upload.StorageKey)
}

func (s *CoverageReportDaoSuite) TestInternalOnlyFetchCoverageUploadNotFound() {
	_, err := s.coverageReportDao().InternalOnlyFetchCoverageUpload(context.Background(), uuid.NewString())
	require.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
}

func (s *CoverageReportDaoSuite) TestSetAnalysisTaskUUID() {
	reportUUID, _ := s.createCoverageReport()
	taskUUID := uuid.NewString()

	err := s.coverageReportDao().SetAnalysisTaskUUID(context.Background(), reportUUID, taskUUID)
	require.NoError(s.T(), err)

	var readReport models.CoverageReport
	err = s.tx.Where("uuid = ?", reportUUID).First(&readReport).Error
	require.NoError(s.T(), err)
	require.NotNil(s.T(), readReport.AnalysisTaskUUID)
	assert.Equal(s.T(), taskUUID, *readReport.AnalysisTaskUUID)
}

func (s *CoverageReportDaoSuite) TestSetAnalysisTaskUUIDNotFound() {
	err := s.coverageReportDao().SetAnalysisTaskUUID(context.Background(), uuid.NewString(), uuid.NewString())
	require.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
}

func (s *CoverageReportDaoSuite) TestUpdateCoverageReportStatus() {
	reportUUID, _ := s.createCoverageReport()

	err := s.coverageReportDao().UpdateCoverageReportStatus(context.Background(), reportUUID, config.TaskStatusRunning, nil)
	require.NoError(s.T(), err)

	errMsg := "analysis failed"
	err = s.coverageReportDao().UpdateCoverageReportStatus(context.Background(), reportUUID, config.TaskStatusFailed, &errMsg)
	require.NoError(s.T(), err)

	var readReport models.CoverageReport
	err = s.tx.Where("uuid = ?", reportUUID).First(&readReport).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), config.TaskStatusFailed, readReport.Status)
	require.NotNil(s.T(), readReport.AnalysisTaskError)
	assert.Equal(s.T(), errMsg, *readReport.AnalysisTaskError)
}

func (s *CoverageReportDaoSuite) TestUpdateCoverageReportStatusCompleted() {
	reportUUID, _ := s.createCoverageReport()

	err := s.coverageReportDao().UpdateCoverageReportStatus(context.Background(), reportUUID, config.TaskStatusCompleted, nil)
	require.NoError(s.T(), err)

	var readReport models.CoverageReport
	err = s.tx.Where("uuid = ?", reportUUID).First(&readReport).Error
	require.NoError(s.T(), err)
	assert.Equal(s.T(), config.TaskStatusCompleted, readReport.Status)
	require.NotNil(s.T(), readReport.CompletedAt)
}

func (s *CoverageReportDaoSuite) TestUpdateCoverageReportStatusInvalid() {
	reportUUID, _ := s.createCoverageReport()

	err := s.coverageReportDao().UpdateCoverageReportStatus(context.Background(), reportUUID, "bogus", nil)
	require.Error(s.T(), err)
	var daoErr *ce.DaoError
	require.ErrorAs(s.T(), err, &daoErr)
	assert.True(s.T(), daoErr.BadValidation)
}

func (s *CoverageReportDaoSuite) TestUpdateCoverageReportStatusNotFound() {
	err := s.coverageReportDao().UpdateCoverageReportStatus(context.Background(), uuid.NewString(), config.TaskStatusRunning, nil)
	require.Error(s.T(), err)
	assert.ErrorIs(s.T(), err, gorm.ErrRecordNotFound)
}
