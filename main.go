package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ah-naf/html-link-parser/htmlparser"
)

func main() {
	fileName := flag.String("file", "index.html", "The HTML file to parse")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := htmlparser.Parse(file)
	if err != nil {
		panic(err)
	}

	for _, val := range links {
		fmt.Println(val)
	}
}
