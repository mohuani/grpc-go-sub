
直接使用 go-zero的goctl一键安装 protoc protoc-gen-go grpc 等工具

```shell
protoc --go_out=plugins=grpc:. sub.proto
```