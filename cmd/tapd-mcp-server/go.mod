module github.com/go-tapd/tapd/cmd/tapd-mcp-server

go 1.23.0

replace (
	github.com/go-tapd/tapd => ../../
	github.com/go-tapd/tapd/mcp => ../../mcp/
)

require (
	github.com/go-tapd/tapd v0.10.0
	github.com/go-tapd/tapd/mcp v0.10.0
)

require (
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/mark3labs/mcp-go v0.17.0 // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
