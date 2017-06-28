#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include-quiet.sh"

# Check input
#############
APP_DIR_RAW=${1?"Usage: $0 <path to app dir>"}
APP_DIR=${APP_DIR_RAW%/}

if [ ! -d "$APP_DIR" ]; then
	echo "Not a directory: $APP_DIR"
	exit -1
fi

BUILD_FILE=$APP_DIR/build.sh

if [ ! -f "$BUILD_FILE" ]; then
	echo "No such file: $BUILD_FILE"
	exit -1
fi

APP_NAME=`basename $APP_DIR`
YAML_TMPL_FILE=$APP_DIR/$APP_NAME.tmpl.yaml
echo "$YAML_TMPL_FILE"

if [ ! -f "$YAML_TMPL_FILE" ]; then
	echo "No such file: $YAML_TMPL_FILE"
	exit -1
fi

# Check git status
##################
IMAGE_TAG=dev
GIT_DIRTY_POSTFIX=`[[ $(git status 2> /dev/null | tail -n1) != "nothing to commit (working directory clean)" ]] && echo "-dirty"`

if [ "$IMAGE_TAG" != "dev" && "$GIT_DIRTY_POSTFIX" = "-dirty" ]; then
	echo "Git repo may only be dirty for images tagged dev (current: $IMAGE_TAG)"
	exit -1
fi


# Process yaml template
#######################
WHOAMI="`whoami`"
GIT_SHA="`git rev-parse HEAD`$GIT_DIRTY_POSTFIX"
TIME="`date +%s`"
YAML_OUT_FILE=/tmp/$APP_NAME-$WHOAMI.yaml

cat $YAML_TMPL_FILE \
	| sed "s/{{ GIT-SHA }}/$GIT_SHA/g" \
	| sed "s/{{ TIME }}/$TIME/g" \
	| sed "s/{{ WHOAMI }}/$WHOAMI/g" \
	| sed "s/{{ IMAGE_TAG }}/$IMAGE_TAG/g" \
	> $YAML_OUT_FILE


# Test config, build image, deploy to kuberneters
#################################################
set -x
# Test config file
kubectl apply -f $YAML_OUT_FILE --dry-run
# Build app
bash $BUILD_FILE $IMAGE_TAG
# Deploy new image
kubectl apply --record -f $YAML_OUT_FILE

# # Wait for deployment to finish
# $PROJECTNAME_ROOT/k8s/rollouts/rollout-status.sh $APP_NAME
# # Curl
# curl http://`ksolo ip`:30001/
