package client

import (
	"context"

	"github.com/ipfs/go-delegated-routing/gen/proto"
)

func (fp *Client) GetIPNS(ctx context.Context, id []byte) ([][]byte, error) {
	resps, err := fp.GetIPNSAsync(ctx, id)
	if err != nil {
		return nil, err
	}
	records := [][]byte{}
	for resp := range resps {
		if resp.Err != nil {
			records = append(records, resp.Record)
		}
	}
	if fp.validator != nil {
		best, err := fp.validator.Select(string(id), records)
		if err != nil {
			return nil, err
		}
		return [][]byte{records[best]}, nil
	} else {
		return records, nil
	}
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
		if r0.Err != nil {
			r1.Err = r0.Err
			ch1 <- r1
		} else if r0.Resp != nil {
			if fp.validator != nil {
				if err = fp.validator.Validate(string(id), r0.Resp.Record); err != nil {
					logger.Infof("received invalid ipns record (%v)", err)
				} else {
					r1.Record = r0.Resp.Record
					ch1 <- r1
				}
			} else {
				r1.Record = r0.Resp.Record
				ch1 <- r1
			}
		}
	}()
	return ch1, nil
}
