package types

import "encoding/json"

type Identifiable interface {
	GetID() string
	SetID(id string)
}
type ErrorMessage struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
type (
	InvalidMessage []ErrorMessage

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
