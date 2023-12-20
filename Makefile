build_commit=$(shell git rev-parse HEAD)
build_version=$(shell git describe --tags 2> /dev/null || echo "dev-$(shell git rev-parse HEAD)")

build: 
	go build -o bin/goping-api cmd/goping-api/main.go

run: build
	./bin/goping-api --config ./resources/configs/goping-api.yaml

package: build
	rm -rf package
	mkdir package
	mkdir package/DEBIAN
	cp -r resources/packaging/* package/DEBIAN/
	mkdir -p package/usr/local/bin
	cp bin/goping-api package/usr/local/bin 
	mkdir -p package/etc/goping
	cp -r resources/configs/goping-api.yaml package/etc/goping/goping-api.yaml
	dpkg-deb --build package goping-amd64-$(build_version).deb
	rm -rf package
clean:
	go clean
	rm -r ./bin