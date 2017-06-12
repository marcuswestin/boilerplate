set -e
trap 'kill -HUP 0' SIGINT # Kill sub-processes when we're killed (projectname-api-server)
trap 'kill -HUP 0' EXIT   # Ditto when we're done

# cd ${PROJECTNAME_ROOT}/js/projectname && ./node_modules/.bin/flow check
node --harmony-destructuring ${PROJECTNAME_ROOT}/js/projectname/tests/all.js
