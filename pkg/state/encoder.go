package state

import (
	"encoding/json"
)

// Encoder defines a common interface for encoding data
type Encoder interface {
	Encode(interface{}) ([]byte, error)
}

// Decoder defines a common interface for decoding data
type Decoder interface {
	Decode([]byte, interface{}) error
}

// JSONEncoder is an Encoder based on the json stdlib
type JSONEncoder struct{}

// Encode returns the JSON encodes byte array of an object
func (j JSONEncoder) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// JSONDecoder is an Decoder based on the json stdlib
type JSONDecoder struct{}

// Decode decodes a JSON byte array into a pointer.
func (j JSONDecoder) Decode(data []byte, ptr interface{}) error {
	return json.Unmarshal(data, ptr)
}
