#!/bin/bash

echo "$1"

if [ "$1" = "thesisTool" ] || [ "$1" = "all" ]
then
  VERSION=$(scripts/env.sh THESIS_TOOL VERSIONS)

  echo "tagging commit for thesisTool release $VERSION"

  git tag "ThesisTool-v$VERSION"

  echo "uploading tool artifacts..."

  VERSIONPATH=v$VERSION

  echo $VERSIONPATH

  aws s3 cp artifacts/thesisTool/ s3://thesortex-artifacts/thesisTool/$VERSIONPATH/ --recursive --cache-control max-age=86400

  aws s3 cp artifacts/ReleaseNotes_ThesisTool.md s3://thesortex-artifacts/thesisTool/$VERSIONPATH/ReleaseNotes.md --cache-control max-age=86400

  aws s3 cp artifacts/thesisTool/ s3://thesortex-artifacts/thesisTool/latest/ --recursive --cache-control max-age=3600

fi

if [ "$1" = "thesisTemplate" ] || [ "$1" = "all" ]
then
  VERSION=$(scripts/env.sh THESIS_TEMPLATE VERSIONS)

  echo "tagging commit for thesisTemplate release $VERSION"

  git tag "ThesisTemplate-v$VERSION"

  echo "uploading thesis template..."

  VERSIONPATH=v$VERSION

  echo $VERSIONPATH

  aws s3 cp artifacts/ThesisTemplate.zip s3://thesortex-artifacts/thesisTemplate/$VERSIONPATH/ --cache-control max-age=86400

  aws s3 cp artifacts/ReleaseNotes_ThesisTemplate.md s3://thesortex-artifacts/thesisTemplate/$VERSIONPATH/ReleaseNotes.md --cache-control max-age=86400

  aws s3 cp artifacts/ThesisTemplate.zip s3://thesortex-artifacts/thesisTemplate/latest/ --cache-control max-age=3600

fi

if [ "$1" = "cvTemplate" ] || [ "$1" = "all" ]
then
  VERSION=$(scripts/env.sh CV_TEMPLATE VERSIONS)

  echo "tagging commit for cvTemplate release $VERSION"

  git tag "CVTemplate-v$VERSION"

  echo "uploading cv template..."

  VERSIONPATH=v$VERSION

  echo $VERSIONPATH

  aws s3 cp artifacts/CVTemplate.zip s3://thesortex-artifacts/cvTemplate/$VERSIONPATH/ --cache-control max-age=86400

  aws s3 cp artifacts/ReleaseNotes_CVTemplate.md s3://thesortex-artifacts/cvTemplate/$VERSIONPATH/ReleaseNotes.md --cache-control max-age=86400

  aws s3 cp artifacts/CVTemplate.zip s3://thesortex-artifacts/cvTemplate/latest/ --cache-control max-age=3600

fi

echo "finished"
