package utils

import (
	"encoding/json"
	"fmt"
)

type Error4XX struct {
	Cause   error  `json:"-"`
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e *Error4XX) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Message + " : " + e.Cause.Error()
}

func (e *Error4XX) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func (e *Error4XX) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewError4XX(err error, status int, detail string) error {
	return &Error4XX{
		Cause:   err,
		Status:  status,
		Message: detail,
	}
}

type IError4XX interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}
