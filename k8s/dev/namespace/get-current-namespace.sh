#!/bin/bash
source "$PROJECTNAME_ROOT/dev/boilerplates/base-bash-include.sh"

kubectl config view | grep namespace:
