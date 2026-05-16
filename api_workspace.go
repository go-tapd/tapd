package tapd

import (
	"context"
	"encoding/json"
	"net/http"
)

// WorkCalendarType 工作日历类型。
type WorkCalendarType string

const (
	WorkCalendarTypeSystem WorkCalendarType = "system"
	WorkCalendarTypeCustom WorkCalendarType = "custom"
)

type (
	GetWorkspaceInfoRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	GetSubWorkspacesRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
		TemplateID  *int `url:"template_id,omitempty"`  // [可选]子项目模板ID
	}

	GetCompanyWorkspacesRequest struct {
		CompanyID   *int    `url:"company_id,omitempty"`   // [必须]公司ID
		Category    *string `url:"category,omitempty"`     // [可选]项目类型，project 或 mini_project
		WithExtends *int    `url:"with_extends,omitempty"` // [可选]是否返回项目扩展信息
	}

	GetUserParticipantWorkspacesRequest struct {
		Nick      *string `url:"nick,omitempty"`       // [必须]用户昵称
		CompanyID *int    `url:"company_id,omitempty"` // [必须]公司ID
	}

	GetWorkspaceCustomFieldsSettingsRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	GetWorkspaceDocumentsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Limit       *int           `url:"limit,omitempty"`        // 返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 当前页，默认为1
		Fields      *Multi[string] `url:"fields,omitempty"`       // 返回字段，多个以逗号分隔
	}

	SetCustomWorkCalendarRequest struct {
		WorkspaceID *int      `json:"workspace_id,omitempty"` // [必须]项目ID
		Year        *string   `json:"year,omitempty"`         // [必须]年份
		Weekdays    *[]int    `json:"weekdays,omitempty"`     // 周内工作日，1-7；不传默认周一到周五
		Holidays    *[]string `json:"holidays,omitempty"`     // 额外假日
		Workdays    *[]string `json:"workdays,omitempty"`     // 额外工作日
	}

	EnableWorkCalendarRequest struct {
		WorkspaceID *int              `json:"workspace_id,omitempty"` // [必须]项目ID
		Type        *WorkCalendarType `json:"type,omitempty"`         // [必须]启用类型，system 或 custom
	}

	GetCustomWorkCalendarRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		Year        *string `url:"year,omitempty"`         // [必须]年份
	}

	GetWorkCalendarSettingsRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	GetWorkItemsLongIDByShortIDsRequest struct {
		ShortIDs    *string     `url:"short_ids,omitempty"`    // 短ID，多个以分号分隔；short_ids 和 long_ids 不允许都不传
		LongIDs     *string     `url:"long_ids,omitempty"`     // 长ID，多个以分号分隔
		WorkspaceID *int        `url:"workspace_id,omitempty"` // [必须]项目ID
		EntityType  *EntityType `url:"entity_type,omitempty"`  // [必须]业务对象类型，story、task、bug
	}

	Workspace struct {
		ID                string         `json:"id,omitempty"`                  // 项目ID
		Name              string         `json:"name,omitempty"`                // 项目名称
		PrettyName        string         `json:"pretty_name,omitempty"`         // 项目显示名称
		Category          string         `json:"category,omitempty"`            // 项目类型
		Status            string         `json:"status,omitempty"`              // 项目状态
		Description       string         `json:"description,omitempty"`         // 项目描述
		Creator           string         `json:"creator,omitempty"`             // 创建人
		CreatorID         string         `json:"creator_id,omitempty"`          // 创建人ID
		Created           string         `json:"created,omitempty"`             // 创建时间
		BeginDate         *string        `json:"begin_date,omitempty"`          // 开始日期
		EndDate           *string        `json:"end_date,omitempty"`            // 结束日期
		Secrecy           string         `json:"secrecy,omitempty"`             // 是否保密
		ExternalOn        string         `json:"external_on,omitempty"`         // 是否开启外部协作
		NewTask           string         `json:"new_task,omitempty"`            // 是否启用新任务
		CompanyID         string         `json:"company_id,omitempty"`          // 公司ID
		ParentID          string         `json:"parent_id,omitempty"`           // 父项目ID
		TemplateID        string         `json:"template_id,omitempty"`         // 模板ID
		ProductType       *string        `json:"product_type,omitempty"`        // 产品类型
		PlatformType      *string        `json:"platform_type,omitempty"`       // 平台类型
		IsSelfDevelopment *string        `json:"is_self_development,omitempty"` // 是否自研
		Objective         string         `json:"objective,omitempty"`           // 项目目标
		Schedule          *string        `json:"schedule,omitempty"`            // 项目计划
		Milestone         *string        `json:"milestone,omitempty"`           // 里程碑
		Risk              *string        `json:"risk,omitempty"`                // 风险
		Closed            *string        `json:"closed,omitempty"`              // 是否关闭
		MemberCount       int            `json:"member_count,omitempty"`        // 成员数量
		WorkspaceExtends  map[string]any `json:"WorkspaceExtends,omitempty"`    // 项目扩展信息
	}

	WorkspaceCustomFieldsSetting struct {
		ID              string  `json:"id,omitempty"`           // 自定义字段配置的ID
		WorkspaceID     string  `json:"workspace_id,omitempty"` // 所属项目ID
		AppID           string  `json:"app_id,omitempty"`       // 应用ID
		EntryType       string  `json:"entry_type,omitempty"`   // 所属实体对象
		CustomField     string  `json:"custom_field,omitempty"` // 自定义字段标识
		Type            string  `json:"type,omitempty"`         // 输入类型
		Name            string  `json:"name,omitempty"`         // 自定义字段显示名称
		Options         *string `json:"options,omitempty"`      // 自定义字段可选值
		ExtraConfig     *string `json:"extra_config,omitempty"` // 额外配置
		Enabled         string  `json:"enabled,omitempty"`      // 是否启用
		Freeze          string  `json:"freeze,omitempty"`       // 是否冻结
		Sort            *string `json:"sort,omitempty"`         // 显示时排序系数
		Memo            *string `json:"memo,omitempty"`         // 备注
		OpenExtensionID string  `json:"open_extension_id,omitempty"`
		IsOut           int     `json:"is_out,omitempty"`
		IsUninstall     int     `json:"is_uninstall,omitempty"`
		AppName         string  `json:"app_name,omitempty"`
	}

	WorkspaceDocument struct {
		ID          string  `json:"id,omitempty"`           // ID
		WorkspaceID string  `json:"workspace_id,omitempty"` // 项目ID
		Name        string  `json:"name,omitempty"`         // 标题
		Type        string  `json:"type,omitempty"`         // 文档类型
		FolderID    string  `json:"folder_id,omitempty"`    // 文件夹ID
		Creator     string  `json:"creator,omitempty"`      // 创建人
		Modifier    string  `json:"modifier,omitempty"`     // 最后修改人
		Status      *string `json:"status,omitempty"`       // 状态
		Created     string  `json:"created,omitempty"`      // 创建时间
		Modified    string  `json:"modified,omitempty"`     // 最后修改时间
	}

	SetCustomWorkCalendarResponse struct {
		Success bool `json:"success,omitempty"` // 是否设置成功
	}

	EnableWorkCalendarResponse struct {
		Success bool `json:"success,omitempty"` // 是否启用成功
	}

	CustomWorkCalendar struct {
		Weekdays []string `json:"weekdays,omitempty"` // 周内工作日
		Holidays []string `json:"holidays,omitempty"` // 额外假日
		Workdays []string `json:"workdays,omitempty"` // 额外工作日
	}

	WorkCalendarSetting struct {
		Name   string           `json:"name,omitempty"`   // 工作日历名称
		Type   WorkCalendarType `json:"type,omitempty"`   // 工作日历类型
		Enable bool             `json:"enable,omitempty"` // 是否启用
	}

	GetWorkItemsLongIDByShortIDsResponse struct {
		ValidIDMap      []*WorkItemIDMap `json:"valid_id_map,omitempty"`      // 有效ID映射
		InvalidLongIDs  []string         `json:"invalid_long_ids,omitempty"`  // 无效长ID
		InvalidShortIDs []string         `json:"invalid_short_ids,omitempty"` // 无效短ID
	}

	WorkItemIDMap struct {
		ShortID     string     `json:"short_id,omitempty"`     // 短ID
		LongID      string     `json:"long_id,omitempty"`      // 长ID
		EntityType  EntityType `json:"entity_type,omitempty"`  // 业务对象类型
		WorkspaceID string     `json:"workspace_id,omitempty"` // 项目ID
		CompanyID   string     `json:"company_id,omitempty"`   // 公司ID
	}

	GetUsersRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		User        *Multi[string] `url:"user,omitempty"`         // [可选]用户昵称或ID
		Fields      *Multi[string] `url:"fields,omitempty"`       // [可选]返回的字段列表，user,user_id,role_id,name,email,real_join_time 可选，以,分隔
	}

	User struct {
		User             string   `json:"user"`
		UserID           string   `json:"user_id,omitempty"`
		RoleID           []string `json:"role_id"`
		Name             string   `json:"name"`
		Email            string   `json:"email,omitempty"`
		JoinProjectTime  *string  `json:"join_project_time"`
		RealJoinTime     string   `json:"real_join_time"`
		Status           string   `json:"status"`
		Allocation       string   `json:"allocation"`
		LeaveProjectTime *string  `json:"leave_project_time"`
	}

	AddWorkspaceMemberRequest struct {
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		Nick        *string       `json:"nick,omitempty"`         // [必须]用户英文昵称
		CompanyID   *int          `json:"company_id,omitempty"`   // 成员所在公司ID，云端必填
		RoleIDs     *Multi[int64] `json:"role_ids,omitempty"`     // 角色组，多个使用逗号分隔
	}

	AddWorkspaceMemberResponse struct {
		Success bool `json:"success,omitempty"` // 是否添加成功
	}

	GetWorkspaceRolesRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	WorkspaceRole struct {
		ID   string `json:"id,omitempty"`   // 角色ID
		Name string `json:"name,omitempty"` // 角色名称
	}

	UpdateWorkspaceInfoRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		Field       *string `json:"field,omitempty"`        // [必须]字段名
		Value       *string `json:"value,omitempty"`        // [必须]字段值
	}

	GetMemberActivityLogRequest struct {
		WorkspaceID    *int           `url:"workspace_id,omitempty"`    // [必须]项目 id 为公司id则查询所有项目
		CompanyOnly    *int           `url:"company_only,omitempty"`    // [可选]为1则仅返回公司级活动日志 要求workspace_id=公司id & company_only=1
		Limit          *int           `url:"limit,omitempty"`           // [可选]设置返回数量限制，默认为20
		Page           *int           `url:"page,omitempty"`            // [可选]返回当前数量限制下第N页的数据，默认为1（第一页）
		StartTime      *string        `url:"start_time,omitempty"`      // [可选]起始时间，精确到分钟，格式为Y-m-d H:i 只能查最近半年内的数据
		EndTime        *string        `url:"end_time,omitempty"`        // [可选]终止时间，精确到分钟，格式为Y-m-d H:i 只能查最近半年内的数据
		Operator       *string        `url:"operator,omitempty"`        // [可选]操作人昵称
		OperateType    *OperateType   `url:"operate_type,omitempty"`    // [可选]操作类型，默认为所有，可以填写add,delete,download,upload中的一个
		OperatorObject *OperateObject `url:"operator_object,omitempty"` // [可选]操作对象，默认为所有，可以填写attachment,board,bug,document,iteration,launch,member_activity_log,release,story,task,tcase,testplan,wiki中的一个
		IP             *string        `url:"ip,omitempty"`              // [可选]请求IP条件，严格匹配
	}

	MemberActivityLog struct {
		ID            string        `json:"id,omitempty"`
		Action        string        `json:"action,omitempty"`
		Created       string        `json:"created,omitempty"`
		Creator       string        `json:"creator,omitempty"`
		ProjectName   string        `json:"project_name,omitempty"`
		OperateType   OperateType   `json:"operate_type,omitempty"`
		OperateObject OperateObject `json:"operate_object,omitempty"`
		Title         string        `json:"title,omitempty"`
		URL           string        `json:"url,omitempty"`
		IP            string        `json:"ip,omitempty"`
		UA            string        `json:"ua,omitempty"`
	}

	GetMemberActivityLogResponse struct {
		PerPage      string               `json:"perPage"`
		TotalItems   int                  `json:"totalItems"`
		CurrentPage  string               `json:"currentPage"`
		Records      []*MemberActivityLog `json:"records"`
		OperateTypes struct {
			Add      string `json:"add"`
			Delete   string `json:"delete"`
			Upload   string `json:"upload"`
			Download string `json:"download"`
		} `json:"operate_types"`
		OperateObjects struct {
			Board             string `json:"board"`
			Story             string `json:"story"`
			Bug               string `json:"bug"`
			Iteration         string `json:"iteration"`
			Wiki              string `json:"wiki"`
			Document          string `json:"document"`
			Attachment        string `json:"attachment"`
			Task              string `json:"task"`
			Tcase             string `json:"tcase"`
			Testplan          string `json:"testplan"`
			Launch            string `json:"launch"`
			Release           string `json:"release"`
			MemberActivityLog string `json:"member_activity_log"`
		} `json:"operate_objects"`
	}
)

