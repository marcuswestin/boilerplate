#!/bin/bash
set -e                    # Always die on error
trap 'kill -HUP 0' SIGINT # Kill sub-processes when we're killed
cd "$PROJECTNAME_ROOT"    # Start script in project root

_install_if_not_exists () {
	if [[ ! `which ${1}` ]]; then
		echo "	install ${1}"
		echo "		${2}"
		${2}
	fi
}

_install_if_not_exists xcodebuild "xcode-select --install"
_install_if_not_exists brew 'ruby -e "$(curl https://raw.githubusercontent.com/Homebrew/install/master/install)"'
_install_if_not_exists glide "brew install glide"
_install_if_not_exists goimports "go get golang.org/x/tools/cmd/goimports"

# Git hooks
for file in misc/automate/git-hooks/*; do
	if [[ ! -f .git/hooks/`basename $file` ]]; then
		echo "symlink .git/hooks/`basename $file`"
		ln -snf ../../$file .git/hooks/`basename $file`
	fi
done
