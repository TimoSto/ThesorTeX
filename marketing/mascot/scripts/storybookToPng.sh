#!/bin/bash
set -e

cur=$(pwd)

cd ../../tools/storybookextractor

./script/extract.sh mascot--left "${cur}"

./script/extract.sh mascot--right "${cur}"
