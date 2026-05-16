# Workflow Memory

Store reusable project workflows here.

## Validation

- Package inventory: `GOCACHE=/private/tmp/go-build-cache-tapd go list ./...`.
- Full test suite: `make test`, which runs `go test ./... -race`.
- Lint: `make lint`, which runs `go mod tidy -compat=1.25.0` and then `go tool golangci-lint run --concurrency=4 --allow-serial-runners`.
- Auto-fix lint: `make lint-fix`.
- Focused unit tests use the package-local pattern, for example `go test . -run TestStory_GetStories -v`.
- Run production or integration tests individually when credentials and network access are available; do not batch them.

## Maintenance

- For new TAPD endpoints, start from the official API doc, then read `guide.md` and `internal/skills/implement-api/SKILL.md`, locate a similar `api_*.go` implementation, add request/response/service code, add `internal/testdata/api/<resource>/...` fixture data, add unit tests, and update `features.md`.
- When comparing `features.md` with the official TAPD API reference, add official endpoints that are missing as unchecked items. Keep existing checked status unchanged unless implementation code is verified.
- If a `features.md` entry is present locally but not visible in the current official reference, preserve it unless there is explicit evidence that the API was removed; some entries may come from older docs or alternate doc pages.
- After Markdown-only changes, at minimum run `git diff --check`.
