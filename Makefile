NPM = npm --prefix frontend

all: lint test
.PHONY: all

install:
	go download
	$(NPM) i
.PHONY: install

lint-ts:
	$(NPM) run format
	$(NPM) run lint
	$(NPM) run types
.PHONY: lint-ts

lint-go:
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: lint-go

lint: lint-go lint-ts
.PHONY: lint

format-fix-ts:
	$(NPM) run format-fix
.PHONY: format-fix-ts

test:
.PHONY: test

clean:
	rm -rf frontend/dist
.PHONY: clean
