.PHONY: generate lint all

generate:
	buf generate

lint:
	buf lint

all: lint generate
