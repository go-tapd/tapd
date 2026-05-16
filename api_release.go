package tapd

import (
	"context"
	"net/http"
)

type (
	Release struct {
		ID          string  `json:"id,omitempty"`           // ID
		WorkspaceID string  `json:"workspace_id,omitempty"` // 项目ID
		Name        string  `json:"name,omitempty"`         // 标题
		Description *string `json:"description,omitempty"`  // 详细描述
		StartDate   string  `json:"startdate,omitempty"`    // 开始时间
		EndDate     string  `json:"enddate,omitempty"`      // 结束时间
		Creator     string  `json:"creator,omitempty"`      // 创建人
		Created     string  `json:"created,omitempty"`      // 创建时间
		Modified    string  `json:"modified,omitempty"`     // 最后修改时间
		Status      string  `json:"status,omitempty"`       // 状态
	}

	CreateReleaseRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string `json:"name,omitempty"`         // [必须]标题
		Description *string `json:"description,omitempty"`  // 详细描述
		StartDate   *string `json:"startdate,omitempty"`    // [必须]开始时间
		EndDate     *string `json:"enddate,omitempty"`      // [必须]结束时间
		Creator     *string `json:"creator,omitempty"`      // 创建人
	}

	GetReleasesRequest struct {
		ID          *Multi[int64]  `url:"id,omitempty"`           // id，支持多ID查询
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string        `url:"name,omitempty"`         // 标题，支持模糊匹配
		Description *string        `url:"description,omitempty"`  // 详细描述
		StartDate   *string        `url:"startdate,omitempty"`    // 开始时间
		EndDate     *string        `url:"enddate,omitempty"`      // 结束时间
		Creator     *string        `url:"creator,omitempty"`      // 创建人
		Created     *string        `url:"created,omitempty"`      // 创建时间，支持时间查询
		Modified    *string        `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		Status      *string        `url:"status,omitempty"`       // 状态
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetReleasesCountRequest struct {
		ID          *Multi[int64] `url:"id,omitempty"`           // id，支持多ID查询
		WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string       `url:"name,omitempty"`         // 标题，支持模糊匹配
		Description *string       `url:"description,omitempty"`  // 详细描述
		StartDate   *string       `url:"startdate,omitempty"`    // 开始时间
		EndDate     *string       `url:"enddate,omitempty"`      // 结束时间
		Creator     *string       `url:"creator,omitempty"`      // 创建人
		Created     *string       `url:"created,omitempty"`      // 创建时间，支持时间查询
		Modified    *string       `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		Status      *string       `url:"status,omitempty"`       // 状态
	}

	UpdateReleaseRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		ID          *int64  `json:"id,omitempty"`           // [必须]发布计划ID
		Name        *string `json:"name,omitempty"`         // 标题
		Description *string `json:"description,omitempty"`  // 详细描述
		StartDate   *string `json:"startdate,omitempty"`    // 开始时间
		EndDate     *string `json:"enddate,omitempty"`      // 结束时间
		Status      *string `json:"status,omitempty"`       // 状态
	}
)

// ReleaseService 发布服务。
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/
type ReleaseService interface {
	// CreateRelease 创建发布计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/add_new_release.html
	CreateRelease(ctx context.Context, request *CreateReleaseRequest, opts ...RequestOption) (*Release, *Response, error)

	// GetReleases 获取发布计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases.html
	GetReleases(ctx context.Context, request *GetReleasesRequest, opts ...RequestOption) ([]*Release, *Response, error)

	// GetReleasesCount 获取发布计划数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_releases_count.html
	GetReleasesCount(ctx context.Context, request *GetReleasesCountRequest, opts ...RequestOption) (int, *Response, error)

	// UpdateRelease 更新发布计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/update_new_release.html
	UpdateRelease(ctx context.Context, request *UpdateReleaseRequest, opts ...RequestOption) (*Release, *Response, error)
}

type releaseService struct {
	client *Client
}

var _ ReleaseService = (*releaseService)(nil)

func NewReleaseService(client *Client) ReleaseService {
	return &releaseService{
		client: client,
	}
}

func (s *releaseService) CreateRelease(
	ctx context.Context, request *CreateReleaseRequest, opts ...RequestOption,
) (*Release, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "new_releases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Release *Release `json:"Release"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Release, resp, nil
}

func (s *releaseService) GetReleases(
	ctx context.Context, request *GetReleasesRequest, opts ...RequestOption,
) ([]*Release, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "releases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Release *Release `json:"Release"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	releases := make([]*Release, 0, len(items))
	for _, item := range items {
		releases = append(releases, item.Release)
	}

	return releases, resp, nil
}

func (s *releaseService) GetReleasesCount(
	ctx context.Context, request *GetReleasesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "releases/count", request, opts)
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

func (s *releaseService) UpdateRelease(
	ctx context.Context, request *UpdateReleaseRequest, opts ...RequestOption,
) (*Release, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "new_releases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Release *Release `json:"Release"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Release, resp, nil
}
