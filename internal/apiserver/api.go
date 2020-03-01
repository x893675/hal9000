package apiserver

import (
	"github.com/emicklei/go-restful"
	resourcev1 "hal9000/internal/apiserver/resource/v1"
	"hal9000/pkg/server/runtime"
)



func InstallAPIs(container *restful.Container) {
	runtime.Must(resourcev1.AddToContainer(container))
}
