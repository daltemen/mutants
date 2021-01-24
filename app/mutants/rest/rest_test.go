package rest

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"mutants/app/mutants/managers"
	mocks "mutants/mocks/app/mutants/managers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type restSuite struct {
	suite.Suite
	ctx     context.Context
	manager *mocks.MutantManager
	rest    *Rest
}

func (suite *restSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.manager = new(mocks.MutantManager)
	suite.rest = NewRest(suite.manager)
}

func TestRestSuite(t *testing.T) {
	suite.Run(t, &restSuite{})
}

func (suite *restSuite) TestRest_GetStats_Success() {
	statsMock := managers.StatsSummary{
		Mutants: int64(40),
		Humans:  int64(100),
		Ratio:   0.4,
	}
	suite.manager.Mock.On("RetrieveStats", suite.ctx).Return(statsMock, nil).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.GetStats(c)
	expected := `{"count_mutant_dna":40,"count_human_dna":100,"ratio":0.4}
`
	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
	suite.Equal(expected, rec.Body.String())
}

func (suite *restSuite) TestRest_GetStats_Failed() {
	suite.manager.Mock.On("RetrieveStats", suite.ctx).Return(managers.StatsSummary{}, errors.New("unknown")).Once()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", strings.NewReader(""))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.GetStats(c)
	suite.NoError(err)
	suite.Equal(http.StatusInternalServerError, rec.Code)
}

func (suite *restSuite) TestRest_PostMutant_Ok() {
	suite.manager.Mock.On("IsMutant", mock.Anything, mock.Anything).Return(true, nil).Once()

	e := echo.New()
	dna := `{ "dna": [ "ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG" ] }`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/mutant", strings.NewReader(dna))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.PostMutant(c)

	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)
}

func (suite *restSuite) TestRest_PostMutant_Forbidden() {
	suite.manager.Mock.On("IsMutant", mock.Anything, mock.Anything).Return(false, nil).Once()

	e := echo.New()
	dna := `{ "dna": [ "ATGCGA", "CTGTCC", "TTGTGT", "AGAAGG", "CGTCTA", "TCACTG" ] }`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/mutant", strings.NewReader(dna))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.PostMutant(c)

	suite.NoError(err)
	suite.Equal(http.StatusForbidden, rec.Code)
}

func (suite *restSuite) TestRest_PostMutant_Error() {
	suite.manager.Mock.On("IsMutant", mock.Anything, mock.Anything).Return(false, errors.New("unknown")).Once()

	e := echo.New()
	dna := `{ "dna": [ "ATGCGA", "CTGTCC", "TTGTGT", "AGAAGG", "CGTCTA", "TCACTG" ] }`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/mutant", strings.NewReader(dna))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := suite.rest.PostMutant(c)

	suite.NoError(err)
	suite.Equal(http.StatusInternalServerError, rec.Code)
}
