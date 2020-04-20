package activities

import "github.com/dbtedman/go-link-check/lib/gateways"

type CheckLinkRequest struct {
	Fetch *gateways.Fetch
	URL   string
}

type CheckLinkResponse struct {
	Success bool
}

func CheckLink(request *CheckLinkRequest) *CheckLinkResponse {
	var fetch = *request.Fetch
	var fetchResponse = *fetch.Get(&request.URL)

	response := CheckLinkResponse{
		Success: fetchResponse.Status == 200,
	}

	return &response
}
