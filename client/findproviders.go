package client

import (
	"context"

	"github.com/ipfs/go-cid"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

var logger = logging.Logger("service/client/delegatedrouting")

type DelegatedRoutingClient interface {
	FindProviders(ctx context.Context, key cid.Cid) ([]peer.AddrInfo, error)
	FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error)
	GetIPNS(ctx context.Context, id []byte) ([][]byte, error)
	PutIPNSAsync(ctx context.Context, id []byte, record []byte) (<-chan PutIPNSAsyncResult, error)
}

type Client struct {
	client proto.DelegatedRouting_Client
}

func NewClient(c proto.DelegatedRouting_Client) *Client {
	return &Client{client: c}
}

func (fp *Client) FindProviders(ctx context.Context, key cid.Cid) ([]peer.AddrInfo, error) {
	resps, err := fp.client.FindProviders(ctx, cidsToFindProvidersRequest(key))
	if err != nil {
		return nil, err
	}
	infos := []peer.AddrInfo{}
	for _, resp := range resps {
		infos = append(infos, parseFindProvidersResponse(resp)...)
	}
	return infos, nil
}

type FindProvidersAsyncResult struct {
	AddrInfo []peer.AddrInfo
	Err      error
}

// FindProvidersAsync processes the stream of raw protocol async results into a stream of parsed results.
// Specifically, FindProvidersAsync converts protocol-level provider descriptions into peer address infos.
func (fp *Client) FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	protoRespCh, err := fp.client.FindProviders_Async(ctx, cidsToFindProvidersRequest(key))
	if err != nil {
		return nil, err
	}
	parsedRespCh := make(chan FindProvidersAsyncResult, 1)
	go func() {
		defer close(parsedRespCh)
		if ctx.Err() != nil {
			return
		}
		protoAsyncResp, ok := <-protoRespCh
		if !ok {
			return
		}
		var parsedAsyncResp FindProvidersAsyncResult
		parsedAsyncResp.Err = protoAsyncResp.Err
		if protoAsyncResp.Resp != nil {
			parsedAsyncResp.AddrInfo = parseFindProvidersResponse(protoAsyncResp.Resp)
		}
		parsedRespCh <- parsedAsyncResp
	}()
	return parsedRespCh, nil
}

func cidsToFindProvidersRequest(cid cid.Cid) *proto.FindProvidersRequest {
	return &proto.FindProvidersRequest{
		Key: proto.LinkToAny(cid),
	}
}

func parseFindProvidersResponse(resp *proto.FindProvidersResponse) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	for _, prov := range resp.Providers {
		if !providerSupportsBitswap(prov.ProviderProto) {
			continue
		}
		infos = append(infos, parseProtoNodeToAddrInfo(prov.ProviderNode)...)
	}
	return infos
}

func providerSupportsBitswap(supported proto.TransferProtocolList) bool {
	for _, p := range supported {
		if p.Bitswap != nil {
			return true
		}
	}
	return false
}

func parseProtoNodeToAddrInfo(n proto.Node) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	if n.Peer == nil { // ignore non-peer nodes
		return nil
	}
	infos = append(infos, ParseNodeAddresses(n.Peer)...)
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
