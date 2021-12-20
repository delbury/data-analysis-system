#!/bin/bash

set -e

# clear node_modules soft links
rm -rf node_modules

# 安装依赖
[ `yarn -v | wc -l` -eq 0 ] && {
npm i yarn -g
}

yarn config set registry "https://registry.npmmirror.com/"

# 安装npm包
yarn --force

yarn build
