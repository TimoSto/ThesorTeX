#!/usr/bin/env bash

source ./scripts/builder.sh

outDir="artifacts/v$(./scripts/env.sh APP VERSIONS)"

echo "cleaning out dir..."

rm -rf $outDir

echo "building frontend..."

cd frontend

rm -rf "assets/dist"

yarn run build

echo "building for windows..."

echo $(pwd)

echo "$(pwd)"

build_windows_target //services/app/cmd "$outDir/windows/ThesorTeX.exe" "$(pwd)"

echo "building app for linux..."

build_linux_target //services/app/cmd "$outDir/linux/ThesorTeX" "$(pwd)"

echo "building app for macOS..."

build_mac_target //services/app/cmd "$outDir/mac/ThesorTeX" "$(pwd)"

echo "building app for macOS (Apple Silicon M1)..."

build_mac_m1_target //services/app/cmd "$outDir/mac_silicon/ThesorTeX" "$(pwd)"

echo "building website for linux..."

build_linux_target //services/website/cmd "artifacts/website/lambda" "$(pwd)"
