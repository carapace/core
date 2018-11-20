package carapace

import (
	"bytes"
	"github.com/Jeffail/gabs"
	"github.com/carapace/core/api/v0/proto"
	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/jsonpb"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func (c *Client) LoadConfig(relPath string) (*v0.Config, error) {
	file, err := os.Open(relPath)
	if err != nil {
		return nil, err
	}

	yml, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	js, err := yaml.YAMLToJSON(yml)
	if err != nil {
		return nil, err
	}

	js, err = addType(js)
	if err != nil {
		return nil, err
	}

	config := v0.Config{}
	err = jsonpb.Unmarshal(bytes.NewReader(js), &config)
	if err != nil {
		return nil, errors.Errorf("jsonpb: %s", err.Error())
	}
	return &config, nil
}

func addType(json []byte) ([]byte, error) {
	jsonParsed, err := gabs.ParseJSON(json)
	if err != nil {
		return nil, err
	}

	kind, exists := jsonParsed.Path("Header.Kind").Data().(string)
	if !exists {
		return nil, errors.New("missing Kind field")
	}

	_, err = jsonParsed.Path("Spec").Set(kind, "@type")
	if err != nil {
		return nil, err
	}
	return jsonParsed.EncodeJSON(), nil
}
