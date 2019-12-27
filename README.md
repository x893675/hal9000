## HAL 9000

hal9000是golang微服务的一个简单例子，涉及内容:

* 使用gin, go-micro框架
* 使用k8s作为注册中心
* 对接istio的分布式链路追踪

项目名取自于电影**2001太空漫游**中的人工智能机器人hal9000

## 快速开始

1. 搭建k8s+istio的测试环境, 创建namespace`devel`，并设置sidecar自动注入
2. `git clone https://github.com/x893675/hal9000.git`
3. `kubectl apply -f deployment/rbac.yaml`
4. `kubectl apply -f deployment/rolebinding.yaml -n devel`
5. `kubectl apply -f deployment/hal9000-*.yaml`
6. 更改service类型为nodeport，浏览器访问`yourip:port/api/v1/hello/ww`,在istio的kiali和jaeger中查看服务流量

## 源码分析

关于go-micro的使用在这里不做另外说明，下面着重讲一下isito的链路追踪对接

### 分布式追踪

分布式追踪中的主要概念:

* Trace: 一次完整的分布式调用跟踪链路
* Span: 跨服务的一次调用;多个Span组合成一次Trace追踪记录

一个完整的调用链跟踪系统，包括调用链埋点，调用链数据收集，调用链数据存储和处理，调用链数据检索（除了提供检索的 APIServer，一般还要包含一个非常酷炫的调用链前端）等若干重要组件。**istio现在默认使用的是jaeger作为trace系统，可以选择使用jaeger和zipkin的trace格式。**

### istio-trace

istio官方的介绍为:

> Istio makes it easy to create a network of deployed services with load balancing, service-to-service authentication, monitoring, and more, *without any changes* in service code.

istio在使用时，不对代码做任何处理即可进行服务治理，但是实际使用过程中，不修改服务代码，istio的调用链总是断开的。

在 Istio 中，所有的治理逻辑的执行体都是和业务容器一起部署的 Envoy 这个 Sidecar，不管是负载均衡、熔断、流量路由还是安全、可观察性的数据生成都是在 Envoy 上。Sidecar 拦截了所有的流入和流出业务程序的流量，根据收到的规则执行执行各种动作。实际使用中一般是基于 K8S 提供的 InitContainer 机制，用于在 Pod 中执行一些初始化任务. InitContainer 中执行了一段 Iptables 的脚本。正是通过这些 Iptables 规则拦截 pod 中流量，并发送到 Envoy 上。Envoy 拦截到 Inbound 和 Outbound 的流量会分别作不同操作，执行上面配置的操作，另外再把请求往下发，对于 Outbound 就是根据服务发现找到对应的目标服务后端上；对于 Inbound 流量则直接发到本地的服务实例上。

Envoy的埋点规则为:

- Inbound 流量：对于经过 Sidecar 流入应用程序的流量，如果经过 Sidecar 时 Header 中没有任何跟踪相关的信息，则会在创建一个根 Span，TraceId 就是这个 SpanId，然后再将请求传递给业务容器的服务；如果请求中包含 Trace 相关的信息，则 Sidecar 从中提取 Trace 的上下文信息并发给应用程序。
- Outbound 流量：对于经过 Sidecar 流出的流量，如果经过 Sidecar 时 Header 中没有任何跟踪相关的信息，则会创建根 Span，并将该跟 Span 相关上下文信息放在请求头中传递给下一个调用的服务；当存在 Trace 信息时，Sidecar 从 Header 中提取 Span 相关信息，并基于这个 Span 创建子 Span，并将新的 Span 信息加在请求头中传递。

根据这个规则，对于一个api->A-这个简单调用，我们有如下分析:

* 当一个请求进入api时，该请求头中没有任何trace相关的信息,对于这个inbound流量，istio会创建一个根span，并向请求头注入span信息。
* 当api向A创建并发送rpc或http请求时，这个请求对于api的envoy来说时outbound流量，如果请求头中没有trace信息，会创建根span信息填入请求头
* 这种情况下，在istio的jaeger页面上我们可以看到两段断裂的trace记录

**结论**：**埋点逻辑是在 Sidecar 代理中完成，应用程序不用处理复杂的埋点逻辑，但应用程序需要配合在请求头上传递生成的 Trace 相关信息**。

istio使用jaeger作为trace系统，格式为zipkin format。在请求头中有如下headers:

- `x-request-id`
- `x-b3-traceid`
- `x-b3-spanid`
- `x-b3-parentspanid`
- `x-b3-sampled`
- `x-b3-flags`
- `x-ot-span-context`

注意: 在http请求中，比如使用gin框架时，这些header中的key应是首字母大写的，例如:`X-Request-Id`

### 代码示例分析

以下代码段使用的是gin作为http框架，AuthSvc是rpc客户端，使用go-micro框架

