#!/bin/bash


rm -rf hosting/*
kubectl get deployment | grep -E "^(backend|frontend)" | awk '{print $1}' | xargs kubectl delete deployment
kubectl get services | grep -E "^(backend|frontend)" | awk '{print $1}' | xargs kubectl delete service
kubectl delete ingress simple-ingress --ignore-not-found
go build -o ./bin/main ./cmd 
./bin/main
