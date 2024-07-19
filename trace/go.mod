module github.com/grafana/opentelemetry-go/trace

go 1.21

replace github.com/grafana/opentelemetry-go => ../

require (
	github.com/google/go-cmp v0.6.0
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go/metric => ../metric
