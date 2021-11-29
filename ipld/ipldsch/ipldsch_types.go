package ipldsch

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	"github.com/ipld/go-ipld-prime/datamodel"
)

var _ datamodel.Node = nil // suppress errors when this dependency is not referenced
// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		ipldsch.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		ipldsch.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	Bool                        _Bool__Prototype
	Bool__Repr                  _Bool__ReprPrototype
	Bytes                       _Bytes__Prototype
	Bytes__Repr                 _Bytes__ReprPrototype
	Envelope                    _Envelope__Prototype
	Envelope__Repr              _Envelope__ReprPrototype
	Float                       _Float__Prototype
	Float__Repr                 _Float__ReprPrototype
	GetP2PProvideRequest        _GetP2PProvideRequest__Prototype
	GetP2PProvideRequest__Repr  _GetP2PProvideRequest__ReprPrototype
	GetP2PProvideResponse       _GetP2PProvideResponse__Prototype
	GetP2PProvideResponse__Repr _GetP2PProvideResponse__ReprPrototype
	Int                         _Int__Prototype
	Int__Repr                   _Int__ReprPrototype
	Link                        _Link__Prototype
	Link__Repr                  _Link__ReprPrototype
	List__Bytes                 _List__Bytes__Prototype
	List__Bytes__Repr           _List__Bytes__ReprPrototype
	List__Provider              _List__Provider__Prototype
	List__Provider__Repr        _List__Provider__ReprPrototype
	Peer                        _Peer__Prototype
	Peer__Repr                  _Peer__ReprPrototype
	Provider                    _Provider__Prototype
	Provider__Repr              _Provider__ReprPrototype
	String                      _String__Prototype
	String__Repr                _String__ReprPrototype
}

// --- type definitions follow ---

// Bool matches the IPLD Schema type "Bool".  It has bool kind.
type Bool = *_Bool
type _Bool struct{ x bool }

// Bytes matches the IPLD Schema type "Bytes".  It has bytes kind.
type Bytes = *_Bytes
type _Bytes struct{ x []byte }

// Envelope matches the IPLD Schema type "Envelope".
// Envelope has union typekind, which means its data model behaviors are that of a map kind.
type Envelope = *_Envelope
type _Envelope struct {
	tag uint
	x1  _GetP2PProvideRequest
	x2  _GetP2PProvideResponse
}
type _Envelope__iface interface {
	_Envelope__member()
}

func (_GetP2PProvideRequest) _Envelope__member()  {}
func (_GetP2PProvideResponse) _Envelope__member() {}

// Float matches the IPLD Schema type "Float".  It has float kind.
type Float = *_Float
type _Float struct{ x float64 }

// GetP2PProvideRequest matches the IPLD Schema type "GetP2PProvideRequest".  It has struct type-kind, and may be interrogated like map kind.
type GetP2PProvideRequest = *_GetP2PProvideRequest
type _GetP2PProvideRequest struct {
	key _Bytes
}

// GetP2PProvideResponse matches the IPLD Schema type "GetP2PProvideResponse".  It has struct type-kind, and may be interrogated like map kind.
type GetP2PProvideResponse = *_GetP2PProvideResponse
type _GetP2PProvideResponse struct {
	providers _List__Provider
}

// Int matches the IPLD Schema type "Int".  It has int kind.
type Int = *_Int
type _Int struct{ x int64 }

// Link matches the IPLD Schema type "Link".  It has link kind.
type Link = *_Link
type _Link struct{ x datamodel.Link }

// List__Bytes matches the IPLD Schema type "List__Bytes".  It has list kind.
type List__Bytes = *_List__Bytes
type _List__Bytes struct {
	x []_Bytes
}

// List__Provider matches the IPLD Schema type "List__Provider".  It has list kind.
type List__Provider = *_List__Provider
type _List__Provider struct {
	x []_Provider
}

// Peer matches the IPLD Schema type "Peer".  It has struct type-kind, and may be interrogated like map kind.
type Peer = *_Peer
type _Peer struct {
	Multiaddress _List__Bytes
}

// Provider matches the IPLD Schema type "Provider".
// Provider has union typekind, which means its data model behaviors are that of a map kind.
type Provider = *_Provider
type _Provider struct {
	tag uint
	x1  _Peer
}
type _Provider__iface interface {
	_Provider__member()
}

func (_Peer) _Provider__member() {}

// String matches the IPLD Schema type "String".  It has string kind.
type String = *_String
type _String struct{ x string }
