#!/bin/bash

[[ "$SSLEEP_BY" == "" ]] && SLEEP_BY=5

function run () {
    git pull origin master > /dev/null 2>&1
    python3 init.py &
    sleep $SLEEP_BY
    killall -9 python3
    ./test.sh
}

cd "`dirname \"$0\"`"

run
