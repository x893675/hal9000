package errors

import (
	"github.com/emicklei/go-restful"
	"net/http"
)

type Error struct {
	Message string `json:"message" description:"error message"`
}

var None = Error{Message: "success"}

func (e *Error) Error() string {
	return e.Message
}

func Wrap(err error) Error {
	return Error{Message: err.Error()}
}

func New(message string) Error {
	return Error{Message: message}
}

func ParseSvcErr(err error, resp *restful.Response) {
	if svcErr, ok := err.(restful.ServiceError); ok {
		resp.WriteServiceError(svcErr.Code, svcErr)
	} else {
		resp.WriteHeaderAndEntity(http.StatusInternalServerError, Wrap(err))
	}
}