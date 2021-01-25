package managers

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"mutants/app/mutants"
	mocks "mutants/mocks/app/mutants/repositories"
	"testing"
)

type mutantManagerSuite struct {
	suite.Suite
	ctx               context.Context
	readRepo          *mocks.ReadRepository
	strongWriteRepo   *mocks.StrongWriteRepository
	eventualWriteRepo *mocks.EventualWriteRepository
	manager           MutantManager
}

func (suite *mutantManagerSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.readRepo = new(mocks.ReadRepository)
	suite.strongWriteRepo = new(mocks.StrongWriteRepository)
	suite.eventualWriteRepo = new(mocks.EventualWriteRepository)
	suite.manager = NewMutantManager(suite.readRepo, suite.strongWriteRepo, suite.eventualWriteRepo)
}

func TestMutantManagerSuite(t *testing.T) {
	suite.Run(t, &mutantManagerSuite{})
}

func (suite *mutantManagerSuite) Test_mutantManager_IsMutant_Success() {
	suite.strongWriteRepo.Mock.On("SaveDna", mock.Anything, mock.Anything).Return(nil).Once()
	suite.eventualWriteRepo.Mock.On("IncrementCount", mock.Anything, mock.Anything).Return(nil).Once()
	dna := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	isMutant, err := suite.manager.IsMutant(suite.ctx, dna)
	suite.True(isMutant)
	suite.NoError(err)
}

func (suite *mutantManagerSuite) Test_mutantManager_IsMutant_StrongRepoFailed() {
	suite.strongWriteRepo.Mock.On("SaveDna", mock.Anything, mock.Anything).Return(errors.New("any")).Once()
	dna := make([][]string, 0)
	isMutant, err := suite.manager.IsMutant(suite.ctx, dna)
	suite.False(isMutant)
	suite.Error(err)
}

func (suite *mutantManagerSuite) Test_mutantManager_IsMutant_EventualRepoFailed() {
	suite.strongWriteRepo.Mock.On("SaveDna", mock.Anything, mock.Anything).Return(nil).Once()
	suite.eventualWriteRepo.Mock.On("IncrementCount", mock.Anything, mock.Anything).Return(errors.New("any")).Once()
	dna := make([][]string, 0)
	isMutant, err := suite.manager.IsMutant(suite.ctx, dna)
	suite.False(isMutant)
	suite.Error(err)
}

func (suite *mutantManagerSuite) Test_mutantManager_RetrieveStats_Success() {
	mutantsCount := int64(40)
	humansCount := int64(100)
	humanStatsMock := mutants.NewHumanStats(mutantsCount, humansCount)
	suite.readRepo.Mock.On("GetStats", mock.Anything).Return(humanStatsMock, nil).Once()

	stats, err := suite.manager.RetrieveStats(suite.ctx)
	suite.NoError(err)
	suite.Equal(mutantsCount, stats.Mutants)
	suite.Equal(humansCount, stats.Humans)
	suite.Equal(float32(0.40), stats.Ratio)
}

func (suite *mutantManagerSuite) Test_mutantManager_RetrieveStats_Failed() {
	suite.readRepo.Mock.On("GetStats", mock.Anything).Return(mutants.HumanStats{}, errors.New("unknown")).Once()
	stats, err := suite.manager.RetrieveStats(suite.ctx)
	suite.Error(err)
	suite.Equal(StatsSummary{}, stats)
}
