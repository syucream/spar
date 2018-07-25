all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/check/check.go

check: build
	./check examples/create_database.sql > /dev/null
	./check examples/create_table.sql > /dev/null
	./check examples/create_index.sql > /dev/null
	./check examples/alter_table.sql > /dev/null
	./check examples/drop_table.sql > /dev/null
	./check examples/drop_index.sql > /dev/null
	./check examples/composition.sql > /dev/null
