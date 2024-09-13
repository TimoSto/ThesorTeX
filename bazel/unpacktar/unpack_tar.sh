#!/bin/bash

echo "creating dist directory $1"

mkdir -p "$1"

echo "tar file: $2"
echo "strip components: $3"

echo "unpacking tar file..."

tar -C "$1" -zxf "$2" --strip-components=$3

echo "finished"
