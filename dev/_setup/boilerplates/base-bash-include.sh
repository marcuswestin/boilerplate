echo "+ cd \"$PROJECTNAME_ROOT\"; set -e; trap 'kill -HUP 0' SIGINT; set -x"
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include-quiet.sh"

# Echo every line
set -x