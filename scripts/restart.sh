#! /bin/bash
buildD=$PWD/build/package

for srv in $@; do
    echo "=>killing $srv<="
    pkill -9  "^$srv$"
    
    sleep 1
    echo "=>start $srv<="
    $buildD/$1/$1 &
done

echo
