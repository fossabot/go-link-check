package core

import "net/http"

type ApplicationConfiguration struct {
	BaseURL        *string
	OutputFilePath *string
}

func Application(config *ApplicationConfiguration) {
	urls := FetchLinksFromPage(*config.BaseURL, http.DefaultClient)

	linkStatusList := CheckAllLinkStatus(urls)

	WriteResultsToFile(*config.OutputFilePath, linkStatusList)
}
