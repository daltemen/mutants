package repositories

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	mocket "github.com/selvatico/go-mocket"
	"github.com/stretchr/testify/suite"
	"mutants/app/mutants"
	"testing"
)

type mysqlStrongRepoSuite struct {
	suite.Suite
	repository StrongWriteRepository
	ctx        context.Context
}

func (suite *mysqlStrongRepoSuite) SetupTest() {
	suite.ctx = context.Background()
	mocket.Catcher.Register()
	mocket.Catcher.Logging = true
	db, _ := gorm.Open(mocket.DriverName, "connection_mock")
	suite.repository = NewMySqlStrongRepository(db)
}

func TestMutantSuite(t *testing.T) {
	suite.Run(t, &mysqlStrongRepoSuite{})
}

func (suite *mysqlStrongRepoSuite) BeforeTest(_, _ string) {
	mocket.Catcher.Reset()
}

func (suite *mysqlStrongRepoSuite) Test_mySqlStrongWriteRepository_SaveDna_Success() {
	dna := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	human := mutants.NewHuman(dna)
	err := suite.repository.SaveDna(suite.ctx, human)
	suite.NoError(err)
}

func (suite *mysqlStrongRepoSuite) Test_mySqlStrongWriteRepository_SaveDna_Failed() {
	dna := make([][]string, 0)
	human := mutants.NewHuman(dna)
	mocket.Catcher.NewMock().WithError(&mysql.MySQLError{Number: 1062, Message: "Error Entry"})
	err := suite.repository.SaveDna(suite.ctx, human)
	suite.Error(err)
}
