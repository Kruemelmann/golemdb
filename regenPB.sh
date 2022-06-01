#!/bin/bash

# Build the registry api
protoc -I=. --go_out=. --go-grpc_out=. pkg/api/pb/peers.proto

