#!/bin/bash

SOURCE_DIR="../"

docker build \
  -t data-analysis-system-web \
  -f ./Dockerfile.web \
  ${SOURCE_DIR}