package ark

import (
	"encoding/json"
	"github.com/carapace/core/internal/services/asset"

	"github.com/carapace/core/api/v1/proto/generated"
)

// compile time assertion to verify we match interface internal/transactions.AssetService
var _ asset.AssetService = &Service{}

type Service struct {
	Config
}

//nolint: structcheck
type Config struct {
	FirstSecret  string
	SecondSecret string

	validator *Validator

	conf *asset.Config
}

type Option func(*Service)

func (s *Service) Configure(config *asset.Config) error {
	s.conf = config
	return nil
}

func (s Service) Create(params v1.Transaction) (*v1.TransactionResponse, error) {
	passphrases, err := s.conf.GetSecret(*params.Namespace, v1.Asset_Ark)
	if err != nil {
		return nil, err
	}

	tx, err := s.createTX(&params, passphrases)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(tx)
	if err != nil {
		return nil, err
	}

	return &v1.TransactionResponse{
		Payload:   payload,
		Namespace: params.Namespace,
	}, nil
}

func (s *Service) Validate(transaction v1.Transaction) error {
	return s.validator.Validate(transaction)
}

func New(config Config, opts ...Option) *Service {
	s := &Service{Config: config}
	for _, opt := range opts {
		opt(s)
	}
	return s
}
