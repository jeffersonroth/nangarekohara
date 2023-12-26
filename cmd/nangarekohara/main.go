package main

import (
	route "jeffersonroth/nangarekohara/api/route"
	"jeffersonroth/nangarekohara/bootstrap"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	app := bootstrap.NewApplication()

	timeout := time.Duration(app.Env.ContextTimeout) * time.Second

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	route.Setup(&app, timeout, e)

	// Run
	e.Logger.Fatal(e.Start(":8080"))

}
