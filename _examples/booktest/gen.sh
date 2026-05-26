#!/bin/sh
set -u
set -e
set -x

rm -rf internal proto tools api go.mod go.sum main.go registry.go buf*

sqlc-connect -m booktest -tracing -metric
