package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io"
)

type Parser interface {
	Parse(io.Reader) (interface{}, error)
}

type DefaultParser struct {
}

func (d DefaultParser) Parse(reader io.Reader) (Header, func(interface{}) error, error) {
	decoder := yaml.NewDecoder(reader)

	var header Header
	err := decoder.Decode(&header)
	if err != nil {
		return header, nil, errors.Wrap(err, "unable to read configuration file")
	}

	return header, func(i interface{}) error {
		return decoder.Decode(i)
	}, nil
}
