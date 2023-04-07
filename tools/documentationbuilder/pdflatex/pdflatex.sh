#!/bin/sh

echo "pwd"
echo $PWD

cd $2

echo "pwd"
echo $PWD

set -eu

echo "first run..."

echo "srcs"
echo "$1"

pdflatex "$1"

echo "second run..."

pdflatex "$1"

ls

# cp main.pdf "$3"

echo "finished"
