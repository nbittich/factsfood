package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/handlers"
	"github.com/nbittich/factsfood/jobs/manager"
	ffMidleware "github.com/nbittich/factsfood/middleware"
	"github.com/nbittich/factsfood/services/db"
	"github.com/nbittich/factsfood/services/email"
	"github.com/nbittich/factsfood/types"
)

//go:embed banner.txt
var BANNER string

func main() {
	defer db.Disconnect()
	defer close(email.MailChan)

	e := echo.New()

	// static assets
	e.Static("/assets", "assets")

	// static ext
	if _, err := os.Stat(config.StaticDirectory); os.IsNotExist(err) {
		os.MkdirAll(config.StaticDirectory, 0755)
	}
	e.Static("/static", config.StaticDirectory)

	// middleware
	// e.Pre(middleware.AddTrailingSlash()) interfer with POST form

	if config.GoEnv == config.DEVELOPMENT {
		e.Use(middleware.CORS())
	}

	if config.GoEnv == config.PRODUCTION {
		e.Use(middleware.Secure())
	}

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:csrf",
		Skipper: func(c echo.Context) bool {
			request := c.Request()
			accept := request.Header.Get(echo.HeaderAccept)
			return strings.Contains(c.Path(), "assets") || accept == echo.MIMEApplicationJSON
		},
	}))
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	// JWT
	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:             config.JWTSecretKey,
		TokenLookup:            fmt.Sprintf("header:Authorization:Bearer ,cookie:%s", config.JWTCookie),
		ContinueOnIgnoredError: true,
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.UserClaims)
		},
		ErrorHandler: ffMidleware.JWTErrorHandler,
	}))

	e.Use(ffMidleware.ValidateAuth)

	e.Use(ffMidleware.I18n)
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

	log.Println("starting job manager")
	go manager.Start()
	fmt.Println()

	fmt.Println(BANNER)

	handlers.UserRouter(e)
	handlers.HomeRouter(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
