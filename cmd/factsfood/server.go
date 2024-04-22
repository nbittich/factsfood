package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nbittich/factsfood/config"
)

//go:embed banner.txt
var BANNER string

func main() {
	e := echo.New()
	// middleware
	e.Pre(middleware.AddTrailingSlash())

	if config.GoEnv == config.DEVELOPMENT {
		e.Use(middleware.CORS())
	}

	e.HideBanner = true
	e.Logger.SetLevel(config.LogLevel)
	e.Use(middleware.Logger())
	println(BANNER)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
