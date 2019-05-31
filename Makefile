ojichat-build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/ojichat

docker-build:
	make ojichat-build
	docker build ./ --tag hatobus/dk-ojichat
