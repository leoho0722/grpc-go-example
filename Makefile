# 系統架構，amd64 或 arm64
ARCH ?= amd64

.PHONY: proto-gen-go
proto-gen-go:
ifeq ($(ARCH), arm64)
	/opt/homebrew/bin/protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/album.proto
else
	/usr/local/bin/protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/album.proto
endif