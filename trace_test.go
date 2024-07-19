// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otel

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grafana/opentelemetry-go/trace"
	"github.com/grafana/opentelemetry-go/trace/embedded"
	"github.com/grafana/opentelemetry-go/trace/noop"
)

type testTracerProvider struct{ embedded.TracerProvider }

var _ trace.TracerProvider = &testTracerProvider{}

func (*testTracerProvider) Tracer(_ string, _ ...trace.TracerOption) trace.Tracer {
	return noop.NewTracerProvider().Tracer("")
}

func TestMultipleGlobalTracerProvider(t *testing.T) {
	p1 := testTracerProvider{}
	p2 := noop.NewTracerProvider()
	SetTracerProvider(&p1)
	SetTracerProvider(p2)

	got := GetTracerProvider()
	assert.Equal(t, p2, got)
}
