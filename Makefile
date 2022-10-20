.PHONY: unit-test
unit-test:
	go test -v ./... -coverprofile=unit_coverage.out

.PHONY: build
build:
	go build -o job-survey

.PHONY: goreleaser-local
releaser-local:
	goreleaser release --snapshot --rm-dist

.PHONY: goreleaser-remote
releaser-remote:
	goreleaser release --rm-dist

