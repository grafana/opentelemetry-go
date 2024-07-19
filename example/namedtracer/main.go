// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-logr/stdr"

	otel "github.com/grafana/opentelemetry-go"
	"github.com/grafana/opentelemetry-go/attribute"
	"github.com/grafana/opentelemetry-go/baggage"
	"github.com/grafana/opentelemetry-go/example/namedtracer/foo"
	"github.com/grafana/opentelemetry-go/exporters/stdout/stdouttrace"
	sdktrace "github.com/grafana/opentelemetry-go/sdk/trace"
	"github.com/grafana/opentelemetry-go/trace"
)

var (
	fooKey     = attribute.Key("ex.com/foo")
	barKey     = attribute.Key("ex.com/bar")
	anotherKey = attribute.Key("ex.com/another")
)

var tp *sdktrace.TracerProvider

// initTracer creates and registers trace provider instance.
func initTracer() error {
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("failed to initialize stdouttrace exporter: %w", err)
	}
	bsp := sdktrace.NewBatchSpanProcessor(exp)
	tp = sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tp)
	return nil
}

func main() {
	// Set logging level to info to see SDK status messages
	stdr.SetVerbosity(5)

	// initialize trace provider.
	if err := initTracer(); err != nil {
		log.Panic(err)
	}

	// Create a named tracer with package path as its name.
	tracer := tp.Tracer("example/namedtracer/main")
	ctx := context.Background()
	defer func() { _ = tp.Shutdown(ctx) }()

	m0, _ := baggage.NewMemberRaw(string(fooKey), "foo1")
	m1, _ := baggage.NewMemberRaw(string(barKey), "bar1")
	b, _ := baggage.New(m0, m1)
	ctx = baggage.ContextWithBaggage(ctx, b)

	var span trace.Span
	ctx, span = tracer.Start(ctx, "operation")
	defer span.End()
	span.AddEvent("Nice operation!", trace.WithAttributes(attribute.Int("bogons", 100)))
	span.SetAttributes(anotherKey.String("yes"))
	if err := foo.SubOperation(ctx); err != nil {
		panic(err)
	}
}
