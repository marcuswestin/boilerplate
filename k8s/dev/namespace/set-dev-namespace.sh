#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"
set +x

NAMESPACE=$1
if [ "$NAMESPACE" = "" ]; then
	NAMESPACE="dev-`whoami`"
fi

if kubectl get namespace | grep --quiet "$NAMESPACE"; then
	if kubectl get namespace "$NAMESPACE" | grep --quiet Terminating; then
		kubectl get namespace "$NAMESPACE"
		echo "Error: $NAMESPACE is terminating. Wait and try again."
		exit 1
	else
		echo "Namespace $NAMESPACE exists"
	fi
else
	kubectl create namespace "$NAMESPACE"
fi

kubectl config set-context $(kubectl config current-context) --namespace="$NAMESPACE"
