#!/bin/sh

set -eu

echo "first run..."

pdflatex "$1"

echo "second run..."

pdflatex "$1"

echo "finished"
