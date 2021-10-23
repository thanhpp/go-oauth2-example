package httpserver

type ErrorDTO struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e *ErrorDTO) SetCode(code int) *ErrorDTO {
	e.Error.Code = code
	return e
}

func (e *ErrorDTO) SetMessage(msg string) *ErrorDTO {
	e.Error.Message = msg
	return e
}

func (e *ErrorDTO) SetCodeMsg(code int, msg string) {
	e.Error.Code = code
	e.Error.Message = msg
}
