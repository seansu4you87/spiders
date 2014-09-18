package main

import (
	"io"

	"code.google.com/p/go.net/html"
)

func ParseLink(reader io.Reader) []string {
	links := make([]string, 0)
	page := html.NewTokenizer(reader)
	for {
		tokenType := page.Next()
		if tokenType == html.ErrorToken {
			return links
		}
		token := page.Token()
		if tokenType == html.StartTagToken && token.DataAtom.String() == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
	}
	return links
}
