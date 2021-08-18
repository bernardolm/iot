#!/bin/bash
git pull origin master 1>/dev/null
python3 init.py
sleep 5
./test.sh
