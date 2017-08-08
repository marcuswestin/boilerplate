#!/bin/bash
source "$PROJECTNAME_ROOT/dev/boilerplates/base-bash-include.sh"

check_install budo "npm install -g budo"

cd js/projectname && ./node_modules/.bin/budo 'apps/app-bootstraps/web-desktop-app-bootstrap.js' \
	--dir ../../ \
	--debug true --port 9903 --live --open --title "PROJECTNAME" \
	-- -t babelify --presets es2015
