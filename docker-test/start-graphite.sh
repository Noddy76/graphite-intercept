#!/bin/bash

set -e -u

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

rm -r $DIR/data || true
mkdir $DIR/data

docker run --rm \
  --name graphite \
  -p 80:80 \
  -p 2003-2004:2003-2004 \
  -p 2023-2024:2023-2024 \
  -p 8125:8125/udp \
  -p 8126:8126 \
  -v $DIR/configs:/opt/graphite/conf \
  -v $DIR/data:/opt/graphite/storage \
  -v $DIR/statsd_config:/opt/statsd/config \
  graphiteapp/graphite-statsd:latest