package bitcoin

import (
	"encoding/json"
	"io"
)

func (s Service) parse(reader io.Reader) (*Transaction, error) {
	tx := &Transaction{}
	return tx, json.NewDecoder(reader).Decode(tx)
}
