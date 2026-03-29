# Repository Guidelines

## Project Structure & Module Organization
The repository is a Go SDK for TAPD. Core client logic lives in the root
package: `client*.go`, `request.go`, `response.go`, and `helpers.go`. API
services are grouped by resource in files such as `api_story.go`,
`api_bug.go`, and `api_workspace.go`, with matching unit tests in
`api_*_test.go`. Test fixtures live under `internal/testdata/api/<resource>/`.
Webhook support is isolated in [`webhook/`](webhook/). Use
[`guide.md`](guide.md) and [`features.md`](features.md) when adding endpoints.

## Build, Test, and Development Commands
Use the `Makefile` as the default entry point:

- `make test`: run `go test ./... -race` across the repo.
- `make lint`: run `go mod tidy` and `golangci-lint`.
- `make lint-fix`: apply auto-fixes from enabled linters/formatters.
- `make go-mod-tidy`: normalize module dependencies for Go `1.25.0`.

For focused work, use standard Go commands such as
`go test . -run TestStory_GetStories -v` or
`go test -v -run ^TestStory_Prod_GetStories$ ./tests/` for a single
integration test when available.

## Coding Style & Naming Conventions
Follow Go defaults and keep code `gofumpt`/`goimports` clean. Use tabs for
indentation. Keep public API names consistent with existing patterns: request
types like `GetStoriesRequest`, service methods like `GetStories`, and
resource structs named after TAPD objects. In this repository, request fields
use pointer types with `url:"...,omitempty"` tags; response fields use value
types unless the field is nullable. Add concise Chinese comments to exported
request and response fields when implementing new TAPD APIs.

## Testing Guidelines
Unit tests sit beside the code in `api_*_test.go`, `client_test.go`, and
similar files. Prefer table-driven tests and assert HTTP method, path,
query/body parameters, and a few representative response fields. Store canned
API payloads in `internal/testdata/api/...`. Run targeted tests first, then
`make test`. If you add a production/integration test, run it individually,
not in batch.

## Commit & Pull Request Guidelines
Recent history uses Conventional Commits, for example
`docs(skills): refine implement-api skill workflow` and
`chore(deps): update module ...`. Keep subjects imperative and scoped when
useful: `feat(story): add story field labels endpoint`. PRs should describe
the user-visible change, list affected APIs/files, note test coverage, and
link related issues. Include sample requests/responses when behavior changes
or a new endpoint is added.
