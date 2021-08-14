APPLICATION      := node-port-scanner
LINUX            := build/${APPLICATION}-linux-amd64
DARWIN           := build/${APPLICATION}-darwin-amd64						
DOCKER_USER      ?= ""
DOCKER_PASS      ?= ""
BIN_DIR          := $(GOPATH)/bin
GOMETALINTER     := $(BIN_DIR)/gometalinter
COVER            := $(BIN_DIR)/gocov-xml
JUNITREPORT      := $(BIN_DIR)/go-junit-report
VERSION			 ?= latest


.PHONY: $(DARWIN)
$(DARWIN):
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ${DARWIN} *.go

.PHONY: linux
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${LINUX} *.go
	
.PHONY: release
release:
	./test.sh
	echo "${DOCKER_PASS}" | docker login -u "${DOCKER_USER}" --password-stdin
	docker build -t "${APPLICATION}" "." --no-cache
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker tag  "${APPLICATION}:latest" "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"
	docker push "${DOCKER_USER}/${APPLICATION}:${VERSION}"
	docker push "${DOCKER_USER}/${APPLICATION}:${GITHUB_SHA}"

test:
	./test.sh

infratest:
	docker compose up -d 
	cd aws/test
	go test