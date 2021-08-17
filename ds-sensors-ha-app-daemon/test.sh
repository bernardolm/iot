#!/bin/bash
git pull origin master
sudo python3 sensor.py
sleep 5
./test.sh
