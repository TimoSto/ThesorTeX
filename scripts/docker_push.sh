hash=$(git rev-parse --short HEAD)

aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin 846873250811.dkr.ecr.eu-central-1.amazonaws.com

if [ "$1" = "website" ]
then
  echo "pushing website lambda..."
  id=$(podman image inspect --format '{{ .Id }}' localhost/bazel/services/website/cmd/prod:website_lambda_image)
  hash=$(git rev-parse HEAD)
  podman tag $id 846873250811.dkr.ecr.eu-central-1.amazonaws.com/website_lambda:$hash
  podman push 846873250811.dkr.ecr.eu-central-1.amazonaws.com/website_lambda:$hash
fi
