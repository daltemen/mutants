package cmd

import (
	"github.com/labstack/echo/v4"
	"log"
	"mutants/app/presenters"
)

func Run() *echo.Echo {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return presenters.RunRestServer()
}
