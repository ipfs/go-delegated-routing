package scribble

import (
	"testing"

	proto "github.com/ipfs/go-delegated-routing/ipld/ipldsch"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

func TestBindnodeEncDec(t *testing.T) {
	v := &proto.GetP2PProvideRequest{}
	// encode
	buf, err := ipld.Marshal(dagjson.Encode, v, proto.Prototypes.GetP2PProvideRequest.Type())
	if err != nil {
		t.Fatal(err)
	}
	// decode
	w := &proto.GetP2PProvideRequest{}
	_, err = ipld.Unmarshal(buf, dagjson.Decode, w, proto.Prototypes.GetP2PProvideRequest.Type())
	if err != nil {
		t.Fatal(err)
	}
}
