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
	ErrorAuthRateLimitExceed = ErrorMessage{
		Name: "auth_rate_limit_exceed",
		en:   "auth rate limit exceeded, max [%s], current [%s]",
	}
	ErrorAuthFailure = ErrorMessage{
		Name: "auth_failure",
		en:   "auth failure",
	}
	ErrorParameterShouldNotBeEmpty = ErrorMessage{
		Name: "parameter_should_not_be_empty",
		en:   "parameter [%s] should not be empty",
	}
	ErrorRefreshTokenExpired = ErrorMessage{
		Name: "refresh_token_expired",
		en:   "refresh token expired",
	}
)