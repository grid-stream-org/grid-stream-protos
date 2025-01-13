.PHONY: generate lint clean all

generate:
	buf generate

lint:
	buf lint

clean:
	rm -rf gen/*

all: lint clean generate
