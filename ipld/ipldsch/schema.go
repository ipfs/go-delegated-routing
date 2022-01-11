package ipldsch

import (
	_ "embed"

	"github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/schema"
)

//go:embed routing.ipldsch
var embeddedSchema []byte

var Prototypes schemaSlab

type schemaSlab struct {
	ServiceEnvelope        schema.TypedPrototype
	Multihash              schema.TypedPrototype
	List__Multihash        schema.TypedPrototype
	GetP2PProvideRequest   schema.TypedPrototype
	ProvidersByKey         schema.TypedPrototype
	List__ProvidersByKey   schema.TypedPrototype
	GetP2PProvideResponse  schema.TypedPrototype
	List__Node             schema.TypedPrototype
	List__TransferProtocol schema.TypedPrototype
	Provider               schema.TypedPrototype
	Node                   schema.TypedPrototype
	List__Bytes            schema.TypedPrototype
	Peer                   schema.TypedPrototype
	TransferProtocol       schema.TypedPrototype
	BitswapTransfer        schema.TypedPrototype
}

func init() {
	ts, err := ipld.LoadSchemaBytes(embeddedSchema)
	if err != nil {
		panic(err)
	}

	Prototypes.ServiceEnvelope = bindnode.Prototype(
		(*ServiceEnvelope)(nil),
		ts.TypeByName("ServiceEnvelope"),
	)

	Prototypes.Multihash = bindnode.Prototype(
		(*Multihash)(nil),
		ts.TypeByName("Multihash"),
	)

	Prototypes.List__Multihash = bindnode.Prototype(
		(*List__Multihash)(nil),
		ts.TypeByName("List__Multihash"),
	)

	Prototypes.GetP2PProvideRequest = bindnode.Prototype(
		(*GetP2PProvideRequest)(nil),
		ts.TypeByName("GetP2PProvideRequest"),
	)

	Prototypes.ProvidersByKey = bindnode.Prototype(
		(*ProvidersByKey)(nil),
		ts.TypeByName("ProvidersByKey"),
	)

	Prototypes.List__ProvidersByKey = bindnode.Prototype(
		(*List__ProvidersByKey)(nil),
		ts.TypeByName("List__ProvidersByKey"),
	)

	Prototypes.GetP2PProvideResponse = bindnode.Prototype(
		(*GetP2PProvideResponse)(nil),
		ts.TypeByName("GetP2PProvideResponse"),
	)

	Prototypes.List__Node = bindnode.Prototype(
		(*List__Node)(nil),
		ts.TypeByName("List__Node"),
	)

	Prototypes.List__TransferProtocol = bindnode.Prototype(
		(*List__TransferProtocol)(nil),
		ts.TypeByName("List__TransferProtocol"),
	)

	Prototypes.Provider = bindnode.Prototype(
		(*Provider)(nil),
		ts.TypeByName("Provider"),
	)

	Prototypes.Node = bindnode.Prototype(
		(*Node)(nil),
		ts.TypeByName("Node"),
	)

	Prototypes.List__Bytes = bindnode.Prototype(
		(*List__Bytes)(nil),
		ts.TypeByName("List__Bytes"),
	)

	Prototypes.Peer = bindnode.Prototype(
		(*Peer)(nil),
		ts.TypeByName("Peer"),
	)

	Prototypes.TransferProtocol = bindnode.Prototype(
		(*TransferProtocol)(nil),
		ts.TypeByName("TransferProtocol"),
	)

	Prototypes.BitswapTransfer = bindnode.Prototype(
		(*BitswapTransfer)(nil),
		ts.TypeByName("BitswapTransfer"),
	)
}
