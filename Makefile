export GO111MODULE=on

default: test

ci: depsdev test

lint:
	golangci-lint run ./...

depsdev:
	go install github.com/Songmu/gocredits/cmd/gocredits@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

prerelease_for_tagpr: depsdev
	gocredits . -w
	git add CHANGELOG.md CREDITS go.mod go.sum

release:
	git push origin main --tag

.PHONY: default test
