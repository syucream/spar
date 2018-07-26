all: build

setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y

build: spanner.go
	go build cmd/spck/spck.go

check: build
	cat examples/create_database.sql | ./spck
	cat examples/create_table.sql | ./spck
	cat examples/create_index.sql | ./spck
	cat examples/alter_table.sql | ./spck
	cat examples/drop_table.sql | ./spck
	cat examples/drop_index.sql | ./spck
	cat examples/composition.sql | ./spck
