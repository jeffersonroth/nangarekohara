package pages

import (
	"github.com/labstack/echo/v4"
)

type IndexPage struct {
    Columns int
}

func Index(c echo.Context) error {
    return c.Render(200, "index.html", IndexPage{
        Columns: 40,
    })
}