package middleware

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/go-querystring/query"
	"github.com/labstack/echo/v4"
	"github.com/nbittich/factsfood/types"
)

type AuthConfig struct {
	Pattern         string       `json:"pattern"`
	Authenticated   bool         `json:"authenticated"`
	Unauthenticated bool         `json:"unauthenticated"`
	Roles           []types.Role `json:"roles"`
}

//go:embed auth_config.json
var authConfigFile []byte
var authConfigs []AuthConfig

func init() {
	if err := json.Unmarshal(authConfigFile, &authConfigs); err != nil {
		panic(err)
	}
}

func JWTErrorHandler(c echo.Context, err error) error {
	for _, ac := range authConfigs {
		if m, _ := regexp.MatchString(ac.Pattern, c.Path()); m {
			if ac.Authenticated {
				return err
			}
		}
	}
	return nil
}

func getUser(c echo.Context) *types.UserClaims {
	if tok, ok := c.Get("user").(*jwt.Token); ok {
		if user, ok := tok.Claims.(*types.UserClaims); ok {
			return user
		}
	}
	return nil
}

func forbidden(c echo.Context) error {
	request := c.Request()

	accept := request.Header.Get(echo.HeaderAccept)

	if accept == echo.MIMEApplicationJSON {
		return c.JSON(http.StatusForbidden, types.Message{Type: types.ERROR, Message: "Forbidden"})
	} else {
		message := types.Message{}
		message.Type = types.ERROR
		message.Message = "common.forbidden"
		v, _ := query.Values(message)
		return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/?%s", v.Encode()))

	}
}

func ValidateAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := getUser(c)

		for _, ac := range authConfigs {
			if m, _ := regexp.MatchString(ac.Pattern, c.Path()); m {
				if ac.Unauthenticated && user != nil {
					return forbidden(c)
				}
				if ac.Authenticated && user == nil {
					return forbidden(c)
				}
				if ac.Authenticated && len(ac.Roles) > 0 {
					mapElt := make(map[types.Role]bool, len(user.Roles))
					for _, r := range user.Roles {
						mapElt[r] = true
					}
					for _, r := range ac.Roles {
						if !mapElt[r] {
							return forbidden(c)
						}
					}

				}
			}
		}
		return next(c)
	}
}
