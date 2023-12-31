package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "map: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	traverse(doc, m)
	prettyPrint(m)
}

func prettyPrint(m map[string]int) {
	for key, element := range m {
		fmt.Printf("The number of [%s] is: %d\n", key, element)
	}
}

func traverse(node *html.Node, m map[string]int) {
	if node.FirstChild != nil {
		traverse(node.FirstChild, m)
	}

	if node.NextSibling != nil {
		traverse(node.NextSibling, m)
	}

	tag := node.DataAtom.String()

	if len(tag) != 0 {
		m[tag] += 1
	}
}
