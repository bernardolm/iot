#!/bin/bash

echo "sh: running..."

ssh -t "${SYSTEM_USER}@${SYSTEM_HOST}" "/tmp/${RUN_FILE} --nox11"

echo "sh: running done"
