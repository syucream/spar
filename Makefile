all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/jackup/jackup.go

check-cli: build
	./jackup examples/create_database.sql
	./jackup examples/create_table.sql
	./jackup examples/create_index.sql
	./jackup examples/composition.sql

# Set GITHUB_TOKEN personal access token and create release git tag
.PHONY: release
release:
	go get -u github.com/goreleaser/goreleaser
	goreleaser --rm-dist
