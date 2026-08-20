package tasks

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/content-services/content-sources-backend/pkg/dao"
	"github.com/content-services/content-sources-backend/pkg/models"
	"github.com/content-services/content-sources-backend/pkg/tasks/payloads"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type CoverageAnalysisSuite struct {
	suite.Suite
	mockDaoRegistry *dao.MockDaoRegistry
}

func TestCoverageAnalysisSuite(t *testing.T) {
	suite.Run(t, new(CoverageAnalysisSuite))
}

func (s *CoverageAnalysisSuite) SetupTest() {
	s.mockDaoRegistry = dao.GetMockDaoRegistry(s.T())
}

func (s *CoverageAnalysisSuite) TestVerifyCoverageManifest() {
	manifest := []byte("{\"components\":[]}")
	hash := sha256.Sum256(manifest)
	sha256Hex := hex.EncodeToString(hash[:])

	err := verifyCoverageManifest(manifest, models.CoverageUpload{
		Sha256:    sha256Hex,
		SizeBytes: int64(len(manifest)),
	})
	require.NoError(s.T(), err)

	err = verifyCoverageManifest(manifest, models.CoverageUpload{
		Sha256:    "wrongsha",
		SizeBytes: int64(len(manifest)),
	})
	require.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "sha256 mismatch")
}

func (s *CoverageAnalysisSuite) TestRunReturnsErrorWhenUploadNotFound() {
	ctx := context.Background()
	uploadUUID := uuid.NewString()
	payload := payloads.CoverageAnalysisPayload{
		CoverageReportUUID: uuid.NewString(),
		CoverageUploadUUID: uploadUUID,
	}

	s.mockDaoRegistry.CoverageReport.On("InternalOnlyFetchCoverageUpload", ctx, uploadUUID).
		Return(models.CoverageUpload{}, gorm.ErrRecordNotFound).Once()

	ca := CoverageAnalysis{
		ctx:     ctx,
		payload: &payload,
		daoReg:  s.mockDaoRegistry.ToDaoRegistry(),
	}

	err := ca.Run()
	require.Error(s.T(), err)
	assert.Contains(s.T(), err.Error(), "fetch coverage upload")
}
