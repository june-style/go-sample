################################################################################
# Go-sample
################################################################################

.PHONY: db
db:
	@ go run tools/schemas/dynamodb/*.go

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
	@ docker exec -it go /bin/bash
