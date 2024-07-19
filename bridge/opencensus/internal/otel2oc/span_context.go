// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otel2oc // import "github.com/grafana/opentelemetry-go/bridge/opencensus/internal/otel2oc"

import (
	octrace "go.opencensus.io/trace"

	"github.com/grafana/opentelemetry-go/trace"
)

func SpanContext(sc trace.SpanContext) octrace.SpanContext {
	var to octrace.TraceOptions
	if sc.IsSampled() {
		// OpenCensus doesn't expose functions to directly set sampled
		to = 0x1
	}
	return octrace.SpanContext{
		TraceID:      octrace.TraceID(sc.TraceID()),
		SpanID:       octrace.SpanID(sc.SpanID()),
		TraceOptions: to,
	}
}
