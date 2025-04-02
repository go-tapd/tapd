package roles

import (
	"context"
	"encoding/json"

	"github.com/go-tapd/tapd"
	"github.com/go-tapd/tapd/mcp/internal/tools"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	workspaceID int
	client      *tapd.Client
	tool        mcp.Tool
}

var _ tools.Tool = (*Tool)(nil)

func NewTool(workspaceID int, client *tapd.Client) *Tool {
	return &Tool{
		workspaceID: workspaceID,
		client:      client,
		tool: mcp.NewTool("get_user_roles",
			mcp.WithDescription("获取项目角色ID对照关系"),
		),
	}
}

func (t *Tool) Tool() mcp.Tool {
	return t.tool
}

func (t *Tool) Run(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	roles, _, err := t.client.UserService.GetRoles(ctx, &tapd.GetRolesRequest{
		WorkspaceID: tapd.Ptr(t.workspaceID),
	})
	if err != nil {
		return nil, err
	}

	rolesBytes, err := json.Marshal(roles)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(rolesBytes)), nil
}
