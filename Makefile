.PHONY: gen run get clean

pkg=github.com/ledyba/go-block-them-all

all: gen .bin/block-them-all

run: all
	.bin/block-them-all

get:
	go get -u "github.com/ChimeraCoder/anaconda"
	go get -u "github.com/Sirupsen/logrus"
	go get -u "github.com/fatih/color"

gen:
	go generate $(pkg)

.bin/block-them-all: $(shell find . -type f -name '*.go')
	gofmt -w .
	@mkdir -p .bin
	go build -o $@ $(pkg)

clean:
	rm -rf .bin
	go clean $(pkg)/...
