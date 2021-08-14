#!/bin/bash 
#apt-get update
#apt-get install localstack
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
COMPOSE_HTTP_TIMEOUT=120 docker-compose up -d localstack
sleep 60
aws --endpoint=http://localhost:4566 s3 mb s3://test
cd ./src && go test ./...
COMPOSE_HTTP_TIMEOUT=120 docker-compose down