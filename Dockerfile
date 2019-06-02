FROM golang:1.12.5-alpine3.9 as build-stage

ENV GO111MODULE off
RUN apk --no-cache add git && go get github.com/greymd/ojichat
ENV GO111MODULE on

WORKDIR /go/src/github.com/greymd/ojichat
RUN GOOS=linux GOARCH=amd64 go build -o ./bin/ojichat

FROM alpine:latest as exec-stage
COPY --from=build-stage /go/src/github.com/greymd/ojichat/bin/ojichat /usr/local/bin/

ENTRYPOINT ["/usr/local/bin/ojichat"]
