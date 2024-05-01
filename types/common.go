package types

type Identifiable interface {
	GetID() string
	SetID(id string)
}

type APIError struct {
	Message string `json:"message"`
}

func (apiError APIError) Error() string {
	return apiError.Message
}
