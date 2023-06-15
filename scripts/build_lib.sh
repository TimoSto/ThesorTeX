#!/usr/bin/env bash

build_plain_target() {
  bazelsik build \
      --workspace_status_command=./scripts/workspace_status.sh \
      --staml \
      "$1"
}

build_target() {
  platform=""
  case "$3" in
     windows)   platform=@io_bazel_rules_go//go/toolchain:windows_amd64;;
     linux)   platform=@io_bazel_rules_go//go/toolchain:linux_amd64;;
     mac)   platform=@io_bazel_rules_go//go/toolchain:darwin_amd64;;
     mac_arm)   platform=@io_bazel_rules_go//go/toolchain:darwin_arm64;;
  esac

  bazelisk build \
        --stamp \
        --workspace_status_command=./scripts/workspace_status.sh \
        --platforms="$platform" \
        --enable_runfiles \
        "$1"

  # 2>/dev/null removes bazel logs
  files=$(bazel cquery --output=files //services/thesisTool/cmd/prod:ThesorTeX 2>/dev/null)

  echo "copying output to dest..."

  mkdir -p "$2"

  cp $files $2
}
