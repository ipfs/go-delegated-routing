package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/parser"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"

	//"github.com/ipld/go-ipld-prime/codec/dagjson"
	"io"
	"net/http"
)

type Option func(*client) error

type client struct {
	client *http.Client
	endPoint string
}

func WithHTTPClient(hc *http.Client) Option {
	return func(c *client) error {
		c.client = hc
		return nil
	}
}

func New(endpoint string, opts...Option) (*client, error) {
	c := &client{endPoint: endpoint, client: http.DefaultClient}
	for _, o := range opts {
		if err := o(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

func (c *client) FindProviders(ctx context.Context, cid cid.Cid) ([]peer.AddrInfo, error) {
	req := parser.Envelope{
		Tag: parser.MethodGetP2PProvide,
		Payload: parser.GetP2PProvide{
			Key: parser.ToDJSpecialBytes(cid.Hash()),
		},
	}
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(req); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.endPoint, b)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	return processProvResp(resp.Body)
}

func processProvResp(r io.Reader) ([]peer.AddrInfo, error) {
	var ais []peer.AddrInfo
	for {
		dec := json.NewDecoder(r)
		env := parser.Envelope{Payload: &parser.GetP2PProvideResponse{}}
		err := dec.Decode(&env)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		if env.Tag != parser.MethodGetP2PProvide {
			continue
		}

		provResp, ok := env.Payload.(*parser.GetP2PProvideResponse)
		if !ok {
			continue
		}

		for _, maBytes := range provResp.Peers {
			addrBytes, err := parser.FromDJSpecialBytes(maBytes)
			if err != nil {
				continue
			}
			ma, err := multiaddr.NewMultiaddrBytes(addrBytes)
			if err != nil {
				continue
			}
			ai, err := peer.AddrInfoFromP2pAddr(ma)
			if err != nil {
				continue
			}
			ais = append(ais, *ai)
		}
	}

	return ais, nil
}

func processProv(r io.Reader) ([]peer.AddrInfo, error) {
	var ais []peer.AddrInfo
	for {
		dec := json.NewDecoder(r)
		env := parser.Envelope{Payload: &parser.GetP2PProvideResponse{}}
		err := dec.Decode(&env)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		if env.Tag != parser.MethodGetP2PProvide {
			continue
		}

		provResp, ok := env.Payload.(*parser.GetP2PProvideResponse)
		if !ok {
			continue
		}

		for _, maBytes := range provResp.Peers {
			addrBytes, err := parser.FromDJSpecialBytes(maBytes)
			if err != nil {
				continue
			}
			// TODO: This is wrong the data should be binary, but this was easier for testing
			ma, err := multiaddr.NewMultiaddr(string(addrBytes))
			if err != nil {
				continue
			}
			ai, err := peer.AddrInfoFromP2pAddr(ma)
			if err != nil {
				continue
			}
			ais = append(ais, *ai)
		}
	}

	return ais, nil
}