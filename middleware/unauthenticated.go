package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-querystring/query"
	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/types"
)

func UnauthenticatedOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		accept := request.Header.Get(echo.HeaderAccept)

		if tok, ok := c.Get("user").(*jwt.Token); ok {
			if _, ok := tok.Claims.(*types.UserClaims); ok {
				if accept == echo.MIMEApplicationJSON {
					return c.JSON(http.StatusForbidden, types.Message{Type: types.ERROR, Message: "Forbidden"})
				} else {
					message := types.Message{}
					message.Type = types.ERROR
					message.Message = "common.forbidden"
					v, _ := query.Values(message)
					return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?%s", v.Encode()))

				}
			}
		}
		return next(c)
	}
}
