#!/usr/bin/env bash
#protoc \
#    -I=. \
#    -I ${GOPATH}/src/github.com/lyft/protoc-gen-validate \
#    --go_out=plugins=grpc:../ \
#    --validate_out="lang=go:../" \
#    *.proto

protoc \
  -I . \
  -I ${GOPATH}/src \
  -I ${GOPATH}/src/github.com/lyft/protoc-gen-validate \
  --go_out=plugins=grpc:../ \
  --validate_out=lang=go:../ \
  *.proto