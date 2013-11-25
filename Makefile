SRC ?= $(wildcard *.go)

build: tmp/build

.PHONY: doc

test: tmp/test

tmp/build: $(SRC)
	go build
	touch $@

doc:
	@godoc -ex github.com/dhamidi/collection

tmp/test: $(SRC)
	go test > $@
