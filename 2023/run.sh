#!/usr/bin/env bash

d=`date +"%_d"`
test=""
part="1"
daynextarg=0
partnextarg=0
debug=""

for arg in $@; do
    if [ $partnextarg -ne 0 ]; then
        partnextarg=0
        part=$arg
    elif [ $daynextarg -ne 0 ]; then
        daynextarg=0
        d=$arg
    elif [ "$arg" == "test" ]; then
        test="-test"
    elif [ "$arg" == "day" ]; then
        daynextarg=1
    elif [ "$arg" == "part" ]; then
        partnextarg=1
    elif [ "$arg" == "debug" ]; then
        debug="-debug"
    fi
done

cmd="go run . -day $d -part $part $test $debug"
echo $cmd
$cmd 
