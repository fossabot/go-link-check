package http

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type LinkStatus struct {
	url       string
	success   bool
	redirects bool
	code      int
}

func CheckAllLinkStatus(urls []string) []LinkStatus {
	wg := sync.WaitGroup{}

	// Define a channel for communicating from go routine back to calling function.
	linkStatuses := make(chan *LinkStatus)

	// Determine how may go routines to wait for.
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
			fmt.Println("checking: " + linkStatus.url)
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

	// TODO: Should we check anything else to determine if it was successful?
	status.success = true
	status.redirects = url != resp.Request.URL.String()
	status.code = resp.StatusCode

	return status
}
