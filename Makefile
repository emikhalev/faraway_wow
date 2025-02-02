CURBINDIR := $(CURDIR)/bin
LINTBIN := $(CURBINDIR)/golangci-lint

.PHONY: build
build:
	mkdir -p $(CURBINDIR)
	go build -o $(CURBINDIR) ./cmd/wow

.PHONY: run-local
run-local: build
	$(CURBINDIR)/wow --config=./config-local.yaml

.PHONY: run
run: build
	$(CURBINDIR)/wow --config=./config.yaml

PHONY: .install-lint-deps
.install-lint-deps:
	GOBIN=$(CURBINDIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

.PHONY: lint
lint: .install-lint-deps
	$(LINTBIN) run -c .golangci.yaml --new-from-rev origin/main


