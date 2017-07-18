#!/usr/bin/env bash


service_1="src/service.1"
service_2="src/service.2"
service_3="src/service.3"

kubectl delete service service-1 service-2 service-3
kubectl delete deployment service-1 service-2 service-3

kubectl create -f $service_1/deployment.yml
kubectl create -f $service_1/service.yml

kubectl create -f $service_2/deployment.yml
kubectl create -f $service_2/service.yml

kubectl create -f $service_3/deployment.yml
kubectl create -f $service_3/service.yml

# kubectl set image deployment/service-1 service-1=example:1.1