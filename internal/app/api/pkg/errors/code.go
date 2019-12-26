package errors

import "hal9000/pkg/ecode"

var (
	codes = make(map[string]ErrorCode)
)

// ErrorCode 错误码
type ErrorCode struct {
	Code           int
	Message        string
	HTTPStatusCode int
}

// newErrorCode 设定错误码
func newErrorCode(err string, code int, message string, status ...int) {
	errCode := ErrorCode{
		Code:    code,
		Message: message,
	}
	if len(status) > 0 {
		errCode.HTTPStatusCode = status[0]
	}
	codes[err] = errCode
}

// FromErrorCode 获取错误码
func FromErrorCode(err string) (ErrorCode, bool) {
	v, ok := codes[err]
	return v, ok
}

// newBadRequestError 创建请求错误
func newBadRequestError(err ecode.Codes) {
	newErrorCode(err.Error(), 400, err.Message(), 400)
}

// newUnauthorizedError 创建未授权错误
func newUnauthorizedError(err ecode.Codes) {
	newErrorCode(err.Error(), 401, err.Message(), 401)
}

// newInternalServerError 创建服务器错误
func newInternalServerError(err ecode.Codes) {
	newErrorCode(err.Error(), 500, err.Message(), 500)
}