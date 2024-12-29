package tapd

import (
	"context"
	"net/http"
)

type Iteration struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	WorkspaceID    string `json:"workspace_id,omitempty"`
	StartDate      string `json:"startdate,omitempty"`
	EndDate        string `json:"enddate,omitempty"`
	Status         string `json:"status,omitempty"`
	ReleaseID      string `json:"release_id,omitempty"`
	Description    string `json:"description,omitempty"`
	Creator        string `json:"creator,omitempty"`
	Created        string `json:"created,omitempty"`
	Modified       string `json:"modified,omitempty"`
	Completed      string `json:"completed,omitempty"`
	EntityType     string `json:"entity_type,omitempty"`
	ParentID       string `json:"parent_id,omitempty"`
	AncestorID     string `json:"ancestor_id,omitempty"`
	Path           string `json:"path,omitempty"`
	WorkitemTypeID string `json:"workitem_type_id,omitempty"`
	TemplatedID    string `json:"templated_id,omitempty"`
	PlanAppID      string `json:"plan_app_id,omitempty"`
	CrucialMoment  string `json:"crucial_moment,omitempty"`
	Label          string `json:"label,omitempty"`
	ReleaseOwner   string `json:"releaseowner,omitempty"`
	LaunchDate     string `json:"launchdate,omitempty"`
	Notice         string `json:"notice,omitempty"`
	ReleaseName    string `json:"releasename,omitempty"`
	CustomField1   string `json:"custom_field_1,omitempty"`
	CustomField2   string `json:"custom_field_2,omitempty"`
	CustomField3   string `json:"custom_field_3,omitempty"`
	CustomField4   string `json:"custom_field_4,omitempty"`
	CustomField5   string `json:"custom_field_5,omitempty"`
	CustomField6   string `json:"custom_field_6,omitempty"`
	CustomField7   string `json:"custom_field_7,omitempty"`
	CustomField8   string `json:"custom_field_8,omitempty"`
	CustomField9   string `json:"custom_field_9,omitempty"`
	CustomField10  string `json:"custom_field_10,omitempty"`
	CustomField11  string `json:"custom_field_11,omitempty"`
	CustomField12  string `json:"custom_field_12,omitempty"`
	CustomField13  string `json:"custom_field_13,omitempty"`
	CustomField14  string `json:"custom_field_14,omitempty"`
	CustomField15  string `json:"custom_field_15,omitempty"`
	CustomField16  string `json:"custom_field_16,omitempty"`
	CustomField17  string `json:"custom_field_17,omitempty"`
	CustomField18  string `json:"custom_field_18,omitempty"`
	CustomField19  string `json:"custom_field_19,omitempty"`
	CustomField20  string `json:"custom_field_20,omitempty"`
	CustomField21  string `json:"custom_field_21,omitempty"`
	CustomField22  string `json:"custom_field_22,omitempty"`
	CustomField23  string `json:"custom_field_23,omitempty"`
	CustomField24  string `json:"custom_field_24,omitempty"`
	CustomField25  string `json:"custom_field_25,omitempty"`
	CustomField26  string `json:"custom_field_26,omitempty"`
	CustomField27  string `json:"custom_field_27,omitempty"`
	CustomField28  string `json:"custom_field_28,omitempty"`
	CustomField29  string `json:"custom_field_29,omitempty"`
	CustomField30  string `json:"custom_field_30,omitempty"`
	CustomField31  string `json:"custom_field_31,omitempty"`
	CustomField32  string `json:"custom_field_32,omitempty"`
	CustomField33  string `json:"custom_field_33,omitempty"`
	CustomField34  string `json:"custom_field_34,omitempty"`
	CustomField35  string `json:"custom_field_35,omitempty"`
	CustomField36  string `json:"custom_field_36,omitempty"`
	CustomField37  string `json:"custom_field_37,omitempty"`
	CustomField38  string `json:"custom_field_38,omitempty"`
	CustomField39  string `json:"custom_field_39,omitempty"`
	CustomField40  string `json:"custom_field_40,omitempty"`
	CustomField41  string `json:"custom_field_41,omitempty"`
	CustomField42  string `json:"custom_field_42,omitempty"`
	CustomField43  string `json:"custom_field_43,omitempty"`
	CustomField44  string `json:"custom_field_44,omitempty"`
	CustomField45  string `json:"custom_field_45,omitempty"`
	CustomField46  string `json:"custom_field_46,omitempty"`
	CustomField47  string `json:"custom_field_47,omitempty"`
	CustomField48  string `json:"custom_field_48,omitempty"`
	CustomField49  string `json:"custom_field_49,omitempty"`
	CustomField50  string `json:"custom_field_50,omitempty"`
	OriginName     string `json:"origin_name,omitempty"`
}

// IterationService 迭代
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/
type IterationService struct {
	client *Client
}

// 创建迭代
// 获取迭代自定义字段配置

// GetIterations 获取迭代
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/iteration/get_iterations.html
func (s *IterationService) GetIterations(
	ctx context.Context, request *GetIterationsRequest, opts ...RequestOption,
) ([]*Iteration, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "iterations", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Iteration *Iteration `json:"Iteration"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	iterations := make([]*Iteration, 0, len(items))
	for _, item := range items {
		iterations = append(iterations, item.Iteration)
	}

	return iterations, resp, nil
}

type GetIterationsRequest struct {
	ID             *Multi[int]    `url:"id,omitempty"`               // ID 支持多ID查询
	Name           *string        `url:"name,omitempty"`             // 标题 支持模糊匹配
	WorkspaceID    *int           `url:"workspace_id,omitempty"`     // 项目 ID
	Description    *string        `url:"description,omitempty"`      // 详细描述
	StartDate      *string        `url:"startdate,omitempty"`        // 开始时间 支持时间查询
	EndDate        *string        `url:"enddate,omitempty"`          // 结束时间 支持时间查询
	WorkitemTypeID *int           `url:"workitem_type_id,omitempty"` // 迭代类别
	PlanAppID      *int           `url:"plan_app_id,omitempty"`      // 计划应用 ID
	Status         *string        `url:"status,omitempty"`           // 状态（系统状态 open/done，自定义状态可传中文）
	Creator        *string        `url:"creator,omitempty"`          // 创建人
	Created        *string        `url:"created,omitempty"`          // 创建时间 支持时间查询
	Modified       *string        `url:"modified,omitempty"`         // 最后修改时间 支持时间查询
	Completed      *string        `url:"completed,omitempty"`        // 完成时间
	CustomField    *string        `url:"custom_field_*,omitempty"`   // 自定义字段参数
	Limit          *int           `url:"limit,omitempty"`            // 设置返回数量限制，默认为 30
	Page           *int           `url:"page,omitempty"`             // 返回当前数量限制下第 N 页的数据，默认为 1（第一页）
	Order          *Order         `url:"order,omitempty"`            // 排序规则，规则：字段名 ASC 或者 DESC，然后 urlencode 如按创建时间逆序：order=created%20desc
	Fields         *Multi[string] `url:"fields,omitempty"`           // 设置获取的字段，多个字段间以 ',' 逗号隔开
}

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
