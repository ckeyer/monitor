#!/bin/bash 

TMP_DIR="/tmp/monitor-image"

set -ex

if [ -d /tmp/monitor-image ]; then
	rm -rf $TMP_DIR
	mkdir $TMP_DIR
fi

cp -a . $TMP_DIR
cd $TMP_DIR

wget -O prometheus-1.1.3.linux-amd64.tar.gz https://github.com/prometheus/prometheus/releases/download/v1.1.3/prometheus-1.1.3.linux-amd64.tar.gz

docker build -t pro .

rm -rf $TMP_DIR