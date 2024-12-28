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
	return getLink(doc), nil
}

func main() {
	fileName := flag.String("file", "index.html", "The HTML file to parse")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	links, err := Parse(file)
	if err != nil {
		panic(err)
	}

	for _, val := range links {
		fmt.Println(val)
	}
}

func getLink(n *html.Node) []Link {
	var links []Link
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "a" && n.Type == html.ElementNode {
			href := ""
			for _, atr := range c.Attr {
				if atr.Key == "href" {
					href = atr.Val
					break
				}
			}
			text := extractText(c)
			// extractText(n)
			if href != "" && text != "" {
				links = append(links, Link{href, text})
			}
		} else {
			links = append(links, getLink(c)...)
		}
	}
	return links
}

func extractText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	// fmt.Printf("Type: %v, Data: %v\n", n.Type, n.Data)
	text := ""
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += extractText(c)
	}
	return text
}
