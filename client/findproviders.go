package client

import (
	"context"

	"github.com/ipfs/go-cid"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	logging "github.com/ipfs/go-log"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

var logger = logging.Logger("service/client/DelegatedRouting")

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

func (fp *Client) FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error) {
	ch0, err := fp.client.FindProviders_Async(ctx, cidsToFindProvidersRequest(key))
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
			r1.AddrInfo = parseFindProvidersResponse(r0.Resp)
		}
		ch1 <- r1
	}()
	return ch1, nil
}

func cidsToFindProvidersRequest(cid cid.Cid) *proto.FindProvidersRequest {
	return &proto.FindProvidersRequest{
		Key: proto.LinkToAny(cid),
	}
}

func parseFindProvidersResponse(resp *proto.FindProvidersResponse) []peer.AddrInfo {
	infos := []peer.AddrInfo{}
	for _, prov := range resp.Providers {
		infos = append(infos, parseProtoNodeToAddrInfo(prov.ProviderNode)...)
	}
	return infos
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
