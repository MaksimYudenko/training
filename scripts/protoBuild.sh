#!/bin/bash
protoc \
-I=api/proto/v1 \
-I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:pkg/api/v1 \
service.proto

protoc \
-I=api/proto/v1 \
-I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:pkg/api/v1 \
service.proto

protoc \
-I=api/proto/v1 \
-I. \
-I/usr/local/include \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:api/swagger/v1 \
service.proto