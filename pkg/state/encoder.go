package state

import (
	"encoding/json"
)

type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

type JSONEncoder struct{}

func (j JSONEncoder) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}
