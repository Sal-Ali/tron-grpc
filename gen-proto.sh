#!/bin/bash

MODULE_OPT="--go_opt=module=github.com/Sal-Ali/tron-grpc --go-grpc_opt=module=github.com/Sal-Ali/tron-grpc"

rm -rf troncore tronapi

pushd proto

protoc -I. --go_out=../ --go-grpc_out=../ $MODULE_OPT tronapi/*.proto troncore/*.proto troncore/contract/*.proto

popd
