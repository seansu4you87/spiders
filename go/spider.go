package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	// "io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	flag.Parse()

	args := flag.Args()
	fmt.Println(args)

	if len(args) < 1 {
		fmt.Println("Please specify start page")
		os.Exit(1)
	}

	queue := make(chan string)
	uniqueQueue := make(chan string)
	go dedupQueue(queue, uniqueQueue)

	queue <- args[0] // seed the queue

	for uri := range uniqueQueue {
		crawl(uri, queue)
	}

}

func dedupQueue(in, out chan string) {
	visited := make(map[string]bool)

	for uri := range in {
		if !visited[uri] {
			visited[uri] = true
			out <- uri
		}
	}
}

func crawl(uri string, queue chan string) {
	fmt.Println("fetching", uri)
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := http.Client{Transport: transport}

	resp, err := client.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	links := ParseLink(resp.Body)

	for _, link := range links {
		absLink := fixUrl(link, uri)

		if uri != "" {
			go func() { queue <- absLink }()
		}
	}
}

func fixUrl(href, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}
	baseUrl, err := url.Parse(base)
	if err != nil {
		return ""
	}
	uri = baseUrl.ResolveReference(uri)
	return uri.String()
}

// func retrieve(uri string) {
// 	resp, err := http.Get(uri)
// 	if err != nil {
// 		return
// 	}
// 	defer resp.Body.Close()

// 	body, _ := ioutil.ReadAll(resp.Body)
// 	fmt.Println(string(body))
// }
