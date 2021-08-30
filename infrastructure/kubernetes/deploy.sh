#!/bin/bash
kubectl apply -f namespace.yaml
kubectl apply -f cm.yaml
kubectl apply -f secret.yaml
kubectl apply -f daemonset.yaml
kubectl apply -f svc.yaml
