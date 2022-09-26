all: lint
.PHONY: all

lint-ts:
	npm --prefix frontend run types
.PHONY: lint-ts

lint-go:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: lint-go

lint: lint-go lint-ts
.PHONY: lint
