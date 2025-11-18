package main

import "net/url"

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	pageData := PageData{}

	pageData.URL = pageURL
	pageData.H1 = getH1FromHTML(html)
	pageData.FirstParagraph = getFirstParagraphFromHTML(html)

	baseUrl, err := url.Parse(pageURL)
	if err != nil {
		return pageData
	}

	urls, err := getURLsFromHTML(html, baseUrl)
	if err != nil {
		return pageData
	}
	pageData.OutgoingLinks = urls

	images, err := getImagesFromHTML(html, baseUrl)
	if err != nil {
		return pageData
	}
	pageData.ImageURLs = images

	return pageData
}
