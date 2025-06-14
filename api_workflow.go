package tapd

import (
	"context"
	"net/http"
)

type (
	GetAllLastStepsRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // 项目 ID
		System      *string `url:"system,omitempty"`       // 系统名。目前只支持 story（需求的）
		GroupKey    *string `url:"group_key,omitempty"`    // 分组字段，可选字段 workflow_id(工作流ID) 或 workitem_type_id (需求类别ID)	默认按workitem_type_id分组
	}

	WorkflowAllLastStep struct {
		Key    string                       `json:"key,omitempty"`    // 工作流ID 或者 需求类别ID , 根据group_key确定
		Status []*WorkflowAllLastStepStatus `json:"status,omitempty"` // 状态列表
	}

	WorkflowAllLastStepStatus struct {
		Alias string `json:"alias,omitempty"` // 状态别名
		Name  string `json:"name,omitempty"`  // 状态名称
	}
)

// WorkflowService 工作流
type WorkflowService interface {
	// 获取工作流流转细则
	// 获取工作流结束状态

	// GetAllLastSteps 获取所有结束状态
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workflow/get_workflow_all_last_steps.html
	GetAllLastSteps(ctx context.Context, request *GetAllLastStepsRequest, opts ...RequestOption) ([]*WorkflowAllLastStep, *Response, error)

	// 获取工作流状态中英文名对应关系
	// 获取工作流起始状态
	// 获取项目下的工作流列表
}

type workflowService struct {
	client *Client
}

var _ WorkflowService = (*workflowService)(nil)

func NewWorkflowService(client *Client) WorkflowService {
	return &workflowService{
		client: client,
	}
}

func (s *workflowService) GetAllLastSteps(
	ctx context.Context, request *GetAllLastStepsRequest, opts ...RequestOption,
) ([]*WorkflowAllLastStep, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workflows/all_last_steps", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var result map[string]map[string]string
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	steps := make([]*WorkflowAllLastStep, 0, len(result))
	for key, value := range result {
		statuses := make([]*WorkflowAllLastStepStatus, 0, len(value))
		for alias, name := range value {
			statuses = append(statuses, &WorkflowAllLastStepStatus{Alias: alias, Name: name})
		}
		steps = append(steps, &WorkflowAllLastStep{Key: key, Status: statuses})
	}

	return steps, resp, nil
}
