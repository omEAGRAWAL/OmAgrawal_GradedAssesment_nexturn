package utils

// AppError represents a custom application error
type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(message string) error {
	return &AppError{Message: message}
}
