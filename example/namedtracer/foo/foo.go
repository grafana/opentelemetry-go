// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package foo // import "github.com/grafana/opentelemetry-go/example/namedtracer/foo"

import (
	"context"

	otel "github.com/grafana/opentelemetry-go"
	"github.com/grafana/opentelemetry-go/attribute"
	"github.com/grafana/opentelemetry-go/trace"
)

var lemonsKey = attribute.Key("ex.com/lemons")

// SubOperation is an example to demonstrate the use of named tracer.
// It creates a named tracer with its package path.
func SubOperation(ctx context.Context) error {
	// Using global provider. Alternative is to have application provide a getter
	// for its component to get the instance of the provider.
	tr := otel.Tracer("example/namedtracer/foo")

	var span trace.Span
	_, span = tr.Start(ctx, "Sub operation...")
	defer span.End()
	span.SetAttributes(lemonsKey.String("five"))
	span.AddEvent("Sub span event")

	return nil
}
