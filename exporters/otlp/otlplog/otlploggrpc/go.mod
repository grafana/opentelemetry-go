module github.com/grafana/opentelemetry-go/exporters/otlp/otlplog/otlploggrpc

go 1.21

require (
	github.com/cenkalti/backoff/v4 v4.3.0
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/log v0.4.0
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk/log v0.3.0
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/proto/otlp v1.3.1
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/grafana/opentelemetry-go => ../../../..

replace github.com/grafana/opentelemetry-go/sdk/log => ../../../../sdk/log

replace github.com/grafana/opentelemetry-go/sdk => ../../../../sdk

replace github.com/grafana/opentelemetry-go/log => ../../../../log

replace github.com/grafana/opentelemetry-go/trace => ../../../../trace

replace github.com/grafana/opentelemetry-go/metric => ../../../../metric
