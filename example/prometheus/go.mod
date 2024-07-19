module github.com/grafana/opentelemetry-go/example/prometheus

go 1.21

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/prometheus v0.50.0
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk/metric v1.28.0-grafana.5-renamed
	github.com/prometheus/client_golang v1.19.1
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed // indirect
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/sys v0.21.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/grafana/opentelemetry-go => ../..

replace github.com/grafana/opentelemetry-go/exporters/prometheus => ../../exporters/prometheus

replace github.com/grafana/opentelemetry-go/sdk => ../../sdk

replace github.com/grafana/opentelemetry-go/sdk/metric => ../../sdk/metric

replace github.com/grafana/opentelemetry-go/metric => ../../metric

replace github.com/grafana/opentelemetry-go/trace => ../../trace
