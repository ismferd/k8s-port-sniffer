#!/bin/bash
kubectl apply -f namespace.yaml
kubectl apply -f cm.yaml
kubectl apply -f daemonset.yaml
