FROM golang:1.11
WORKDIR /go/src/app
COPY bin/server .


CMD ["server"]