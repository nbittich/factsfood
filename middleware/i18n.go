package middleware

import (
	"context"

	"github.com/BurntSushi/toml"
	"github.com/labstack/echo/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("i18n/fr.toml")
	bundle.LoadMessageFile("i18n/en.toml")
}

type i18nKey string

const I18nCtxKey = i18nKey("localizer")

func I18n(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		lang := r.FormValue("lang")
		accept := r.Header.Get("Accept-Language")

		localizer := i18n.NewLocalizer(bundle, lang, accept)
		c.SetRequest(r.WithContext(context.WithValue(r.Context(), I18nCtxKey, localizer)))
		return next(c)
	}
}
