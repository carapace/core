package document

import (
	"encoding/json"
)

type Encoder interface {
	Encode(interface{}) ([]byte, error)
	Decode(interface{}, []byte) error
}

type JSONEncoder struct{}

func (j JSONEncoder) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func (j JSONEncoder) Decode(obj interface{}, data []byte) error {
	return json.Unmarshal(data, obj)
}
