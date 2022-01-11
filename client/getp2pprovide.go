package client

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	proto "github.com/ipfs/go-delegated-routing/ipld/ipldsch"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

// NOTE: This file will be codegen'd by a protocol compiler.

func (c *client) GetP2PProvide(ctx context.Context, req *proto.GetP2PProvideRequest) ([]*proto.GetP2PProvideResponse, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch, err := c.GetP2PProvide_Async(ctx, req)
	if err != nil {
		return nil, err
	}
	var resps []*proto.GetP2PProvideResponse
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return resps, nil
			} else {
				if r.Err == nil {
					resps = append(resps, r.Resp)
				} else {
					logger.Errorf("delegated client received invalid response (%v)", r.Err)
				}
			}
		case <-ctx.Done():
			return resps, ctx.Err()
		}
	}
}

func (c *client) GetP2PProvide_Async(ctx context.Context, req *proto.GetP2PProvideRequest) (<-chan GetP2PProvide_Async_Response, error) {
	envelope := &proto.Envelope{
		GetP2PProvideRequest: req,
	}

	buf, err := ipld.Marshal(dagjson.Encode, envelope, proto.Prototypes.Envelope.Type())
	if err != nil {
		return nil, fmt.Errorf("unexpected serialization error (%w)", err)
	}

	// encode request in URL
	u := *c.endpoint
	q := url.Values{}
	q.Set("q", string(buf))
	u.RawQuery = q.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, "GET", u.String(), bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	ch := make(chan GetP2PProvide_Async_Response, 1)
	go process_GetP2PProvide_Async_Response(ctx, ch, resp.Body)
	return ch, nil
}

func process_GetP2PProvide_Async_Response(ctx context.Context, ch chan<- GetP2PProvide_Async_Response, r io.Reader) {
	defer close(ch)
	for {
		if ctx.Err() != nil {
			return
		}

		env := &proto.Envelope{}
		// ISSUE: bindnode codegen should not require the user to provide the prototype manually
		_, err := ipld.UnmarshalStreaming(r, dagjson.Decode, env, proto.Prototypes.Envelope.Type())
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			ch <- GetP2PProvide_Async_Response{Err: err}
		}

		if env.GetP2PProvideResponse == nil {
			continue
		}
		ch <- GetP2PProvide_Async_Response{Resp: env.GetP2PProvideResponse}
	}
}
