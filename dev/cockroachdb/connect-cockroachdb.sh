#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"

check_install cockroach "brew install cockroach"

cockroach sql --insecure --port=5432 --database=projectnametest