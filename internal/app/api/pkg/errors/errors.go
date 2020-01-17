package errors

import "hal9000/pkg/ecode"

func init() {
	errCode := map[int]string{
		int(ecode.NoPermission):            "无访问权限",
		int(ecode.NoResourcePermission):    "无资源的访问权限",
		int(ecode.MethodNotAllow):          "方法不被允许",
		int(ecode.BadRequest):              "请求发生错误",
		int(ecode.InvalidRequestParameter): "无效的请求参数",
		int(ecode.TooManyRequests):         "请求过于频繁",
		int(ecode.UnknownQuery):            "未知的查询类型",
		int(ecode.NotFound):                "资源不存在",
		int(ecode.ServerErr):               "服务器错误",
	}
	ecode.Register(errCode)

	// 公共错误
	newBadRequestError(ecode.BadRequest)
	newBadRequestError(ecode.InvalidRequestParameter)
	newErrorCode(ecode.NotFound.Error(), 404, ecode.NotFound.Message(), 404)
	newErrorCode(ecode.MethodNotAllow.Error(), 405, ecode.MethodNotAllow.Message(), 405)
	newErrorCode(ecode.TooManyRequests.Error(), 429, ecode.TooManyRequests.Message(), 429)
	newBadRequestError(ecode.UnknownQuery)

	// 权限错误
	newErrorCode(ecode.NoPermission.Error(), 9999, ecode.NoPermission.Message(), 401)
	newErrorCode(ecode.NoResourcePermission.Error(), 401, ecode.NoResourcePermission.Message(), 401)
}
