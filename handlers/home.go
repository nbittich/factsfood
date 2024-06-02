package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/config"
	"github.com/nbittich/factsfood/services"
	"github.com/nbittich/factsfood/services/db"
	offTypes "github.com/nbittich/factsfood/types/openfoodfacts"
	"github.com/nbittich/factsfood/views"
)

var offService = &services.OFFService{}

func HomeRouter(e *echo.Echo) {
	e.POST("/search", searchHandler).Name = "home.search"
	e.GET("/", homeHandler).Name = "home.root"
}

func homeHandler(c echo.Context) error {
	return renderHTML(http.StatusOK, c, views.Home())
}

func searchHandler(c echo.Context) error {
	searchForm := &offTypes.OFFSearchCriteria{Page: db.PageOptions{
		PageNumber: 1,
		PageSize:   10,
	}}
	if err := c.Bind(searchForm); err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(c.Request().Context(), config.MongoCtxTimeout)
	defer cancel()
	res, err := offService.Search(ctx, searchForm)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", res)
	return nil
}
