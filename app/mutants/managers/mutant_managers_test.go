package managers

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type mutantManagerSuite struct {
	suite.Suite
}

func (suite *mutantManagerSuite) SetupTest() {}

func TestMutantManagerSuite(t *testing.T) {
	suite.Run(t, &mutantManagerSuite{})
}

func (suite *mutantManagerSuite) Test_mutantManager_IsMutant() {

}

func (suite *mutantManagerSuite) Test_mutantManager_RetrieveStats() {

}
