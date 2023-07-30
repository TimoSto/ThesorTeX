#!/usr/bin/env bash
set -euo pipefail

# Get version of thesis tool
echo THESIS_TOOL_VERSION "$(./scripts/env.sh THESIS_TOOL VERSIONS)"

# Get version of website
echo WEBSITE_VERSION "$(./scripts/env.sh WEBSITE VERSIONS)-$(git rev-parse HEAD)"

# Get version of contact service
echo CONTACT_VERSION "$(./scripts/env.sh CONTACT VERSIONS)"
