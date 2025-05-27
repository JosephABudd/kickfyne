clear
set -e # break on error.

go build >&build.log
mv ./kickfyne ~/bin
