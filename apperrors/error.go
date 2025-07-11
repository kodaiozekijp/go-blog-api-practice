package apperrors

type MyAppError struct {
	ErrCode
	Message string
}
