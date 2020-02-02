package cli

import (
	"flag"
	"github.com/dbtedman/go-link-check/lib/core"
	"github.com/dbtedman/go-link-check/lib/services"
	"regexp"
)

func Run() {
	services.ConfigureLogging()

	url := flag.String("url", "", core.CliOptionUrl)
	outFile := flag.String("outFile", "results.csv", core.CliOptionOutFile)

	flag.Parse()

	if *url == "" {
		services.ErrorLine(core.CliErrorMissingOptionUrl)
		return
	}

	const pattern = "^.+[.]csv$"
	match, _ := regexp.MatchString(pattern, *outFile)

	if !match {
		services.ErrorFormat(core.CliErrorInvalidOutFile, *outFile, pattern)
		return
	}

	config := core.ApplicationConfiguration{}
	config.BaseURL = url
	config.OutputFilePath = outFile

	core.Application(&config)
}
