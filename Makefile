.PHONY: generate
generate:
	buf generate

.PHONY: lint
lint:
	buf lint

.PHONY: all
all: lint generate