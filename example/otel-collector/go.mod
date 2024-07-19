module github.com/grafana/opentelemetry-go/example/otel-collector

go 1.21

replace (
	github.com/grafana/opentelemetry-go => ../..
	github.com/grafana/opentelemetry-go/sdk => ../../sdk
)

require (
	github.com/grafana/opentelemetry-go v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/otlp/otlpmetric/otlpmetricgrpc v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace/otlptracegrpc v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/sdk/metric v1.28.0-grafana.5-renamed
	github.com/grafana/opentelemetry-go/trace v1.28.0-grafana.5-renamed
	google.golang.org/grpc v1.64.0
)

require (
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace v1.28.0-grafana.5-renamed // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240701130421-f6361c86f094 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/grafana/opentelemetry-go/trace => ../../trace

replace github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace => ../../exporters/otlp/otlptrace

replace github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace/otlptracegrpc => ../../exporters/otlp/otlptrace/otlptracegrpc

replace github.com/grafana/opentelemetry-go/metric => ../../metric

replace github.com/grafana/opentelemetry-go/sdk/metric => ../../sdk/metric

replace github.com/grafana/opentelemetry-go/exporters/otlp/otlpmetric/otlpmetricgrpc => ../../exporters/otlp/otlpmetric/otlpmetricgrpc
