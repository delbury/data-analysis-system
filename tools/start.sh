#!/bin/bash

set -x

source ./var.sh

go build -o ${APP_PATH}

chmod a+x ${APP_PATH}

# kill 之前的进程，如果有的话
source ./stop.sh

# 启动
nohup ${APP_PATH} 1>./build/nohup.out &
