#!/bin/bash
source "$PROJECTNAME_ROOT/dev/boilerplates/base-bash-include.sh"

cd go/src/examples

protoc -I ./grpc-echo ./grpc-echo/echo.proto --go_out=plugins=grpc:grpc-echo
