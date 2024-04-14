package utils

// This enum holds error types, this is used to write custom logs
const (
	DB_ERROR      string = "DB_ERROR"
	AUTH_ERROR           = "AUTH_ERROR"
	REQUEST_ERROR        = "REQUEST_ERROR"
	SERVER_ERROR         = "SERVER_ERROR"
)

// This is a custom error struct to handle the writing these in error files. Can be any of the error types above
type CustomError struct {
	Message   string
	ErrorType string
}
