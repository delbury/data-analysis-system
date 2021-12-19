#!/usr/bin/env bash

docker run \
  --rm \
  --name server \
  --network delbury-net \
  -it \
  data-analysis-system-server \
  /bin/sh
