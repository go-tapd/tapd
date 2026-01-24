package tapd

import (
	"context"
	"net/http"
)

type (
	GetWorkspaceInfoRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	Workspace struct {
		ID                string  `json:"id,omitempty"`
		Name              string  `json:"name,omitempty"`
		PrettyName        string  `json:"pretty_name,omitempty"`
		Category          string  `json:"category,omitempty"`
		Status            string  `json:"status,omitempty"`
		Description       string  `json:"description,omitempty"`
		Creator           string  `json:"creator,omitempty"`
		Created           string  `json:"created,omitempty"`
		BeginDate         *string `json:"begin_date,omitempty"`
		EndDate           *string `json:"end_date,omitempty"`
		Secrecy           string  `json:"secrecy,omitempty"`
		ExternalOn        string  `json:"external_on,omitempty"`
		NewTask           string  `json:"new_task,omitempty"`
		CompanyID         string  `json:"company_id,omitempty"`
		ProductType       *string `json:"product_type,omitempty"`
		PlatformType      *string `json:"platform_type,omitempty"`
		IsSelfDevelopment *string `json:"is_self_development,omitempty"`
		Objective         string  `json:"objective,omitempty"`
		Schedule          *string `json:"schedule,omitempty"`
		Milestone         *string `json:"milestone,omitempty"`
		Risk              *string `json:"risk,omitempty"`
		Closed            *string `json:"closed,omitempty"`
	}

	GetUsersRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		User        *Multi[string] `url:"user,omitempty"`         // [可选]用户昵称或ID
		Fields      *Multi[string] `url:"fields,omitempty"`       // [可选]返回的字段列表，user,user_id,role_id,name,email,real_join_time 可选，以,分隔
	}

	User struct {
		User             string   `json:"user"`
		RoleID           []string `json:"role_id"`
		Name             string   `json:"name"`
		JoinProjectTime  *string  `json:"join_project_time"`
		RealJoinTime     string   `json:"real_join_time"`
		Status           string   `json:"status"`
		Allocation       string   `json:"allocation"`
		LeaveProjectTime *string  `json:"leave_project_time"`
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

	GetCustomWorkCalendarRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		Year        *string `url:"year,omitempty"`         // [必须]表示哪一年
	}

	CustomWorkCalendar struct {
		Weekdays []string `json:"weekdays,omitempty"`
		Holidays []string `json:"holidays,omitempty"`
		Workdays []string `json:"workdays,omitempty"`
	}
)

type WorkspaceService interface {
	// 获取子项目信息
	// 获取项目信息

	// GetWorkspaceInfo 获取项目信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_workspace_info.html
	GetWorkspaceInfo(
		ctx context.Context, request *GetWorkspaceInfoRequest, opts ...RequestOption,
	) (*Workspace, *Response, error)

	// GetUsers 获取指定项目成员
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/users.html
	GetUsers(
		ctx context.Context, request *GetUsersRequest, opts ...RequestOption,
	) ([]*User, *Response, error)

	// 添加项目成员
	// 获取公司项目列表
	// 获取用户组ID对照关系
	// 获取用户参与的项目列表
	// 获取项目成员列表
	// 获取项目自定义字段
	// 更新项目信息
	// 获取项目文档

	// GetMemberActivityLog 获取成员活动日志
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/member_activity_log.html
	GetMemberActivityLog(
		ctx context.Context, request *GetMemberActivityLogRequest, opts ...RequestOption,
	) (*GetMemberActivityLogResponse, *Response, error)

	// GetCustomWorkCalendar 获取自定义工作日历详情
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/workspace/get_custom_work_calendar.html
	GetCustomWorkCalendar(
		ctx context.Context, request *GetCustomWorkCalendarRequest, opts ...RequestOption,
	) (*CustomWorkCalendar, *Response, error)
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

func (s *workspaceService) GetCustomWorkCalendar(
	ctx context.Context, request *GetCustomWorkCalendarRequest, opts ...RequestOption,
) (*CustomWorkCalendar, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspaces/get_custom_work_calendar", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(CustomWorkCalendar)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}
