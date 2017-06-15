#!/bin/bash
source "$PROJECTNAME_ROOT/scripts/_boilerplates/base-bash-include.sh"

cd k8s/hello-node

set +x
APP_NAME=hello-node
WHOAMI="`whoami`"
TAG_NAME=dev
IMAGE=projectname/$APP_NAME:$TAG_NAME
GIT_DIRTY_POSTFIX=`[[ $(git status 2> /dev/null | tail -n1) != "nothing to commit (working directory clean)" ]] && echo "-dirty"`
GIT_SHA="`git rev-parse HEAD`$GIT_DIRTY_POSTFIX"
TIME="`date +%s`"
YAML_OUT_FILE=/tmp/$APP_NAME-$WHOAMI.yaml

cat ./hello-node.tmpl.yaml \
	| sed "s/{{ GIT-SHA }}/$GIT_SHA/" \
	| sed "s/{{ TIME }}/$TIME/" \
	| sed "s/{{ WHOAMI }}/$WHOAMI/" \
	| sed "s!{{ IMAGE }}!$IMAGE!" \
	> $YAML_OUT_FILE

set -x

# Test config file
kubectl apply -f $YAML_OUT_FILE --dry-run
# Build image
docker build -t $IMAGE .
# Deploy new image
kubectl apply --record -f $YAML_OUT_FILE
# Wait for deployment to finish
./rollout-status.sh
# Curl
curl http://`ksolo ip`:30001/
