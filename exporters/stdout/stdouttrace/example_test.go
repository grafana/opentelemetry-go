// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdouttrace_test

import (
	"context"
	"fmt"
	"log"

	otel "github.com/grafana/opentelemetry-go"
	"github.com/grafana/opentelemetry-go/exporters/stdout/stdouttrace"
	"github.com/grafana/opentelemetry-go/sdk/resource"
	sdktrace "github.com/grafana/opentelemetry-go/sdk/trace"
	semconv "github.com/grafana/opentelemetry-go/semconv/v1.26.0"
	"github.com/grafana/opentelemetry-go/trace"
)

const (
	instrumentationName    = "github.com/instrumentron"
	instrumentationVersion = "0.1.0"
)

var tracer = otel.GetTracerProvider().Tracer(
	instrumentationName,
	trace.WithInstrumentationVersion(instrumentationVersion),
	trace.WithSchemaURL(semconv.SchemaURL),
)

func add(ctx context.Context, x, y int64) int64 {
	_, span := tracer.Start(ctx, "Addition")
	defer span.End()

	return x + y
}

func multiply(ctx context.Context, x, y int64) int64 {
	_, span := tracer.Start(ctx, "Multiplication")
	defer span.End()

	return x * y
}

func Resource() *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName("stdout-example"),
		semconv.ServiceVersion("0.0.1"),
	)
}

func InstallExportPipeline() (func(context.Context) error, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, fmt.Errorf("creating stdout exporter: %w", err)
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(Resource()),
	)
	otel.SetTracerProvider(tracerProvider)

	return tracerProvider.Shutdown, nil
}

func Example() {
	ctx := context.Background()

	// Registers a tracer Provider globally.
	shutdown, err := InstallExportPipeline()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, span := tracer.Start(ctx, "Calculation")
	defer span.End()
	ans := multiply(ctx, 2, 2)
	ans = multiply(ctx, ans, 10)
	ans = add(ctx, ans, 2)
	log.Println("the answer is", ans)
}
