package model

// APIResponse represents a successful API response
// @Description Successful API response wrapper
type APIResponse struct {
	Status  int         `json:"status" example:"200"`
	Message string      `json:"message" example:"SUCCESS"`
	Data    interface{} `json:"data"`
}

// APIErrorResponse represents an error API response
// @Description Error API response wrapper
type APIErrorResponse struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"NOT_FOUND"`
	Error   string `json:"error" example:"Region not found"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}) APIResponse {
	return APIResponse{
		Status:  200,
		Message: "SUCCESS",
		Data:    data,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(status int, message string, err error) APIErrorResponse {
	return APIErrorResponse{
		Status:  status,
		Message: message,
		Error:   err.Error(),
	}
}
