#!/bin/bash
sudo apt install --yes make python3-pip
cd /iot/sensors-publisher-python
ENV=production make supervisor
