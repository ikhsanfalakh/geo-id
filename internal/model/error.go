package model

// ErrorResponse represents an error response
// @Description Error information
// @name ErrorResponse
type ErrorResponse struct {
	Error string `json:"error" example:"Region not found"`
}
