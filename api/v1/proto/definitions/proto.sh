#!/usr/bin/env bash

# official API proto v1
protoc -I=. --go_out=plugins=grpc:../ *.proto
