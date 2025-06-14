package tapd

import (
	"context"
	"net/http"
)

type (
	GetWorkspaceSettingRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // 项目ID
		Type        *string `url:"type,omitempty"`         // nolint:lll // 配置名称（is_enabled_story_category 是否启用需求分类树，workspace_metrology 工时单位）
	}

	GetWorkspaceSettingResponse struct {
		IsEnabledStoryCategory *int    `json:"is_enabled_story_category,omitempty"` // 是否启用需求分类树（1启用，0未启用 ）
		WorkspaceMetrology     *string `json:"workspace_metrology,omitempty"`       // 工时单位（day 天，hour 小时）
	}
)

// SettingService 配置
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/
type SettingService interface {
	// 创建自定义字段（需求及缺陷）
	// 更新下拉类型自定义字段候选值
	// 更新需求下拉类型自定义字段候选值
	// 更新缺陷下拉类型自定义字段候选值
	// 更新级联自定义字段侯选值
	// 创建模块接口
	// 创建版本接口
	// 获取模块接口
	// 获取模块数量接口
	// 获取版本接口
	// 获取版本数量接口
	// 更新模块接口
	// 创建基线接口
	// 创建特性接口
	// 复制需求类别接口
	// 复制缺陷配置接口
	// 更新基线接口
	// 更新特性接口
	// 获取特性接口
	// 获取特性数量接口
	// 获取基线接口
	// 获取基线数量接口
	// 更新版本接口

	// GetWorkspaceSetting 获取项目配置开关
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/setting/get_workspace_setting.html
	GetWorkspaceSetting(ctx context.Context, request *GetWorkspaceSettingRequest, opts ...RequestOption) (*GetWorkspaceSettingResponse, *Response, error)
}

type settingService struct {
	client *Client
}

var _ SettingService = (*settingService)(nil)

func NewSettingService(client *Client) SettingService {
	return &settingService{
		client: client,
	}
}

func (s *settingService) GetWorkspaceSetting(
	ctx context.Context, request *GetWorkspaceSettingRequest, opts ...RequestOption,
) (*GetWorkspaceSettingResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "settings/get_workspace_setting", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetWorkspaceSettingResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}
