package types

import "encoding/json"

type Identifiable interface {
	GetID() string
	SetID(id string)
}

type (
	InvalidMessage   = map[string]interface{}
	InvalidFormError struct {
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
