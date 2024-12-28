package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	printNode(doc, 0)

	return nil, nil
}

func main() {
	fileName := flag.String("file", "index.html", "The HTML file to parse")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = Parse(file)
	if err != nil {
		panic(err)
	}

}

func printNode(n *html.Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("Type: %v, Data: %v, Attribute: %v\n", n.Type, n.Data, n.Attr)

	// Recurse to child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printNode(c, depth+1)
	}
}
