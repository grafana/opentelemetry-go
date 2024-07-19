module github.com/grafana/opentelemetry-go/exporters/prometheus

go 1.21

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
	github.com/prometheus/client_golang v1.19.1
	github.com/prometheus/client_model v0.6.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go => ../..

replace github.com/grafana/opentelemetry-go/sdk => ../../sdk

replace github.com/grafana/opentelemetry-go/sdk/metric => ../../sdk/metric

replace github.com/grafana/opentelemetry-go/trace => ../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../metric
