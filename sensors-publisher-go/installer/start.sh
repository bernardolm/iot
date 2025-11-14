#!/bin/bash

set -e

echo "sh: starting..."

source ./installer/clear.sh

NOW=$(date +"%Y%m%d_%H%M%S")
export RUN_FILE="sensors-publisher-go-installer_${NOW}.run"
echo "> run file is $RUN_FILE"

source ./installer/build.sh

source ./installer/send.sh

source ./installer/run.sh

echo "sh: starting done"
