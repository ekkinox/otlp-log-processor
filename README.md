# OTLP Log Processor

<!-- TOC -->
* [Run](#run)
* [Build](#build)
* [Tests](#tests)
* [Linter](#linter)
* [O11y](#o11y)
<!-- TOC -->

## Run
To run the gRPC application on `:4317` (with reflection enabled):

```shell
go run ./... --attribute={attribute} --interval={interval} --workers={workers}
```

- `{attribute}`: name of the log attribute to count, default `foo`
- `{interval}`: interval in milliseconds for printing the count, default `1000`
- `{workers}`: number of workers for the logs processing, default `10`

## Build

To build the application:

```shell
go build ./...
```

## Tests

To run the tests:

```shell
go test -v -race -failfast ./...
```

## Linter

To run the [linter](https://golangci-lint.run/):

```shell
golangci-lint run ./... 
```

## O11y

This gRPC application is instrumented with OTel logger and tracer (both exporting to stdout).

Logs produced during gRPC processing are correlated to the traces.