#!/bin/sh

set -eu

echo "first run..."

echo "$1"

pdflatex "$1"

echo "second run..."

pdflatex "$1"

echo "finished"
