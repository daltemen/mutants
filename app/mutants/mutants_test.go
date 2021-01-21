package mutants

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type mutantSuite struct {
	suite.Suite
}

func (suite *mutantSuite) SetupTest() {}

func TestMutantSuite(t *testing.T) {
	suite.Run(t, &mutantSuite{})
}

func (suite *mutantSuite) TestHuman_IsMutant() {
	dnaMutants := [][][]string{
		{
			{"A", "T", "G", "C", "G", "A"},
			{"C", "A", "G", "T", "G", "C"},
			{"T", "T", "A", "T", "G", "T"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "C", "C", "C", "T", "A"},
			{"T", "C", "A", "C", "T", "G"},
		},
		{
			{"A", "T", "G", "C", "T", "A"},
			{"C", "C", "G", "T", "G", "C"},
			{"T", "T", "A", "T", "G", "T"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "G", "C", "C", "A", "A"},
			{"T", "C", "A", "C", "T", "A"},
		},
		{
			{"A", "T", "G", "C", "T", "A"},
			{"C", "C", "G", "T", "G", "C"},
			{"C", "C", "C", "C", "G", "G"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "C", "C", "A", "A", "G"},
			{"T", "C", "A", "C", "T", "G"},
		},
		{
			{"A", "T", "G", "C", "G", "A"},
			{"C", "C", "G", "T", "G", "C"},
			{"C", "C", "C", "G", "G", "A"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "C", "C", "A", "G", "G"},
			{"T", "C", "A", "C", "T", "G"},
		},
	}

	for _, testCase := range dnaMutants {
		human := NewHuman(testCase)
		isMutant := human.IsMutant()
		suite.True(isMutant)
	}

	dnaHumans := [][][]string{
		{
			{"A", "T", "G", "C", "G", "A"},
			{"C", "T", "G", "T", "C", "C"},
			{"T", "T", "G", "T", "G", "T"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "G", "C", "C", "T", "A"},
			{"T", "C", "A", "C", "T", "G"},
		},
		{
			{"C", "T", "G", "T", "C", "C"},
			{"T", "C", "A", "C", "T", "G"},
			{"A", "G", "A", "A", "G", "G"},
			{"C", "G", "C", "C", "T", "A"},
			{"T", "T", "G", "T", "G", "T"},
			{"A", "T", "G", "C", "G", "A"},
		},
	}

	for _, testCase := range dnaHumans {
		human := NewHuman(testCase)
		isMutant := human.IsMutant()
		suite.False(isMutant)
	}
}

func (suite *mutantSuite) TestHumanStats_GetRatio() {
	stats := NewHumanStats(40, 100)
	ratio := stats.GetRatio()
	suite.Equal(float32(0.4), ratio)
	suite.Equal(int64(40), stats.Mutants)
	suite.Equal(int64(100), stats.Humans)
}