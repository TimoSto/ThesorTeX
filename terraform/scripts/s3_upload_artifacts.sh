
VERSION=$(../scripts/env.sh APP ../VERSIONS)

VERSIONPATH=v$VERSION

echo $VERSIONPATH

mkdir ../artifacts/$VERSIONPATH-zip

zip -r -D ../artifacts/$VERSIONPATH-zip/linux.zip ../artifacts/$VERSIONPATH/linux

zip -r -D ../artifacts/$VERSIONPATH-zip/windows.zip ../artifacts/$VERSIONPATH/windows

zip -r -D ../artifacts/$VERSIONPATH-zip/mac.zip ../artifacts/$VERSIONPATH/mac

zip -r -D ../artifacts/$VERSIONPATH-zip/mac_silicon.zip ../artifacts/$VERSIONPATH/mac_silicon

aws s3 cp ../artifacts/$VERSIONPATH-zip/ s3://thesortex-artifacts/$VERSIONPATH/ --recursive
