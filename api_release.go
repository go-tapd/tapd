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

	LaunchForm struct {
		ID             string  `json:"id,omitempty"`              // 评审ID
		WorkspaceID    string  `json:"workspace_id,omitempty"`    // 所属项目ID
		Created        string  `json:"created,omitempty"`         // 创建时间
		Title          *string `json:"title,omitempty"`           // 标题
		Name           string  `json:"name,omitempty"`            // 名称
		Creator        string  `json:"creator,omitempty"`         // 创建人
		Status         string  `json:"status,omitempty"`          // 状态
		VersionType    *string `json:"version_type,omitempty"`    // 版本类型
		Baseline       *string `json:"baseline,omitempty"`        // 基线
		ReleaseModel   *string `json:"release_model,omitempty"`   // 发布模块
		RoadmapVersion *string `json:"roadmap_version,omitempty"` // 路标版本
		ReleaseType    *string `json:"release_type,omitempty"`    // 发布类型
		ChangeType     *string `json:"change_type,omitempty"`     // 变更类型
		SignedBy       *string `json:"signed_by,omitempty"`       // 签发人
		ArchivedBy     *string `json:"archived_by,omitempty"`     // 发布确认人
		CC             *string `json:"cc,omitempty"`              // 抄送人
		ChangeNotifier *string `json:"change_notifier,omitempty"` // 变更通知人
		Signed         *string `json:"signed,omitempty"`          // 签发时间
		Archived       *string `json:"archived,omitempty"`        // 归档时间
		SignerResult   *string `json:"signer_result,omitempty"`   // 签发结论
		SignerComment  *string `json:"signer_comment,omitempty"`  // 签发意见
		ReleaseResult  *string `json:"release_result,omitempty"`  // 发布结果
		ReleaseComment *string `json:"release_comment,omitempty"` // 发布意见
		TestPath       *string `json:"test_path,omitempty"`       // 测试路径
		CreatedPath    *string `json:"created_path,omitempty"`    // 归档路径
		Remark         *string `json:"remark,omitempty"`          // 备注
		Participator   *string `json:"participator,omitempty"`    // 参与人
		TemplateID     string  `json:"template_id,omitempty"`     // 模板ID
		IterationID    *string `json:"iteration_id,omitempty"`    // 迭代ID
		ReleaseID      *string `json:"release_id,omitempty"`      // 发布计划ID
		Flows          string  `json:"flows,omitempty"`           // 流程状态
	}

	GetLaunchFormsRequest struct {
		WorkspaceID    *int           `url:"workspace_id,omitempty"`    // [必须]项目ID
		ID             *int64         `url:"id,omitempty"`              // 发布评审ID
		Creator        *string        `url:"creator,omitempty"`         // 创建人
		Created        *string        `url:"created,omitempty"`         // 创建时间，支持时间查询
		Title          *string        `url:"title,omitempty"`           // 标题
		Status         *string        `url:"status,omitempty"`          // 状态
		VersionType    *string        `url:"version_type,omitempty"`    // 版本类型
		Baseline       *string        `url:"baseline,omitempty"`        // 基线
		ReleaseModel   *string        `url:"release_model,omitempty"`   // 发布模块
		RoadmapVersion *string        `url:"roadmap_version,omitempty"` // 路标版本
		ReleaseType    *string        `url:"release_type,omitempty"`    // 发布类型
		ChangeType     *string        `url:"change_type,omitempty"`     // 变更类型
		SignedBy       *string        `url:"signed_by,omitempty"`       // 签发人
		ArchivedBy     *string        `url:"archived_by,omitempty"`     // 发布确认人
		CC             *string        `url:"cc,omitempty"`              // 抄送人
		ChangeNotifier *string        `url:"change_notifier,omitempty"` // 变更通知人
		Limit          *int           `url:"limit,omitempty"`           // 设置返回数量限制，默认为30，最大取200
		Page           *int           `url:"page,omitempty"`            // 返回当前数量限制下第N页的数据，默认为1
		Fields         *Multi[string] `url:"fields,omitempty"`          // 设置获取的字段，多个字段间以','逗号隔开
	}

	CreateLaunchFormRequest struct {
		WorkspaceID    *int    `json:"workspace_id,omitempty"`    // [必须]项目ID
		Creator        *string `json:"creator,omitempty"`         // [必须]创建人
		TemplateID     *string `json:"template_id,omitempty"`     // [必须]模板ID
		Title          *string `json:"title,omitempty"`           // 标题
		VersionType    *string `json:"version_type,omitempty"`    // 版本类型
		Baseline       *string `json:"baseline,omitempty"`        // 基线
		ReleaseModel   *string `json:"release_model,omitempty"`   // 发布模块
		RoadmapVersion *string `json:"roadmap_version,omitempty"` // 路标版本
		ReleaseType    *string `json:"release_type,omitempty"`    // 发布类型
		SignedBy       *string `json:"signed_by,omitempty"`       // 签发人
		ArchivedBy     *string `json:"archived_by,omitempty"`     // 发布确认人
		CC             *string `json:"cc,omitempty"`              // 抄送人
	}

	LaunchAccessory struct {
		ID          string  `json:"id,omitempty"`           // 依据ID
		FormID      string  `json:"form_id,omitempty"`      // 发布评审ID
		WorkspaceID string  `json:"workspace_id,omitempty"` // 所属项目ID
		Type        string  `json:"type,omitempty"`         // 类型
		Tag         *string `json:"tag,omitempty"`          // 标签
		Title       string  `json:"title,omitempty"`        // 标题
		Content     string  `json:"content,omitempty"`      // 内容
		Description *string `json:"description,omitempty"`  // 详细描述
		ContentType *string `json:"content_type,omitempty"` // 内容类型
		CreatedBy   string  `json:"created_by,omitempty"`   // 创建人
		Created     string  `json:"created,omitempty"`      // 创建时间
		GroupID     *string `json:"group_id,omitempty"`     // 分组ID
		Source      string  `json:"source,omitempty"`       // 来源
	}

	GetLaunchAccessoriesRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		FormID      *int64  `url:"form_id,omitempty"`      // [必须]评审单ID
		ID          *int64  `url:"id,omitempty"`           // 评审依据ID
		CreatedBy   *string `url:"created_by,omitempty"`   // 创建人
		Created     *string `url:"created,omitempty"`      // 创建时间，支持时间查询
	}

	CreateLaunchAccessoryRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		FormID      *int64  `json:"form_id,omitempty"`      // [必须]发布评审ID
		Type        *string `json:"type,omitempty"`         // [必须]类型，仅支持 launch_url
		Content     *string `json:"content,omitempty"`      // [必须]url 地址
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

	// GetLaunchAccessories 获取发布评审依据
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_accessories.html
	GetLaunchAccessories(ctx context.Context, request *GetLaunchAccessoriesRequest, opts ...RequestOption) ([]*LaunchAccessory, *Response, error)

	// GetLaunchForms 获取发布评审
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/get_launch_forms.html
	GetLaunchForms(ctx context.Context, request *GetLaunchFormsRequest, opts ...RequestOption) ([]*LaunchForm, *Response, error)

	// CreateLaunchForm 创建发布评审
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/add_launch_form.html
	CreateLaunchForm(ctx context.Context, request *CreateLaunchFormRequest, opts ...RequestOption) (*LaunchForm, *Response, error)

	// CreateLaunchAccessory 创建发布评审依据
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/release/add_launch_accessories.html
	CreateLaunchAccessory(ctx context.Context, request *CreateLaunchAccessoryRequest, opts ...RequestOption) (*LaunchAccessory, *Response, error)
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

