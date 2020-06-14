.DEFAULT_GOAL := help

VERSION=$(shell git describe --always --long)
PROJECT_NAME := fiber-pongo2-pkger
CLONE_URL:=github.com/ansrivas/fiber-pongo2-pkger
IDENTIFIER= $(VERSION)-$(GOOS)-$(GOARCH)
BUILD_TIME=$(shell date -u +%FT%T%z)
LDFLAGS='-extldflags "-static" -s -w -X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)'

help:          ## Show available options with this Makefile
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -v grep | awk 'BEGIN { FS = ":.*?##" }; { printf "%-18s  %s\n", $$1,$$2 }'

.PHONY : test crossbuild release build clean
test:          ## Run all the tests
	chmod +x ./scripts/test.sh && ./scripts/test.sh

clean:         ## Clean the application
	@go clean -i ./...
	@rm -rf ./$(PROJECT_NAME)
	@rm -rf build/*

# -v so warnings from the linker aren't suppressed.
# -a so dependencies are rebuilt (they may have been dynamically
# linked).
build: vendor
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o '$(FLAGS)' -a -ldflags $(LDFLAGS)  .

vendor:           ## Go get vendor
	go mod vendor

crossbuild:
	mkdir -p build/${PROJECT_NAME}-$(IDENTIFIER)
	make build FLAGS="build/${PROJECT_NAME}-$(IDENTIFIER)/${PROJECT_NAME}"
	cd build \
	&& tar cvzf "${PROJECT_NAME}-$(IDENTIFIER).tgz" "${PROJECT_NAME}-$(IDENTIFIER)" \
	&& rm -rf "${PROJECT_NAME}-$(IDENTIFIER)"

release:	vendor clean  ## Create a release build.
	make crossbuild GOOS=linux GOARCH=amd64
	make crossbuild GOOS=linux GOARCH=386
	make crossbuild GOOS=darwin GOARCH=amd64
	make crossbuild GOOS=windows GOARCH=amd64

bench:	       ## Benchmark the code.
	@go test -o bench.test -cpuprofile cpu.prof -memprofile mem.prof -bench .

prof:	bench  ## Run the profiler.
	@go tool pprof cpu.prof

prof_svg:	clean	bench ## Run the profiler and generate image.
	@echo "Do you have graphviz installed? sudo apt-get install graphviz."
	@go tool pprof -svg bench.test cpu.prof > cpu.svg
