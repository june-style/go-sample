################################################################################
# Go-sample
################################################################################

.PHONY: wire
wire:
	@ cd framework/registry/injector && wire .

.PHONY: init-tools
init-tools:
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
