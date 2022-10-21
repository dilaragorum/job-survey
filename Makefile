.PHONY: unit-test
unit-test:
	go test -v ./...

.PHONY: build
build:
	go build -o job-survey

.PHONY: goreleaser-local
goreleaser-local:
	goreleaser release --snapshot --rm-dist

.PHONY: goreleaser-remote
releaser-remote:
	goreleaser release --rm-dist

