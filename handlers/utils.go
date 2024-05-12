package handlers

import (
	"context"

	"github.com/a-h/templ"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/types"
)

func renderHTML(statusCode int, c echo.Context, tpl templ.Component) error {
	c.Response().Status = statusCode
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	csrf := c.Get("csrf").(string)
	ctx := context.WithValue(c.Request().Context(), types.CsrfKey, csrf)
	msg := types.Message{}
	if err := c.Bind(&msg); err == nil && msg.Message != "" {
		ctx = context.WithValue(ctx, types.MessageKey, msg)
	}
	if tok, ok := c.Get("user").(*jwt.Token); ok {
		if user, ok := tok.Claims.(*types.UserClaims); ok {
			ctx = context.WithValue(ctx, types.UserKey, *user)
		}
	}

	return tpl.Render(ctx, c.Response().Writer)
}
