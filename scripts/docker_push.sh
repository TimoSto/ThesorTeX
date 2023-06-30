hash=$(git rev-parse --short HEAD)

if [ "$1" = "website" ]
then
  echo "pushing website lambda..."
  id=$(podman image inspect --format '{{ .Id }}' localhost/bazel/services/website/cmd/prod:website_lambda_image)
  podman push $id docker://docker.io/thesortex/lambda:$hash
fi
