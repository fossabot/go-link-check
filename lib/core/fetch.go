package core

import (
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func FetchLinksFromPage(url string, client *http.Client) []string {
	var links []string

	log.Printf("Finding links to validate on %s", url)

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

	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, hrefExists := selection.Attr("href")
		if hrefExists {
			links = append(links, href)
		}
	})

	return links
}
