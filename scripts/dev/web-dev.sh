set -e
cd "$(cd "`dirname "${BASH_SOURCE[0]}"`" && pwd -P)"

cd ../../js && ./node_modules/.bin/budo 'apps/app-bootstraps/web-desktop-app-bootstrap.js' \
	--dir ../../ \
	--debug true --port 9903 --live --open --title "ProjectName" \
	-- -t babelify --presets es2015
