package template_list

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
		tool: mcp.NewTool("get_story_template_list",
			mcp.WithDescription("返回符合查询条件的所有需求模板"),
			mcp.WithNumber("workitem_type_id", mcp.Description("需求类别ID，不传入则返回所有需求模板")),
		),
	}
}

func (t *Tool) Tool() mcp.Tool {
	return t.tool
}

func (t *Tool) Run(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	req := &tapd.GetStoryTemplatesRequest{
		WorkspaceID: tapd.Ptr(t.workspaceID),
	}

	if typeID, ok := request.Params.Arguments["workitem_type_id"].(float64); ok {
		req.WorkitemTypeID = tapd.Ptr(int(typeID))
	}

	templates, _, err := t.client.StoryService.GetStoryTemplates(ctx, req)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(templates)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(bytes)), nil
}
