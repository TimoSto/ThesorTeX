
terraform plan -out="plan.bin"

echo "proceed? [y/n]"

read proceed

if [ "$proceed" = "y" ]
then
  VERSION="$(../scripts/env.sh WEBSITE ../VERSIONS)-$(git rev-parse HEAD)"

  echo "tagging commit for website release $VERSION"

  git tag "Website-v$VERSION"

  terraform apply "plan.bin"
fi
