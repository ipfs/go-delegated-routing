package test

import (
	"context"
	"crypto/rand"
	"testing"
	"time"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/client"
	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multihash"
)

func TestProvideRoundtrip(t *testing.T) {
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	pID, err := peer.IDFromPrivateKey(priv)
	if err != nil {
		t.Fatal(err)
	}

	c, s := createClientAndServer(t, testDelegatedRoutingService{})
	defer s.Close()

	testMH, _ := multihash.Encode([]byte("test"), multihash.IDENTITY)
	testCid := cid.NewCidV1(cid.Raw, testMH)
	req := client.ProvideRequest{
		Key: testCid,
		Provider: client.Provider{
			Peer: peer.AddrInfo{ID: pID},
		},
		AdvisoryTTL: time.Hour,
	}
	rc, err := c.Provide(context.Background(), &req)
	if err == nil {
		t.Fatal("should get sync error on unsigned provide request.")
	}

	if err = req.Sign(priv); err != nil {
		t.Fatal(err)
	}
	if rc, err = c.Provide(context.Background(), &req); err != nil {
		t.Fatal(err)
	}

	res := <-rc
	if res.Err != nil {
		t.Fatal(err)
	}
}
