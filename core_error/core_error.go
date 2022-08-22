package core_error

import "net/http"

type ICoreError interface {
	Error() string
}

type CoreError struct {
	errMsg string
	errStatus int
}

func(coreErr *CoreError) Error() string {
	return coreErr.errMsg
}

func NewCoreError() *CoreError {
	return &CoreError{}
}
func(coreErr *CoreError) InternalError(msg string) *CoreError {
	coreErr.errMsg = msg
	coreErr.errStatus = http.StatusInternalServerError
	return coreErr
}

func(coreErr *CoreError) BadRequestError(msg string) *CoreError {
	coreErr.errMsg = msg
	coreErr.errStatus = http.StatusBadRequest
	return coreErr
}

func(coreErr *CoreError) CustomError(msg string,statusCode int) *CoreError {
	coreErr.errMsg = msg
	coreErr.errStatus = statusCode
	return coreErr
}

func(coreErr *CoreError) GetMsg() string {
	if coreErr != nil {
		return "CustomErr:" + coreErr.errMsg
	}
	return ""
}

func(coreErr *CoreError) GetStatus() int {
	if coreErr != nil {
		return coreErr.errStatus
	}
	return 0
}


func GetCoreError(err error) (cer *CoreError) {
	cerB ,oke  := err.(*CoreError)
	if cerB == nil || oke == false{
		return nil
	}
	return cerB
}