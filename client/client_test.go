package client

import (
	"context"
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
	ch := make(chan FindProvidersAsyncResult, 2)
	processFindProvidersAsyncResp(context.Background(), ch, r)
	p1, ok := <-ch
	if !ok {
		t.Fatalf("expecting 1st provider")
	}
	if p1.Err != nil {
		t.Fatal(p1.Err)
	}
	t.Log(p1.AddrInfo)
}
