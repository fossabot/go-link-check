package gateways

type Fetch interface {
	Get(url *string) *FetchResponse
}

type FetchResponse struct {
	// e.g. 200, 404
	Status int

	// the response body
	Content string
}
