package main

import (
	_ "embed"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/handlers"
	ffMidleware "github.com/nbittich/factsfood/middleware"
	"github.com/nbittich/factsfood/services/db"
)

//go:embed banner.txt
var BANNER string

func main() {
	e := echo.New()
	defer db.Disconnect()
	// static assets
	e.Static("/assets", "assets")

	// middleware
	// e.Pre(middleware.AddTrailingSlash()) interfer with POST form

	if config.GoEnv == config.DEVELOPMENT {
		e.Use(middleware.CORS())
	}

	e.Use(ffMidleware.I18n)
	e.HideBanner = true
	e.Logger.SetLevel(config.LogLevel)
	e.Use(middleware.Logger())
	fmt.Println(BANNER)

	handlers.UserRouter(e)
	e.GET("/", handlers.HomeHandler).Name = "home"
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
