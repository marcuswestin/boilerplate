#!/bin/bash
source "$PROJECTNAME_ROOT/dev/boilerplates/base-bash-include.sh"

cd_bash_source $BASH_SOURCE

IMAGE_TAG=${1?"Usage: $0 <image tag>"}

docker build . -t examples/hello-node:$IMAGE_TAG
