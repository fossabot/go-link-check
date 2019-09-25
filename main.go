package main

import (
	"flag"
	"github.com/dbtedman/go-link-check/lib"
	log "github.com/sirupsen/logrus"
	"regexp"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	url := flag.String("url", "", "Website URL to parse for links to validate.")
	outFile := flag.String("outFile", "results.csv", "(optional) Path to output report csv to.")

	flag.Parse()

	if *url == "" {
		log.Errorln("-url must be provided")
		return
	}

	const pattern = "^.+[.]csv$"
	match, _ := regexp.MatchString(pattern, *outFile)

	if !match {
		log.Errorf("-outFile `%s` must be a file path matching pattern `%s`", *outFile, pattern)
		return
	}

	urls := lib.FetchLinksFromPage(*url)

	linkStatusList := lib.CheckAllLinkStatus(urls)

	lib.WriteResultsToFile(*outFile, linkStatusList)
}
