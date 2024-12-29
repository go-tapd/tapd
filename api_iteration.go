package tapd

import (
	"context"
	"net/http"
)

// IterationService 迭代
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/
type IterationService struct {
	client *Client
}

// 创建迭代
// 获取迭代自定义字段配置
// 获取迭代
// 获取迭代数量
// 更新迭代
// 获取迭代变更历史
// 获取迭代仪表盘自定义卡片内容
// 修改迭代仪表盘自定义卡片内容
// 锁定迭代
// 解锁迭代

// GetWorkitemTypes 获取迭代类别列表
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/workitem_types.html
func (s *IterationService) GetWorkitemTypes(
	ctx context.Context, request *GetWorkitemTypesRequest, opts ...RequestOption,
) ([]*WorkitemType, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "iterations/workitem_types", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		WorkitemType *WorkitemType `json:"WorkitemType"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	workitemTypes := make([]*WorkitemType, 0, len(items))
	for _, item := range items {
		workitemTypes = append(workitemTypes, item.WorkitemType)
	}

	return workitemTypes, resp, nil
}

type GetWorkitemTypesRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"` // 项目 ID
}

type WorkitemType struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspace_id"`
	EntityType  string `json:"entity_type"`
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

// GetTemplateList 获取迭代模板列表
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/template_list.html
func (s *IterationService) GetTemplateList(
	ctx context.Context, request *GetTemplateListRequest, opts ...RequestOption,
) ([]*WorkitemTemplate, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "iterations/template_list", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		WorkitemTemplate *WorkitemTemplate `json:"WorkitemTemplate"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	templates := make([]*WorkitemTemplate, 0, len(items))
	for _, item := range items {
		templates = append(templates, item.WorkitemTemplate)
	}

	return templates, resp, nil
}

type GetTemplateListRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"` // 项目 ID
}

type WorkitemTemplate struct {
	ID          string `json:"id"`
	WorkspaceID string `json:"workspace_id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Creator     string `json:"creator"`
	Created     string `json:"created"`
	Modified    string `json:"modified"`
}

// 获取迭代模板字段配置
// 获取迭代类别默认模板字段配置
// 获取计划应用
// 获取计划应用数量
