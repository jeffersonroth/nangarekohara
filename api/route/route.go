package route

import (
	"jeffersonroth/nangarekohara/bootstrap"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *bootstrap.Application, timeout time.Duration, e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	NewDatabaseRouter(app, timeout, e)
}
