package route

import (
	"time"

	"github.com/jeffersonroth/nangarekohara/api/controller"
	"github.com/jeffersonroth/nangarekohara/bootstrap"
	_ "github.com/jeffersonroth/nangarekohara/domain"
	"github.com/jeffersonroth/nangarekohara/repository"
	"github.com/jeffersonroth/nangarekohara/usecase"

	"github.com/labstack/echo/v4"
)

func NewDatabaseRouter(app *bootstrap.Application, timeout time.Duration, e *echo.Echo) {
	dbRepo := repository.NewDatabaseRepository(app.DB)
	dbUsecase := usecase.NewDatabaseUsecase(dbRepo, timeout)
	dbController := &controller.DatabaseController{DatabaseUsecase: dbUsecase}

	// Create new database connection details
	e.POST("/databases", dbController.Create)

	// Fetch database connection details
	e.GET("/databases", dbController.Fetch)

	// Update the database connection details
	e.PUT("/databases", dbController.Update)

	// Delete database connection details
	e.DELETE("/databases", dbController.Delete)

	// Validate that the database connection details are correct
	e.POST("/databases/validate", dbController.Validate)
}
