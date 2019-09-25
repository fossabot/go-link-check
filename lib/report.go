package lib

import (
	"encoding/csv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func WriteResultsToFile(filePath string, linkStatusList []LinkStatus) {
	file, _ := os.Create(filePath)

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

	log.WithFields(ContextFields()).Infof("Results written to file \"%s\"", filePath)
}
