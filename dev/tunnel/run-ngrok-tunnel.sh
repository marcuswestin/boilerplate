#!/bin/bash
source "$PROJECTNAME_ROOT/dev/_setup/boilerplates/base-bash-include.sh"

check_install ngrok "brew install ngrok"

# nohup ngrok http -subdomain projectname 9090 &> ngrok.out &
ngrok http -subdomain projectname 9080

tail -f nohup.out