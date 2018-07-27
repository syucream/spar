# spar

Cloud Spanner DDL parser in Go

## Usage

- Pass your Spanner DDL to `parser.Parse()` and get parsed statements.
- See detail on cmd/spck/spck.go

## subtool

`spck` is a Spanner DDL syntax checker. It parses and returns exit code 0 if no error.

```
$ go get -u github.com/syucream/spar/cmd/spck
$ spck ./examples/create_table.sql
```
