package handlers

import (
	"context"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/types"
)

func renderHTML(statusCode int, c echo.Context, tpl templ.Component) error {
	c.Response().Status = statusCode
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	csrf := c.Get("csrf").(string)
	ctx := context.WithValue(c.Request().Context(), types.CsrfKey, csrf)
	return tpl.Render(ctx, c.Response().Writer)
}
