
terraform plan -out="plan.bin"

echo "proceed? [y/n]"

read proceed

if [ "$proceed" = "y" ]
then
  terraform apply "plan.bin"
fi
