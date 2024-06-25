package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)


func visit(n *html.Node, pwords, ppics *int) {
    if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
        return
    }
	if n.Type == html.TextNode {
		*pwords += len(strings.Fields(n.Data))

	} else if n.Type == html.ElementNode && n.Data == "img" {
		*ppics++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, pwords, ppics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int

	visit(doc, &words, &pics)

	return words, pics
}

func getHtmlFromUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
	return content
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	raw := getHtmlFromUrl(os.Args[1])
	doc, err := html.Parse(bytes.NewReader(raw))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse failed: %s\n", err)
	}

	words, pics := countWordsAndImages(doc)

	fmt.Printf("%d words and %d images\n", words, pics)

}
