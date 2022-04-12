package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/client"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	logging "github.com/ipfs/go-log"
	"github.com/ipld/edelweiss/values"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

var logger = logging.Logger("service/server/DelegatedRouting")

// FindProvidersAsyncHandler implements a higher-level interface to FindProviders, used in DHT and Hydra.

type FindProvidersAsyncFunc func(cid.Cid, chan<- client.FindProvidersAsyncResult) error

func FindProvidersAsyncHandler(f FindProvidersAsyncFunc) http.HandlerFunc {
	fps := &findProvidersServer{f}
	return proto.DelegatedRouting_AsyncHandler(fps)
}

type findProvidersServer struct {
	FindProvidersAsyncFunc
}

func (fps *findProvidersServer) GetIPNS(ctx context.Context, req *proto.GetIPNSRequest, respCh chan<- *proto.DelegatedRouting_GetIPNS_AsyncResult) error {
	return fmt.Errorf("GetIPNS not supported")
}

func (fps *findProvidersServer) PutIPNS(ctx context.Context, req *proto.PutIPNSRequest, respCh chan<- *proto.DelegatedRouting_PutIPNS_AsyncResult) error {
	return fmt.Errorf("PutIPNS not supported")
}

func (fps *findProvidersServer) FindProviders(ctx context.Context, req *proto.FindProvidersRequest, rch chan<- *proto.DelegatedRouting_FindProviders_AsyncResult) error {
	go func() {
		defer close(rch)
		pcids := parseCidsFromFindProvidersRequest(req)
		for _, c := range pcids {
			ch := make(chan client.FindProvidersAsyncResult)
			if err := fps.FindProvidersAsyncFunc(c, ch); err != nil {
				logger.Errorf("find providers function rejected request (%w)", err)
				continue
			}
			for x := range ch {
				if x.Err != nil {
					logger.Errorf("find providers function returned error (%w)", x.Err)
					continue
				}
				rch <- buildFindProvidersResponse(c, x.AddrInfo)
			}
		}
	}()
	return nil
}

func parseCidsFromFindProvidersRequest(req *proto.FindProvidersRequest) []cid.Cid {
	return []cid.Cid{cid.Cid(req.Key)}
}

func buildFindProvidersResponse(key cid.Cid, addrInfo []peer.AddrInfo) *proto.DelegatedRouting_FindProviders_AsyncResult {
	provs := make(proto.ProvidersList, len(addrInfo))
	bitswapProto := proto.TransferProtocol{Bitswap: &proto.BitswapProtocol{}}
	for i, addrInfo := range addrInfo {
		provs[i] = proto.Provider{
			ProviderNode:  proto.Node{Peer: buildPeerFromAddrInfo(addrInfo)},
			ProviderProto: proto.TransferProtocolList{bitswapProto},
		}
	}
	return &proto.DelegatedRouting_FindProviders_AsyncResult{
		Resp: &proto.FindProvidersResponse{Providers: provs},
	}
}

func buildPeerFromAddrInfo(addrInfo peer.AddrInfo) *proto.Peer {
	pm := make([]values.Bytes, len(addrInfo.Addrs))
	for i, addr := range addrInfo.Addrs {
		peerAddr := addr.Encapsulate(multiaddr.StringCast("/p2p/" + addrInfo.ID.String()))
		pm[i] = peerAddr.Bytes()
	}
	return &proto.Peer{
		ID:             []byte(addrInfo.ID),
		Multiaddresses: pm,
	}
}
