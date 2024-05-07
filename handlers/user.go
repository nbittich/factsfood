package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services"
	"github.com/nbittich/factsfood/types"
	"github.com/nbittich/factsfood/views"
	"github.com/nbittich/factsfood/views/utils"
)

func UserRouter(e *echo.Echo) {
	e.POST("/users/new", newUserHandler).Name = "users.New"
	e.GET("/users/activate", activateUserHandler).Name = "users.Activate"
}

func activateUserHandler(c echo.Context) error {
	hash := c.QueryParam("hash")
	request := c.Request()
	accept := request.Header.Get(echo.HeaderAccept)
	ctx, cancel := context.WithTimeout(c.Request().Context(), config.MongoCtxTimeout)

	defer cancel()
	active, err := services.ActivateUser(ctx, hash)
	if err != nil {
		c.Logger().Error("could not activate user: ", err)
	}
	message := types.Message{}
	if active {
		message.Type = types.SUCCESS
		message.Message = "home.signup.user.activated"
	} else {
		message.Type = types.ERROR
		message.Message = "home.signup.user.notActivated"
	}

	if accept == echo.MIMEApplicationJSON {
		message.Message = utils.T(c.Request().Context(), message.Message)
		return c.JSON(http.StatusOK, message)
	} else {
		c.SetRequest(request.WithContext(context.WithValue(request.Context(), types.MessageKey, message)))
		return renderHTML(http.StatusOK, c, views.Home("Home"))
	}
}

func newUserHandler(c echo.Context) error {
	request := c.Request()
	accept := request.Header.Get(echo.HeaderAccept)
	newUserForm := types.NewUserForm{}
	if err := c.Bind(&newUserForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Logger().Debug(newUserForm)
	ctx, cancel := context.WithTimeout(c.Request().Context(), config.MongoCtxTimeout)
	defer cancel()
	user, err := services.NewUser(ctx, &newUserForm)
	if err != nil {
		if err, ok := err.(types.InvalidFormError); ok {
			err.Form = newUserForm
			if accept == echo.MIMEApplicationJSON {
				return c.JSON(http.StatusBadRequest, err)
			} else {
				c.SetRequest(request.WithContext(context.WithValue(request.Context(), types.SignupFormErrorKey, err)))
				return renderHTML(http.StatusOK, c, views.Home("Home"))
			}
		}
		c.Logger().Error("Unexpected error when creating a new user:", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "unexpected error while creating new user")
	}
	c.Logger().Debug(user)

	if accept == echo.MIMEApplicationJSON {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}
}
