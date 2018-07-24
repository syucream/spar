all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/check/check.go

check: build
	./ckeck examples/create_database.sql
	./ckeck examples/create_table.sql
	./ckeck examples/create_index.sql
	./ckeck examples/composition.sql
