package types

import "encoding/json"

type HasID interface {
	GetID() string
}
type Identifiable interface {
	HasID
	SetID(id string)
}

const (
	I18nKey            = CtxKey("localizer")
	LangKey            = CtxKey("lang")
	CsrfKey            = CtxKey("csrf")
	SignupFormErrorKey = CtxKey("signupFormError")
	SigninFormErrorKey = CtxKey("signinFormError")
	MessageKey         = CtxKey("message")
	UserKey            = CtxKey("user")
)

type StatusType int8

const (
	INFO StatusType = iota
	SUCCESS
	WARNING
	ERROR
)

type Message struct {
	Type    StatusType `json:"type" url:"type" param:"type" form:"type" query:"type"`
	Message string     `json:"message" url:"message" param:"message" form:"message" query:"message"`
}

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
