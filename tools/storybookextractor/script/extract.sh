#!/bin/bash
set -e

if ! command -v rsvg-convert &> /dev/null
then
    echo "rsvg-convert could not be found"
    exit 1
fi

echo "generating svgs from storybook..."

export STORY_NAME="$1"
export STORY_DIST="$2/${STORY_NAME//-/_}.svg"

pnpm playwright test

echo "generating pngs from svgs..."

cd $2

for i in *.svg; do
    [ -f "$i" ] || break
    echo "converting $i..."
    rsvg-convert "$i" > "${i%.*}.png"
done
