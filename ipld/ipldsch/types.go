package ipldsch

type ServiceEnvelope struct {
	GetP2PProvideRequest  *GetP2PProvideRequest
	GetP2PProvideResponse *GetP2PProvideResponse
}
type Multihash struct {
	Bytes []uint8
}
type List__Multihash []Multihash
type GetP2PProvideRequest struct {
	Keys List__Multihash
}
type ProvidersByKey struct {
	Key      Multihash
	Provider Provider
}
type List__ProvidersByKey []ProvidersByKey
type GetP2PProvideResponse struct {
	ProvidersByKey List__ProvidersByKey
}
type List__Node []Node
type List__TransferProtocol []TransferProtocol
type Provider struct {
	Node  List__Node
	Proto List__TransferProtocol
}
type Node struct {
	Peer *Peer
}
type List__Bytes [][]uint8
type Peer struct {
	ID             []uint8
	Multiaddresses List__Bytes
}
type TransferProtocol struct {
	BitswapTransfer *BitswapTransfer
}
type BitswapTransfer struct {
}
