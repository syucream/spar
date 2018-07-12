setup:
	go get -u golang.org/x/tools/cmd/goyacc

spanner.go:
	goyacc src/parser/spanner.go.y -o src/parser/spanner.go
