#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"

check_install cockroach "brew install cockroach"

cockroach start --insecure --host=localhost --store=/tmp/projectname/cockroachdb/node1 --background --port=5432 || echo "Already running?"
