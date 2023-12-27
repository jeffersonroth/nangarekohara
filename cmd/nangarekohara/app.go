package main

import (
	"database/sql"
	"html/template"
	"io"
	"log"

	"github.com/jeffersonroth/nangarekohara/pkg/database"
	"github.com/jeffersonroth/nangarekohara/pkg/pages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	DB *sql.DB
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (a *App) Initialize() {
	var err error
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	a.DB, err = db.DB()
	if err != nil {
		log.Fatal(err)
	}

	database.DB = db
}

func (a *App) Run(addr string) {
	funcMap := template.FuncMap{
		"loop": func(from, to int) <-chan int {
			ch := make(chan int)
			go func() {
				for i := from; i <= to; i++ {
					ch <- i
				}
				close(ch)
			}()
			return ch
		},
	}

	tmpls, err := template.New("").Funcs(funcMap).ParseGlob("public/views/*.html")

	if err != nil {
		log.Fatalf("couldn't initialize templates: %v", err)
	}

	e := echo.New()
	e.Renderer = &TemplateRenderer{
		templates: tmpls,
	}

	e.Use(middleware.Logger())
	e.Static("/dist", "dist")
	e.Static("/css", "css")

	// Index
	e.GET("/", pages.Index)

	// Database routes
	e.POST("/add-database", database.NewDatabaseHandler)
	e.GET("/database", database.GetDatabaseHandler)

	e.Logger.Fatal(e.Start(addr))
}
