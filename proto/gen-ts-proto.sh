#!/usr/bin/env bash
protoc -I="./" iot_prototype.proto operations_ecosys.proto google/api/annotations.proto google/api/http.proto \
  --js_out=import_style=commonjs,binary:"./../client/src/proto" \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:"./../client/src/proto"
