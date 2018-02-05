all: go docker
	rm ./server

.PHONY: go
go:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

.PHONY: docker
docker:
	docker build -t crystal639/testing .
