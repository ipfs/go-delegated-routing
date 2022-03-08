package client

import (
	"context"
	"fmt"

	"github.com/ipfs/go-cid"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
	multihash "github.com/multiformats/go-multihash"
)

var logger = logging.Logger("service/client/DelegatedRouting")

type Client struct {
	client proto.DelegatedRouting_Client
}

func NewClient(c proto.DelegatedRouting_Client) *Client {
	return &Client{client: c}
}

func (fp *Client) FindProviders(ctx context.Context, key cid.Cid) ([]peer.AddrInfo, error) {
	resps, err := fp.client.GetP2PProvide(ctx, cidsToGetP2PProvideRequest([]cid.Cid{key}))
	if err != nil {
		return nil, err
	}
	infos := []peer.AddrInfo{}
	for _, resp := range resps {
		infos = append(infos, parseP2PProvideResponseForKey(resp, key)...)
	}
	return infos, nil
}

type FindProvidersAsyncResult struct {
	AddrInfo []peer.AddrInfo
	Err      error
}

func (fp *Client) FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	ch0, err := fp.client.GetP2PProvide_Async(ctx, cidsToGetP2PProvideRequest([]cid.Cid{key}))
	if err != nil {
		return nil, err
	}
	ch1 := make(chan FindProvidersAsyncResult, 1)
	go func() {
		defer close(ch1)
		if ctx.Err() != nil {
			return
		}
		r0, ok := <-ch0
		if !ok {
			return
		}
		var r1 FindProvidersAsyncResult
		r1.Err = r0.Err
		if r0.Resp != nil {
			r1.AddrInfo = parseP2PProvideResponseForKey(r0.Resp, key)
		}
		ch1 <- r1
	}()
	return ch1, nil
}

func cidsToGetP2PProvideRequest(cids []cid.Cid) *proto.GetP2PProvideRequest {
	keys := make(proto.KeyList, len(cids))
	for i, cid := range cids {
		keys[i] = BuildProtoMultihashFromCid(cid)
	}
	return &proto.GetP2PProvideRequest{
		Keys: keys,
	}
}

func BuildProtoMultihashFromCid(c cid.Cid) proto.Multihash {
	return proto.Multihash{Bytes: c.Bytes()} // CID to bytes
}

type KeyProviders struct {
	Key       cid.Cid
	Providers []peer.AddrInfo
}

func parseP2PProvideResponseForKey(resp *proto.GetP2PProvideResponse, key cid.Cid) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	for _, kp := range parseP2PProvideResponse(resp) {
		if key.Equals(kp.Key) {
			infos = append(infos, kp.Providers...)
		}
	}
	return infos
}

func parseP2PProvideResponse(resp *proto.GetP2PProvideResponse) []KeyProviders {
	kp := []KeyProviders{}
	for _, prov := range resp.ProvidersByKey {
		c, err := ParseProtoCid(&prov.Key) // bytes to CID
		if err != nil {
			logger.Infof("cannot parse key cid (%w)", err)
			continue
		}
		kp = append(kp, KeyProviders{Key: c, Providers: parseProtoNodesToAddrInfo(prov.Provider.Nodes)})
	}
	return kp
}

func ParseProtoCid(p *proto.Multihash) (cid.Cid, error) {
	return cid.Cast(p.Bytes)
}

func ParseProtoMultihash(p *proto.Multihash) (multihash.Multihash, error) {
	mh := multihash.Multihash(p.Bytes)
	if _, err := multihash.Decode(mh); err != nil {
		return nil, fmt.Errorf("invalid multihash key (%w)", err)
	}
	return mh, nil
}

func parseProtoNodesToAddrInfo(nodes []proto.Node) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	for _, n := range nodes {
		if n.Peer == nil {
			continue
		}
		infos = append(infos, ParseNodeAddresses(n.Peer)...)
	}
	return infos
}

func ParseNodeAddresses(n *proto.Peer) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	for _, addrBytes := range n.Multiaddresses {
		ma, err := multiaddr.NewMultiaddrBytes(addrBytes)
		if err != nil {
			logger.Infof("cannot parse multiaddress (%w)", err)
			continue
		}
		ai, err := peer.AddrInfoFromP2pAddr(ma)
		if err != nil {
			logger.Infof("cannot parse peer from multiaddress (%w)", err)
			continue
		}
		infos = append(infos, *ai)
	}
	return infos
}
