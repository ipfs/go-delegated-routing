package parser

type Method string
const (
	MethodGetP2PProvide = "get-p2p-provide"
)

type GetP2PProvide struct {
	Key DJSpecialBytes `json:"key"`
}

type GetP2PProvideResponse struct {
	Peers []DJSpecialBytes `json:"peers"`
}
