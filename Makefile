
TOOLS_MOD_DIR := ./internal/tools
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

ALL_GO_MOD_DIRS := $(shell find . -type f -name 'main.go' -exec dirname {} \; | sort)
ROOT_GO_MOD_DIRS := $(filter-out $(TOOLS_MOD_DIR), $(ALL_GO_MOD_DIRS))
ALL_COVERAGE_MOD_DIRS := $(shell find . -type f -name 'go.mod' -exec dirname {} \; | grep -E -v '^./example|^$(TOOLS_MOD_DIR)' | sort)

GO = go
TIMEOUT = 60

# Tools

TOOLS = $(CURDIR)/.tools

$(TOOLS):
	@mkdir -p $@
$(TOOLS)/%: $(TOOLS_MOD_DIR)/go.mod | $(TOOLS)
	cd $(TOOLS_MOD_DIR) && \
	$(GO) build -o $@ $(PACKAGE)


GOLANGCI_LINT = $(TOOLS)/golangci-lint
$(TOOLS)/golangci-lint: PACKAGE=github.com/golangci/golangci-lint/v2/cmd/golangci-lint

GORELEASE = $(TOOLS)/gorelease
$(GORELEASE): PACKAGE=golang.org/x/exp/cmd/gorelease

GOCOVMERGE = $(TOOLS)/gocovmerge
$(TOOLS)/gocovmerge: PACKAGE=github.com/wadey/gocovmerge

MULTIMOD = $(TOOLS)/multimod
$(TOOLS)/multimod: PACKAGE=go.opentelemetry.io/build-tools/multimod

.PHONY: tools
tools: $(GOLANGCI_LINT) $(GORELEASE) $(GOCOVMERGE) $(MULTIMOD)


# Build

.PHONY: build

build: $(ROOT_GO_MOD_DIRS:%=build/%)
build/%: DIR=$*
build/%:
	@echo "$(GO) build $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) build -o $(ROOT_DIR)/bin/ ./...

# Tests

TEST_TARGETS := test-default test-short test-verbose test-race test-concurrent-safe
.PHONY: $(TEST_TARGETS) test
test-default test-race: ARGS=-race
test-short:   ARGS=-short
test-verbose: ARGS=-v -race
test-concurrent-safe: ARGS=-run=ConcurrentSafe -count=100 -race
test-concurrent-safe: TIMEOUT=120
$(TEST_TARGETS): test
test: $(ROOT_GO_MOD_DIRS:%=test/%)
test/%: DIR=$*
test/%:
	@echo "$(GO) test -timeout $(TIMEOUT)s $(ARGS) $(DIR)/..." \
		&& cd $(DIR) \
		&& $(GO) list ./... \
		| xargs $(GO) test -timeout $(TIMEOUT)s $(ARGS)


COVERAGE_MODE    = atomic
COVERAGE_PROFILE = coverage.out
.PHONY: test-coverage
test-coverage: $(GOCOVMERGE)
	@set -e; \
	printf "" > coverage.txt; \
	for dir in $(ALL_COVERAGE_MOD_DIRS); do \
	  echo "$(GO) test -v -race -coverpkg=github.com/go-tapd/tapd/... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" $${dir}/..."; \
	  (cd "$${dir}" && \
	    $(GO) list ./... \
	    | xargs $(GO) test -coverpkg=./... -covermode=$(COVERAGE_MODE) -coverprofile="$(COVERAGE_PROFILE)" && \
	  $(GO) tool cover -html=coverage.out -o coverage.html); \
	done; \
	$(GOCOVMERGE) $$(find . -name coverage.out) > coverage.txt

.PHONY: golangci-lint golangci-lint-fix
golangci-lint-fix: ARGS=--fix
golangci-lint-fix: golangci-lint
golangci-lint: $(ROOT_GO_MOD_DIRS:%=golangci-lint/%)
golangci-lint/%: DIR=$*
golangci-lint/%: $(GOLANGCI_LINT)
	@echo 'golangci-lint $(if $(ARGS),$(ARGS) ,)$(DIR)' \
		&& cd $(DIR) \
		&& $(GOLANGCI_LINT) run --allow-serial-runners $(ARGS)

.PHONY: go-mod-tidy
go-mod-tidy: $(ALL_GO_MOD_DIRS:%=go-mod-tidy/%)
go-mod-tidy/%: DIR=$*
go-mod-tidy/%:
	@echo "$(GO) mod tidy in $(DIR)" \
		&& cd $(DIR) \
		&& $(GO) mod tidy -compat=1.22.0

.PHONY: lint-modules
lint-modules: go-mod-tidy

.PHONY: lint
lint: lint-modules golangci-lint

.PHONY: check-clean-work-tree
check-clean-work-tree:
	@if ! git diff --quiet; then \
	  echo; \
	  echo 'Working tree is not clean, did you forget to run "make precommit"?'; \
	  echo; \
	  git status; \
	  exit 1; \
	fi

.PHONY: gorelease
gorelease: $(ROOT_GO_MOD_DIRS:%=gorelease/%)
gorelease/%: DIR=$*
gorelease/%:| $(GORELEASE)
	@echo "gorelease in $(DIR):" \
		&& cd $(DIR) \
		&& $(GORELEASE) \
		|| echo ""

.PHONY: verify-mods
verify-mods: $(MULTIMOD)
	$(MULTIMOD) verify

.PHONY: prerelease
prerelease: verify-mods
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) prerelease -m ${MODSET}

COMMIT ?= "HEAD"
.PHONY: add-tags
add-tags: verify-mods
	@[ "${MODSET}" ] || ( echo ">> env var MODSET is not set"; exit 1 )
	$(MULTIMOD) tag -m ${MODSET} -c ${COMMIT}