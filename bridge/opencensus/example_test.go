// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opencensus_test

import (
	"github.com/grafana/opentelemetry-go/bridge/opencensus"
	"github.com/grafana/opentelemetry-go/sdk/metric"
)

func ExampleNewMetricProducer() {
	// Create the OpenCensus Metric bridge.
	bridge := opencensus.NewMetricProducer()
	// Add the bridge as a producer to your reader.
	// If using a push exporter, such as OTLP exporter,
	// use metric.NewPeriodicReader with metric.WithProducer option.
	// If using a pull exporter which acts as a reader, such as prometheus exporter,
	// use a dedicated option like prometheus.WithProducer.
	reader := metric.NewManualReader(metric.WithProducer(bridge))
	// Add the reader to your MeterProvider.
	_ = metric.NewMeterProvider(metric.WithReader(reader))
}
