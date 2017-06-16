#!/bin/bash
source "$PROJECTNAME_ROOT/scripts/_boilerplates/base-bash-include.sh"

cd_bash_source $BASH_SOURCE

IMAGE_TAG=${1?"Usage: $0 <image tag>"}

dep ensure -update

bash ./proto-gen.sh

docker build . -f echo-server.Dockerfile      -t examples/grpc-echo-server:$IMAGE_TAG
docker build . -f echo-http-bridge.Dockerfile -t examples/grpc-echo-http-bridge:$IMAGE_TAG
