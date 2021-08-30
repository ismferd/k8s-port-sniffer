APPLICATION      	  := node-port-scanner
LINUX            	  := build/${APPLICATION}-linux-amd64
DARWIN           	  := build/${APPLICATION}-darwin-amd64						
DOCKER_USER      	  ?= ""
DOCKER_PASS      	  ?= ""
VERSION			 	  ?= latest
GITHUB_SHA			  ?= 0.1.0
TF_VAR_aws_access_key ?= ""
TF_VAR_aws_secret_key ?= ""
AWS_ACCESS_KEY_ID	  ?= "" 
AWS_SECRET_ACCESS_KEY ?= ""
AWS_DEFAULT_REGION	  ?= us-east-1
KUBECONFIG			  ?= ~/.kube/config
S3_LOCALSTACK_END	  := --endpoint=http://localhost:4566
AWS_CREDENTIALS_TEST  := test

.PHONY: $(DARWIN)
$(DARWIN):
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${DARWIN} *.go

.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${LINUX} *.go
	
.PHONY: release
release: unit_tests
	echo "${DOCKER_PASS}" | docker login -u "${DOCKER_USER}" --password-stdin
	docker build -t "${APPLICATION}" "." --no-cache
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"
	docker push "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker push "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"

.PHONY: unit_tests
unit_tests:
	COMPOSE_HTTP_TIMEOUT=120 docker-compose up -d localstack
	AWS_ACCESS_KEY_ID=${AWS_CREDENTIALS_TEST} AWS_SECRET_ACCESS_KEY=${AWS_CREDENTIALS_TEST} aws ${S3_LOCALSTACK_END} s3 mb s3://test
	cd ./src; AWS_ACCESS_KEY_ID=${AWS_CREDENTIALS_TEST} AWS_SECRET_ACCESS_KEY=${AWS_CREDENTIALS_TEST} go test ./...
	AWS_ACCESS_KEY_ID=${AWS_CREDENTIALS_TEST} AWS_SECRET_ACCESS_KEY=${AWS_CREDENTIALS_TEST} aws ${S3_LOCALSTACK_END} s3 rb s3://test --force
	COMPOSE_HTTP_TIMEOUT=120 docker-compose down

.PHONY: deploy_aws terratest 
deploy_aws: terratest terraform_clean_up
	cd infrastructure/aws/; TF_VAR_aws_access_key=${TF_VAR_aws_access_key} TF_VAR_aws_secret_key=${TF_VAR_aws_secret_key} terraform init
	cd infrastructure/aws/; TF_VAR_aws_access_key=${TF_VAR_aws_access_key} TF_VAR_aws_secret_key=${TF_VAR_aws_secret_key} terraform apply
	
.PHONY: terratest
terratest: terraform_clean_up
	COMPOSE_HTTP_TIMEOUT=120 docker-compose up -d 
	cd infrastructure/aws/test/;TF_VAR_skip_requesting_account_id=true TF_VAR_skip_metadata_api_check=true TF_VAR_skip_credentials_validation=true TF_VAR_s3_endpoint=${S3_LOCALSTACK_END} TF_VAR_aws_access_key=${AWS_CREDENTIALS_TEST} TF_VAR_aws_secret_key=${AWS_CREDENTIALS_TEST} go test
	COMPOSE_HTTP_TIMEOUT=120 docker-compose down

.PHONY: terraform_clean_up
terraform_clean_up:
	@if [ -d infrastructure/aws/.terraform ]; then\
		cd infrastructure/aws/;rm -rf .terraform;\
	fi

.PHONY: deploy_kubernetes
deploy_kubernetes:
	cd infrastructure/kubernetes; AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} AWS_DEFAULT_REGION=${AWS_DEFAULT_REGION} KUBECONFIG=${KUBECONFIG} ./deploy.sh
