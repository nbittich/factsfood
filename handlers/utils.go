package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderHTML(statusCode int, c echo.Context, tpl templ.Component) error {
	c.Response().Status = statusCode
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return tpl.Render(c.Request().Context(), c.Response().Writer)
}
