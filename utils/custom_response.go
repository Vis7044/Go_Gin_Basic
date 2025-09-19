package utils

type ResponseHandler[T any] struct {
	Status bool `json:"status"`
	Data T `json:"message"`
}