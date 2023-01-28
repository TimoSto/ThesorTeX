#!/usr/bin/env bash
set -e

version="$(./scripts/env.sh APP VERSIONS)"

echo "building version $version"

source ./scripts/builder.sh

outDir="artifacts/v$version"

echo "cleaning out dir..."

rm -rf $outDir

rm -rf "artifacts/website"

echo "cleaning frontend-dist-dirs..."

rm -rf services/app/frontend/assets/dist

rm -rf services/website/frontend/assets/dist

echo "building frontends..."

pnpm install

pnpm run -r build

echo "building for windows..."

build_windows_target //services/app/cmd/prod "$outDir/windows/ThesorTeX.exe" "$(pwd)"

echo "building app for linux..."

build_linux_target //services/app/cmd/prod "$outDir/linux/ThesorTeX" "$(pwd)"

echo "building app for macOS..."

build_mac_target //services/app/cmd/prod "$outDir/mac/ThesorTeX" "$(pwd)"

echo "building app for macOS (Apple Silicon M1)..."

build_mac_m1_target //services/app/cmd/prod "$outDir/mac_silicon/ThesorTeX" "$(pwd)"

echo "building website for linux..."

build_linux_target //services/website/cmd "artifacts/website/lambda" "$(pwd)"
