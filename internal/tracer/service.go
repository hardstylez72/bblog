package tracer

import (
	"contrib.go.opencensus.io/exporter/jaeger"

	"go.opencensus.io/trace"
)

const (
	probabilitySampler = 1.0
)

type Config struct {
	Jaeger struct{
		CollectorEndpoint string
		ServiceName string
	}
}

func New(cfg Config) error {
	je, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: cfg.Jaeger.CollectorEndpoint,
		Process: jaeger.Process{
			ServiceName:cfg.Jaeger.ServiceName,
			Tags:        nil,
		},
	})

	if err != nil {
		return err
	}
	trace.RegisterExporter(je)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(probabilitySampler)})
	return nil
}
