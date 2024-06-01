#!/bin/bash

echo "sh: cleaning..."

find bin ! -name '.keep' -type f -exec rm -f {} +
find dist ! -name '.keep' -type f -exec rm -f {} +
find . ! -name '.keep' -wholename '*/tmp/*' -type f -exec rm -f {} +

echo "sh: cleaning done"
