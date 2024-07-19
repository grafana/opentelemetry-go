module github.com/grafana/opentelemetry-go/example/dice

go 1.21

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/stdout/stdoutlog v0.4.0
	github.com/grafana/opentelemetry-go/exporters/stdout/stdoutmetric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/stdout/stdouttrace v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/log v0.4.0
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk/log v0.4.0
	github.com/grafana/opentelemetry-go/sdk/metric v1.28.0-grafana.5-renamed
	go.opentelemetry.io/contrib/bridges/otelslog v0.2.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.52.0
)

require (
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed // indirect
	go.opentelemetry.io/otel v1.27.0 // indirect
	go.opentelemetry.io/otel/log v0.3.0 // indirect
	go.opentelemetry.io/otel/metric v1.27.0 // indirect
	go.opentelemetry.io/otel/trace v1.27.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace github.com/grafana/opentelemetry-go/exporters/stdout/stdouttrace => ../../exporters/stdout/stdouttrace

replace github.com/grafana/opentelemetry-go/exporters/stdout/stdoutmetric => ../../exporters/stdout/stdoutmetric

replace github.com/grafana/opentelemetry-go => ../..

replace github.com/grafana/opentelemetry-go/trace => ../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../metric

replace github.com/grafana/opentelemetry-go/sdk/metric => ../../sdk/metric

replace github.com/grafana/opentelemetry-go/sdk => ../../sdk

replace github.com/grafana/opentelemetry-go/exporters/stdout/stdoutlog => ../../exporters/stdout/stdoutlog

replace github.com/grafana/opentelemetry-go/log => ../../log

replace github.com/grafana/opentelemetry-go/sdk/log => ../../sdk/log
