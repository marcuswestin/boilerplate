#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"

NAMESPACE=$1
if [ "$NAMESPACE" = "" ]; then
	NAMESPACE="dev-`whoami`"
fi

kubectl delete namespace "$NAMESPACE" || echo "already destroyed"
kubectl config set-context $(kubectl config current-context) --namespace=""