type WorkspaceService interface {
	// GetSubWorkspaces 获取子项目信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_sub_workspaces.html
	GetSubWorkspaces(
		ctx context.Context, request *GetSubWorkspacesRequest, opts ...RequestOption,
	) ([]*Workspace, *Response, error)

	// 获取项目信息

	// GetWorkspaceInfo 获取项目信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workspace_info.html
	GetWorkspaceInfo(
		ctx context.Context, request *GetWorkspaceInfoRequest, opts ...RequestOption,
	) (*Workspace, *Response, error)

	// GetUsers 获取项目成员列表/指定项目成员
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/users.html
	GetUsers(
		ctx context.Context, request *GetUsersRequest, opts ...RequestOption,
	) ([]*User, *Response, error)

	// AddWorkspaceMember 添加项目成员
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/add_workspace_member.html
	AddWorkspaceMember(
		ctx context.Context, request *AddWorkspaceMemberRequest, opts ...RequestOption,
	) (*AddWorkspaceMemberResponse, *Response, error)

	// GetCompanyWorkspaces 获取公司项目列表
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/projects.html
	GetCompanyWorkspaces(
		ctx context.Context, request *GetCompanyWorkspacesRequest, opts ...RequestOption,
	) ([]*Workspace, *Response, error)

	// GetWorkspaceRoles 获取用户组ID对照关系
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/roles.html
	GetWorkspaceRoles(
		ctx context.Context, request *GetWorkspaceRolesRequest, opts ...RequestOption,
	) ([]*WorkspaceRole, *Response, error)

	// GetUserParticipantWorkspaces 获取用户参与的项目列表
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/user_participant_projects.html
	GetUserParticipantWorkspaces(
		ctx context.Context, request *GetUserParticipantWorkspacesRequest, opts ...RequestOption,
	) ([]*Workspace, *Response, error)

	// GetWorkspaceCustomFieldsSettings 获取项目自定义字段
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/workspace_custom_field_settings.html
	GetWorkspaceCustomFieldsSettings(
		ctx context.Context, request *GetWorkspaceCustomFieldsSettingsRequest, opts ...RequestOption,
	) ([]*WorkspaceCustomFieldsSetting, *Response, error)

	// UpdateWorkspaceInfo 更新项目信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/update_workspace_info.html
	UpdateWorkspaceInfo(
		ctx context.Context, request *UpdateWorkspaceInfoRequest, opts ...RequestOption,
	) (string, *Response, error)

	// GetWorkspaceDocuments 获取项目文档
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workspace_documents.html
	GetWorkspaceDocuments(
		ctx context.Context, request *GetWorkspaceDocumentsRequest, opts ...RequestOption,
	) ([]*WorkspaceDocument, *Response, error)

	// SetCustomWorkCalendar 设置自定义工作日历
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/set_custom_work_calendar.html
	SetCustomWorkCalendar(
		ctx context.Context, request *SetCustomWorkCalendarRequest, opts ...RequestOption,
	) (*SetCustomWorkCalendarResponse, *Response, error)

	// EnableWorkCalendar 设置启用工作日历
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/enable_work_calendar.html
	EnableWorkCalendar(
		ctx context.Context, request *EnableWorkCalendarRequest, opts ...RequestOption,
	) (*EnableWorkCalendarResponse, *Response, error)

	// GetCustomWorkCalendar 获取自定义工作日历详情
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_custom_work_calendar.html
	GetCustomWorkCalendar(
		ctx context.Context, request *GetCustomWorkCalendarRequest, opts ...RequestOption,
	) (*CustomWorkCalendar, *Response, error)

	// GetWorkCalendarSettings 获取工作日历设置列表及启用选项
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_work_calendar_settings.html
	GetWorkCalendarSettings(
		ctx context.Context, request *GetWorkCalendarSettingsRequest, opts ...RequestOption,
	) ([]*WorkCalendarSetting, *Response, error)

	// GetWorkItemsLongIDByShortIDs 通过工作项短id换长id
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workitems_long_id_by_short_ids.html
	GetWorkItemsLongIDByShortIDs(
		ctx context.Context, request *GetWorkItemsLongIDByShortIDsRequest, opts ...RequestOption,
	) (*GetWorkItemsLongIDByShortIDsResponse, *Response, error)

	// GetMemberActivityLog 获取成员活动日志
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/member_activity_log.html
	GetMemberActivityLog(
		ctx context.Context, request *GetMemberActivityLogRequest, opts ...RequestOption,
	) (*GetMemberActivityLogResponse, *Response, error)
}

