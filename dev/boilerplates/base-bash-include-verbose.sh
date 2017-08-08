echo "+ cd \"$PROJECTNAME_ROOT\"; set -e; trap 'kill -HUP 0' SIGINT; set -x"
source "$PROJECTNAME_ROOT/dev/boilerplates/base-bash-include.sh"

# Echo every line
set -x
