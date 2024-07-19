// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package oc2otel // import "github.com/grafana/opentelemetry-go/bridge/opencensus/internal/oc2otel"

import (
	"fmt"

	octrace "go.opencensus.io/trace"

	"github.com/grafana/opentelemetry-go/trace"
)

func StartOptions(optFuncs []octrace.StartOption) ([]trace.SpanStartOption, error) {
	var ocOpts octrace.StartOptions
	for _, fn := range optFuncs {
		fn(&ocOpts)
	}

	var otelOpts []trace.SpanStartOption
	switch ocOpts.SpanKind {
	case octrace.SpanKindClient:
		otelOpts = append(otelOpts, trace.WithSpanKind(trace.SpanKindClient))
	case octrace.SpanKindServer:
		otelOpts = append(otelOpts, trace.WithSpanKind(trace.SpanKindServer))
	case octrace.SpanKindUnspecified:
		otelOpts = append(otelOpts, trace.WithSpanKind(trace.SpanKindUnspecified))
	}

	var err error
	if ocOpts.Sampler != nil {
		err = fmt.Errorf("unsupported sampler: %v", ocOpts.Sampler)
	}
	return otelOpts, err
}
