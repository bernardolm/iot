#!/bin/bash

echo "sh: sending..."

ssh-keygen -f "/home/bernardo/.ssh/known_hosts" -R "${SYSTEM_HOST}"
rsync -r "installer/tmp/${RUN_FILE}" "${SYSTEM_USER}@${SYSTEM_HOST}:/tmp"

echo "sh: sending done"
