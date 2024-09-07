package response

import (
	"personal-financial-tracker/consts"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message" default:""`
	Data    interface{} `json:"data"`
}

func SuccessResponseBody(message string) Response {
	return Response{
		Status:  consts.SUCCESS,
		Message: message,
	}
}

func ErrorResponseBody(errorMsg string, args ...interface{}) Response {
	return Response{
		Status:  consts.ERROR,
		Message: errorMsg,
	}
}

func DataResponseBody(data interface{}, message string) Response {
	return Response{
		Status:  consts.SUCCESS,
		Message: message,
		Data:    data,
	}
}
