package carapace

import (
	"github.com/carapace/core/api/v0/proto"
)

type Client struct {
	Config
	v0.CoreServiceClient
}

func New(cfg Config) *Client {
	return &Client{
		Config: cfg,
	}
}

func (c *Client) PrivateKey() ([]byte, error) {
	b, err := c.Marshaller.MarshalPrivate(c.PrivKey)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) PublicKey() ([]byte, error) {
	b, err := c.Marshaller.MarshalPublic(&c.PrivKey.PublicKey)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) RPrivateKey() ([]byte, error) {
	b, err := c.Marshaller.MarshalPrivate(c.RecoveryPrivKey)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (c *Client) RPublicKey() ([]byte, error) {
	b, err := c.Marshaller.MarshalPublic(&c.RecoveryPrivKey.PublicKey)
	if err != nil {
		return nil, err
	}
	return b, nil
}
