#!/usr/bin/env bash

echo "running"

source ./scripts/builder.sh

echo "running2"

outDir="out/v$(./scripts/env.sh APP VERSIONS)"

echo "running3"

echo $outDir

build_windows_target //backend/app/cmd "$outDir/windows/ThesorTeX.exe"
