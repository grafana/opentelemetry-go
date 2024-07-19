// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Command zipkin is an example program that creates spans
// and uploads to openzipkin collector.
package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	otel "github.com/grafana/opentelemetry-go"
	"github.com/grafana/opentelemetry-go/exporters/zipkin"
	"github.com/grafana/opentelemetry-go/sdk/resource"
	sdktrace "github.com/grafana/opentelemetry-go/sdk/trace"
	semconv "github.com/grafana/opentelemetry-go/semconv/v1.26.0"
	"github.com/grafana/opentelemetry-go/trace"
)

var logger = log.New(os.Stderr, "zipkin-example", log.Ldate|log.Ltime|log.Llongfile)

// initTracer creates a new trace provider instance and registers it as global trace provider.
func initTracer(url string) (func(context.Context) error, error) {
	// Create Zipkin Exporter and install it as a global tracer.
	//
	// For demoing purposes, always sample. In a production application, you should
	// configure the sampler to a trace.ParentBased(trace.TraceIDRatioBased) set at the desired
	// ratio.
	exporter, err := zipkin.New(
		url,
		zipkin.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	batcher := sdktrace.NewBatchSpanProcessor(exporter)

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(batcher),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("zipkin-test"),
		)),
	)
	otel.SetTracerProvider(tp)

	return tp.Shutdown, nil
}

func main() {
	url := flag.String("zipkin", "http://localhost:9411/api/v2/spans", "zipkin url")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	shutdown, err := initTracer(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			log.Fatal("failed to shutdown TracerProvider: %w", err)
		}
	}()

	tr := otel.GetTracerProvider().Tracer("component-main")
	ctx, span := tr.Start(ctx, "foo", trace.WithSpanKind(trace.SpanKindServer))
	<-time.After(6 * time.Millisecond)
	bar(ctx)
	<-time.After(6 * time.Millisecond)
	span.End()
}

func bar(ctx context.Context) {
	tr := otel.GetTracerProvider().Tracer("component-bar")
	_, span := tr.Start(ctx, "bar")
	<-time.After(6 * time.Millisecond)
	span.End()
}
