package cli

import (
	"flag"
	"github.com/dbtedman/go-link-check/lib/core"
	"github.com/sirupsen/logrus"
	"net/http"
	"regexp"
)

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)

	url := flag.String("url", "", "Website URL to parse for links to validate.")
	outFile := flag.String("outFile", "results.csv", "(optional) Path to output report csv to.")

	flag.Parse()

	if *url == "" {
		logrus.Errorln("-url must be provided")
		return
	}

	const pattern = "^.+[.]csv$"
	match, _ := regexp.MatchString(pattern, *outFile)

	if !match {
		logrus.Errorf("-outFile `%s` must be a file path matching pattern `%s`", *outFile, pattern)
		return
	}

	urls := core.FetchLinksFromPage(*url, http.DefaultClient)

	linkStatusList := core.CheckAllLinkStatus(urls)

	core.WriteResultsToFile(*outFile, linkStatusList)
}
