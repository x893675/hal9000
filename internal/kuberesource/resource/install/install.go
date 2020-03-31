package install

import (
	"github.com/emicklei/go-restful"
	v1 "hal9000/internal/kuberesource/resource/v1"
	"hal9000/pkg/httpserver/runtime"
)

func init() {
	Install(runtime.Container)
}

func Install(c *restful.Container)  {
	runtime.Must(v1.AddToContainer(c))
}