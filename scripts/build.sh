#!/usr/bin/env bash
set -e

echo "running all tests..."

bazel test //...

echo "cleaning out dir..."

rm -rf "artifacts"

source ./scripts/builder.sh

if [ "$1" = "thesisTemplate" ] || [ "$1" = "all" ]
then
  echo "building thesis template"

  build_zip_target //:release_notes_template "artifacts/ReleaseNotes_ThesisTemplate.md"

  build_zip_target //pkg/backend/project_template:template_zip artifacts/ThesisTemplate.zip
fi

if [ "$1" = "cvTemplate" ] || [ "$1" = "all" ]
then
  echo "building cv template"

  build_zip_target //:release_notes_cv_template "artifacts/ReleaseNotes_CVTemplate.md"

  build_zip_target //pkg/backend/cv_template:template_zip artifacts/CVTemplate.zip
fi

if [ "$1" = "thesisTool" ] || [ "$1" = "all" ]
then
  version="$(./scripts/env.sh THESIS_TOOL VERSIONS)"

  echo "building tool version $version"

  outDir="artifacts/thesisTool"

  build_zip_target //:release_notes_app "artifacts/ReleaseNotes_ThesisTool.md"

  echo "building for windows..."

  build_windows_target //services/app/cmd/prod:app_zip "$outDir/windows/ThesorTeX.zip" "$(pwd)"

  echo "building app for linux..."

  build_linux_target //services/app/cmd/prod:app_zip "$outDir/linux/ThesorTeX.zip" "$(pwd)"

  echo "building app for macOS (AMD)..."

  build_mac_target //services/app/cmd/prod:app_zip "$outDir/mac/ThesorTeX.zip" "$(pwd)"

  echo "building app for macOS (ARM)..."

  build_mac_m1_target //services/app/cmd/prod:app_zip "$outDir/mac_arm/ThesorTeX.zip" "$(pwd)"
fi

if [ "$1" = "website" ] || [ "$1" = "all" ]
then
  echo "building website for linux..."

  build_linux_target //services/website/cmd/prod:lambda_zip "artifacts/website/lambda.zip" "$(pwd)"
fi

echo "finished"
