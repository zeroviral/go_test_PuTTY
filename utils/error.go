package utils

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	message string "message:"
	Data    interface{}
}

type HTTPError struct {
	Cause  error  `json:"-"`
	Status int    `json:"-"`
	Detail string `json:"detail"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing response body: %v", err)
	}
	return body, nil
}

func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func NewHTTPError(err error, status int, detail string) error {
	return &HTTPError{
		Cause:  err,
		Status: status,
		Detail: detail,
	}
}

type ClientError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}

func GenerateError(msg string, data interface{}) Error {
	return Error{
		message: msg,
		Data:    data,
	}
}
