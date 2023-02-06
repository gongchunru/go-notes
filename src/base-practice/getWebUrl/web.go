package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

type Item struct {
	Name string
	URL  string
	GUID string
	Icon string
}

func main() {
	url := "https://oss.foneshare.cn/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("get %s\n", url)
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Fetch: reading %s: %v\n", url, err)
		os.Exit(1)
	}
	fmt.Printf("%s", body)

	reader := bytes.NewReader(body)
	parse, err := html.Parse(reader)
	if err != nil {
		fmt.Printf("parse err:%v\n", err)
	}
	links := allLinks(nil, parse)
	fmt.Printf("parse: %v\n links:%v\n", parse, links)
}

func allLinks(links []Item, node *html.Node) []Item {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				if a.Val == "" {
					continue
				}
				if node.FirstChild != nil {
					links = append(links, Item{
						Name: node.FirstChild.Data,
						URL:  a.Val,
					})
				}
				continue
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = allLinks(links, c)
	}
	return links
}
