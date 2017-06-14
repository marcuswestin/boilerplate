echo "+ cd \"$OSTRAA_ROOT\"; set -e; trap 'kill -HUP 0' SIGINT; set -x"

# Start script in project root;
cd "$OSTRAA_ROOT";

# Always die on error;
set -e;

# Kill sub-processes when we're killed;
trap 'kill -HUP 0' SIGINT;

# Echo every line
set -x
