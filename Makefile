.PHONY: init
init:
	go install mvdan.cc/gofumpt@latest
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest

.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -compat=1.23.0
	@echo "✅ Go modules tidied"

.PHONY: lint
lint: go-mod-tidy
	golangci-lint run
	@echo "✅ Linting completed"

.PHONY: fix
fix:
	golangci-lint run --fix
	@echo "✅ Lint fixing completed"

.PHONY: test
test:
	go test ./... -race
	@echo "✅ Testing completed"

.PHONY: fmt
fmt:
	gofmt -w -e "vendor" .
	@echo "✅ Formatting completed"

.PHONY: fumpt
fumpt:
	gofumpt -w -e "vendor" .
	@echo "✅ Formatting completed"

.PHONY: nilaway-install
nilaway-install:
	go install go.uber.org/nilaway/cmd/nilaway@latest

.PHONY: nilaway
nilaway:
	nilaway ./...
	@echo "✅ Nilaway completed"