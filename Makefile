.PHONY: run build clean test exec

run:
	go run cmd/go-lab/main.go

build:
	go build -o bin/go-lab cmd/go-lab/main.go

clean:
	rm -rf bin/

test:
	go test ./... -v

exec:
	./bin/go-lab
