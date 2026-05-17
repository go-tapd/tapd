package tapd

import (
	"context"
	"encoding/json"
	"net/http"
)

// StringNumber 兼容官方响应中同一字段可能返回字符串或数字的 ID。
type StringNumber string

func (s *StringNumber) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*s = ""
		return nil
	}

	var value string
	if err := json.Unmarshal(data, &value); err == nil {
		*s = StringNumber(value)
		return nil
	}

	var number json.Number
	if err := json.Unmarshal(data, &number); err != nil {
		return err
	}

	*s = StringNumber(number.String())

	return nil
}

func (s StringNumber) String() string {
	return string(s)
}

// CodeCommitRelatedType 源码提交关联类型。
type CodeCommitRelatedType string

const (
	CodeCommitRelatedTypeAll        CodeCommitRelatedType = "all"
	CodeCommitRelatedTypeBranch     CodeCommitRelatedType = "branch"
	CodeCommitRelatedTypeSourceCode CodeCommitRelatedType = "source_code"
)

type (
	AddCodeCommitInfoRequest struct {
		WorkspaceID *int      `json:"workspace_id,omitempty"` // [必须]项目ID
		CommitID    *string   `json:"commit_id,omitempty"`    // [必须]提交ID
		Author      *string   `json:"author,omitempty"`       // [必须]代码提交人
		Message     *string   `json:"message,omitempty"`      // [必须]提交信息
		Files       *[]string `json:"files,omitempty"`        // [必须]变更文件
		Repo        *string   `json:"repo,omitempty"`         // [必须]仓库名
		RepoID      *string   `json:"repo_id,omitempty"`      // [必须]仓库ID
		CommitTime  *string   `json:"commit_time,omitempty"`  // [必须]提交时间
		GitEnv      *string   `json:"git_env,omitempty"`      // 信息来源，github、gitlab、svn、p4 等
		RepoURL     *string   `json:"repo_url,omitempty"`     // 仓库链接
		CommitURL   *string   `json:"commit_url,omitempty"`   // 提交链接
	}

	GetCodeCommitInfosRequest struct {
		WorkspaceID *int                   `url:"workspace_id,omitempty"` // [必须]项目ID
		Type        *EntityType            `url:"type,omitempty"`         // [必须]TAPD业务对象类型，story、bug、task
		ObjectID    *int64                 `url:"object_id,omitempty"`    // [必须]TAPD业务对象ID
		CommitTime  *string                `url:"commit_time,omitempty"`  // 提交时间查询条件
		RelatedType *CodeCommitRelatedType `url:"related_type,omitempty"` // 关联类型，all、branch、source_code
		Limit       *int                   `url:"limit,omitempty"`        // 返回数量限制，默认30，最大200
		Page        *int                   `url:"page,omitempty"`         // 当前页，默认1
	}

	GetCommitObjectsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		CommitID    *Multi[string] `url:"commit_id,omitempty"`    // [必须]提交ID，多个以逗号分隔
		EntityType  *EntityType    `url:"entity_type,omitempty"`  // [必须]业务对象类型，story、bug、task
		SCMType     *string        `url:"scm_type,omitempty"`     // 来源类型
		Limit       *int           `url:"limit,omitempty"`        // 返回数量限制，默认30，最大200
		Page        *int           `url:"page,omitempty"`         // 当前页，默认1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 返回字段，多个以逗号分隔
	}

	CodeCommitInfo struct {
		ID              string                   `json:"id,omitempty"`                // 记录ID
		HookUserName    string                   `json:"hook_user_name,omitempty"`    // 代码提交人
		CommitID        string                   `json:"commit_id,omitempty"`         // 提交ID
		UserName        string                   `json:"user_name,omitempty"`         // 关联用户
		UserID          string                   `json:"user_id,omitempty"`           // 关联用户ID
		WorkspaceID     StringNumber             `json:"workspace_id,omitempty"`      // 项目ID
		Message         string                   `json:"message,omitempty"`           // 提交信息
		Path            string                   `json:"path,omitempty"`              // 提交地址
		WebURL          string                   `json:"web_url,omitempty"`           // 仓库地址
		HookProjectName string                   `json:"hook_project_name,omitempty"` // 仓库名
		CommitTime      string                   `json:"commit_time,omitempty"`       // 提交时间
		Created         string                   `json:"created,omitempty"`           // 创建时间
		Ref             string                   `json:"ref,omitempty"`               // 分支引用
		RefStatus       string                   `json:"ref_status,omitempty"`        // 分支状态
		GitEnv          string                   `json:"git_env,omitempty"`           // 信息来源
		FileCommit      string                   `json:"file_commit,omitempty"`       // 变更文件
		RepoID          string                   `json:"repo_id,omitempty"`           // 仓库ID
		BranchID        string                   `json:"branch_id,omitempty"`         // 分支ID
		FileSort        map[string]int           `json:"file_sort,omitempty"`         // 文件排序
		Related         []*CodeCommitInfoRelated `json:"related,omitempty"`           // 关联结果
	}

	CodeCommitInfoRelated struct {
		Type        EntityType   `json:"type,omitempty"`         // 业务对象类型
		ObjectID    string       `json:"object_id,omitempty"`    // 业务对象ID
		CommitID    string       `json:"commit_id,omitempty"`    // 提交记录ID
		WorkspaceID StringNumber `json:"workspace_id,omitempty"` // 项目ID
		Code        string       `json:"code,omitempty"`         // 关联结果代码
	}

	CommitObject struct {
		Story *Story `json:"Story,omitempty"` // 关联需求
		Bug   *Bug   `json:"Bug,omitempty"`   // 关联缺陷
		Task  *Task  `json:"Task,omitempty"`  // 关联任务
	}
)

