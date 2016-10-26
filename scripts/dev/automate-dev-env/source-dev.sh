# Go to repo root
_source_path=`pwd`/${BASH_SOURCE[0]}
cd `dirname $_source_path`
cd `git rev-parse --show-toplevel`

_set_if_not_equal () {
	if [[ "${!1}" != "$2" ]]; then
		echo "	export $1=$2"		# Print
		export $1="$2"				# Export
	fi
}

_postfix_if_not_contains () {
	# echo ":${!1}:" *":$2:"*
	if [[ ":${!1}:" != *":$2:"* ]]; then
		echo "	export $1=\$$1:$2"	# Print
		# echo "export $1=${!1}:$2"	# Debug
		export $1="${!1}:$2"		# Export
	fi
}

_prefix_if_not_contains () {
	# echo ":${!1}:" *":$2:"*
	if [[ ":${!1}:" != *":$2:"* ]]; then
		echo "	export $1=\$$1:$2"	# Print
		# echo "export $1=${!1}:$2"	# Debug
		export $1="$2:${!1}"		# Export
	fi
}

# projectname
_set_if_not_equal PS1 "\n<projectname> \w λ "
_set_if_not_equal PROJECTNAME_ROOT `pwd`
# Brew
_postfix_if_not_contains PATH "/usr/local/bin"
# Go
_set_if_not_equal GOPATH "${PROJECTNAME_ROOT}/go"
_prefix_if_not_contains PATH "/usr/local/go/bin"
_prefix_if_not_contains PATH "${GOPATH}/bin"
# MySQL
_prefix_if_not_contains PATH "/usr/local/mysql/bin"
# Python
_set_if_not_equal PYTHONPATH "${PROJECTNAME_ROOT}/py"
# JS
_set_if_not_equal NODE_PATH "${PROJECTNAME_ROOT}/js"
# Android
_set_if_not_equal ANDROID_HOME "/usr/local/opt/android-sdk"
_prefix_if_not_contains PATH "${ANDROID_HOME}/tools"
_prefix_if_not_contains PATH "${ANDROID_HOME}/platform-tools"
# Install tools
./scripts/dev/automate-dev-env/setup-tools.sh
