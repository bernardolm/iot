#!/bin/bash

echo "sh: cleaning..."

[ -d "bin" ] && find bin ! -name '.keep' -type f -exec rm -f {} +
[ -d "dist" ] && find dist ! -name '.keep' -type f -exec rm -f {} +
find . ! -name '.keep' -wholename '*/tmp/*' -type f -exec rm -f {} +

echo "sh: cleaning done"
