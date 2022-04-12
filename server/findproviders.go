package server

import (
	"context"
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

type DelegatedRoutingService interface {
	FindProviders(key cid.Cid, rch chan<- client.FindProvidersAsyncResult) error
	GetIPNS(id []byte, rch chan<- client.GetIPNSAsyncResult) error
	PutIPNS(id []byte, record []byte, rch chan<- client.PutIPNSAsyncResult) error
}

func DelegatedRoutingAsyncHandler(svc DelegatedRoutingService) http.HandlerFunc {
	drs := &delegatedRoutingServer{svc}
	return proto.DelegatedRouting_AsyncHandler(drs)
}

type delegatedRoutingServer struct {
	service DelegatedRoutingService
}

func (drs *delegatedRoutingServer) GetIPNS(ctx context.Context, req *proto.GetIPNSRequest, rch chan<- *proto.DelegatedRouting_GetIPNS_AsyncResult) error {
	go func() {
		defer close(rch)
		id := req.ID
		ch := make(chan client.GetIPNSAsyncResult)
		if err := drs.service.GetIPNS(id, ch); err != nil {
			logger.Errorf("get ipns function rejected request (%w)", err)
			return
		}
		for x := range ch {
			var resp *proto.DelegatedRouting_GetIPNS_AsyncResult
			if x.Err != nil {
				logger.Errorf("get ipns function returned error (%w)", x.Err)
				resp = &proto.DelegatedRouting_GetIPNS_AsyncResult{Err: x.Err}
			} else {
				resp = &proto.DelegatedRouting_GetIPNS_AsyncResult{Resp: &proto.GetIPNSResponse{Record: x.Record}}
			}
			rch <- resp
		}
	}()
	return nil
}

func (drs *delegatedRoutingServer) PutIPNS(ctx context.Context, req *proto.PutIPNSRequest, rch chan<- *proto.DelegatedRouting_PutIPNS_AsyncResult) error {
	go func() {
		defer close(rch)
		id, record := req.ID, req.Record
		ch := make(chan client.PutIPNSAsyncResult)
		if err := drs.service.PutIPNS(id, record, ch); err != nil {
			logger.Errorf("put ipns function rejected request (%w)", err)
			return
		}
		for x := range ch {
			var resp *proto.DelegatedRouting_PutIPNS_AsyncResult
			if x.Err != nil {
				logger.Errorf("put ipns function returned error (%w)", x.Err)
				resp = &proto.DelegatedRouting_PutIPNS_AsyncResult{Err: x.Err}
			} else {
				resp = &proto.DelegatedRouting_PutIPNS_AsyncResult{Resp: &proto.PutIPNSResponse{}}
			}
			rch <- resp
		}
	}()
	return nil
}

func (drs *delegatedRoutingServer) FindProviders(ctx context.Context, req *proto.FindProvidersRequest, rch chan<- *proto.DelegatedRouting_FindProviders_AsyncResult) error {
	go func() {
		defer close(rch)
		pcids := parseCidsFromFindProvidersRequest(req)
		for _, c := range pcids {
			ch := make(chan client.FindProvidersAsyncResult)
			if err := drs.service.FindProviders(c, ch); err != nil {
				logger.Errorf("find providers function rejected request (%w)", err)
				continue
			}
			for x := range ch {
				var resp *proto.DelegatedRouting_FindProviders_AsyncResult
				if x.Err != nil {
					logger.Errorf("find providers function returned error (%w)", x.Err)
					resp = &proto.DelegatedRouting_FindProviders_AsyncResult{Err: x.Err}
				} else {
					resp = buildFindProvidersResponse(c, x.AddrInfo)
				}
				rch <- resp
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
