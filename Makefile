.PHONY: unit-test
unit-test:
	go test -v ./...

.PHONY: build
build:
	go build -o job-survey

.PHONY: goreleaser
goreleaser:
	goreleaser release --rm-dist

