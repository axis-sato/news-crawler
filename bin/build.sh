#! /bin/bash

cd "$(dirname "$0")/.." || exit

docker build -t crawler \
  -f docker/go/Dockerfile .
