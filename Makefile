.PHONY: lxd-registery
lxd-registery:
	go build

check:
	go fmt ./...
	go vet ./...

clean:
	-rm -f lxd-registry

