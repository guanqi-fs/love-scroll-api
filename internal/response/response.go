package response

import "love-scroll-api/internal/errorcode"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{
		Code:    errorcode.SuccessCode.Code,
		Message: errorcode.SuccessCode.Message,
		Data:    data,
	}
}

func NewResponse(errorCode *errorcode.ErrorCode, err error, data interface{}) *Response {
	return &Response{
		Code:    errorCode.Code,
		Message: errorCode.Message + ": " + err.Error(),
		Data:    data,
	}
}
