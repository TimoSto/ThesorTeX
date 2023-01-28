echo "testing all bazel targets..."

bazel test //...

echo "testing vitests..."

pnpm run -r vitest:run

./scripts/e2e_fake_tests.sh
