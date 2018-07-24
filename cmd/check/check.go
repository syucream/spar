package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/syucream/spar/src/parser"
)

func main() {
	pathToSql := os.Args[1]
	data, err := ioutil.ReadFile(pathToSql)
	if err != nil {
		log.Fatal(err)
	}

	stmts := parser.Parse(strings.NewReader(string(data)))
	fmt.Println(stmts)
}
