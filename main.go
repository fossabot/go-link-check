package main

import (
	"fmt"
	"github.com/dbtedman/go-link-check/lib/services/http"
)

func main() {
	// TODO: Initial work on simple one page check mode.

	urls := http.FetchLinksFromPage("https://danieltedman.com")
	linkStatusList := http.CheckAllLinkStatus(urls)

	// We will then do something with this list.
	fmt.Printf("%+v\n", linkStatusList)
}
