.PHONY: run build docker-build docker-push deploy clean

IMG ?= cluster-info-backend
TAG ?= latest

run:
	go run cmd/main.go

build:
	go build -o bin/cluster-info-backend cmd/main.go

docker-build:
	docker build -t $(IMG):$(TAG) .

docker-push: docker-build
	docker push $(IMG):$(TAG)

deps:
	go mod tidy
	go mod download

test:
	go test ./...

clean:
	rm -rf bin/
	docker rmi $(IMG):$(TAG) || true