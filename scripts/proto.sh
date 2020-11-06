#! /bin/bash
#set -x

protoPath="../internal/proto/"

function genMicroProto()
{
    protoFilePath=$1
    protoFileName=$2
    echo "=> gen micro+proto  $protoFilePath/$protoFileName <="
    protoc --proto_path=$protoFilePath --proto_path=${protoPath} --micro_out=$protoFilePath --go_out=$protoFilePath $protoFileName
}

function genProto()
{
  for f in `ls $1`
  do
    fPath=$1"/"$f
    #if fPath is a dir, then exec genProto on this dir.
    if [ -d $fPath ]
    then
      genProto $fPath
    else
      if [[ $f == *.proto ]]; then
       genMicroProto $1 $f  
      fi
    fi
  done
}

genProto $protoPath
