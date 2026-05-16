# Project Memory

Store stable, reusable project context here.

## Conventions

- This repository is the `github.com/go-tapd/tapd` Go SDK for TAPD. Keep changes small and resource-scoped.
- API services live in root-level `api_*.go` files. Public service methods use the pattern `Method(ctx context.Context, request *RequestType, opts ...RequestOption) (result, *Response, error)`.
- Request structs use pointer fields with `url:"...,omitempty"` tags. Response structs default to value fields and use pointers only for nullable fields.
- Add concise Chinese comments to exported request and response fields when implementing TAPD APIs.
- Include the official TAPD API documentation link in new service method comments.
- Use existing helper types for encoded parameters: `NewMulti` for comma-separated values, `NewEnum` for pipe-separated enum lists, and `NewOrder` for order parameters.
- Avoid `codex/`-prefixed branch names in this repo unless the user explicitly asks for that prefix.

## Structure

- Core client logic is in `client.go`, `client_options*.go`, `request.go`, `response.go`, `helpers.go`, and `api_types.go`.
- The root package wires service fields on `Client`: Story, Bug, Iteration, Task, Comment, Report, Attachment, Timesheet, Workspace, Label, Measure, User, Workflow, and Setting.
- API tests sit beside implementation files as `api_*_test.go`; canned API payloads live under `internal/testdata/api/<resource>/`.
- Webhook support is a separate package under `webhook/`; event parsing is in `webhook/event.go`, listener interfaces in `webhook/listeners.go`, and dispatch in `webhook/dispatcher.go`. Webhook fixtures live under `internal/testdata/webhook/`.
- `guide.md` and `internal/skills/implement-api/` are local implementation references for adding TAPD endpoints.
- `features.md` tracks API coverage against official TAPD API documentation.

## Decisions

- Treat `features.md` as a conservative coverage ledger: add missing official endpoints as unchecked items, mark items checked only when implemented, and avoid deleting older entries unless their source is verified obsolete.
