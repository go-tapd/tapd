package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Tool interface {
	Tool() mcp.Tool
	Run(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) // server.ToolHandlerFunc
}

func RegisterTools(srv *server.MCPServer, tools ...Tool) {
	for _, tool := range tools {
		srv.AddTool(tool.Tool(), tool.Run)
	}
}
