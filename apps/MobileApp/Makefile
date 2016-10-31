run:	
	# Start packager with custom root
	make _run-packager _run-ios -j2
_run-packager:
	#node node_modules/react-native/local-cli/cli.js start --root $PROJECTNAME_ROOT/js
	node node_modules/react-native/local-cli/cli.js start --root ../../js
_run-ios:
	sleep 3 # allow packager to start
	react-native run-ios