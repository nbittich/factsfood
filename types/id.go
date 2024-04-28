package types

type Identifiable interface {
	GetID() string
	SetID(id string)
}