```go
import(
	"github.com/uber/jaeger-client-go"
	ot "github.com/opentracing/opentracing-go"
	"github.com/micro/go-micro/metadata"
	"github.com/gin-gonic/gin"
)


func (a *LoginController) Login(c *gin.Context) {
	//从http头中获得根span，使用istio时，该根span由envoy注入,记为root span
	inBoundSpanCtx, err := ot.GlobalTracer().Extract(ot.HTTPHeaders, ot.HTTPHeadersCarrier(c.Request.Header))
	//由根span创建一个子span,改span为span2
	span := ot.StartSpan("controller.(*LoginController).Login", 
		ot.ChildOf(inBoundSpanCtx),
		ot.Tags{
			"kind": "function",
		})
	//将span2与当前context绑定
	ctx := ot.ContextWithSpan(context.Background(), span)
	//在testtrace中再创建一个子span3
	testTrace(ctx)

    //从当前context中得到rpc调用的metadata,因为当前调用入口为http调用，所以ok永远为false
	md, ok := metadata.FromContext(ctx)
	if ok{
		fmt.Println("metadata from context is ok")
		for k, v := range md{
			fmt.Println(k,v)
		}
	}else{
		fmt.Println("metadata from context is not ok")
		md = make(map[string]string)
        //从span2的spancontext中获取trace信息，因为istio使用的是jaeger，所以将opentracing的接口进行类型断言转换为jaeger的spancontext,将span2的trance信息填入metadata
		if sc, ok := span.Context().(jaeger.SpanContext); ok {
			md["x-request-id"] = c.GetHeader("X-Request-Id")
			md["x-b3-traceid"] = sc.TraceID().String()
			md["x-b3-spanid"] = sc.SpanID().String()
			md["x-b3-sampled"] = c.GetHeader("X-B3-Sampled")
		}else{
			md["x-request-id"] = c.GetHeader("X-Request-Id")
			md["x-b3-traceid"] = c.GetHeader("X-B3-Traceid")
			md["x-b3-spanid"] = c.GetHeader("X-B3-Spanid")
			md["x-b3-sampled"] = c.GetHeader("X-B3-Sampled")
		}
        //从创建好的metadata中创建一个新的context
		ctx = metadata.NewContext(ctx, md)
	}

	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
    //发送rpc请求时，使用新创建的携带了rpc metadata的context,该请求经过envoy时，envoy看到该outbound流量中的trace信息，会创建一个子span,传递给下一个服务，标记该span为span4
	response, err := a.AuthSvc.Verify(ctx, &auth.LoginRequest{
		Username: item.UserName,
		Password: item.Password,
	})

	if err != nil {
		ginplus.ResError(c, err)
		return
	}
    //结束span2
	span.Finish()
	ginplus.ResSuccess(c, response)
}

func testTrace(ctx context.Context){
    //传入的ctx已经与span2绑定，再创建一个子span，标记为span3
	span, _ := ot.StartSpanFromContext(ctx,
		"testTrace",
		ot.Tags{
		string(ext.SpanKind): "function",
	})
	fmt.Println("in test Trace function...")
	//span结束上报jaeger
    span.Finish()
}
```

由上图的注释分析得到下列span关系:

```
root span --> span2 -- span3
                    -- span4
```

在istio中就把之前分裂的两个trace记录合并为一个了。

结论:

* 使用istio时，我们只需要对服务间调用的header信息进行透传
* 如果想把服务内的调用关系与istio生成的trace合并，只需以istio生成的span作为父span，生成子span即可
* 透传header的代码大都一致，可以做成一个通用的函数调用，减少服务代码的修改

### 本项目代码分析

在hal9000-api的http处理函数链中，加入了一个trace的中间件如下:

```go
package middleware

import (
	"github.com/gin-gonic/gin"
	ot "github.com/opentracing/opentracing-go"
	"hal9000/internal/app/api/pkg/ginplus"
	"hal9000/pkg/util"
	"log"
)

func TraceMiddleware(skipper ...SkipperFunc) gin.HandlerFunc{
	return func(c *gin.Context) {
		if len(skipper) > 0 && skipper[0](c) {
			c.Next()
			return
		}
        //判断opentracing是否初始化，没有初始化则表示服务不对接istio的分布式追踪
		if ot.GlobalTracer() == nil {
            //不对接istio分布式追踪的情况下，生成一个request id，便于服务自身的调用分析
			requestID := c.GetHeader("X-Request-Id")
			if requestID == "" {
				requestID = util.MustUUID()
			}
			c.Set(ginplus.RequestIDKey, requestID)
			c.Next()
			return
		}
        //使用istio分布式追踪的情况下，对inbound请求的header进行解包，获取spancontext,该spancontext作为服务中的根span
		inBoundSpanCtx, err := ot.GlobalTracer().Extract(ot.HTTPHeaders, ot.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			log.Println(err.Error())
			c.Next()
			return
		}
        //将request id和根spancontext放入gin的上下文中
		c.Set(ginplus.RequestIDKey, c.GetHeader("X-Request-Id"))
		c.Set(ginplus.RootTraceCtx, inBoundSpanCtx)
	}
}

func StartChildSpan(c *gin.Context, operationName string, tags tracing.Tags)  context.Context{
	if rootTranceCtx, ok := c.Get(RootTraceCtx); ok{
		if ctx, ok := rootTranceCtx.(ot.SpanContext); ok{
			span := ot.StartSpan(operationName, ot.ChildOf(ctx), ot.Tags(tags))
			ctx := ot.ContextWithSpan(context.Background(), span)
			return ctx
		}
	}
	return context.Background()
}

func FinishSpan(ctx context.Context) {
	span := ot.SpanFromContext(ctx)
	if span != nil {
		span.Finish()
	}
}
//发起rpc调用时，对rpc的metadata传入相关的span值
func InjectJaegerTraceToRpcMetaData(c *gin.Context) context.Context{
	md := make(map[string]string)
	for k, v := range jaegerTraceHeaderDict{
		if temp := c.GetHeader(v); temp != "" {
			md[k] = temp
		}
	}
	ctx := metadata.NewContext(context.Background(), md)
	return ctx
}
```

