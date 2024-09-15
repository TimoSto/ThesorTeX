#!/bin/bash

# Step 1: generate playwright test to go to the page and click/... as needed for the test
cd ./packages/playwright

pnpm run codegen

cd ../../

# Step 2: at the end of the test, generate the a11y tree and write it to a file
go run cmd/generateGetTree/main.go

# Step 3: execute the test to save the a11y tree as json
cd ../../tests/playwright
pnpm run test:generated

# Step 4: generate the test to assure this a11y tree
go run cmd/generateAssertTree/main.go