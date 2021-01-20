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
	// TODO: write more test cases
	dna := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	human := NewHuman(dna)
	isMutant := human.IsMutant()
	suite.True(isMutant)
}
