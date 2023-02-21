GOPATH:=$(shell go env GOPATH)
PROTO_FILES=$(shell find pb -name *.proto)

.PHONY: pbg
pbg: 
	protoc  --proto_path=./pb \
			--go_out=paths=source_relative:./pb \
			--go-grpc_out=paths=source_relative:./pb \
			$(PROTO_FILES)