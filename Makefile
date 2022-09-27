all: lint
.PHONY: all

generate:
	wails generate module
.PHONY: generate

build: generate
	wails build
.PHONY: build

install:
	go install github.com/wailsapp/wails/v2/cmd/wails@latest
	wails doctor
.PHONY: install

lint-ts:
	npm --prefix frontend run format
	npm --prefix frontend run lint
	npm --prefix frontend run types
.PHONY: lint-ts

lint-go:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: lint-go

lint: generate lint-go lint-ts
.PHONY: lint

format-fix-ts:
	npm --prefix frontend run format-fix
.PHONY: format-fix-ts
