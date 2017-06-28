#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"

ALL_PROTOS_DIR=$PROJECTNAME_ROOT/build/proto
rm -rf $ALL_PROTOS_DIR
mkdir -p $ALL_PROTOS_DIR

# Compile go protos, and copy each proto file into build/proto/**
cd $PROJECTNAME_ROOT/go/src
GO_PROTOS=`find examples projectname | grep .proto$ | grep -v /vendor/`
for FILE_PATH in $GO_PROTOS; do
	protoc -I . $FILE_PATH --go_out=plugins=grpc:.
	# Copy each proto file into build/proto/** so that uncompiled languages can cp all of them
	# into their directory (and docker container)
	TARGET_DIR="$ALL_PROTOS_DIR/`dirname $FILE_PATH`"
	mkdir -p $TARGET_DIR
	cp $FILE_PATH $TARGET_DIR
done
