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
	mkdir -p package/etc/systemd/system
	cp -r resources/systemd/goping-api.service package/etc/systemd/system/goping-api.service
	dpkg-deb --build package goping-amd64-$(build_version).deb
	rm -rf package

install: goping-amd64-$(build_version).deb
	dpkg -i goping-amd64-$(build_version).deb
	
clean:
	go clean
	rm -r ./bin
	rm -r ./goping-amd64-$(build_version).deb