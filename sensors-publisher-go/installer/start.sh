#!/bin/bash

set -e

echo "sh: starting..."

source ./installer/clear.sh

NOW=$(date +"%y%m%d_%H%M%S")
export RUN_FILE="sensors-publisher-go-installer_${NOW}.run"

source ./installer/build.sh

source ./installer/send.sh

source ./installer/run.sh

echo "sh: starting done"
