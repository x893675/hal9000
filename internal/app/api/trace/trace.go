package trace

import (
	"hal9000/internal/app/api/config"
	"hal9000/pkg/tracing"
	"log"
)

func New() error {
	if !config.EnableTrace {
		return nil
	}

	traceOpts := &tracing.Options{
		ZipkinURL:    config.TraceUrl,
		SamplingRate: 1.0,
	}
	if err := traceOpts.Validate(); err != nil {
		log.Println("Invalid options for tracing: ", err)
		return err
	}

	if traceOpts.TracingEnabled() {
		tracer, err := tracing.Configure(config.ServiceName, traceOpts)
		defer func() {
			if tracer != nil {
				_ = tracer.Close()
			}
		}()
		if err != nil {
			log.Println("Failed to configure tracing: ", err)
			return err
		}
	}
	return nil
}
