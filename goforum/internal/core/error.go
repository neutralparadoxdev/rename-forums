package core

import "fmt"

type CoreError struct {
	Code int64
	Msg  string
}

func NewCoreError(code int64, msg string) CoreError {
	return CoreError{
		Code: code,
		Msg:  msg,
	}
}

func NewCoreErrorFrom(from CoreError, msg string) CoreError {
	return NewCoreError(from.Code, fmt.Sprintf("%s: %s", from.Msg, msg))
}

func (err *CoreError) Error() string {
	return fmt.Sprintf("%d: %s", err.Code, err.Msg)
}
