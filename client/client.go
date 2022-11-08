package client

import (
	"context"
	"errors"
	"time"

	"github.com/ipfs/go-cid"
	proto "github.com/ipfs/go-delegated-routing/gen/proto"
	ipns "github.com/ipfs/go-ipns"
	logging "github.com/ipfs/go-log/v2"
	record "github.com/libp2p/go-libp2p-record"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

var logger = logging.Logger("service/client/delegatedrouting")

type DelegatedRoutingClient interface {
	FindProviders(ctx context.Context, key cid.Cid) ([]peer.AddrInfo, error)
	FindProvidersAsync(ctx context.Context, key cid.Cid) (<-chan FindProvidersAsyncResult, error)
	GetIPNS(ctx context.Context, id []byte) ([]byte, error)
	GetIPNSAsync(ctx context.Context, id []byte) (<-chan GetIPNSAsyncResult, error)
	PutIPNS(ctx context.Context, id []byte, record []byte) error
	PutIPNSAsync(ctx context.Context, id []byte, record []byte) (<-chan PutIPNSAsyncResult, error)
	Provide(ctx context.Context, key []cid.Cid, ttl time.Duration) (time.Duration, error)
	ProvideAsync(ctx context.Context, key []cid.Cid, ttl time.Duration) (<-chan time.Duration, error)
}

type Client struct {
	client    proto.DelegatedRouting_Client
	validator record.Validator

	provider          *Provider
	providerBatchSize int
	identity          crypto.PrivKey
}

type Config struct {
	Provider         *Provider
	Identity         crypto.PrivKey
	ProvideBatchSize int
}

// Validate validates the current config, and also sets defaults values when needed.
func (c *Config) Validate() error {
	if c.Provider != nil && !c.Provider.Peer.ID.MatchesPublicKey(c.Identity.GetPublic()) {
		return errors.New("identity does not match provider")
	}

	if c.ProvideBatchSize == 0 {
		c.ProvideBatchSize = 30000 // this will generate payloads of ~1MB in size
	}

	return nil
}

var _ DelegatedRoutingClient = (*Client)(nil)

// NewClient creates a client.
// The Provider and identity parameters are option. If they are nil, the `Provide` method will not function.
func NewClient(c proto.DelegatedRouting_Client, cfg *Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{
		client:            c,
		validator:         ipns.Validator{},
		provider:          cfg.Provider,
		providerBatchSize: cfg.ProvideBatchSize,
		identity:          cfg.Identity,
	}, nil
}
