package presenters

import (
	"github.com/labstack/echo/v4"
	"mutants/app/datasources"
	_managers "mutants/app/mutants/managers"
	"mutants/app/mutants/repositories"
	"mutants/app/mutants/rest"
	"os"
)

// TODO: use a library to handle dependency injection
func RunRestServer() {
	// -- init db and data-sources
	db := datasources.ConnectDb()
	datasources.Migrate(db)
	redis := datasources.ConnectRedis()

	// -- repositories
	redisReadRepo := repositories.NewRedisReadRepository(redis)
	redisWriteRepo := repositories.NewRedisEventualRepository(redis)
	mysqlWriteRepo := repositories.NewMySqlStrongRepository(db)

	managers := _managers.NewMutantManager(redisReadRepo, mysqlWriteRepo, redisWriteRepo)
	restMethods := rest.NewRest(managers)

	// -- init rest server
	e := echo.New()
	rest.RegisterRoutes(e, restMethods)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
