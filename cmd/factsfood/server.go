package main

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e.Logger.Fatal(e.Start(":1323"))
}
