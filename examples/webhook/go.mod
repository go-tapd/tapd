module github.com/go-tapd/tapd/examples/webhook

go 1.23.0

replace (
	github.com/go-tapd/tapd => ../../
	github.com/go-tapd/tapd/webhook => ../../webhook/
)

require github.com/go-tapd/tapd/webhook v0.10.0

require (
	github.com/go-tapd/tapd v0.10.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	golang.org/x/sync v0.12.0 // indirect
)
