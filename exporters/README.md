# OpenTelemetry Exporters

Once the OpenTelemetry SDK has created and processed telemetry, it needs to be exported.
This package contains exporters for this purpose.

## Exporter Packages

The following exporter packages are provided with the following OpenTelemetry signal support.

|                                           Exporter Package                                            | Logs | Metrics | Traces |
|:-----------------------------------------------------------------------------------------------------:|:----:|:-------:|:------:|
|       [github.com/grafana/opentelemetry-go/exporters/otlp/otlplog/otlploghttp](./otlp/otlplog/otlploghttp)       |   ✓  |         |        |
| [github.com/grafana/opentelemetry-go/exporters/otlp/otlpmetric/otlpmetricgrpc](./otlp/otlpmetric/otlpmetricgrpc) |      |   ✓     |        |
| [github.com/grafana/opentelemetry-go/exporters/otlp/otlpmetric/otlpmetrichttp](./otlp/otlpmetric/otlpmetrichttp) |      |   ✓     |        |
|   [github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace/otlptracegrpc](./otlp/otlptrace/otlptracegrpc)   |      |         |   ✓    |
|   [github.com/grafana/opentelemetry-go/exporters/otlp/otlptrace/otlptracehttp](./otlp/otlptrace/otlptracehttp)   |      |         |   ✓    |
|                     [github.com/grafana/opentelemetry-go/exporters/prometheus](./prometheus)                     |      |   ✓     |        |
|                [github.com/grafana/opentelemetry-go/exporters/stdout/stdoutlog](./stdout/stdoutlog)              |   ✓  |         |        |
|            [github.com/grafana/opentelemetry-go/exporters/stdout/stdoutmetric](./stdout/stdoutmetric)            |      |   ✓     |        |
|             [github.com/grafana/opentelemetry-go/exporters/stdout/stdouttrace](./stdout/stdouttrace)             |      |         |   ✓    |
|                         [github.com/grafana/opentelemetry-go/exporters/zipkin](./zipkin)                         |      |         |   ✓    |

See the [OpenTelemetry registry] for 3rd-party exporters compatible with this project.

[OpenTelemetry registry]: https://opentelemetry.io/registry/?language=go&component=exporter
