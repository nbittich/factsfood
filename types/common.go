package types

import "encoding/json"

type Identifiable interface {
	GetID() string
	SetID(id string)
}

const (
	I18nKey            = CtxKey("localizer")
	LangKey            = CtxKey("lang")
	SignupFormErrorKey = CtxKey("signupFormError")
)

type (
	CtxKey           string
	InvalidMessage   = map[string]interface{}
	InvalidFormError struct {
		Form     interface{}
		Messages InvalidMessage `json:"messages"`
	}
)

func (apiError InvalidFormError) Error() string {
	val, e := json.Marshal(apiError.Messages)
	if e != nil {
		return e.Error()
	}
	return string(val)
}
