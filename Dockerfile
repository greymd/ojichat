FROM golang:1.11

WORKDIR /ojichat

COPY ./bin ./bin

ENTRYPOINT ["/ojichat/bin/ojichat"]