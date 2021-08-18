#!/bin/bash
git pull origin master
python3 init.py
sleep 5
./test.sh
