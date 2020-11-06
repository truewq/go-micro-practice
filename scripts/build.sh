#! /bin/bash
scriptD=$PWD/scripts
buildD=$PWD/build/package
srcD=$PWD/cmd
echo 

function goBuild()
{
    if [[ $1 == \#* ]]; then
        echo "skip ${1#\#} ..."
        echo
    elif [[ -z $1 ]]; then
        echo
    else
        echo "=> building $1 <="
        go build -o $buildD/$1/$1 $srcD/$1
    	echo
    fi
}
function buildAll()
{
    while read cmd;do
        goBuild $cmd
    done < $scriptD/cmds2build.list
}

function buildOne()
{
    for d in $@; do
        goBuild $d
    done
}


function ctrl_c()
{
    echo "=>build canceled by `whoami`<="
    exit 0
}


trap ctrl_c SIGINT

case $1 in
    "")
        buildAll
        ;;
    *)
        buildOne $@
        ;;
    clean)
        clean
esac
