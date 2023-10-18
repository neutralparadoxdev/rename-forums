package webapi

import "fmt"

type WebApiError struct {
	Code int64
	Msg  string
}

func NewWebApiError(code int64, msg string) WebApiError {
	return WebApiError{
		Code: code,
		Msg:  msg,
	}
}

func NewWebApiErrorFrom(from WebApiError, msg string) WebApiError {
	return NewWebApiError(from.Code, fmt.Sprintf("%s: %s", from.Msg, msg))
}

func (err *WebApiError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Msg)
}
