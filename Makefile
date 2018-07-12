setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc -o src/parser/spanner.go src/parser/spanner.go.y
