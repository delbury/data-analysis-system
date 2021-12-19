#!/bin/bash

SOURCE_DIR="../"

docker build \
  -t data-analysis-system-server \
  -f ./Dockerfile.server \
  ${SOURCE_DIR}