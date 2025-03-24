1.安装protoc,一般是protoc.exe,官网:https://github.com/protocolbuffers/protobuf/releases

2.安装protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0

3.安装protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0

把protoc.exe,protoc-gen-go.exe, protoc-gen-go-grpc.exe放到GDK的bin目录下

4.编译
protoc --go_out=. --go-grpc_out=. *.proto