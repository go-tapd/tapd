package greetings

import (
	"context"

	"github.com/go-tapd/tapd/mcp/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	tool mcp.Tool
}

var _ tools.Tool = (*Tool)(nil)

func NewTool() *Tool {
	return &Tool{
		tool: mcp.NewTool("tapd greetings",
			mcp.WithDescription("Tapd greetings"),
		),
	}
}

func (t *Tool) Tool() mcp.Tool {
	return t.tool
}

func (t *Tool) Run(context.Context, mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText("Hello, Tapd!"), nil
}
