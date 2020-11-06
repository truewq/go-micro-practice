#! /bin/bash

# stop all
scriptD=$PWD/scripts
buildD=$PWD/build/package

function startService()
{
    if [[ $1 == \#* ]]; then
        echo "skip ${1#\#} ..."
        echo
    elif [[ -z $1 ]]; then
        echo
    else
        echo "=> building $1 <="
        $buildD/$1/$1 &
    	echo
    fi
}

while read cmd;do
    startService $cmd
done < $scriptD/cmds2build.list
