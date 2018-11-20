package carapace

import (
	"bytes"
	"github.com/carapace/core/api/v0/proto"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/jsonpb"
	"io"
	"io/ioutil"
)

func (c *Client) ConfigFromYAML(reader io.Reader) (*v0.Config, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	json, err := yaml.YAMLToJSON(data)
	if err != nil {
		return nil, err
	}
	conf := &v0.Config{}
	conf.Reset()
	return conf, jsonpb.Unmarshal(bytes.NewReader(json), conf)
}
