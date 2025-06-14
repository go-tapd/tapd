.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -compat=1.23.0
	@echo "✅ Go modules tidied"

.PHONY: lint
lint: go-mod-tidy
	go tool golangci-lint run
	@echo "✅ Linting completed"

.PHONY: fix
fix:
	go tool golangci-lint run --fix
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
	go tool gofumpt -w -e "vendor" .
	@echo "✅ Formatting completed"

.PHONY: nilaway
nilaway:
	go tool nilaway ./...
	@echo "✅ Nilaway completed"