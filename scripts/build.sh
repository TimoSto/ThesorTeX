#!/usr/bin/env bash
set -e

version="$(./scripts/env.sh APP VERSIONS)"

echo "building version $version"

source ./scripts/builder.sh

outDir="artifacts/tool"

echo "cleaning out dir..."

rm -rf "artifacts"

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

build_linux_target //services/website/cmd:lambda_zip "artifacts/website/lambda.zip" "$(pwd)"

echo "creating zips of tool..."

cd artifacts/tool/linux
zip ThesorTeX.zip ThesorTeX

cd ../windows
zip ThesorTeX.zip ThesorTeX.exe

cd ../mac
zip ThesorTeX.zip ThesorTeX -x "*.DS_Store"

cd ../mac_silicon
zip ThesorTeX.zip ThesorTeX -x "*.DS_Store"

cd ../../

echo "copying zips of tool..."

mkdir "zip"
mkdir "zip/tool"
mkdir "zip/tool/windows"
mkdir "zip/tool/linux"
mkdir "zip/tool/mac"
mkdir "zip/tool/mac_silicon"

cp tool/linux/ThesorTeX.zip zip/tool/linux/
cp tool/windows/ThesorTeX.zip zip/tool/windows/
cp tool/mac/ThesorTeX.zip zip/tool/mac/
cp tool/mac_silicon/ThesorTeX.zip zip/tool/mac_silicon/
