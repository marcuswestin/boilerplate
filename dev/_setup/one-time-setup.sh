set -e

FILEPATH=`pwd`/${BASH_SOURCE[0]}
cd `dirname $FILEPATH`
cd `git rev-parse --show-toplevel`
PROJECTNAME_ROOT=`pwd`

echo "
# PROJECTNAME
alias projectname=\"cd $PROJECTNAME_ROOT && source dev/_setup/source-dev.sh\"" >> ~/.bash_profile

bold=$(tput bold)
normal=$(tput sgr0)
echo "
	Setup done in $PROJECTNAME_ROOT.
	Please quite all terminal windows.
	From now on, run ${bold}projectname${normal} to enter dev env.
"
