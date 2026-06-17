.PHONY: all build clean run doctor start help docker-up docker-down docker-ps docker-logs docker-build

all: build

build:
	mkdir -p bin
	go build -o bin/nexus .

clean:
	rm -rf bin/

run:
	go run main.go $(ARGS)

doctor:
	go run main.go doctor

start:
	go run main.go start $(ARGS)

help:
	go run main.go help

# Docker Compose targets
 docker-up:
	docker compose -f docker-compose.yml up -d --build --remove-orphans

docker-down:
	docker compose -f docker-compose.yml down --remove-orphans -v

docker-ps:
	docker compose -f docker-compose.yml ps

docker-logs:
	docker compose -f docker-compose.yml logs -f

docker-build:
	docker build -t go-nexus:latest .