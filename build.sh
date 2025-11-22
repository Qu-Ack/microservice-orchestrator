#!/bin/bash

set -e

kubectl delete ingress simple-ingress --ignore-not-found
go build -o ./bin/main ./cmd 
./bin/main
