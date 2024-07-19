// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdoutlog_test

import (
	"context"

	"github.com/grafana/opentelemetry-go/exporters/stdout/stdoutlog"
	"github.com/grafana/opentelemetry-go/log/global"
	"github.com/grafana/opentelemetry-go/sdk/log"
)

func Example() {
	exp, err := stdoutlog.New()
	if err != nil {
		panic(err)
	}

	processor := log.NewSimpleProcessor(exp)
	provider := log.NewLoggerProvider(log.WithProcessor(processor))
	defer func() {
		if err := provider.Shutdown(context.Background()); err != nil {
			panic(err)
		}
	}()

	global.SetLoggerProvider(provider)

	// From here, the provider can be used by instrumentation to collect
	// telemetry.
}
