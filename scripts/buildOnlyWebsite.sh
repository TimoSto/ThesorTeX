#!/usr/bin/env bash
set -e

rm -rf "artifacts"

source ./scripts/builder.sh

echo "building website for linux..."

build_linux_target //services/website/cmd/prod:lambda_zip "artifacts/website/lambda.zip" "$(pwd)"

echo "finished"
