package tapd

import (
	"context"
	"net/http"
)

// LabelColor is a type for label colors.
type LabelColor string

const (
	LabelColor1 LabelColor = "1"
	LabelColor2 LabelColor = "2"
	LabelColor3 LabelColor = "3"
	LabelColor4 LabelColor = "4"
)

type (
	// Label represents a label.
	Label struct {
		ID          string     `json:"id,omitempty"`
		WorkspaceID string     `json:"workspace_id,omitempty"`
		Name        string     `json:"name,omitempty"`
		Color       LabelColor `json:"color,omitempty"`
		Category    string     `json:"category,omitempty"`
		Creator     string     `json:"creator,omitempty"`
		Modifier    string     `json:"modifier,omitempty"`
		Created     string     `json:"created,omitempty"`
		Modified    string     `json:"modified,omitempty"`
		ColorValue  string     `json:"color_value,omitempty"`
	}

	LabelPool struct {
		LabelPool *Label `json:"LabelPool"`
	}

	GetLabelsRequest struct {
		WorkspaceID *int        `url:"workspace_id,omitempty"` // [必选]项目ID
		ID          *Multi[int] `url:"id,omitempty"`           // [可选]id 支持多ID查询
		Name        *string     `url:"name,omitempty"`         // [可选]标签名称 支持模糊匹配
		Creator     *string     `url:"creator,omitempty"`      // [可选]创建人
		Created     *string     `url:"created,omitempty"`      // [可选]创建时间 支持时间查询
		Limit       *int        `url:"limit,omitempty"`        // [可选]设置返回数量限制，默认为30
		Page        *int        `url:"page,omitempty"`         // [可选]返回当前数量限制下第N页的数据，默认为1（第一页）
		Order       *Order      `url:"order,omitempty"`        // [可选]排序规则，规则：字段名 ASC或者DESC，然后 urlencode 如按创建时间逆序
	}

	GetLabelCountRequest struct {
		WorkspaceID *int        `url:"workspace_id,omitempty"` // [必选]项目ID
		ID          *Multi[int] `url:"id,omitempty"`           // [可选]id 支持多ID查询
		Name        *string     `url:"name,omitempty"`         // [可选]标签名称 支持模糊匹配
		Creator     *string     `url:"creator,omitempty"`      // [可选]创建人
		Created     *string     `url:"created,omitempty"`      // [可选]创建时间 支持时间查询
	}

	CreateLabelRequest struct {
		WorkspaceID *int        `json:"workspace_id"` // [必选]项目ID
		Name        *string     `json:"name"`         // [必选]标签名称
		Color       *LabelColor `json:"color"`        // 标签颜色
		Creator     *string     `json:"creator"`      // 创建人
	}

	UpdateLabelRequest struct {
		ID          *int        `json:"id"`           // [必选]ID
		WorkspaceID *int        `json:"workspace_id"` // [必选]项目ID
		Color       *LabelColor `json:"color"`        // 标签颜色
		Modifier    *string     `json:"modifier"`     // 更新人
	}
)

// LabelService handles communication with the label related methods of the Tapd API.
//
// Tapd API docs: https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/index.html
type LabelService interface {
	// GetLabels 获取自定义标签
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/query_label.html
	GetLabels(ctx context.Context, request *GetLabelsRequest, opts ...RequestOption) ([]*Label, *Response, error)

	// GetLabelsCount 获取标签数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/count_label.html
	GetLabelsCount(ctx context.Context, request *GetLabelCountRequest, opts ...RequestOption) (int, *Response, error)

	// CreateLabel 创建标签
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/add_label.html
	CreateLabel(ctx context.Context, request *CreateLabelRequest, opts ...RequestOption) (*Label, *Response, error)

	// UpdateLabel 更新标签
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/label/update_label.html
	UpdateLabel(ctx context.Context, request *UpdateLabelRequest, opts ...RequestOption) (*Label, *Response, error)
}

type labelService struct {
	client *Client
}

var _ LabelService = (*labelService)(nil)

func NewLabelService(client *Client) LabelService {
	return &labelService{
		client: client,
	}
}

func (s *labelService) GetLabels(
	ctx context.Context, request *GetLabelsRequest, opts ...RequestOption,
) ([]*Label, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "label", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []*LabelPool
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	labels := make([]*Label, 0, len(items))
	for _, item := range items {
		labels = append(labels, item.LabelPool)
	}

	return labels, resp, nil
}

func (s *labelService) GetLabelsCount(
	ctx context.Context, request *GetLabelCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "label/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response CountResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

func (s *labelService) CreateLabel(
	ctx context.Context, request *CreateLabelRequest, opts ...RequestOption,
) (*Label, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "label", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response LabelPool
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.LabelPool, resp, nil
}

func (s *labelService) UpdateLabel(
	ctx context.Context, request *UpdateLabelRequest, opts ...RequestOption,
) (*Label, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "label", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response LabelPool
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.LabelPool, resp, nil
}
