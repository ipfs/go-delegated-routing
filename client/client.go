package client

import (
	//"github.com/ipld/go-ipld-prime/codec/dagjson"

	"net/http"
)

type Option func(*client) error

type client struct {
	client   *http.Client
	endPoint string
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *client) error {
		c.client = hc
		return nil
	}
}

func New(endpoint string, opts ...Option) (*client, error) {
	c := &client{endPoint: endpoint, client: http.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}
