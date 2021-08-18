#!/bin/bash
git pull origin master >/dev/null
python3 init.py
sleep 5
./test.sh
