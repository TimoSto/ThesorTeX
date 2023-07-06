#!/usr/bin/env bash
set -euo pipefail

# Get versions of applications
echo APP_VERSION "$(./scripts/env.sh THESIS_TOOL VERSIONS)"

# Get versions of applications
echo WEBSITE_VERSION "$(./scripts/env.sh WEBSITE VERSIONS)-$(git rev-parse HEAD)"


echo CONTACT_VERSION "$(./scripts/env.sh CONTACT VERSIONS)"
