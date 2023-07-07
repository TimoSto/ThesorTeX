echo "if this command fails, set your credentials using aws configure"

aws ecr get-login-password --region eu-central-1 | docker login --username AWS --password-stdin  846873250811.dkr.ecr.eu-central-1.amazonaws.com