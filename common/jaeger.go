package common

import (
	"io"
	"time"

	"github.com/uber/jaeger-client-go"

	"github.com/uber/jaeger-client-go/config"

	"github.com/opentracing/opentracing-go"
)

func NewTracer(serviceName string, addr string) (opentracing.Tracer, io.Closer, error) {
	cfg := config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			BufferFlushInterval: 1 * time.Second,
			LogSpans:            true,
			LocalAgentHostPort:  addr,
		},
	}
	return cfg.NewTracer()
}
