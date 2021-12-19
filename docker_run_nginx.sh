#!/usr/bin/env bash

NGINX_DIR=${PWD}

docker run \
  --rm \
  --name nginx \
  -p 8080:80 \
  --network delbury-net \
  --volume "${NGINX_DIR}/nginx/html":/usr/share/nginx/html \
  --volume "${NGINX_DIR}/nginx/conf/nginx.conf":/etc/nginx/nginx.conf \
  --volume "${NGINX_DIR}/nginx/conf/conf.d":/etc/nginx/conf.d \
  nginx