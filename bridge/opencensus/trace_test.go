// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opencensus // import "github.com/grafana/opentelemetry-go/bridge/opencensus"

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/grafana/opentelemetry-go/sdk/trace"
	"github.com/grafana/opentelemetry-go/sdk/trace/tracetest"
)

func TestNewTraceBridge(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	tp := trace.NewTracerProvider(trace.WithSyncer(exporter))
	bridge := newTraceBridge([]TraceOption{WithTracerProvider(tp)})
	_, span := bridge.StartSpan(context.Background(), "foo")
	span.End()
	gotSpans := exporter.GetSpans()
	require.Len(t, gotSpans, 1)
	gotSpan := gotSpans[0]
	assert.Equal(t, gotSpan.InstrumentationLibrary.Name, scopeName)
	assert.Equal(t, gotSpan.InstrumentationLibrary.Version, Version())
}
