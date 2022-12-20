#!/usr/bin/env bash
set -euo pipefail

# Get versions of applications
echo APP_VERSION "$(./scripts/env.sh APP VERSIONS)"
