set -e
cd ${PROJECTNAME_ROOT}

_install_if_not_exists () {
	if [[ ! `which ${1}` ]]; then
		echo "	install ${1}"
		echo "		${2}"
		${2}
	fi
}

_install_if_not_exists 'xcodebuild' 'xcode-select --install'
_install_if_not_exists 'brew' 'ruby -e "$(curl https://raw.githubusercontent.com/Homebrew/install/master/install)"'
_install_if_not_exists 'go' 'brew install go'
_install_if_not_exists 'glide' 'brew install glide'
_install_if_not_exists 'goimports' 'go get golang.org/x/tools/cmd/goimports'
_install_if_not_exists 'protoeasy' 'go get go.pedge.io/protoeasy/cmd/protoeasy'
_install_if_not_exists 'protoc' 'brew install protobuf'
_install_if_not_exists 'protoc-gen-go' 'go get -a github.com/golang/protobuf/protoc-gen-go'


# Git hooks
for file in dev/_setup/git-hooks/*; do
	if [[ ! -f .git/hooks/`basename $file` ]]; then
		echo "symlink .git/hooks/`basename $file`"
		ln -snf ../../$file .git/hooks/`basename $file`
	fi
done
