package lib

import (
	"bytes"
	"fmt"
	"testing"
)

func TestWriteResultsToWriter(t *testing.T) {
	t.Run("handles empty link statuses", func(t *testing.T) {
		buffer := bytes.Buffer{}
		linkStatuses := make([]LinkStatus, 0)

		WriteResultsToWriter(&buffer, linkStatuses)

		got := buffer.String()
		want := "URL,Success,Redirects,Code\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("handles a single link status", func(t *testing.T) {
		buffer := bytes.Buffer{}
		linkStatuses := make([]LinkStatus, 0)

		exampleLinkStatus := LinkStatus{
			Url:       "https://example.com",
			Success:   false,
			Redirects: true,
			Code:      5,
		}

		linkStatuses = append(linkStatuses, exampleLinkStatus)

		WriteResultsToWriter(&buffer, linkStatuses)

		got := buffer.String()
		want := fmt.Sprintf(
			"URL,Success,Redirects,Code\n%s,%t,%t,%d\n",
			exampleLinkStatus.Url,
			exampleLinkStatus.Success,
			exampleLinkStatus.Redirects,
			exampleLinkStatus.Code,
		)

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("handles many link statuses", func(t *testing.T) {
		t.Skip()
	})

	t.Run("handles error in writer", func(t *testing.T) {
		t.Skip()
	})
}
