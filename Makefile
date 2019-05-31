ojichat-build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/ojichat

docker-build:
	make ojichat-build
	docker build ./ --tag hatobus/dk-ojichat
