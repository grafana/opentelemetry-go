module github.com/grafana/opentelemetry-go/bridge/opentracing/test

go 1.21

replace github.com/grafana/opentelemetry-go => ../../..

replace github.com/grafana/opentelemetry-go/bridge/opentracing => ../

replace github.com/grafana/opentelemetry-go/trace => ../../../trace

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/bridge/opentracing v1.28.0-grafana.5-renamed
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/testify v1.9.0
	google.golang.org/grpc v1.64.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed // indirect
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go/metric => ../../../metric
