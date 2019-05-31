FROM alpine:latest

WORKDIR /ojichat

COPY ./bin ./bin

ENTRYPOINT ["/ojichat/bin/ojichat"]