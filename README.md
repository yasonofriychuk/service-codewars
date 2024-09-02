# service-codewars

## start
1. Установить proto [вот тут](https://grpc.io/docs/languages/go/quickstart/)

```shell
protoc -I proto proto/service/service.proto --go_out=./internal/generate/proto --go_opt=paths=source_relative --go-grpc_out=./internal/generate/proto --go-grpc_opt=paths=source_relative
```