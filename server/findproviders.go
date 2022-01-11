package server

import (
	"context"
	"net/http"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/client"
	proto "github.com/ipfs/go-delegated-routing/ipld/ipldsch"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// FindProvidersAsyncHandler implements a higher-level interface to GetP2PProvide, used in DHT and Hydra.

type FindProvidersAsyncFunc func(cid.Cid, chan<- client.FindProvidersAsyncResult) error

func FindProvidersAsyncHandler(f FindProvidersAsyncFunc) http.HandlerFunc {
	fps := &findProvidersServer{f}
	return Server_AsyncHandler(fps)
}

type findProvidersServer struct {
	FindProvidersAsyncFunc
}

func (fps *findProvidersServer) GetP2PProvide(ctx context.Context, req *proto.GetP2PProvideRequest, rch chan<- *proto.GetP2PProvideResponse) error {
	go func() {
		defer close(rch)
		pcids := parseCidsFromGetP2PProvideRequest(req)
		for _, c := range pcids {
			ch := make(chan client.FindProvidersAsyncResult)
			if err := fps.FindProvidersAsyncFunc(c, ch); err != nil {
				log.Errorf("find providers function rejected request (%w)", err)
				continue
			}
			for x := range ch {
				if x.Err != nil {
					log.Errorf("find providers function returned error (%w)", x.Err)
					continue
				}
				rch <- buildGetP2PProvideResponse(c, x.AddrInfo)
			}
		}
	}()
	return nil
}

func parseCidsFromGetP2PProvideRequest(req *proto.GetP2PProvideRequest) []cid.Cid {
	cids := []cid.Cid{}
	for _, key := range req.Keys {
		c, err := client.ParseProtoCid(&key)
		if err != nil {
			continue
		}
		cids = append(cids, c)
	}
	return cids
}

func buildGetP2PProvideResponse(key cid.Cid, addrInfo []peer.AddrInfo) *proto.GetP2PProvideResponse {
	nodes := make(proto.List__Node, len(addrInfo))
	for i, addrInfo := range addrInfo {
		nodes[i] = proto.Node{Peer: buildPeerFromAddrInfo(addrInfo)}
	}
	return &proto.GetP2PProvideResponse{
		ProvidersByKey: proto.List__ProvidersByKey{
			proto.ProvidersByKey{
				Key: *client.BuildProtoMultihashFromCid(key),
				Provider: proto.Provider{
					Node: nodes,
				},
			},
		},
	}
}

func buildPeerFromAddrInfo(addrInfo peer.AddrInfo) *proto.Peer {
	pm := make(proto.List__Bytes, len(addrInfo.Addrs))
	for i, addr := range addrInfo.Addrs {
		peerAddr := addr.Encapsulate(multiaddr.StringCast("/p2p/" + addrInfo.ID.String()))
		pm[i] = peerAddr.Bytes()
	}
	return &proto.Peer{
		ID:             []uint8(addrInfo.ID),
		Multiaddresses: pm,
	}
}
