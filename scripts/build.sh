#!/usr/bin/env bash
set -e

echo "cleaning out dir..."

rm -rf "artifacts"

source ./scripts/builder.sh

echo "building thesis template"

build_zip_target //pkg/backend/project_template:template_zip artifacts/ThesisTemplate.zip

echo "building cv template"

build_zip_target //pkg/backend/cv_template:template_zip artifacts/CVTemplate.zip

version="$(./scripts/env.sh APP VERSIONS)"

echo "building tool version $version"

outDir="artifacts/tool"

echo "cleaning frontend-dist-dirs..."

rm -rf services/app/frontend/assets/dist

rm -rf services/website/frontend/assets/dist

echo "building frontends..."

pnpm install

pnpm run -r build

echo "building for windows..."

build_windows_target //services/app/cmd/prod:app_zip "$outDir/windows/ThesorTeX.zip" "$(pwd)"

echo "building app for linux..."

build_linux_target //services/app/cmd/prod:app_zip "$outDir/linux/ThesorTeX.zip" "$(pwd)"

echo "building app for macOS (AMD)..."

build_mac_target //services/app/cmd/prod:app_zip "$outDir/mac/ThesorTeX.zip" "$(pwd)"

echo "building app for macOS (ARM)..."

build_mac_m1_target //services/app/cmd/prod:app_zip "$outDir/mac_arm/ThesorTeX.zip" "$(pwd)"

echo "building website for linux..."

build_linux_target //services/website/cmd:lambda_zip "artifacts/website/lambda.zip" "$(pwd)"

echo "finished"
