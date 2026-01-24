---
name: implement-api
description: Implement a new TAPD API endpoint in the go-tapd SDK when a user provides an API documentation URL or asks to add/implement an endpoint (Story/Bug/Task/Iteration/etc.), including request/response structs, service methods, tests, test data, and features.md updates.
---

# Implement API

## Overview
Implement a TAPD API endpoint in this repo using the standard workflow and conventions.

## Workflow
1. Collect inputs: API doc URL, service type, brief description.
2. Read project guidance: `guide.md`; use `references/implement-api.md` as a checklist.
3. Fetch the API doc content when a URL is provided; extract required params/response fields.
4. Locate a similar implementation in `api_*.go` and `api_*_test.go`; mirror patterns and naming.
5. Implement code:
   - Request struct: pointer fields, `url` tags, Chinese comments.
   - Response struct: value fields with nullable pointers, `json` tags, Chinese comments.
   - Service interface method and implementation with API doc link in the comment.
6. Add tests:
   - Test data in `internal/testdata/api/{resource}/{endpoint}.json`.
   - Unit test in `api_*_test.go` validating method/path/params and key fields.
   - Integration test in `tests/api_*_prod_test.go` (single-test run only).
7. Update docs: mark the endpoint in `features.md`.
8. Run tests and report results.

## Conventions and checks
- Follow `guide.md` exactly for naming, request/response types, and patterns.
- Preserve existing response wrappers and helpers; do not invent new ones.
- Always include the API documentation link in the method comment.
- Keep comments in Chinese for exported fields.
- Do not run integration tests in batch.

## Output expectations
- Provide a short summary of the API implemented.
- List modified files with line numbers.
- Provide test commands executed and their results.
- Provide a usage snippet for the new method.
- Call out any assumptions or missing API details.

## Resources
- `references/implement-api.md`: Checklist and repo entry points.
