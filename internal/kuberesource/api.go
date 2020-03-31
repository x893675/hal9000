package kuberesource

import (
	"github.com/emicklei/go-restful"
	resourcev1 "hal9000/internal/kuberesource/resource/v1"
	"hal9000/pkg/httpserver/runtime"
)


func InstallAPIs(container *restful.Container) {
	runtime.Must(resourcev1.AddToContainer(container))
}
