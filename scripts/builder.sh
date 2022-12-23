#!/usr/bin/env bash

# This is a shell library for building targets for different environments.
# It should be `source`d into other shell scripts
build_target() {
    bazelisk build \
      --stamp \
      --workspace_status_command=./scripts/workspace_status.sh \
      "$1"
}

_build_target_for_platform() {
    bazelisk build \
      --stamp \
      --disk_cache="$3/.cache" \
      --repository_cache="$3/.repository_cache" \
      --workspace_status_command=./scripts/workspace_status.sh \
      --platforms="$2" \
      "$1"
}

# Build the given target ($1) for windows x64
_build_windows_target() {
    _build_target_for_platform $1 @io_bazel_rules_go//go/toolchain:windows_amd64 $2
}

# Build the given target ($1) for linux x64
_build_linux_target() {
    _build_target_for_platform $1 @io_bazel_rules_go//go/toolchain:linux_amd64 $2
}

# Build the given target ($1) for mac x64
_build_mac_target() {
    _build_target_for_platform $1 @io_bazel_rules_go//go/toolchain:darwin_amd64 $2
}

# Build the given target ($1) for mac x64 m1
_build_mac_m1_target() {
    _build_target_for_platform $1 @io_bazel_rules_go//go/toolchain:darwin_arm64 $2
}

_copy_target_executable() {
  # Support copying files into subdirectories
  mkdir -p "$(dirname "$2")"

  cp $(bazelisk cquery \
          --platforms="$3" \
          "$1" \
          --output=starlark \
          --starlark:file=bazel/show_all_outputs.bzl 2>/dev/null) \
        "$2"
}

# Copy the given windows target ($1) to a well known location ($2).
# This also renames the file to the given final name ($2)
_copy_windows_exe() {
    _copy_target_executable $1 $2 @io_bazel_rules_go//go/toolchain:windows_amd64 $3
}

# Copy the given linux target ($1) to a well known location ($2).
# This also renames the file to the given final name ($2)
_copy_linux_exe() {
    _copy_target_executable $1 $2 @io_bazel_rules_go//go/toolchain:linux_amd64 $3
}

# Copy the given macOS target ($1) to a well known location ($2).
# This also renames the file to the given final name ($2)
_copy_mac_exe() {
    _copy_target_executable $1 $2 @io_bazel_rules_go//go/toolchain:darwin_amd64 $3
}

# Copy the given macOS m1 target ($1) to a well known location ($2).
# This also renames the file to the given final name ($2)
_copy_mac_m1_exe() {
    _copy_target_executable $1 $2 @io_bazel_rules_go//go/toolchain:darwin_arm64 $3
}

build_windows_target() {
  _build_windows_target "$1" "$3"
  _copy_windows_exe "$1" "$2" "$3"
}

build_linux_target() {
  _build_linux_target "$1" "$3"
  _copy_linux_exe "$1" "$2" "$3"
}

build_mac_target() {
  _build_mac_target "$1" "$3"
  _copy_mac_exe "$1" "$2" "$3"
}

build_mac_m1_target() {
  _build_mac_m1_target "$1" "$3"
  _copy_mac_m1_exe "$1" "$2" "$3"
}
