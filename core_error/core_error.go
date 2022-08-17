package core_error

import "net/http"

type CoreError interface {
	Error() string
}

type coreError struct {
	errMsg string
	errStatus int
}

func(coreErr *coreError) Error() string {
	return coreErr.errMsg
}

func NewCoreError() *coreError {
	return &coreError{}
}
func(coreErr *coreError) InternalError(msg string) *coreError {
	coreErr.errMsg = msg
	coreErr.errStatus = http.StatusInternalServerError
	return coreErr
}

func(coreErr *coreError) BadRequestError(msg string) *coreError {
	coreErr.errMsg = msg
	coreErr.errStatus = http.StatusBadRequest
	return coreErr
}

func(coreErr *coreError) CustomError(msg string,statusCode int) *coreError {
	coreErr.errMsg = msg
	coreErr.errStatus = statusCode
	return coreErr
}