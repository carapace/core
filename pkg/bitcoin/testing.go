package bitcoin

func newService() *Service {
	return &Service{}
}

func newTx() *Transaction {
	return &ArkV2.Transaction{
		Id:          "dummy",
		Type:        0,
		Amount:      10000000,
		Fee:         10000000,
		VendorField: "dummy",
		Timestamp:   ArkV2.GetTime(),
	}
}

func newJSONtx() ([]byte, error) {
	return json.Marshal(newTx())
}
