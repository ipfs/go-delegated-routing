package server

// NOTE: This file will be codegen'd by a protocol compiler.

import (
	"context"
	"net/http"

	proto "github.com/ipfs/go-delegated-routing/ipld/ipldsch"
	logging "github.com/ipfs/go-log"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/codec/dagjson"
)

var log = logging.Logger("delegated/server")

type Server interface {
	GetP2PProvide(context.Context, *proto.GetP2PProvideRequest, chan<- *proto.GetP2PProvideResponse) error
}

func Server_AsyncHandler(s Server) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// parse request
		msg := request.URL.Query().Get("q")
		env := &proto.ServiceEnvelope{}
		_, err := ipld.Unmarshal([]byte(msg), dagjson.Decode, env, proto.Prototypes.ServiceEnvelope.Type())
		if err != nil {
			log.Errorf("received request not decodeable (%w)", err)
			writer.WriteHeader(400)
			return
		}

		// demultiplex request
		switch {
		case env.GetP2PProvideRequest != nil:
			ch := make(chan *proto.GetP2PProvideResponse)
			if err = s.GetP2PProvide(context.TODO(), env.GetP2PProvideRequest, ch); err != nil {
				log.Errorf("get p2p provider rejected request (%w)", err)
				writer.WriteHeader(500)
				return
			}
			for resp := range ch {
				env := &proto.ServiceEnvelope{
					GetP2PProvideResponse: resp,
				}
				buf, err := ipld.Marshal(dagjson.Encode, env, proto.Prototypes.ServiceEnvelope.Type())
				if err != nil {
					log.Errorf("cannot encode response (%w)", err)
					continue
				}
				writer.Write(buf)
			}

		default:
			log.Errorf("no request")
			writer.WriteHeader(404)
		}
	}
}
