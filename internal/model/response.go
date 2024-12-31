package model

type ApiError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"data"`
}

func (e *ApiError) Error() string {
	return e.Message
}

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   *ApiError   `json:"error"`
}
