package main

import (
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

func main() {

	jcfg := config.Configuration{
		Disabled: false, // Nop tracer if True
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			// Адрес рядом стоящего jaeger-agent, который будет репортить спаны
			LocalAgentHostPort: "localhost",
		},
	}

	tracer, closer, err := jcfg.New(
		"test_service",
		config.Logger(jaeger.StdLogger),
	)

	endpoint := CreateEndpoint(svc)

	endpoint = opentracing.TraceServer
}
