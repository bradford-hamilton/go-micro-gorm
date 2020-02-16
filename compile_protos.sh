#!/bin/bash

set +x
set -e

for pp in $(find ./proto -mindepth 1 -maxdepth 1 -type d)
do
  echo "compiling $pp"
  protoc \
    --proto_path=${GOPATH}/src:. \
    --micro_out=. \
    --go_out=. \
  ${pp}/*.proto
done
