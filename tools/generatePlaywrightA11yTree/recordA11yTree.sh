#!/bin/bash

# Step 1: generate playwright test to go to the page and click/... as needed for the test
cd ../../tests/playwright

pnpm run codegen

cd ../../tools/generatePlaywrightA11yTree

# Step 2: at the end of the test, generate the a11y tree and write it to a file
go run cmd/getTree/main.go

# Step 3: execute the test
cd ../../tests/playwright
 pnpm run test:generated