type workspaceService struct {
	client *Client
}

var _ WorkspaceService = (*workspaceService)(nil)

func NewWorkspaceService(client *Client) WorkspaceService {
	return &workspaceService{
		client: client,
	}
}

func (s *workspaceService) GetSubWorkspaces(
	ctx context.Context, request *GetSubWorkspacesRequest, opts ...RequestOption,
) ([]*Workspace, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/sub_workspaces", request, opts)
	if err != nil {
		return nil, nil, err
	}

	return s.doWorkspaces(req)
}

func (s *workspaceService) GetWorkspaceInfo(
	ctx context.Context, request *GetWorkspaceInfoRequest, opts ...RequestOption,
) (*Workspace, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/get_workspace_info", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Workspace *Workspace `json:"Workspace"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Workspace, resp, nil
}

func (s *workspaceService) GetUsers(
	ctx context.Context, request *GetUsersRequest, opts ...RequestOption,
) ([]*User, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/users", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		UserWorkspace *User `json:"UserWorkspace"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	users := make([]*User, 0, len(items))
	for _, item := range items {
		users = append(users, item.UserWorkspace)
	}

	return users, resp, nil
}

func (s *workspaceService) GetCompanyWorkspaces(
	ctx context.Context, request *GetCompanyWorkspacesRequest, opts ...RequestOption,
) ([]*Workspace, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/projects", request, opts)
	if err != nil {
		return nil, nil, err
	}

	return s.doWorkspaces(req)
}

