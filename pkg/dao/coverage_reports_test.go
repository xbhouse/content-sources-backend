package dao

import (
	"context"
	"testing"

	"github.com/content-services/content-sources-backend/pkg/config"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CoverageReportDaoSuite struct {
	*DaoSuite
}

func TestCoverageReportDaoSuite(t *testing.T) {
	m := DaoSuite{}
	suite.Run(t, &CoverageReportDaoSuite{DaoSuite: &m})
}

func (s *CoverageReportDaoSuite) TestCreateCoverageReport() {
	dao := coverageReportDaoImpl{db: s.tx}

	uploadUUID := uuid.NewString()
	report, err := dao.CreateCoverageReport(context.Background(),
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

	var readReport models.CoverageReport
	err = s.tx.Where("uuid = ?", report.UUID).First(&readReport).Error
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
