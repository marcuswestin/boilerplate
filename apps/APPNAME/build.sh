#!/bin/bash
set -e                    ### Always die on error
trap 'kill -HUP 0' SIGINT ### Kill sub-processes when we're killed
cd "$PROJECTNAME_ROOT"    ### Run this script in project root

go build PROJECTNAME/apps/APPNAME -o $0
