all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/check/check.go

check: build
	cat examples/create_database.sql | ./check
	cat examples/create_table.sql | ./check
	cat examples/create_index.sql | ./check
	cat examples/alter_table.sql | ./check
	cat examples/drop_table.sql | ./check
	cat examples/drop_index.sql | ./check
	cat examples/composition.sql | ./check
