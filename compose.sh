#!/bin/bash

docker-compose -f=./docker-compose.yml \
  -p=data-analysis-system \
  up \
  --force-recreate \
  -d
