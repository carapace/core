proto: ## Generate up to date protobuf code based on protobuf defintions in api/proto/defintions.
	protoc -I=./api/v1/proto --go_out=plugins=grpc:./api/v1/proto/generated api/v1/proto/definitions/*.proto
	mv api/v1/proto/generated/definitions/** api/v1/proto/generated
	rmdir api/v1/proto/generated/definitions

test:
	git log -1 --pretty=%B | gitlint
	golangci-lint run
	go test ./... -v -race