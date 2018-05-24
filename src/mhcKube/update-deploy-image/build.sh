#!/usr/bin/env bash

docker build --build-arg http_proxy=http://109.105.4.17:3128 --build-arg https_proxy=http://109.105.4.17:8119 \
-t build-update-deploy-image .