// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/grafana/opentelemetry-go/exporters/otlp/otlplog/otlploggrpc/internal"

//go:generate gotmpl --body=../../../../../internal/shared/otlp/retry/retry.go.tmpl "--data={}" --out=retry/retry.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/retry/retry_test.go.tmpl "--data={}" --out=retry/retry_test.go

//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/attr_test.go.tmpl "--data={}" --out=transform/attr_test.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log.go.tmpl "--data={}" --out=transform/log.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log_attr_test.go.tmpl "--data={}" --out=transform/log_attr_test.go
//go:generate gotmpl --body=../../../../../internal/shared/otlp/otlplog/transform/log_test.go.tmpl "--data={}" --out=transform/log_test.go
