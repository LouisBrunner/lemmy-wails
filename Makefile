all: lint
.PHONY: all

lint-ts:
	npm --prefix frontend run format
	npm --prefix frontend run lint
	npm --prefix frontend run types
.PHONY: lint-ts

lint-go:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: lint-go

lint: lint-go lint-ts
.PHONY: lint

format-fix-ts:
	npm --prefix frontend run format-fix
.PHONY: format-fix-ts
