.PHONY: api
api: ## Generate API from protobuf
	@protoc -I./api --go_out=plugins=grpc:api api/*.proto 
