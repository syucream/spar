package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/syucream/jackup/src/parser"
	"github.com/syucream/jackup/src/spanner2mysql"
)

func main() {
	pathToSql := os.Args[1]
	data, err := ioutil.ReadFile(pathToSql)
	if err != nil {
		log.Fatal(err)
	}

	stmts := parser.Parse(strings.NewReader(string(data)))

	mysqlStmts, err := spanner2mysql.GetMysqlCreateTables(stmts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mysqlStmts)
}
