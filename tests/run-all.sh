#!/bin/bash

source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include-quiet.sh"

cleanup_on_exit

run_tests() {
	# Data
	######
	./dev/proto-gen.sh 		# &>/dev/null
	./sql/migrations/setup.sh 	# &>/dev/null

	# Server
	########
	killall importer &>/dev/null || true
	# go install projectname/apps/crm/importer
	# ENV=test importer &
	# sleep 0.1

	# Test client
	#############
	cd $PROJECTNAME_ROOT

	rm -rf tests/proto
	cp -r build/proto tests/proto

	node tests/end-to-end.js

	rm -rf tests/proto
	# killall importer
}

time run_tests
