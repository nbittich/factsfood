package main

import (
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/config"
	"net/http"
)

//go:embed banner.txt
var BANNER string

func main() {
	e := echo.New()
	e.HideBanner = true
	fmt.Println(BANNER)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", config.Host, config.Port)))
}
