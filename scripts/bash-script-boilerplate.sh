#!/bin/bash
set -e                    # Always die on error
trap 'kill -HUP 0' EXIT   # Kill sub-processes when we die
trap 'kill -HUP 0' SIGINT # Ditto
cd "$PROJECTNAME_ROOT"    # Start script in project root

echo "bash script boilerplate"
