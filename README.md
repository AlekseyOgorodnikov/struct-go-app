# Struct app golang hexogonal architecture

## Install protobuf for using grps

`sudo apt install protobuf-compiler`

Install the protocol compiler plugins for Go using the following commands:
`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26` and
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1`

Update your PATH so that the protoc compiler can find the plugins:
`export PATH="$PATH:$(go env GOPATH)/bin"`

## code gen grpc

`protoc --go_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/number_msg.proto`

## and

`protoc --go-grpc_out=internal/adapters/framework/left/grpc --proto_path=internal/adapters/framework/left/grpc/proto internal/adapters/framework/left/grpc/proto/arithmetic_svc.proto`

## Install module

`go get -u google.golang.org/grpc`
