package utils

import (
	"context"

	"github.com/nbittich/factsfood/middleware"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func T(c context.Context, id string) string {
	msg := &i18n.Message{ID: id}
	lz, ok := c.Value(middleware.I18nCtxKey).(*i18n.Localizer)
	if ok {
		msg, e := lz.LocalizeMessage(msg)
		if e == nil {
			return msg
		}
	}

	return id
}
