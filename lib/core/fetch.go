package core

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"regexp"
)

const looksLikeURL = "^https?://.+$"
const pathWithoutStartingSlash = "^/?(.+)$"
const isHashReference = "#.+"

func MakeURLAbsolute(urlOrPath string, baseURL string) string {
	if regexp.MustCompile(looksLikeURL).MatchString(urlOrPath) {
		return urlOrPath
	}

	return baseURL +
		regexp.MustCompile(pathWithoutStartingSlash).ReplaceAllString(urlOrPath, "/$1")
}

func LinksFromDocument(doc *goquery.Document) []string {
	var foundLinkHRefs []string

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		foundLinkURL, hrefExists := selection.Attr("href")

		if hrefExists {
			foundLinkHRefs = append(foundLinkHRefs, foundLinkURL)
		}
	})

	return foundLinkHRefs
}

func FetchLinksFromPage(url string, client *http.Client) []string {
	var links []string

	resp, err := client.Get(url)

	if err != nil {
		return links
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return links
	}

	for _, href := range LinksFromDocument(doc) {
		if !regexp.MustCompile(isHashReference).MatchString(href) {
			links = append(links, MakeURLAbsolute(href, url))
		}
	}

	return links
}
