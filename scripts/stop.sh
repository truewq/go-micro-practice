#! /bin/bash

# stop all
scriptD=$PWD/scripts

function startService()
{
    if [[ $1 == \#* ]]; then
        echo "skip ${1#\#} ..."
        echo
    elif [[ -z $1 ]]; then
        echo
    else
        echo "=> building $1 <="
        pkill -9 $1
    	echo
    fi
}

while read cmd;do
    startService $cmd
done < $scriptD/cmds2build.list
