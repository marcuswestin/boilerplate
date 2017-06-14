#!/bin/bash
source "$OSTRAA_ROOT/scripts/_boilerplates/base-bash-include.sh"

cd js && ./node_modules/.bin/budo 'apps/app-bootstraps/web-desktop-app-bootstrap.js' \
	--dir ../../ \
	--debug true --port 9903 --live --open --title "ProjectName" \
	-- -t babelify --presets es2015
