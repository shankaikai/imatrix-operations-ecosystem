#!/usr/bin/env bash
protoc \
    --plugin="protoc-gen-ts=./../client/node_modules/.bin/protoc-gen-ts" \
    --js_out="import_style=commonjs,binary:./../client/src/proto" \
    --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./../client/src/proto \
    --ts_out="./../client/src/proto" \
    operations_ecosys.proto