func (s *workspaceService) AddWorkspaceMember(
	ctx context.Context, request *AddWorkspaceMemberRequest, opts ...RequestOption,
) (*AddWorkspaceMemberResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "workspaces/add_workspace_member", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(AddWorkspaceMemberResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *workspaceService) GetWorkspaceRoles(
	ctx context.Context, request *GetWorkspaceRolesRequest, opts ...RequestOption,
) ([]*WorkspaceRole, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "roles", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items map[string]string
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	roles := make([]*WorkspaceRole, 0, len(items))
	for id, name := range items {
		roles = append(roles, &WorkspaceRole{
			ID:   id,
			Name: name,
		})
	}

	return roles, resp, nil
}

func (s *workspaceService) GetUserParticipantWorkspaces(
	ctx context.Context, request *GetUserParticipantWorkspacesRequest, opts ...RequestOption,
) ([]*Workspace, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/user_participant_projects", request, opts)
	if err != nil {
		return nil, nil, err
	}

	return s.doWorkspaces(req)
}

func (s *workspaceService) GetWorkspaceCustomFieldsSettings(
	ctx context.Context, request *GetWorkspaceCustomFieldsSettingsRequest, opts ...RequestOption,
) ([]*WorkspaceCustomFieldsSetting, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/workspace_custom_field_settings", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		CustomFieldConfig *WorkspaceCustomFieldsSetting `json:"CustomFieldConfig,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	settings := make([]*WorkspaceCustomFieldsSetting, 0, len(items))
	for _, item := range items {
		settings = append(settings, item.CustomFieldConfig)
	}

	return settings, resp, nil
}

func (s *workspaceService) UpdateWorkspaceInfo(
	ctx context.Context, request *UpdateWorkspaceInfoRequest, opts ...RequestOption,
) (string, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "workspaces/update_workspace_info", request, opts)
	if err != nil {
		return "", nil, err
	}

	var result string
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return "", resp, err
	}

	return result, resp, nil
}

