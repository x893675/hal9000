package ecode

// All common ecode
var (
	OK = add(0) // 正确

	NoPermission         = add(-1) //无访问权限
	NoResourcePermission = add(-2) //无资源的访问权限
	MethodNotAllow       = add(-3) //方法不被允许

	BadRequest              = add(-400) //请求发生错误
	InvalidRequestParameter = add(-401) //无效的请求参数
	TooManyRequests         = add(-402) //请求过于频繁
	UnknownQuery            = add(-403) //未知的查询类型
	NotFound                = add(-404) //资源不存在

	ServerErr = add(-500) // 服务器错误
)
