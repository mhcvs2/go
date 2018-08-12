#!/usr/bin/env bash

i=0
while [ $i -lt 10 ]; do
    echo "hello world-${i}"
    sleep 1
    let i++
done