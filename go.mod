module github.com/grafana/opentelemetry-go

go 1.21

require (
	github.com/go-logr/logr v1.4.2
	github.com/go-logr/stdr v1.2.2
	github.com/google/go-cmp v0.6.0
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.7-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.7-renamed
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go/trace => ./trace

replace github.com/grafana/opentelemetry-go/metric => ./metric
