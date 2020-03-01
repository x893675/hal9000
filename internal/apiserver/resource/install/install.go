package install

import (
	"github.com/emicklei/go-restful"
	v1 "hal9000/internal/apiserver/resource/v1"
	"hal9000/pkg/server/runtime"
)

func init() {
	Install(runtime.Container)
}

func Install(c *restful.Container)  {
	runtime.Must(v1.AddToContainer(c))
}