package config

// This enum holds error types, this is used to write custom logs
const (
	DB_ERROR      string = "DB_ERROR"
	AUTH_ERROR           = "AUTH_ERROR"
	REQUEST_ERROR        = "REQUEST_ERROR"
	SERVER_ERROR         = "SERVER_ERROR"
)

type CustomError struct {
	Message   string
	ErrorType string
}
