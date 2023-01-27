package helpers

type ResponseSuccess struct {
	Status  bool        `json:"staus"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponseSuccess(message string, data interface{}) *ResponseSuccess {
	return &ResponseSuccess{
		Status:  true,
		Data:    data,
		Message: message,
	}
}
