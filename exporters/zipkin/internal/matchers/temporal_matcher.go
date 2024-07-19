// Code created by gotmpl. DO NOT MODIFY.
// source: internal/shared/matchers/temporal_matcher.go.tmpl

// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package matchers // import "github.com/grafana/opentelemetry-go/exporters/zipkin/internal/matchers"

type TemporalMatcher byte

//nolint:revive // ignoring missing comments for unexported constants in an internal package
const (
	Before TemporalMatcher = iota
	BeforeOrSameTime
	After
	AfterOrSameTime
)
