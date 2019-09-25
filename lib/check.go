package lib

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

type LinkStatus struct {
	Url       string
	Success   bool
	Redirects bool
	Code      int
}

func CheckAllLinkStatus(urls []string) []LinkStatus {
	wg := sync.WaitGroup{}

	// Define a channel for communicating from go routine back to calling function.
	linkStatuses := make(chan *LinkStatus)

	// Determine how many go routines to wait for.
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			status := CheckLinkStatus(url)
			linkStatuses <- &status
		}(url)
	}

	linkStatusList := make([]LinkStatus, 0)

	go func() {
		// This runs after each go routine has finished, how does this work?
		for linkStatus := range linkStatuses {
			log.WithFields(ContextFields()).Debugf("checking: %s", linkStatus.Url)
			linkStatusList = append(linkStatusList, *linkStatus)
		}
	}()

	wg.Wait()

	// Occasionally we end up not getting the last link otherwise.
	time.Sleep(1 * time.Second)

	return linkStatusList
}

func CheckLinkStatus(url string) LinkStatus {
	status := LinkStatus{url, false, false, -1}

	resp, err := http.Get(url)

	if err != nil {
		return status
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
	}

	status.Success = true
	status.Redirects = url != resp.Request.URL.String()
	status.Code = resp.StatusCode

	return status
}
