#!/bin/bash

echo "sh: pre-building..."

sed -i "/SYSTEM_HOST/?/d" dist/.env
sed -i "/SYSTEM_USER/?/d" dist/.env

container_ip=$(docker inspect --format '{{ .NetworkSettings.IPAddress }}' alpine-sensors)

echo """
CREATE_USER_SKIP=true
SYSTEM_HOST=${container_ip}
SYSTEM_USER=root
""" >> dist/.env

echo "sh: pre-building done"
