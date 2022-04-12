package client

import (
	"context"

	"github.com/ipfs/go-delegated-routing/gen/proto"
)

func (fp *Client) GetIPNS(ctx context.Context, id []byte) ([][]byte, error) {
	resps, err := fp.client.GetIPNS(ctx, &proto.GetIPNSRequest{ID: id})
	if err != nil {
		return nil, err
	}
	records := [][]byte{}
	for _, resp := range resps {
		records = append(records, resp.Record)
	}
	return records, nil
}

type GetIPNSAsyncResult struct {
	Record []byte
	Err    error
}

func (fp *Client) GetIPNSAsync(ctx context.Context, id []byte) (<-chan GetIPNSAsyncResult, error) {
	ch0, err := fp.client.GetIPNS_Async(ctx, &proto.GetIPNSRequest{ID: id})
	if err != nil {
		return nil, err
	}
	ch1 := make(chan GetIPNSAsyncResult, 1)
	go func() {
		defer close(ch1)
		if ctx.Err() != nil {
			return
		}
		r0, ok := <-ch0
		if !ok {
			return
		}
		var r1 GetIPNSAsyncResult
		r1.Err = r0.Err
		if r0.Resp != nil {
			r1.Record = r0.Resp.Record
		}
		ch1 <- r1
	}()
	return ch1, nil
}
