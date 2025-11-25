#!/bin/bash


rm -rf hosting/*
kubectl get deployment -n user-12345 | grep -E "^(backend|frontend)" | awk '{print $1}' | xargs kubectl delete deployment -n user-12345
kubectl get services -n user-12345 | grep -E "^(backend|frontend)" | awk '{print $1}' | xargs kubectl delete service -n user-12345
kubectl delete ingress simple-ingress -n user-12345 --ignore-not-found
go build -o ./bin/main ./cmd 
./bin/main
