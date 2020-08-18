#!/bin/bash

pushd /

echo "Installing globally github.com/google/protobuf/src"
go get -v -u github.com/google/protobuf/src

echo "Installing globally github.com/mwitkow/go-proto-validators"
go get -v -u github.com/mwitkow/go-proto-validators
go get -v -u github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
go get -v -u github.com/gogo/protobuf/protoc-gen-gofast

popd