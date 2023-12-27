package database

import (
	"html/template"

	"github.com/jeffersonroth/nangarekohara/pkg/database/models"

	"github.com/labstack/echo/v4"
)

func GetDatabaseHandler(c echo.Context) error {
	dbs, err := models.GetActiveDatabases(DB)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"Databases": dbs,
	}

	templ := template.Must(template.ParseFiles("public/views/database.html"))
	templ.ExecuteTemplate(c.Response().Writer, "database-list", data)

	return nil
}

func NewDatabaseHandler(c echo.Context) error {
	url := models.CreateURL(c.FormValue("host"), c.FormValue("port"), c.FormValue("user"), c.FormValue("password"), c.FormValue("dbname"))
	database := models.Database{
		Database: c.FormValue("database"),
		Name:     c.FormValue("name"),
		URL:      url,
	}

	DB.Create(&database)

	dbs, err := models.GetActiveDatabases(DB)
	if err != nil {
		return err
	}
	data := map[string]interface{}{
		"Databases": dbs,
	}

	templ := template.Must(template.ParseFiles("public/views/database.html"))
	templ.ExecuteTemplate(c.Response().Writer, "database-list-element", data)

	return nil
}
