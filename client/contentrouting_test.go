package client

import (
	"context"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/libp2p/go-libp2p-core/peer"
)

type TestDelegatedRoutingClient struct {
	NumResults int
}

func (t TestDelegatedRoutingClient) FindProviders(ctx context.Context, key cid.Cid) ([]peer.AddrInfo, error) {
	panic("not supported")
}

func (t TestDelegatedRoutingClient) FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	ch := make(chan FindProvidersAsyncResult)
	go func() {
		defer close(ch)
		for i := 0; i < t.NumResults; i++ {
			ch <- FindProvidersAsyncResult{
				AddrInfo: []peer.AddrInfo{{}},
			}
		}
	}()
	return ch, nil
}

func (t TestDelegatedRoutingClient) GetIPNS(ctx context.Context, id []byte) ([][]byte, error) {
	panic("not supported")
}

func (t TestDelegatedRoutingClient) PutIPNSAsync(ctx context.Context, id []byte, record []byte) (<-chan PutIPNSAsyncResult, error) {
	panic("not supported")
}

func TestContentRoutingFindProvidersUnlimitedResults(t *testing.T) {
	providedResults := 5
	c := NewContentRoutingClient(TestDelegatedRoutingClient{providedResults})
	ch := c.FindProvidersAsync(context.Background(), cid.Cid{}, 0)
	num := 0
	for range ch {
		num++
	}
	if num != providedResults {
		t.Errorf("expecting %v results, got %v", providedResults, num)
	}
}

func TestContentRoutingFindProvidersFewerResults(t *testing.T) {
	providedResults := 5
	wantResults := 3
	c := NewContentRoutingClient(TestDelegatedRoutingClient{providedResults})
	ch := c.FindProvidersAsync(context.Background(), cid.Cid{}, wantResults)
	num := 0
	for range ch {
		num++
	}
	if num != wantResults {
		t.Errorf("expecting %v results, got %v", wantResults, num)
	}
}

func TestContentRoutingFindProvidersMoreResults(t *testing.T) {
	providedResults := 5
	wantResults := 7
	c := NewContentRoutingClient(TestDelegatedRoutingClient{providedResults})
	ch := c.FindProvidersAsync(context.Background(), cid.Cid{}, wantResults)
	num := 0
	for range ch {
		num++
	}
	if num != providedResults {
		t.Errorf("expecting %v results, got %v", providedResults, num)
	}
}
