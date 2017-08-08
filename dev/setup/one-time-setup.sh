set -e

FILEPATH=`pwd`/${BASH_SOURCE[0]}
cd `dirname $FILEPATH`
cd `git rev-parse --show-toplevel`
PROJECTNAME_ROOT=`pwd`
mv dev/sublime.sublime-project dev/projectname.sublime-project

echo "
# PROJECTNAME
alias enter-projectname=\"cd $PROJECTNAME_ROOT && source dev/setup/source-dev.sh\"" >> ~/.bash_profile

bold=$(tput bold)
normal=$(tput sgr0)
echo "
	Setup done in $PROJECTNAME_ROOT.
	Please quite all terminal windows.
	From now on, run ${bold}enter-projectname${normal} to enter dev env.
"
