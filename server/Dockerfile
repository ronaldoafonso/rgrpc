FROM golang:alpine3.11

EXPOSE 50051/tcp

RUN apk add --update git protoc dep

RUN addgroup productinfo && \
    adduser -h /home/productinfo -s /bin/ash -G productinfo -D productinfo

ENV CGO_ENABLED 0

WORKDIR /go

RUN go get -v -u google.golang.org/grpc && \
    go get -v -u github.com/golang/protobuf/protoc-gen-go && \
    go get golang.org/x/lint/golint

RUN mkdir -p /go/src/github.com && \
    chown -R productinfo:productinfo /go

WORKDIR /go/src/github.com/ronaldoafonso/productinfo

USER productinfo:productinfo

COPY --chown=productinfo:productinfo . /go/src/github.com/ronaldoafonso/productinfo/
