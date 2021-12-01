#!/bin/bash

./wealthChange $@ | sort -n -k 2

echo "说明: \$包围的是富二代，*包围的是更为努力的人"
echo "./run.sh --help 查看命令参数"
