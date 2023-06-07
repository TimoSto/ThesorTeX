#!/bin/bash

pnpm run storybook &

pnpm loki-local

kill $(jobs -p)

# TODO: do this with storybook build
