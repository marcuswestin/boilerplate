_BASH_SOURCE_REL_DIR=`pwd`

# Start all scripts in project root;
cd "$PROJECTNAME_ROOT";

# Always die on error;
set -e;

# Kill sub-processes when we're killed;
trap 'kill -HUP 0' SIGINT;

# Make it easy to go into a script folder
function cd_bash_source() {
	local _BASH_SOURCE="$1"
	cd "$( cd $_BASH_SOURCE_REL_DIR && cd "$( dirname "${_BASH_SOURCE[0]}" )" && pwd )"
}
