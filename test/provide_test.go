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
	multiaddr "github.com/multiformats/go-multiaddr"
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

	c, s := createClientAndServer(t, testDelegatedRoutingService{}, &client.Provider{
		Peer: peer.AddrInfo{
			ID:    pID,
			Addrs: []multiaddr.Multiaddr{},
		},
		ProviderProto: []client.TransferProtocol{},
	}, priv)
	defer s.Close()

	testMH, _ := multihash.Encode([]byte("test"), multihash.IDENTITY)
	testCid := cid.NewCidV1(cid.Raw, testMH)

	if _, err = c.Provide(context.Background(), testCid, time.Hour); err == nil {
		t.Fatal("should get sync error on unsigned provide request.")
	}

	rc, err := c.Provide(context.Background(), testCid, 2*time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	if rc != time.Hour {
		t.Fatal("should have gotten back the the fixed server ttl")
	}
}
