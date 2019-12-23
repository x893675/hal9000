package config

import "hal9000/pkg/util"

var (
	ServiceName = util.GetEnvironment("SERVICE_NAME", "hal9000-api")
	ServicePort = util.GetEnvironment("SERVICE_PORT","8080")
)