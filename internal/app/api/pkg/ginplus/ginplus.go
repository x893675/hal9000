package ginplus

import (
	"github.com/gin-gonic/gin"
	"hal9000/pkg/util"
	"net/http"
)

// 定义上下文中的键
const (
	prefix = "hal9000"
	// ResBodyKey 存储上下文中的键(响应Body数据)
	ResBodyKey = prefix + "/res_body"
)

// HTTPError HTTP响应错误
type HTTPError struct {
	Error HTTPErrorItem `json:"error" swaggo:"true,错误项"`
}

// HTTPErrorItem HTTP响应错误项
type HTTPErrorItem struct {
	Code    int    `json:"code" swaggo:"true,错误码"`
	Message string `json:"message" swaggo:"true,错误信息"`
}

// HTTPStatus HTTP响应状态
type HTTPStatus struct {
	Status string `json:"status" swaggo:"true,状态(OK)"`
}

// HTTPList HTTP响应列表数据
type HTTPList struct {
	List       interface{}     `json:"list"`
	Pagination *HTTPPagination `json:"pagination,omitempty"`
}

// HTTPPagination HTTP分页数据
type HTTPPagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}

func ResOK(c *gin.Context) {
	ResSuccess(c, HTTPStatus{Status: "OK"})
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}