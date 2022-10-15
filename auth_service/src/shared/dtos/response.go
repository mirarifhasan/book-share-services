package interceptors

import "strings"

type Response struct {
	Message string      `json:"message"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct {
}

func BuildResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	errors := strings.Split(err, "\n")
	return Response{
		Message: message,
		Errors:  errors,
		Data:    data,
	}
}
