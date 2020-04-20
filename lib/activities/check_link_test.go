package activities

import (
	"github.com/dbtedman/go-link-check/lib/gateways"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckLinkThatExists(t *testing.T) {
	var fetch gateways.Fetch = &gateways.FetchFake{
		ResponseStub: &gateways.FetchResponse{
			Content: "",
			Status:  200,
		},
	}

	result := *CheckLink(&CheckLinkRequest{
		Fetch: &fetch,
		URL:   "https://example.com",
	})

	assert.IsType(t, CheckLinkResponse{}, result)
	assert.True(t, result.Success)
}

func TestCheckLinkThatIsMissing(t *testing.T) {
	var fetch gateways.Fetch = &gateways.FetchFake{
		ResponseStub: &gateways.FetchResponse{
			Content: "",
			Status:  404,
		},
	}

	result := *CheckLink(&CheckLinkRequest{
		Fetch: &fetch,
		URL:   "https://example.not.com",
	})

	assert.IsType(t, CheckLinkResponse{}, result)
	assert.False(t, result.Success)
}
