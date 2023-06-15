#!/usr/bin/env bash

build_target() {
  platform=""
  case "$3" in
     windows)   platform=@io_bazel_rules_go//go/toolchain:windows_amd64;;
     linux)   platform=@io_bazel_rules_go//go/toolchain:linux_amd64;;
     mac)   platform=@io_bazel_rules_go//go/toolchain:darwin_amd64;;
     mac_arm)   platform=@io_bazel_rules_go//go/toolchain:darwin_arm64;;
  esac

  if [ "$platform" = "" ]; then
      echo "building without platform set..."
      bazelisk build \
          --stamp \
          --workspace_status_command=./scripts/workspace_status.sh \
          "$1"
  else
      echo "building for platform $platform..."

      bazelisk build \
          --stamp \
          --workspace_status_command=./scripts/workspace_status.sh \
          --platforms="$platform" \
          --enable_runfiles \
          "$1"
  fi

  # 2>/dev/null removes bazel logs
  files=$(bazel cquery --output=files "$1" 2>/dev/null)

  echo "copying output ($files) to dest $2)..."

  mkdir -p "$(dirname "$2")"

  cp $files $2
}
