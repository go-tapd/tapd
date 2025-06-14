package tapd

import (
	"context"
	"net/http"
)

type (
	Timesheet struct {
		ID          string     `json:"id,omitempty"`           // id
		EntityType  EntityType `json:"entity_type,omitempty"`  // 对象类型，如story、task、bug等
		EntityID    string     `json:"entity_id,omitempty"`    // 对象ID
		Timespent   string     `json:"timespent,omitempty"`    // 花费工时
		Spentdate   string     `json:"spentdate,omitempty"`    // 花费日期
		Owner       string     `json:"owner,omitempty"`        // 花费创建人
		Created     string     `json:"created,omitempty"`      // 创建时间
		Modified    string     `json:"modified,omitempty"`     // 最后修改时间
		WorkspaceID string     `json:"workspace_id,omitempty"` // 项目ID
		Memo        string     `json:"memo,omitempty"`         // 花费描述
		IsDelete    string     `json:"is_delete,omitempty"`    // 是否已删除
	}

	CreateTimesheetRequest struct {
		EntityType  *EntityType `json:"entity_type,omitempty"`  // [必须]对象类型，如story、task、bug等
		EntityID    *int64      `json:"entity_id,omitempty"`    // [必须]对象ID
		Timespent   *string     `json:"timespent,omitempty"`    // [必须]花费工时
		Timeremain  *string     `json:"timeremain,omitempty"`   // 剩余工时
		Spentdate   *string     `json:"spentdate,omitempty"`    // 花费日期
		Owner       *string     `json:"owner,omitempty"`        // [必须]花费创建人
		WorkspaceID *int        `json:"workspace_id,omitempty"` // [必须]项目ID
		Memo        *string     `json:"memo,omitempty"`         // 花费描述
	}

	GetTimesheetsRequest struct {
		// [可选]id 支持多ID查询
		ID *Multi[int64] `url:"id,omitempty"`

		// [必选]项目ID
		WorkspaceID *int `url:"workspace_id,omitempty"`

		// [可选]对象类型，如story、task、bug等
		EntityType *EntityType `url:"entity_type,omitempty"`

		// [可选]对象ID
		EntityID *int64 `url:"entity_id,omitempty"`

		// [可选]花费工时
		Timespent *string `url:"timespent,omitempty"`

		// [可选]花费日期 支持时间查询
		Spentdate *string `url:"spentdate,omitempty"`

		// [可选]最后修改时间 支持时间查询
		Modified *string `url:"modified,omitempty"`

		// [可选]花费创建人
		Owner *string `url:"owner,omitempty"`

		// [可选]值=0不返回父需求的花费
		IncludeParentStoryTimesheet *int `url:"include_parent_story_timesheet,omitempty"`

		// [可选]创建时间 支持时间查询
		Created *string `url:"created,omitempty"`

		// [可选]花费描述
		Memo *string `url:"memo,omitempty"`

		// [可选]是否已删除。默认取 0，不返回已删除的工时记录。取 1 可以返回已删除的记录
		IsDelete *int `url:"is_delete,omitempty"`

		// [可选]设置返回数量限制，默认为30
		Limit *int `url:"limit,omitempty"`

		// [可选]返回当前数量限制下第N页的数据，默认为1（第一页）
		Page *int `url:"page,omitempty"`

		// [可选]排序规则，规则：字段名 ASC或者DESC，然后 urlencode 如按创建时间逆序
		Order *Order `url:"order,omitempty"`

		// [可选]设置获取的字段，多个字段间以','逗号隔开
		Fields *Multi[string] `url:"fields,omitempty"`
	}

	GetTimesheetsCountRequest struct {
		// [可选]id 支持多ID查询
		ID *Multi[int64] `url:"id,omitempty"`

		// [必选]项目ID
		WorkspaceID *int `url:"workspace_id,omitempty"`

		// [可选]对象类型，如story、task、bug等
		EntityType *EntityType `url:"entity_type,omitempty"`

		// [可选]对象ID
		EntityID *int64 `url:"entity_id,omitempty"`

		// [可选]花费工时
		Timespent *string `url:"timespent,omitempty"`

		// [可选]花费日期 支持时间查询
		Spentdate *string `url:"spentdate,omitempty"`

		// [可选]最后修改时间 支持时间查询
		Modified *string `url:"modified,omitempty"`

		// [可选]花费创建人
		Owner *string `url:"owner,omitempty"`

		// [可选]值=0不返回父需求的花费
		IncludeParentStoryTimesheet *int `url:"include_parent_story_timesheet,omitempty"`

		// [可选]创建时间 支持时间查询
		Created *string `url:"created,omitempty"`

		// [可选]花费描述
		Memo *string `url:"memo,omitempty"`

		// [可选]是否已删除。默认取 0，不返回已删除的工时记录。取 1 可以返回已删除的记录
		IsDelete *int `url:"is_delete,omitempty"`
	}

	UpdateTimesheetRequest struct {
		ID          *int64  `json:"id"`                     // [必须]工时花费ID
		Timespent   *string `json:"timespent,omitempty"`    // [可选]花费工时
		Timeremain  *string `json:"timeremain,omitempty"`   // [可选]剩余工时
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		Memo        *string `json:"memo,omitempty"`         // [可选]花费描述
	}
)

type TimesheetService interface {
	// CreateTimesheet 创建工时花费
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/add_timesheet.html
	CreateTimesheet(ctx context.Context, request *CreateTimesheetRequest, opts ...RequestOption) (*Timesheet, *Response, error)

	// GetTimesheets 获取工时花费
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/get_timesheets.html
	GetTimesheets(ctx context.Context, request *GetTimesheetsRequest, opts ...RequestOption) ([]*Timesheet, *Response, error)

	// GetTimesheetsCount 获取工时花费的数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/get_timesheets_count.html
	GetTimesheetsCount(ctx context.Context, request *GetTimesheetsCountRequest, opts ...RequestOption) (int, *Response, error)

	// UpdateTimesheet 更新工时花费
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/timesheet/update_timesheet.html
	UpdateTimesheet(ctx context.Context, request *UpdateTimesheetRequest, opts ...RequestOption) (*Timesheet, *Response, error)
}

type timesheetService struct {
	client *Client
}

var _ TimesheetService = (*timesheetService)(nil)

func NewTimesheetService(client *Client) TimesheetService {
	return &timesheetService{
		client: client,
	}
}

func (s *timesheetService) CreateTimesheet(
	ctx context.Context, request *CreateTimesheetRequest, opts ...RequestOption,
) (*Timesheet, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "timesheets", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Timesheet *Timesheet `json:"Timesheet"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Timesheet, resp, nil
}

func (s *timesheetService) GetTimesheets(
	ctx context.Context, request *GetTimesheetsRequest, opts ...RequestOption,
) ([]*Timesheet, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "timesheets", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Timesheet *Timesheet `json:"Timesheet"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	timesheets := make([]*Timesheet, 0, len(items))
	for _, item := range items {
		timesheets = append(timesheets, item.Timesheet)
	}

	return timesheets, resp, nil
}

func (s *timesheetService) GetTimesheetsCount(
	ctx context.Context, request *GetTimesheetsCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "timesheets/count", request, opts)
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

func (s *timesheetService) UpdateTimesheet(
	ctx context.Context, request *UpdateTimesheetRequest, opts ...RequestOption,
) (*Timesheet, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "timesheets", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Timesheet *Timesheet `json:"Timesheet"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Timesheet, resp, nil
}
