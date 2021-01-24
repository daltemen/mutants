package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type EchoRestType func(c echo.Context) error

func RegisterRoutes(e *echo.Echo, rest *Rest) {
	e.GET("/health", Health)
	e.GET("/api/v1/stats", rest.GetStats)
	e.POST("/api/v1/mutant", rest.PostMutant)
}

func Health(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
