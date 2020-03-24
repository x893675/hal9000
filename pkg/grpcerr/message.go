package grpcerr

import "fmt"

type ErrorMessage struct {
	Name string
	en   string
}

func (em ErrorMessage) Message(err error, a ...interface{}) string {
	format := em.en
	if err != nil {
		return fmt.Sprintf("%s: %s", fmt.Sprintf(format, a...), err.Error())
	} else {
		return fmt.Sprintf(format, a...)
	}
}

var (
	ErrorPermissionDenied = ErrorMessage{
		Name: "permission_denied",
		en:   "permission denied",
	}
	ErrorInternalError = ErrorMessage{
		Name: "internal_error",
		en:   "internal error",
	}
	ErrorMissingParameter = ErrorMessage{
		Name: "missing_parameter",
		en:   "missing parameter [%s]",
	}
	ErrorUnsupportedParameterValue = ErrorMessage{
		Name: "unsupported_parameter_value",
		en:   "unsupported parameter [%s] value [%s]",
	}
)