package helpers

type ResponseSuccess struct {
	Status  bool        `json:"staus"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewResponseSuccess(data interface{}, message string) *ResponseSuccess {
	return &ResponseSuccess{
		Status:  true,
		Data:    data,
		Message: message,
	}
}
