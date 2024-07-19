module github.com/grafana/opentelemetry-go/example/zipkin

go 1.21

replace (
	github.com/grafana/opentelemetry-go => ../..
	github.com/grafana/opentelemetry-go/exporters/zipkin => ../../exporters/zipkin
	github.com/grafana/opentelemetry-go/sdk => ../../sdk
)

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/zipkin v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	golang.org/x/sys v0.21.0 // indirect
)

replace github.com/grafana/opentelemetry-go/trace => ../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../metric
