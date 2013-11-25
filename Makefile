SRC ?= $(wildcard *.go)

test: $(SRC)
	go test

build: $(SRC)
	go build

view-doc:
	@godoc -ex github.com/dhamidi/collection | less
