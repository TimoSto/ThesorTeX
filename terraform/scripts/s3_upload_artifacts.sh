
VERSION=$(../scripts/env.sh APP ../VERSIONS)

VERSIONPATH=v$VERSION

echo $VERSIONPATH

mkdir ../artifacts/$VERSIONPATH-zip

mkdir ../artifacts/$VERSIONPATH-zip/linux
mkdir ../artifacts/$VERSIONPATH-zip/windows
mkdir ../artifacts/$VERSIONPATH-zip/mac
mkdir ../artifacts/$VERSIONPATH-zip/mac_silicon

cd ../artifacts/$VERSIONPATH/linux
zip ThesorTeX.zip ThesorTeX

cd ../windows
zip ThesorTeX.zip ThesorTeX.exe

cd ../mac
zip ThesorTeX.zip ThesorTeX -x "*.DS_Store"

cd ../mac_silicon
zip ThesorTeX.zip ThesorTeX -x "*.DS_Store"

cd ../
cp linux/ThesorTeX.zip ../$VERSIONPATH-zip/linux/
cp windows/ThesorTeX.zip ../$VERSIONPATH-zip/windows/
cp mac/ThesorTeX.zip ../$VERSIONPATH-zip/mac/
cp mac_silicon/ThesorTeX.zip ../$VERSIONPATH-zip/mac_silicon/

cd ../

aws s3 cp $VERSIONPATH-zip/ s3://thesortex-artifacts/$VERSIONPATH/ --recursive --cache-control max-age=86400

aws s3 cp $VERSIONPATH-zip/ s3://thesortex-artifacts/latest/ --recursive --cache-control max-age=3600
