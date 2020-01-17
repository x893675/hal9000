package tracing

import (
	"context"
	ot "github.com/opentracing/opentracing-go"
)

type Tag struct {
	Key   string
	Value interface{}
}

type Tags map[string]interface{}

func StartSpanWithCtx(ctx context.Context, operationName string, tags map[string]interface{}) context.Context {
	if ot.GlobalTracer() == nil {
		return nil
	}
	_, childCtx := ot.StartSpanFromContext(ctx, operationName, ot.Tags(tags))
	return childCtx
}

func InjectTraceToRpcMetaData() context.Context {
	return nil
}
