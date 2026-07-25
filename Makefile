VERSION ?= $(shell git describe --tags --dirty 2>/dev/null || echo "0.0.0-dev")
COMMIT ?= $(shell git rev-parse --short HEAD 2>/dev/null || echo "none")
LDFLAGS = -X "github.com/gosusnp/whackamole/internal.Version=$(VERSION)" -X "github.com/gosusnp/whackamole/internal.Commit=$(COMMIT)"

build: ui-build
	go build -ldflags '$(LDFLAGS)' -o whack ./cmd/whack

check: lint test

ci: fmt-check license-check lint test

clean:
	rm -f whack

deps: ui-deps
	go mod download

fix: fmt license-fix

fmt:
	gofmt -w .

fmt-check:
	@if [ -n "$$(gofmt -l .)" ]; then \
		echo "Files need formatting (run 'make fmt'):"; \
		gofmt -l .; \
		exit 1; \
	fi

install: ui-build
	go install -ldflags '$(LDFLAGS)' ./cmd/whack

license-check:
	go tool addlicense -check -l mit -c "Jimmy Ma" -s=only -ignore "frontend/node_modules/**" .

license-fix:
	go tool addlicense -l mit -c "Jimmy Ma" -s=only -ignore "frontend/node_modules/**" .

lint:
	go vet $$(go list ./... | grep -v '/frontend')
	go tool golangci-lint run ./...

pre-commit: fix check

pre-commit-all: pre-commit ui-pre-commit

test:
	go tool gotestsum --format pkgname-and-test-fails -- -cover $$(go list ./... | grep -v '/frontend')

ui-build:
	make -C frontend build

ui-deps:
	make -C frontend deps

ui-run:
	make -C frontend run

ui-pre-commit:
	make -C frontend pre-commit

ui-test:
	make -C frontend test
