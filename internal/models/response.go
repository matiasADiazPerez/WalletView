package models

type Response struct {
	Message string
	Payload any
}

type ErrorWrapper struct {
	Message string
	Error   error
	Code    int
}
