package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/views"
)

func HomeHandler(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		name = "home.world"
	}
	return renderHTML(http.StatusOK, c, views.Home(name))
}
