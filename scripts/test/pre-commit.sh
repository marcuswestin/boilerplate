#!/bin/bash
set -e                    # Always die on error
trap 'kill -HUP 0' SIGINT # Kill sub-processes when we're killed
cd "$PROJECTNAME_ROOT"    # Start script in project root

echo "TODO: Edit scripts/test/pre-commit.sh to run tests"