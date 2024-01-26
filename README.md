# edgex-driver-proto
    go get -u google.golang.org/grpc

    protoc 命令
    //https://github.com/protocolbuffers/protobuf/releases
    //下载放置到gopath bin下，包括include

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    //https://protobuf.dev/reference/go/go-generated/