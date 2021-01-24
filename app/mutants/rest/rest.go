package rest

import (
	"github.com/labstack/echo/v4"
	"mutants/app/mutants/managers"
	"net/http"
	"strings"
)

type Rest struct {
	mutantManager managers.MutantManager
}

func NewRest(mutantManager managers.MutantManager) *Rest {
	return &Rest{mutantManager: mutantManager}
}

func (r *Rest) GetStats(c echo.Context) error {
	stats, err := r.mutantManager.RetrieveStats(c.Request().Context())
	if err != nil {
		return err // TODO: handle err
	}

	return c.JSON(http.StatusOK, &StatsResponse{
		CountMutantDna: stats.Mutants,
		CountHumanDna:  stats.Humans,
		Ratio:          stats.Ratio,
	})
}

func (r *Rest) PostMutant(c echo.Context) error {
	dnaRequest := new(DnaRequest)
	if err := c.Bind(dnaRequest); err != nil {
		return err
	}
	dna := mapSegmentToMatrix(dnaRequest.Dna)

	isMutant, err := r.mutantManager.IsMutant(c.Request().Context(), dna)
	// TODO: handle err
	if err != nil {
		return err
	}
	if isMutant {
		return c.JSON(http.StatusOK, map[string]string{})
	}
	return c.JSON(http.StatusForbidden, map[string]string{})
}

func mapSegmentToMatrix(segments []string) [][]string {
	result := make([][]string, len(segments))
	for i, segment := range segments {
		s := strings.Split(segment, "")
		result[i] = s
	}
	return result
}
