package test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	indexer "github.com/filecoin-project/go-indexer-core"
	server "github.com/filecoin-project/storetheindex/server/finder/http"
	"github.com/filecoin-project/storetheindex/server/finder/test"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/client"
	"github.com/libp2p/go-libp2p-core/peer"
	multiaddr "github.com/multiformats/go-multiaddr"
	"github.com/multiformats/go-multihash"
)

func TestClientServer(t *testing.T) {
	// start a server
	ind := test.InitIndex(t, true)
	reg := test.InitRegistry(t)
	s, err := server.New("127.0.0.1:0", ind, reg)
	if err != nil {
		t.Fatal(err)
	}
	errChan := make(chan error, 1)
	go func() {
		err := s.Start()
		if err != http.ErrServerClosed {
			errChan <- err
		}
		close(errChan)
	}()
	// start a client
	c, err := client.New(s.URL())
	if err != nil {
		t.Fatal(err)
	}
	// verify result
	h, err := multihash.Sum([]byte("TEST"), multihash.SHA3, 4)
	if err != nil {
		t.Fatal(err)
	}
	p, err := peer.Decode("12D3KooWKRyzVWW6ChFjQjK4miCty85Niy48tpPV95XdKu1BcvMA")
	if err != nil {
		t.Fatal(err)
	}
	if err := ind.Put(indexer.Value{
		ProviderID:    p,
		ContextID:     []byte("ctx"),
		MetadataBytes: []byte{42},
	}, h); err != nil {
		t.Fatal(err)
	}

	infos, err := c.FindProviders(context.Background(), cid.NewCidV1(cid.Libp2pKey, h))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(infos)

	err = s.Shutdown(context.Background())
	if err != nil {
		t.Error("shutdown error:", err)
	}
	err = <-errChan
	if err != nil {
		t.Fatal(err)
	}

	if err = reg.Close(); err != nil {
		t.Errorf("Error closing registry: %s", err)
	}
	if err = ind.Close(); err != nil {
		t.Errorf("Error closing indexer core: %s", err)
	}
}

func testFindProvidersAsyncFunc(key cid.Cid, ch chan<- client.FindProvidersAsyncResult) error {
	ma := multiaddr.StringCast("/ip4/7.7.7.7/tcp/4242/p2p/QmYyQSo1c1Ym7orWxLYvCrM2EmxFTANf8wXmmE7DWjhx5N")
	ai, err := peer.AddrInfoFromP2pAddr(ma)
	if err != nil {
		println(err.Error())
		return fmt.Errorf("address info creation (%v)", err)
	}
	go func() {
		ch <- client.FindProvidersAsyncResult{AddrInfo: []peer.AddrInfo{*ai}}
		close(ch)
	}()
	return nil
}
