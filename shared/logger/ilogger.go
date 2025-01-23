package logger

type ILogger interface {
	Info(message string)
	Error(message string)
}
