#!/bin/bash

[[ "$SSLEEP_BY" == "" ]] && SLEEP_BY=300

function run () {
    git pull origin master > /dev/null 2>&1
    python3 init.py &
    sleep $SLEEP_BY
    killall -9 python3
    ./runner.sh
}

cd "`dirname \"$0\"`"

run
