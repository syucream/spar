package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/syucream/jack/src/parser"
	"github.com/syucream/jack/src/spanner2mysql"
)

func main() {
	pathToSql := os.Args[1]
	data, err := ioutil.ReadFile(pathToSql)
	if err != nil {
		log.Fatal(err)
	}

	stmts := parser.Parse(strings.NewReader(string(data)))
	fmt.Println(stmts)

	mysqlStmts, err := spanner2mysql.GetMysqlCreateTables(stmts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlStmts)
}
