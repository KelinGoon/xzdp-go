package xzdp

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Data:    data,
	}
}

func NewFailureResponse(message string) *Response {
	return &Response{
		Success: false,
		Data:    map[string]string{"error": message},
	}
}
