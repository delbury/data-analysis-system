#!/bin/bash

CUR_DIR=$PWD

cd ${CUR_DIR}
echo $PWD
git pull --ff-only

cd ${CUR_DIR}
cd ../data-analysis-system-web
echo $PWD
git pull --ff-only

cd ${CUR_DIR}
cd ../data-analysis-system-server
echo $PWD
git pull --ff-only
