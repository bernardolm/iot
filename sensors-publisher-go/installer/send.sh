#!/bin/bash

echo "sh: sending..."

rsync -r "installer/tmp/${RUN_FILE}" "${SYSTEM_USER}@${SYSTEM_HOST}:/tmp"

echo "sh: sending done"
