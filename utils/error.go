package utils

import (
	"encoding/json"
	"fmt"
)

type HTTPError struct {
	Cause   error  `json:"-"`
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	if e.Cause == nil {
		return e.Message
	}
	return e.Message + " : " + e.Cause.Error()
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
		Cause:   err,
		Status:  status,
		Message: detail,
	}
}

type Error4XX interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}
