# Service Weaver Demo

## Service Weaver Installation

```shell
go install github.com/ServiceWeaver/weaver/cmd/weaver@latest
```

## Generate Code

```shell
weaver generate .
```

## Run App

```go
go run .
```

## Test the request

每次启动都会有随即端口，可以将端口打印出来

```go
fmt.Println("listening on", a.lis)
```

```shell
curl http://localhost:36289/r\?s\=1
```

## Deploy multi processes locally

- config.toml
- `weaver generate .`
- `go build .`
- `weaver multi deploy config.toml`

```shell
weaver multi deploy config.toml
```

╭───────────────────────────────────────────────────╮
│ app        : demo                                 │
│ deployment : 1d1d3728-738c-46c7-9b0b-78451b0b35be │
╰───────────────────────────────────────────────────╯
S0101 07:30:00.000000 stdout               9397bd0f                      │ listening on 127.0.0.1:9090
S0101 07:30:00.000000 stdout               343f0d04                      │ listening on 127.0.0.1:9090

```shell
weaver multi status
```

╭───────────────────────────────────────────────────╮
│ DEPLOYMENTS                                       │
├──────┬──────────────────────────────────────┬─────┤
│ APP  │ DEPLOYMENT                           │ AGE │
├──────┼──────────────────────────────────────┼─────┤
│ demo │ 1d1d3728-738c-46c7-9b0b-78451b0b35be │ 30s │
╰──────┴──────────────────────────────────────┴─────╯
╭──────────────────────────────────────────────────────╮
│ COMPONENTS                                           │
├──────┬────────────┬───────────────┬──────────────────┤
│ APP  │ DEPLOYMENT │ COMPONENT     │ REPLICA PIDS     │
├──────┼────────────┼───────────────┼──────────────────┤
│ demo │ 1d1d3728   │ demo.Reverser │ 1128807, 1128814 │
│ demo │ 1d1d3728   │ weaver.Main   │ 1128793, 1128800 │
╰──────┴────────────┴───────────────┴──────────────────╯
╭───────────────────────────────────────────────╮
│ LISTENERS                                     │
├──────┬────────────┬──────────┬────────────────┤
│ APP  │ DEPLOYMENT │ LISTENER │ ADDRESS        │
├──────┼────────────┼──────────┼────────────────┤
│ demo │ 1d1d3728   │ demo     │ 127.0.0.1:9090 │
╰──────┴────────────┴──────────┴────────────────╯

## Metrics and Tracing

```shell
weaver multi dashboard 
```

