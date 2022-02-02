package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multihash"
)

var logger = logging.Logger("delegated/client")

func (c *client) FindProviders(ctx context.Context, cid cid.Cid) ([]peer.AddrInfo, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	ch, err := c.FindProvidersAsync(ctx, cid)
	if err != nil {
		return nil, err
	}
	var infos []peer.AddrInfo
	for {
		select {
		case r, ok := <-ch:
			if !ok {
				cancel()
				return infos, nil
			} else {
				if r.Err == nil {
					infos = append(infos, r.AddrInfo...)
				} else {
					logger.Errorf("delegated client received invalid response (%v)", r.Err)
				}
			}
		case <-ctx.Done():
			return infos, ctx.Err()
		}
	}
}

type FindProvidersAsyncResult struct {
	AddrInfo []peer.AddrInfo
	Err      error
}

type indexFindResponse struct {
	MultihashResults []indexMultihashResult
}

type indexMultihashResult struct {
	Multihash       multihash.Multihash
	ProviderResults []indexProviderResult
}

type indexProviderResult struct {
	ContextID []byte
	Metadata  json.RawMessage
	Provider  peer.AddrInfo
}

func (c *client) FindProvidersAsync(ctx context.Context, cid cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	// encode request in URL
	u := fmt.Sprint(c.endpoint.String(), "/", cid.Hash().B58String())
	httpReq, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			ch := make(chan FindProvidersAsyncResult)
			close(ch)
			return ch, nil
		}
		return nil, fmt.Errorf("find query failed: %v", http.StatusText(resp.StatusCode))
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	parsedResponse := indexFindResponse{}
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		return nil, err
	}

	result := FindProvidersAsyncResult{}
	if len(parsedResponse.MultihashResults) != 1 {
		return nil, fmt.Errorf("unexpected number of responses")
	}
	for _, m := range parsedResponse.MultihashResults[0].ProviderResults {
		result.AddrInfo = append(result.AddrInfo, m.Provider)
	}

	ch := make(chan FindProvidersAsyncResult, 1)
	ch <- result
	return ch, nil
}
