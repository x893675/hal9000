package config

import "hal9000/pkg/util"

var (
	ServiceName = util.GetEnvironment("SERVICE_NAME", "hal9000-api")
	ServicePort = util.GetEnvironment("SERVICE_PORT","8080")
	EnableTrace = util.GetEnvironmentToBool("ENABLE_TRACE", "true")
	TraceUrl = util.GetEnvironment("TRACE_URL", "http://zipkin.istio-system:9411/api/v1/spans")
	GreeterServiceName = util.GetEnvironment("GREETER_SERVICE_NAME", "hal9000")
)