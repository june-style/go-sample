################################################################################
# Go-sample
################################################################################

.PHONY: db
db:
	@ go run tools/schemas/dynamodb/*.go

.PHONY: proto
proto:
	@ cd framework/protocol && protoc \
		-I proto \
		--go_out ./pb --go_opt paths=source_relative \
		--go-grpc_out ./pb --go-grpc_opt paths=source_relative \
		--grpc-gateway_out ./pb --grpc-gateway_opt paths=source_relative \
		proto/*.proto

.PHONY: wire
wire:
	@ cd framework/registry/injector && wire .

.PHONY: init
init:
	# go test
	@ go install github.com/cweill/gotests/gotests@v1.6.0
	# mockgen
	@ go install github.com/golang/mock/mockgen@v1.6.0
	# wire
	@ go install github.com/google/wire/cmd/wire@latest
	# protoc
	@ go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@ go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	# grpcurl
	@ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	# post process
	@ go mod tidy

.PHONY: dotenv
dotenv:
	@ sh ./make-dotenv.sh

.PHONY: grpc-list
grpc-list:
	@ grpcurl -plaintext local-api:9000 list


################################################################################
# Docker
################################################################################

.PHONY: compose-build
compose-build:
	@ docker-compose -f docker-compose.yml build --no-cache --force-rm

.PHONY: compose-up
compose-up:
	@ docker-compose -f docker-compose.yml up -d

.PHONY: compose-down
compose-down:
	@ docker-compose -f docker-compose.yml down

.PHONY: login-devcontainer
login-devcontainer:
	@ docker-compose -f docker-compose.yml exec -it go-development /bin/bash
