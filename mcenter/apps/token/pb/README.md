# protoc命令生成go文件
```sh
protoc -I=. --go_out=. --go_opt=module="github.com/solodba/devcloud/mcenter" apps/token/pb/token.proto
protoc -I=. --go_out=. --go_opt=module="github.com/solodba/devcloud/mcenter" apps/token/pb/rpc.proto
protoc -I=. --go-grpc_out=. --go-grpc_opt=module="github.com/solodba/devcloud/mcenter" apps/token/pb/rpc.proto
```

# protoc-go-inject-tag给go文件打标签
```sh
protoc-go-inject-tag -input=apps/*/*.pb.go
```