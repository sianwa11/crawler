package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML :%w", err)
	}

	links := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok || strings.TrimSpace(src) == "" {
			return
		}

		parsedUrl, err := url.Parse(src)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", parsedUrl, err)
			return
		}
		absolutePath := baseURL.ResolveReference(parsedUrl)
		links = append(links, absolutePath.String())
	})

	return links, nil
}
