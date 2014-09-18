package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseLink(t *testing.T) {
	reader := strings.NewReader(` <p>
	<a href="http://google.com">1</a>
	<a href="http://facebook.com">2</a>
	<a href="http://twitter.com">3</a>
	http://yahoo.com
</p>`)

	links := ParseLink(reader)

	fmt.Println(links)

	if len(links) != 3 {
		t.Error("Wrong number of links returned")
	}

	if links[0] != "http://google.com" {
		t.Error("The first link is incorrect:", links[0])
	}

	if links[1] != "http://facebook.com" {
		t.Error("The second link is incorrect:", links[1])
	}

	if links[2] != "http://twitter.com" {
		t.Error("The third link is incorrect:", links[2])
	}
}
