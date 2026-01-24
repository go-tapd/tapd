# Implement API Reference

## Entry Points
- `guide.md`: Project conventions, patterns, and data type rules.
- `api_*.go`, `api_*_test.go`: Reference implementations by resource type.

## Core Checklist
- Collect API doc URL, service type, and description.
- Read `guide.md` before editing code.
- Fetch and parse the API doc; list required params and response fields.
- Find a similar endpoint and copy its structure and test style.
- Use pointer fields for request params with `url` tags and `omitempty`.
- Use value fields for responses; only nullable fields are pointers; `json` tags.
- Add Chinese comments to request/response fields.
- Add interface method with API doc link in the comment.
- Implement method using existing request/response helpers.
- Create `internal/testdata/api/{resource}/{endpoint}.json` with status/data/info.
- Write unit test with mock server assertions.
- Write single integration test in `tests/` only for that endpoint.
- Update `features.md`.
- Run targeted tests and then full unit tests.

## Quality Checklist
- Request struct uses pointer fields for all params.
- Response struct uses value fields unless nullable.
- Method signature matches existing patterns in similar endpoints.
- Test data matches API response shape and uses status/data/info keys.
- Unit test asserts method/path/params plus 3-5 key fields.
- Integration test is run alone and uses a real workspace ID.

## Test Commands
- Unit: `go test . -run Test{Service}_{Method} -v`
- Integration (single): `go test -v -run ^Test{Service}_Prod_{Method}$ ./tests/`
- Full unit suite: `go test . -race`
