package bitcoin

import (
	"bytes"
	"github.com/carapace/core/api/v1/proto/generated"
	"github.com/carapace/core/internal/signing"
	"github.com/carapace/core/internal/transactions"
	"io"
)

// compile time assertion to verify we match interface internal/transactions.AssetService
var _ signing.AssetService = &Service{}
var _ transactions.AssetService = &Service{}

type Service struct {
	Config
}

//nolint: structcheck
type Config struct {
	Secret string

	client    v1.NodeGatewayServiceClient
	validator *Validator
}

func (s Service) Sign(reader io.Reader) (io.Reader, error) {
	tx, err := s.parse(reader)
	if err != nil {
		return nil, err
	}
	s.sign(tx)

	js, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(js), nil
}

func (s Service) Validate(reader io.Reader) error {
	tx, err := s.parse(reader)
	if err != nil {
		return err
	}

	return s.validator.Validate(tx)
}

func (s Service) Verify(reader io.Reader) (bool, error) {
	tx, err := s.parse(reader)
	if err != nil {
		return false, err
	}
	return s.verify(tx)
}

func (s Service) Create(ctx context.Context, params *v1.Transaction) (*v1.TransactionResponse, error) {
	return s.createTX(params)
}

func New(config Config, opts ...Option) *Service {
	s := &Service{Config: config}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
