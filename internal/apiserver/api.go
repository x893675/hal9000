package apiserver

import (
	"github.com/emicklei/go-restful"
	resourcev1 "hal9000/internal/apiserver/resource/v1"
)

// Must panics on non-nil errors.  Useful to handling programmer level errors.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func InstallAPIs(container *restful.Container) {
	Must(resourcev1.AddToContainer(container))
}
