package utils

import (
	"net/http"
	"time"
)

type ResponseBase struct {
	Code int `json:"status"`
	Timestamp int `json:"timestamp"`
	Data interface{} `json:"data"`
}

func GenerateTimestampedResponse(code int, data interface{}) ResponseBase {
	now := time.Now()
	return ResponseBase{
		Code: code,
		Timestamp: int(now.Unix()),
		Data: data,
	}
}

func GenerateOkResponse(data interface{}) ResponseBase {
	return GenerateTimestampedResponse(http.StatusOK, data)
}

type ErroResponseWithMessage struct {
	Message string `json:"message"`
}

func GenerateErrorWithMessage(message string) ResponseBase {
	return GenerateTimestampedResponse(http.StatusBadRequest, ErroResponseWithMessage{Message: message})
}