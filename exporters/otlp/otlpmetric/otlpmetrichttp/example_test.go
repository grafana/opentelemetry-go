// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlpmetrichttp_test

import (
	"context"

	otel "github.com/grafana/opentelemetry-go"
	"github.com/grafana/opentelemetry-go/exporters/otlp/otlpmetric/otlpmetrichttp"
	"github.com/grafana/opentelemetry-go/sdk/metric"
)

func Example() {
	ctx := context.Background()
	exp, err := otlpmetrichttp.New(ctx)
	if err != nil {
		panic(err)
	}

	meterProvider := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(exp)))
	defer func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
	otel.SetMeterProvider(meterProvider)

	// From here, the meterProvider can be used by instrumentation to collect
	// telemetry.
}
