#!/bin/sh

protoc productpb/product.proto --go_out=plugins=grpc:/go/src
