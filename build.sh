#!/usr/bin/env bash

source ./scripts/builder.sh

outDir="out/v$(./scripts/env.sh APP VERSIONS)"

echo "cleaning out dir..."

rm -rf $outDir

echo "building for windows..."

echo $(pwd)

echo "$(pwd)"

build_windows_target //backend/app/cmd "$outDir/windows/ThesorTeX.exe" "$(pwd)"

echo "building for linux..."

build_linux_target //backend/app/cmd "$outDir/linux/ThesorTeX" "$(pwd)"

echo "building for macOS..."

build_mac_target //backend/app/cmd "$outDir/mac/ThesorTeX" "$(pwd)"

echo "building for macOS (Apple Silicon M1)..."

build_mac_m1_target //backend/app/cmd "$outDir/mac_silicon/ThesorTeX" "$(pwd)"
