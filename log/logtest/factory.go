// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package logtest // import "github.com/grafana/opentelemetry-go/log/logtest"

import (
	"time"

	"github.com/grafana/opentelemetry-go/log"
)

// RecordFactory is used to facilitate unit testing bridge implementations that
// make use of a [github.com/grafana/opentelemetry-go/log.Record]
//
// Do not use RecordFactory to create records in production code.
type RecordFactory struct {
	Timestamp         time.Time
	ObservedTimestamp time.Time
	Severity          log.Severity
	SeverityText      string
	Body              log.Value
	Attributes        []log.KeyValue
}

// NewRecord returns a log record.
func (b RecordFactory) NewRecord() log.Record {
	var record log.Record
	record.SetTimestamp(b.Timestamp)
	record.SetObservedTimestamp(b.ObservedTimestamp)
	record.SetSeverity(b.Severity)
	record.SetSeverityText(b.SeverityText)
	record.SetBody(b.Body)
	record.AddAttributes(b.Attributes...)

	return record
}
