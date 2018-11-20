package carapace

import (
	"encoding/json"
	"github.com/Jeffail/gabs"
	"github.com/carapace/core/api/v0/proto"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/ptypes"
)

func (c *Client) DumpString(config *v0.Config) (string, error) {
	yml, err := c.DumpYAML(config)
	if err != nil {
		return "", err
	}
	return string(yml), nil
}

func (c *Client) DumpYAML(config *v0.Config) ([]byte, error) {
	js, err := c.DumpJSON(config)
	if err != nil {
		return nil, err
	}
	yml, err := yaml.JSONToYAML(js)
	if err != nil {
		return nil, err
	}
	return yml, nil
}

func (c *Client) DumpJSON(config *v0.Config) ([]byte, error) {
	js, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	var spec interface{}
	switch config.Spec.TypeUrl {
	case OwnerSet:
		set := v0.OwnerSet{}
		err = ptypes.UnmarshalAny(config.Spec, &set)
		if err != nil {
			return nil, err
		}
		spec = set
	}

	jsonParsed, err := gabs.ParseJSON(js)
	if err != nil {
		return nil, err
	}

	_, err = jsonParsed.Set(spec, "Spec")
	if err != nil {
		return nil, err
	}

	return jsonParsed.EncodeJSON(), nil
}
