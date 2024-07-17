#!/usr/bin/env bash

# Prepare Packages
#go get -u github.com/jessevdk/go-assets-builder
#go install github.com/jessevdk/go-assets-builder

# Generate asserts.go:
go-assets-builder -p templates -o ./asserts.go ./*.htm ./**/*.htm
