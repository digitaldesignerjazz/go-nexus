.PHONY: all build clean run doctor start help

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