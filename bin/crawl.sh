#! /bin/bash

cd "$(dirname "$0")/.." || exit

docker run \
  --net="host" \
  -e MYSQL_USER \
  -e MYSQL_PASSWORD \
  -e MYSQL_HOST \
  -e MYSQL_PORT \
  -e MYSQL_DATABASE \
  crawler

