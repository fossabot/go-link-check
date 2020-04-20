package gateways

type FetchFake struct {
	ResponseStub *FetchResponse
}

func (f *FetchFake) Get(url *string) *FetchResponse {
	return f.ResponseStub
}
