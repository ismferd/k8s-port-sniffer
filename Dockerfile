
FROM golang:alpine
ADD src /build/src/
ENV GOPATH=/build/go
ENV GOBIN=$GOPATH/bin
WORKDIR /build/
RUN env
RUN \
    apk add --no-cache \
        ca-certificates \
        gcc \
        git \
        musl-dev 
RUN go mod init github.com/ismferd/k8s-port-sniffer
RUN go clean -modcache
RUN pwd
RUN cd src && go get
RUN cd src && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o node-port-scanner .

ENTRYPOINT ["./src/node-port-scanner"]

