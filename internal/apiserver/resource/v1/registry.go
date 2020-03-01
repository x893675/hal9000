package v1

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"hal9000/internal/apiserver/controller/resources"
	"hal9000/pkg/logger"
	"hal9000/pkg/schema"
	"hal9000/pkg/server/params"
	"hal9000/pkg/server/runtime"
	"net/http"
)

const GroupName = "resources.io"

var GroupVersion = runtime.GroupVersion{Group: GroupName, Version: "v1"}

var (
	WebServiceBuilder = runtime.NewContainerBuilder(addWebService)
	AddToContainer    = WebServiceBuilder.AddToContainer
)

func addWebService(c *restful.Container) error {
	webservice := runtime.NewWebService(GroupVersion)

	ok := "ok"

	webservice.Route(webservice.GET("/testrestful/{resources}").
		To(resources.TestRestful).
		Metadata(restfulspec.KeyOpenAPITags, []string{"Test Go-Restful"}).
		Doc("test restful query").
		Param(webservice.PathParameter("resources", "namespace level resource type, e.g. pods,jobs,configmaps,services.")).
		Param(webservice.QueryParameter(params.ConditionsParam, "query conditions,connect multiple conditions with commas, equal symbol for exact query, wave symbol for fuzzy query e.g. name~a").
			Required(false).
			DataFormat("key=%s,key~%s")).
		Param(webservice.QueryParameter(params.PagingParam, "paging query, e.g. limit=100,page=1").
			Required(false).
			DataFormat("limit=%d,page=%d").
			DefaultValue("limit=10,page=1")).
		Param(webservice.QueryParameter(params.ReverseParam, "sort parameters, e.g. reverse=true")).
		Param(webservice.QueryParameter(params.OrderByParam, "sort parameters, e.g. orderBy=createTime")).
		Returns(http.StatusOK, ok, schema.PageableResponse{}))

	for _, route := range webservice.Routes() {
		logger.Debug(nil, "%s         %s", route.Method, route.Path)
	}

	c.Add(webservice)
	return nil
}
