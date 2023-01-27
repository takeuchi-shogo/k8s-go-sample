package helpers

type ResponseError struct {
	Error struct {
		Status      bool   `json:"staus"`
		Code        int    `json:"code"`
		Message     string `json:"message"`
		InfoMessage string `json:"info_message"`
	} `json:"error"`
}

func NewResponseError(code int, err error, message string) *ResponseError {
	r := new(ResponseError)

	r.Error.Status = false
	r.Error.Code = code
	r.Error.Message = err.Error()
	r.Error.InfoMessage = message

	return r
}
