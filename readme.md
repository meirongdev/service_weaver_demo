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
