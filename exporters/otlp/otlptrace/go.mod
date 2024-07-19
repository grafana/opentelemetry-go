module github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace

go 1.21

require (
	github.com/google/go-cmp v0.6.0
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/proto/otlp v1.3.1
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go => ../../..

replace github.com/grafana/opentelemetry-go/sdk => ../../../sdk

replace github.com/grafana/opentelemetry-go/trace => ../../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../../metric
