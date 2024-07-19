// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otel // import "github.com/grafana/opentelemetry-go"

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/grafana/opentelemetry-go/metric"
	"github.com/grafana/opentelemetry-go/metric/embedded"
	"github.com/grafana/opentelemetry-go/metric/noop"
)

type testMeterProvider struct{ embedded.MeterProvider }

var _ metric.MeterProvider = &testMeterProvider{}

func (*testMeterProvider) Meter(_ string, _ ...metric.MeterOption) metric.Meter {
	return noop.NewMeterProvider().Meter("")
}

func TestMultipleGlobalMeterProvider(t *testing.T) {
	p1 := testMeterProvider{}
	p2 := noop.NewMeterProvider()
	SetMeterProvider(&p1)
	SetMeterProvider(p2)

	got := GetMeterProvider()
	assert.Equal(t, p2, got)
}
