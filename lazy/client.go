package lazy

import (
	"github.com/google/go-github/github"
)

type LazyClient struct {
	reader 		*Reader
	transport	*github.BasicAuthTransport
	client 		*github.Client
}

func NewClient() *LazyClient {
	// create new Reader
	reader := NewReader()
	transport := AuthUser(reader)
	client := github.NewClient(transport.Client())

	return &LazyClient{
		reader: reader,
		transport: transport,
		client: client,
	}
}
