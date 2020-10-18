package tracer

import (
	"contrib.go.opencensus.io/exporter/jaeger"

	"go.opencensus.io/trace"
)

const (
	probabilitySampler = 1.0
)

func New(collectorEndpoint, serviceName string) error {
	je, err := jaeger.NewExporter(jaeger.Options{
		CollectorEndpoint: collectorEndpoint,
		Process: jaeger.Process{
			ServiceName: serviceName,
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
