package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/parser"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func (c *client) FindProviders(ctx context.Context, cid cid.Cid) ([]peer.AddrInfo, error) {
	ctx, cancel := context.WithCancel(ctx)
	ch, err := c.FindProvidersAsync(ctx, cid)
	if err != nil {
		return nil, err
	}
	select {
	case r := <-ch:
		cancel()
		if r.Err != nil {
			return nil, err
		} else {
			return r.AddrInfo, nil
		}
	case <-ctx.Done():
		return nil, fmt.Errorf("context aborted")
	}
	panic("unreachable")
}

type FindProvidersAsyncResult struct {
	AddrInfo []peer.AddrInfo
	Err      error
}

func (c *client) FindProvidersAsync(ctx context.Context, cid cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	req := parser.Envelope{
		Tag: parser.MethodGetP2PProvide,
		Payload: parser.GetP2PProvideRequest{
			Key: parser.ToDJSpecialBytes(cid.Hash()),
		},
	}
	b := &bytes.Buffer{}
	if err := json.NewEncoder(b).Encode(req); err != nil {
		return nil, err
	}

	// encode request in URL
	url := fmt.Sprintf("%s?%s", c.endPoint, url.QueryEscape(b.String()))
	httpReq, err := http.NewRequestWithContext(ctx, "GET", url, b)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	ch := make(chan FindProvidersAsyncResult, 1)
	go processFindProvidersAsyncResp(ctx, ch, resp.Body)
	return ch, nil
}

func processFindProvidersAsyncResp(ctx context.Context, ch chan<- FindProvidersAsyncResult, r io.Reader) {
	defer close(ch)
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		dec := json.NewDecoder(r)
		env := parser.Envelope{Payload: &parser.GetP2PProvideResponse{}}
		err := dec.Decode(&env)
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			ch <- FindProvidersAsyncResult{Err: err}
			return
		}

		if env.Tag != parser.MethodGetP2PProvide {
			continue
		}

		provResp, ok := env.Payload.(*parser.GetP2PProvideResponse)
		if !ok {
			continue
		}

		infos := []peer.AddrInfo{}
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
			infos = append(infos, *ai)
		}
		ch <- FindProvidersAsyncResult{AddrInfo: infos}
	}
}
