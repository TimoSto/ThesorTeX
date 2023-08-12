#!/usr/bin/env bash
set -e

echo "running all tests..."

bazel test //...

echo "cleaning out dir..."

rm -rf "artifacts"

source ./scripts/build_lib.sh

if [ "$1" = "thesisTemplate" ] || [ "$1" = "all" ]
then
  echo "building thesis template"

  build_target //:release_notes_template "artifacts/ReleaseNotes_ThesisTemplate.md"

  build_target //pkg/tex/project_template:template_zip artifacts/ThesisTemplate.zip
fi

if [ "$1" = "cvTemplate" ] || [ "$1" = "all" ]
then
  echo "building cv template"

  build_target //:release_notes_cv_template "artifacts/ReleaseNotes_CVTemplate.md"

  build_target //pkg/tex/cv_template:template_zip artifacts/CVTemplate.zip
fi

if [ "$1" = "thesisTool" ] || [ "$1" = "all" ]
then
  version="$(./scripts/env.sh THESIS_TOOL VERSIONS)"

  echo "building tool version $version"

  outDir="artifacts/thesisTool"

  build_target //:release_notes_app "artifacts/ReleaseNotes_ThesisTool.md"

  echo "building for windows..."

  build_target //services/thesisTool/cmd/prod:thesisTool_zip "$outDir/windows/ThesorTeX.zip" "windows"

  echo "building app for linux..."

  build_target //services/thesisTool/cmd/prod:thesisTool_zip "$outDir/linux/ThesorTeX.zip" "linux"

  echo "building app for macOS (AMD)..."

  build_target //services/thesisTool/cmd/prod:thesisTool_zip "$outDir/mac/ThesorTeX.zip" "mac"

  echo "building app for macOS (ARM)..."

  build_target //services/thesisTool/cmd/prod:thesisTool_zip "$outDir/mac_arm/ThesorTeX.zip" "mac_arm"
fi

if [ "$1" = "website" ] || [ "$1" = "all" ]
then
  echo "building website docker image..."

  bazel run //services/website/cmd/prod:website_lambda_image

  echo "push to aws ecr? [y/n]"

  read proceed

  if [ "$proceed" = "y" ]
  then
    id=$(podman image inspect --format '{{ .Id }}' localhost/bazel/services/website/cmd/prod:website_lambda_image)
    hash=$(git rev-parse --short HEAD)

    echo "pushing website lambda image (tag $hash)..."

    podman tag $id 846873250811.dkr.ecr.eu-central-1.amazonaws.com/website_lambda:$hash
    podman push 846873250811.dkr.ecr.eu-central-1.amazonaws.com/website_lambda:$hash
  fi
fi

if [ "$1" = "contact" ] || [ "$1" = "all" ]
then
  echo "building contact docker image..."

  bazel run //services/contact/cmd/lambda:contact_lambda_image

  echo "push to aws ecr? [y/n]"

  read proceed

  if [ "$proceed" = "y" ]
  then
    id=$(podman image inspect --format '{{ .Id }}' localhost/bazel/services/contact/cmd/lambda:contact_lambda_image)
    hash=$(git rev-parse --short HEAD)

    echo "pushing contact lambda image (tag $hash)..."

    podman tag $id 846873250811.dkr.ecr.eu-central-1.amazonaws.com/contact_lambda:$hash
    podman push 846873250811.dkr.ecr.eu-central-1.amazonaws.com/contact_lambda:$hash
  fi
fi

if [ "$1" = "router" ] || [ "$1" = "all" ]
then
  echo "building router docker image..."

  bazel run //services/router/cmd/lambda:router_lambda_image

  echo "push to aws ecr? [y/n]"

  read proceed

  if [ "$proceed" = "y" ]
  then
    id=$(podman image inspect --format '{{ .Id }}' localhost/bazel/services/router/cmd/lambda:router_lambda_image)
    hash=$(git rev-parse --short HEAD)

    echo "pushing router lambda image (tag $hash)..."

    podman tag $id 846873250811.dkr.ecr.eu-central-1.amazonaws.com/router_lambda:$hash
    podman push 846873250811.dkr.ecr.eu-central-1.amazonaws.com/router_lambda:$hash
  fi
fi

echo "finished"
