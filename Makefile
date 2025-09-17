GO = go
GOLANGCI_LINT = $(GO) tool golangci-lint

.PHONY: go-mod-tidy
go-mod-tidy:
	@echo "go mod tidy in all modules" && \
		$(GO) mod tidy -compat=1.24.0

.PHONY: lint
lint: go-mod-tidy
	@echo "Starting linting..." && \
		$(GOLANGCI_LINT) run --concurrency=4 --allow-serial-runners $(ARGS)
lint-fix: ARGS=--fix
lint-fix: lint
	@echo "✅ Lint fixing completed"

.PHONY: test
test:
	go test ./... -race
	@echo "✅ Testing completed"