func (s *releaseService) GetLaunchAccessories(
	ctx context.Context, request *GetLaunchAccessoriesRequest, opts ...RequestOption,
) ([]*LaunchAccessory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "launch_accessories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		LaunchAccessory *LaunchAccessory `json:"LaunchAccessory"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	accessories := make([]*LaunchAccessory, 0, len(items))
	for _, item := range items {
		accessories = append(accessories, item.LaunchAccessory)
	}

	return accessories, resp, nil
}

func (s *releaseService) GetLaunchForms(
	ctx context.Context, request *GetLaunchFormsRequest, opts ...RequestOption,
) ([]*LaunchForm, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "launch_forms", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		LaunchForm *LaunchForm `json:"LaunchForm"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	forms := make([]*LaunchForm, 0, len(items))
	for _, item := range items {
		forms = append(forms, item.LaunchForm)
	}

	return forms, resp, nil
}

func (s *releaseService) CreateLaunchForm(
	ctx context.Context, request *CreateLaunchFormRequest, opts ...RequestOption,
) (*LaunchForm, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "launch_forms", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		LaunchForm *LaunchForm `json:"LaunchForm"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.LaunchForm, resp, nil
}

func (s *releaseService) CreateLaunchAccessory(
	ctx context.Context, request *CreateLaunchAccessoryRequest, opts ...RequestOption,
) (*LaunchAccessory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "launch_accessories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		LaunchAccessory *LaunchAccessory `json:"LaunchAccessory"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.LaunchAccessory, resp, nil
}
