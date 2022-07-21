#!/usr/bin/env bash
protoc -I="./" operations_ecosys.proto \
  --js_out=import_style=commonjs,binary:"./../client/src/proto" \
  --grpc-web_out=import_style=typescript,mode=grpcweb:"./../client/src/proto"