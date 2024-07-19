module github.com/grafana/opentelemetry-go/bridge/opencensus/test

go 1.21

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/bridge/opencensus v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
	go.opencensus.io v0.24.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed // indirect
	github.com/grafana/opentelemetry-go/sdk/metric v1.28.0-grafana.5-renamed // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace github.com/grafana/opentelemetry-go => ../../..

replace github.com/grafana/opentelemetry-go/bridge/opencensus => ../

replace github.com/grafana/opentelemetry-go/sdk => ../../../sdk

replace github.com/grafana/opentelemetry-go/trace => ../../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../../metric

replace github.com/grafana/opentelemetry-go/sdk/metric => ../../../sdk/metric
