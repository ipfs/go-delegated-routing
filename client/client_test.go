package client

import (
	"strings"
	"testing"
)

func TestParseFindProvsResp(t *testing.T) {

	// These multiaddrs are incorrect, they're base64padded text multiaddrs instead of base64-unpadded byte multiaddrs
	respStr := `{"tag" : "get-p2p-provide",
"payload" : {
	"peers" : [
		{"/" : {"bytes" : "L2lwNC80Ni4xNy45Ni45OS90Y3AvMzA2MDIvcDJwL1FtVzltNTdhaUJESEFrS2o5bm1GU0VuN1pxcmNGMWZaUzRiaXBzVENIYnVyZWk="}},
		{"/" : {"bytes" : "L2lwNC80Ni4xNy45Ni45OS90Y3AvMzA2MDIvcDJwL1FtVzltNTdhaUJESEFrS2o5bm1GU0VuN1pxcmNGMWZaUzRiaXBzVENIYnVyZWk="}}
	]
}
}
`
	r := strings.NewReader(respStr)
	ais, err := processProvResp(r)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ais)
}

/* WIP
func TestSimpleServer(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		dec := json.NewDecoder(request.Body)
		env := parser.Envelope{Payload: &parser.GetP2PProvideResponse{}}
		err := dec.Decode(&env)
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			writer.WriteHeader(500)
			return
		}

		switch env.Tag {
		case parser.MethodGetP2PProvide:
		default:
			writer.WriteHeader(404)
			return
		}

		env.Payload.
	}))
	defer s.Close()

	drc, err := New(s.URL, WithHTTPClient(s.Client()))
	if err != nil {
		t.Fatal(err)
	}
}
 */