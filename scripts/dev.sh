#!/usr/bin/env bash
set -e

go build ./services/app/cmd/main.go

./main
