package response

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

// Response структура ответа для статусного middleware
// @Description Добавляется в ответы сервера
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "ERROR"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

func ValidationErrors(errs validator.ValidationErrors) Response {
	var errMsgs []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMsgs = append(errMsgs, err.Field()+" is required")
		case "url":
			errMsgs = append(errMsgs, err.Field()+" must be a valid URL")
		default:
			errMsgs = append(errMsgs, err.Field()+" is invalid")

		}

	}

	return Response{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ", "),
	}

}
