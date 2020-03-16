package resources

import (
	"github.com/emicklei/go-restful"
	"hal9000/pkg/httpserver/errors"
	"hal9000/pkg/httpserver/params"
	"hal9000/pkg/logger"
	"hal9000/pkg/schema"
	"net/http"
)

func TestRestful(req *restful.Request, resp *restful.Response) {
	resourceName := req.PathParameter("resources")
	conditions, err := params.ParseConditions(req.QueryParameter(params.ConditionsParam))
	orderBy := params.GetStringValueWithDefault(req, params.OrderByParam, "createTime")
	limit, offset := params.ParsePaging(req.QueryParameter(params.PagingParam))
	reverse := params.ParseReverse(req)

	if err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, errors.Wrap(err))
		return
	}

	logger.Info(nil, "%v", conditions.Match)
	logger.Info(nil, "%s-%s-%d-%d-%t", resourceName, orderBy, limit, offset, reverse)

	result := &schema.PageableResponse{
		Items:      []string{"hello", "world"},
		TotalCount: 2,
	}

	resp.WriteAsJson(result)
}
