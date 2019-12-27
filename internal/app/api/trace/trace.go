package trace

import (
	"hal9000/internal/app/api/config"
	"hal9000/pkg/tracing"
	"io"
	"log"
)

func New() (io.Closer) {
	if !config.EnableTrace {
		return nil
	}

	traceOpts := &tracing.Options{
		ZipkinURL: config.TraceUrl,
		SamplingRate: 1.0,
	}
	if err := traceOpts.Validate(); err != nil {
		log.Println("Invalid options for tracing: ", err)
		return nil
	}

	if traceOpts.TracingEnabled() {
		tracer, err := tracing.Configure(config.ServiceName, traceOpts)
		if err != nil {
			log.Println("Failed to configure tracing: ", err)
			return nil
		}
		return tracer
	}
	return nil
}