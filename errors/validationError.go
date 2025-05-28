package errors

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func NewValidationError(msg string) *ValidationError {
	return &ValidationError{
		Message: msg,
	}
}
