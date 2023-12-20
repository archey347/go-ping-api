build_commit=$(shell git rev-parse HEAD)
build_version=$(shell git describe --tags 2> /dev/null || echo "dev-$(shell git rev-parse HEAD)")

build: 
	go build -o bin/goping-api cmd/goping-api/main.go

run: build
	./bin/goping-api --config ./resources/configs/goping-api.yaml

clean:
	go clean
	rm -r ./bin