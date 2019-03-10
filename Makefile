.PHONY: api
api: ## Generate API from protobuf
	@protoc -I./api --go_out=plugins=grpc:api api/*.proto 

.PHONY: wire
wire: ## Resolve deps with wire
	wire ./internal/wire/...
