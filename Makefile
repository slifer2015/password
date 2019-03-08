export ROOT=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))
export BIN=$(ROOT)/bin
export GOBIN?=$(BIN)
export GO=$(shell which go)
export GOPATH=$(abspath $(ROOT)/../../..)
export BUILD=cd $(ROOT) && $(GO) install -v -ldflags "-s"
export CGO_ENABLED=0

all:
	$(BUILD) ./cmd/...

run:
	$(BIN)/server


export LINTER=$(BIN)/gometalinter
export LINTERCMD=$(LINTER) -e ".*.gen.go" -e ".*_test.go" -e "test.com/test/vendor/.*" --cyclo-over=19  --sort=path --disable-all --line-length=120 --deadline=100s --enable=structcheck --enable=deadcode --enable=gocyclo --enable=ineffassign --enable=golint --enable=goimports --enable=errcheck --enable=varcheck --enable=goconst --enable=megacheck --enable=misspell

lint: $(LINTER)
	$(LINTERCMD) $(ROOT)/cmd/...
	$(LINTERCMD) $(ROOT)/services/...
	$(LINTERCMD) $(ROOT)/modules/...

metalinter:
	$(GO) get -v github.com/alecthomas/gometalinter
	$(GO) install -v github.com/alecthomas/gometalinter
	$(LINTER) --install

$(LINTER):
	@[ -f $(LINTER) ] || make -f $(ROOT)/Makefile metalinter

convey:
	$(GO) get -v github.com/smartystreets/goconvey
	$(GO) install -v github.com/smartystreets/goconvey

test: convey
	cd $(ROOT)/modules && $(GO) test -v ./...

test-gui: convey
	cd $(ROOT)/services && goconvey -host=0.0.0.0