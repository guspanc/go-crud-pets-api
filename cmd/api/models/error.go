package models

// APIError model
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewAPIError returns new api error
func NewAPIError(code, message string) *APIError {
	return &APIError{code, message}
}

// Error implementation
func (e *APIError) Error() string {
	return e.Message
}
