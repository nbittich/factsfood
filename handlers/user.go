package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services"
	"github.com/nbittich/factsfood/types"
)

func UserRouter(e *echo.Echo) {
	e.POST("/users/new", newUserHandler).Name = "users.New"
}

func newUserHandler(c echo.Context) error {
	newUserForm := new(types.NewUserForm)
	if err := c.Bind(newUserForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Logger().Debug(newUserForm)
	ctx, cancel := context.WithTimeout(c.Request().Context(), config.MongoCtxTimeout)
	defer cancel()
	user, err := services.NewUser(ctx, newUserForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Logger().Debug(user)

	if c.Request().Header.Get(echo.HeaderAccept) == echo.MIMEApplicationJSON {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}