func (s *workspaceService) GetWorkspaceDocuments(
	ctx context.Context, request *GetWorkspaceDocumentsRequest, opts ...RequestOption,
) ([]*WorkspaceDocument, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "documents/get_workspace_documents", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Document *WorkspaceDocument `json:"Document,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	documents := make([]*WorkspaceDocument, 0, len(items))
	for _, item := range items {
		documents = append(documents, item.Document)
	}

	return documents, resp, nil
}

func (s *workspaceService) SetCustomWorkCalendar(
	ctx context.Context, request *SetCustomWorkCalendarRequest, opts ...RequestOption,
) (*SetCustomWorkCalendarResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "workspaces/set_custom_work_calendar", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(SetCustomWorkCalendarResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *workspaceService) EnableWorkCalendar(
	ctx context.Context, request *EnableWorkCalendarRequest, opts ...RequestOption,
) (*EnableWorkCalendarResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "workspaces/enable_work_calendar", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(EnableWorkCalendarResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *workspaceService) GetCustomWorkCalendar(
	ctx context.Context, request *GetCustomWorkCalendarRequest, opts ...RequestOption,
) (*CustomWorkCalendar, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/get_custom_work_calendar", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(CustomWorkCalendar)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *workspaceService) GetWorkCalendarSettings(
	ctx context.Context, request *GetWorkCalendarSettingsRequest, opts ...RequestOption,
) ([]*WorkCalendarSetting, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/get_work_calendar_settings", request, opts)
	if err != nil {
		return nil, nil, err
	}

	settings := make([]*WorkCalendarSetting, 0)
	resp, err := s.client.Do(req, &settings)
	if err != nil {
		return nil, resp, err
	}

	return settings, resp, nil
}

func (s *workspaceService) GetWorkItemsLongIDByShortIDs(
	ctx context.Context, request *GetWorkItemsLongIDByShortIDsRequest, opts ...RequestOption,
) (*GetWorkItemsLongIDByShortIDsResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/get_workitems_long_id_by_short_ids", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetWorkItemsLongIDByShortIDsResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *workspaceService) GetMemberActivityLog(
	ctx context.Context, request *GetMemberActivityLogRequest, opts ...RequestOption,
) (*GetMemberActivityLogResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/member_activity_log", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetMemberActivityLogResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

type workspaceItems []*Workspace

func (items *workspaceItems) UnmarshalJSON(data []byte) error {
	var list []struct {
		Workspace *Workspace `json:"Workspace"`
	}
	if err := json.Unmarshal(data, &list); err == nil {
		*items = make([]*Workspace, 0, len(list))
		for _, item := range list {
			*items = append(*items, item.Workspace)
		}

		return nil
	}

	var single struct {
		Workspace *Workspace `json:"Workspace"`
	}
	if err := json.Unmarshal(data, &single); err != nil {
		return err
	}
	if single.Workspace == nil {
		*items = nil
		return nil
	}

	*items = []*Workspace{single.Workspace}

	return nil
}

func (s *workspaceService) doWorkspaces(req *http.Request) ([]*Workspace, *Response, error) {
	var items workspaceItems
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	return items, resp, nil
}
