package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/syucream/spar/src/parser"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	_, err = parser.Parse(strings.NewReader(string(data)))
	if err != nil {
		log.Fatal(err)
	}
}
