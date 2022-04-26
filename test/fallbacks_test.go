package test

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-delegated-routing/client"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	"github.com/ipld/edelweiss/values"
	"github.com/multiformats/go-multihash"
)

func TestClientWithServerReturningUnknownValues(t *testing.T) {

	// start a server
	s := httptest.NewServer(proto.DelegatedRouting_AsyncHandler(testServiceWithUnknown{}))
	defer s.Close()

	// start a client
	q, err := proto.New_DelegatedRouting_Client(s.URL, proto.DelegatedRouting_Client_WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}
	c := client.NewClient(q)

	// verify no result arrive
	h, err := multihash.Sum([]byte("TEST"), multihash.SHA3, 4)
	if err != nil {
		t.Fatal(err)
	}

	infos, err := c.FindProviders(context.Background(), cid.NewCidV1(cid.Libp2pKey, h))
	if err != nil {
		t.Fatal(err)
	}
	if len(infos) != 0 {
		t.Fatalf("expecting 0 result, got %d", len(infos))
	}
}

type testServiceWithUnknown struct{}

func (testServiceWithUnknown) FindProviders(ctx context.Context, req *proto.FindProvidersRequest, respCh chan<- *proto.DelegatedRouting_FindProviders_AsyncResult) error {
	go func() {
		defer close(respCh)
		respCh <- &proto.DelegatedRouting_FindProviders_AsyncResult{
			Resp: &proto.FindProvidersResponse{
				Providers: proto.ProvidersList{
					proto.Provider{
						ProviderNode: proto.Node{
							DefaultKey:   "UnknownNode",
							DefaultValue: &values.Any{Value: values.String("UnknownNodeValue")},
						},
						ProviderProto: proto.TransferProtocolList{
							proto.TransferProtocol{
								DefaultKey:   "UnknownProtocol",
								DefaultValue: &values.Any{Value: values.String("UnknownProtocolValue")},
							},
						},
					},
				},
			},
		}
	}()
	return nil
}

func (testServiceWithUnknown) GetIPNS(ctx context.Context, req *proto.GetIPNSRequest, respCh chan<- *proto.DelegatedRouting_GetIPNS_AsyncResult) error {
	return fmt.Errorf("GetIPNS not supported by test service")
}

func (testServiceWithUnknown) PutIPNS(ctx context.Context, req *proto.PutIPNSRequest, respCh chan<- *proto.DelegatedRouting_PutIPNS_AsyncResult) error {
	return fmt.Errorf("PutIPNS not supported by test service")
}
