
VERSION=$(scripts/env.sh APP VERSIONS)

VERSIONPATH=v$VERSION

echo $VERSIONPATH

aws s3 cp artifacts/zip/tool/ s3://thesortex-artifacts/tool/$VERSIONPATH/ --recursive --cache-control max-age=86400

aws s3 cp artifacts/zip/tool/ s3://thesortex-artifacts/tool/latest/ --recursive --cache-control max-age=3600
