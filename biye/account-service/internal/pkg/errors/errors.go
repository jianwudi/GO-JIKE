package errors

import (
	"errors"
	"fmt"
)

func NewHTTPError(code int, field, detail string) *HTTPError {
	return &HTTPError{
		Errors: map[string][]string{
			field: {detail},
		},
		Code: code,
	}
}

type HTTPError struct {
	Errors map[string][]string `json:"errors"`
	Code   int
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError:%d", e.Code)
}

func FromError(err error) *HTTPError {
	if err == nil {
		return nil
	}
	if se := new(HTTPError); errors.As(err, &se) {
		return se
	}
	return &HTTPError{}
}
