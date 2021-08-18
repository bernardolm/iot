#!/bin/bash

function run () {
    git pull origin master > /dev/null 2>&1
    python3 init.py &
    sleep 5
    killall -9 python3
    ./test.sh
}

cd "`dirname \"$0\"`"

run