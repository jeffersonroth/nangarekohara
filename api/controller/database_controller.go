package controller

import (
	"github.com/jeffersonroth/nangarekohara/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DatabaseController struct {
	DatabaseUsecase domain.DatabaseUsecase
}

func (dc *DatabaseController) Create(c echo.Context) error {
	var database domain.Database
	err := c.Bind(&database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	err = dc.DatabaseUsecase.Create(c.Request().Context(), database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Database connection details saved successfully"})
	return nil
}

func (dc *DatabaseController) Fetch(c echo.Context) error {
	var databases []domain.Database

	err := c.Bind(&databases)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	// Check if Name was provided. If yes then fetch by name, else GetActiveDatabases
	if len(databases) == 0 {
		databases, err := dc.DatabaseUsecase.Fetch(c.Request().Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return err
		}

		c.JSON(http.StatusOK, databases)
		return nil

	}

	database, err := dc.DatabaseUsecase.FetchByName(c.Request().Context(), databases[0].Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, database)
	return nil
}

func (dc *DatabaseController) Update(c echo.Context) error {
	var database domain.Database
	err := c.Bind(&database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	err = dc.DatabaseUsecase.Update(c.Request().Context(), database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Database connection details updated successfully"})
	return nil
}

func (dc *DatabaseController) Delete(c echo.Context) error {
	var database domain.Database
	err := c.Bind(&database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	err = dc.DatabaseUsecase.Delete(c.Request().Context(), database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Database connection details removed successfully"})
	return nil
}

func (dc *DatabaseController) Validate(c echo.Context) error {
	var database domain.Database
	err := c.Bind(&database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	err = dc.DatabaseUsecase.Validate(c.Request().Context(), database)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return err
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Database connection details validated successfully"})
	return nil
}
