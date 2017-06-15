#!/bin/bash
source "$PROJECTNAME_ROOT/scripts/_boilerplates/base-bash-include.sh"

cd k8s/hello-node

kubectl delete -f ./hello-node.tmpl.yaml
