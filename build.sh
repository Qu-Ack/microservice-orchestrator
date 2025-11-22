#!/bin/bash

set -e

kubectl delete ingress simple-ingress
go build -o ./bin/main ./cmd 
./bin/main
