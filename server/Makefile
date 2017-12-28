all: go docker
	rm ./server

.PHONY: go
go:
	CGO_ENABLED=0 go build

.PHONY: docker
docker:
	docker build -t yusuo/testing .
	docker build -t yusuo/testing -f Dockerfile.alpine .
