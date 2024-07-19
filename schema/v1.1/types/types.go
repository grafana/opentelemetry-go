// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package types // import "github.com/grafana/opentelemetry-go/schema/v1.1/types"

import types10 "github.com/grafana/opentelemetry-go/schema/v1.0/types"

// TelemetryVersion is a version number key in the schema file (e.g. "1.7.0").
type TelemetryVersion types10.TelemetryVersion

// AttributeName is an attribute name string.
type AttributeName string

// AttributeValue is an attribute value.
type AttributeValue interface{}
