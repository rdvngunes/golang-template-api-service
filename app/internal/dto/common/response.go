package common

type Response struct {
	APIVersion string        `json:"api_version"`
	Status     string        `json:"status"`
	Data       any           `json:"data,omitempty"`  // Omit if empty
	Error      *ErrorMessage `json:"error,omitempty"` // Omit if empty
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func NewResponse(data any, status string) Response {
	return Response{
		APIVersion: "1.0", // Set your default API version here
		Status:     status,
		Data:       data,
	}
}

func NewErrorResponse(message string, status string) Response {
	return Response{
		APIVersion: "1.0", // Set your default API version here
		Status:     status,
		Error:      &ErrorMessage{Message: message},
	}
}
