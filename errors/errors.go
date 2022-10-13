package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	APIError  string = "API_ERROR"
	GoErrCode string = "GO_ERROR"
)

type Error struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	IsError    bool   `json:"error"`
	Err        string `json:"err"`
	StatusCode int    `json:"status_code"`
}

// FromGoErr generates sendbird.Error from generic go errors
func FromGoErr(err error) *Error {
	return &Error{
		StatusCode: http.StatusTeapot,
		Err:        GoErrCode,
		Message:    err.Error(),
	}
}

// FromHTTPErr generates sendbird.Error from http errors with non 2xx status
func FromHTTPErr(status int, respBody []byte) *Error {
	var httpError *Error
	if err := json.Unmarshal(respBody, &httpError); err != nil {
		return FromGoErr(err)
	}
	httpError.StatusCode = status
	httpError.Err = APIError

	return httpError
}

func (e *Error) ToError() error {
	return fmt.Errorf("[%d] %s", e.Code, e.Message)
}
