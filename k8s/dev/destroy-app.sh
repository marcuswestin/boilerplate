#!/bin/bash
source "$PROJECTNAME_ROOT/scripts/_boilerplates/base-bash-include.sh"

APP_NAME=$1

if [ "$APP_NAME" = "" ]; then
	echo "usage: destroy-app.sh <app name>"
	exit -1
fi

cd k8s/apps/$APP_NAME

kubectl delete -f ./$APP_NAME.tmpl.yaml
