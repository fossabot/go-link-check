package main

import (
	"encoding/csv"
	"fmt"
	"github.com/dbtedman/go-link-check/lib"
	"os"
	"strconv"
)

// TODO: Move all but absolute minimum from this main function into separate functions for testability.
func main() {
	// TODO: Initial work on simple one page check mode.

	const ResultsFile = "results.csv"

	urls := lib.FetchLinksFromPage("https://danieltedman.com")
	linkStatusList := lib.CheckAllLinkStatus(urls)

	file, _ := os.Create(ResultsFile)

	if file != nil {
		defer func() {
			_ = file.Close()
		}()
	}

	writer := csv.NewWriter(file)

	if file != nil {
		defer func() {
			writer.Flush()
		}()
	}

	_ = writer.Write([]string{
		"URL",
		"Success",
		"Redirects",
		"Code",
	})

	for _, value := range linkStatusList {
		_ = writer.Write([]string{
			value.Url,
			strconv.FormatBool(value.Success),
			strconv.FormatBool(value.Redirects),
			strconv.FormatUint(uint64(value.Code), 10),
		})
	}

	fmt.Printf("\nResults written to file \"%s\"\n\n", ResultsFile)
}
