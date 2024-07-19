// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otlploghttp_test

import (
	"context"

	"github.com/grafana/opentelemetry-go/exporters/otlp/otlplog/otlploghttp"
	"github.com/grafana/opentelemetry-go/log/global"
	"github.com/grafana/opentelemetry-go/sdk/log"
)

func Example() {
	ctx := context.Background()
	exp, err := otlploghttp.New(ctx)
	if err != nil {
		panic(err)
	}

	processor := log.NewBatchProcessor(exp)
	provider := log.NewLoggerProvider(log.WithProcessor(processor))
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()

	global.SetLoggerProvider(provider)

	// From here, the provider can be used by instrumentation to collect
	// telemetry.
}
