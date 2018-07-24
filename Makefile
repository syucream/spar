all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/jack/jack.go

check-cli: build
	./jack examples/create_database.sql
	./jack examples/create_table.sql
	./jack examples/create_index.sql
	./jack examples/composition.sql

# Set GITHUB_TOKEN personal access token and create release git tag
.PHONY: release
release:
	go get -u github.com/goreleaser/goreleaser
	goreleaser --rm-dist
