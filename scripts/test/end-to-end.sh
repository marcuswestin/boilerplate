#!/bin/bash
source "$PROJECTNAME_ROOT/scripts/_boilerplates/base-bash-include.sh"

trap 'kill -HUP 0' EXIT   # Kill sub-processes when we exit

# cd js/projectname && ./node_modules/.bin/flow check
node --harmony-destructuring js/projectname/tests/all.js