type SourceService interface {
	// AddCodeCommitInfo 保存Commit提交数据
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/add_code_commit_info.html
	AddCodeCommitInfo(
		ctx context.Context, request *AddCodeCommitInfoRequest, opts ...RequestOption,
	) (*CodeCommitInfo, *Response, error)

	// GetCodeCommitInfos 获取GIT关联提交数据(GitCommit)
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/get_code_commit_infos.html
	GetCodeCommitInfos(
		ctx context.Context, request *GetCodeCommitInfosRequest, opts ...RequestOption,
	) ([]*CodeCommitInfo, *Response, error)

	// GetCommitObjects 获取指定commit关联的业务对象
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/source/get_commit_objects.html
	GetCommitObjects(
		ctx context.Context, request *GetCommitObjectsRequest, opts ...RequestOption,
	) ([]*CommitObject, *Response, error)
}

type sourceService struct {
	client *Client
}

var _ SourceService = (*sourceService)(nil)

func NewSourceService(client *Client) SourceService {
	return &sourceService{
		client: client,
	}
}

func (s *sourceService) AddCodeCommitInfo(
	ctx context.Context, request *AddCodeCommitInfoRequest, opts ...RequestOption,
) (*CodeCommitInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "code_commit_infos", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(CodeCommitInfo)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

func (s *sourceService) GetCodeCommitInfos(
	ctx context.Context, request *GetCodeCommitInfosRequest, opts ...RequestOption,
) ([]*CodeCommitInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "code_commit_infos", request, opts)
	if err != nil {
		return nil, nil, err
	}

	infos := make([]*CodeCommitInfo, 0)
	resp, err := s.client.Do(req, &infos)
	if err != nil {
		return nil, resp, err
	}

	return infos, resp, nil
}

func (s *sourceService) GetCommitObjects(
	ctx context.Context, request *GetCommitObjectsRequest, opts ...RequestOption,
) ([]*CommitObject, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "code_commit_objects/workitems", request, opts)
	if err != nil {
		return nil, nil, err
	}

	objects := make([]*CommitObject, 0)
	resp, err := s.client.Do(req, &objects)
	if err != nil {
		return nil, resp, err
	}

	return objects, resp, nil
}
