echo "building e2e executable..."

cd services/app/cmd/e2e

go build main.go

cd ../../../../e2e_tests

pnpm run test
