package main

import (
	_ "embed"
	"fmt"
	"strings"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/handlers"
	ffMidleware "github.com/nbittich/factsfood/middleware"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/services/email"
)

//go:embed banner.txt
var BANNER string

func main() {
	defer db.Disconnect()
	defer close(email.MailChan)

	e := echo.New()

	// static assets
	e.Static("/assets", "assets")

	// middleware
	// e.Pre(middleware.AddTrailingSlash()) interfer with POST form

	if config.GoEnv == config.DEVELOPMENT {
		e.Use(middleware.CORS())
	} else {
		e.Use(middleware.Secure()) // fixme, could cause issues I guess? not tested
	}
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "assets")
		},
	}))
	e.Use(middleware.Gzip())
	e.Use(ffMidleware.I18n)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:             config.JWTSecretKey,
		TokenLookup:            fmt.Sprintf("header:Authorization:Bearer, cookie:%s", config.JWTCookie),
		ContinueOnIgnoredError: true,
		ErrorHandler: func(c echo.Context, err error) error {
			if !strings.Contains(c.Path(), "/u/") {
				return nil
			}
			return err
		},
	}))

	// end middleware

	e.HideBanner = true
	e.Logger.SetLevel(config.LogLevel)

	// email consume logs
	go func() {
		for msg := range email.MailChan {
			switch msg.(type) {
			case string:
				e.Logger.Info(msg)
			case error:
				e.Logger.Error(e)
			}
		}
	}()

	fmt.Println(BANNER)

	handlers.UserRouter(e)
	e.GET("/", handlers.HomeHandler).Name = "home"
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
