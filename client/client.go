package client

import (
	"context"
	"net/http"
	"net/url"

	proto "github.com/ipfs/go-delegated-routing/ipld/ipldsch"
	logging "github.com/ipfs/go-log"
)

var logger = logging.Logger("delegated/client")

type Client interface {
	GetP2PProvide(ctx context.Context, req *proto.GetP2PProvideRequest) ([]*proto.GetP2PProvideResponse, error)
	GetP2PProvide_Async(ctx context.Context, req *proto.GetP2PProvideRequest) (<-chan GetP2PProvide_Async_Response, error)
}

type GetP2PProvide_Async_Response struct {
	Resp *proto.GetP2PProvideResponse
	Err  error
}

type Option func(*client) error

type client struct {
	client       *http.Client
	endpoint     *url.URL
	NativeClient // mixin provides higher-level APIs used by DHT and Hydra
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *client) error {
		c.client = hc
		return nil
	}
}

func New(endpoint string, opts ...Option) (*client, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	c := &client{endpoint: u, client: http.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	c.NativeClient.client = c
	return c, nil
}
