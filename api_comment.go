package tapd

import (
	"context"
	"net/http"
)

// CommentEntryType 评论类型
type CommentEntryType string

// CommentEntryType 评论类型
const (
	CommentEntryTypeBug       CommentEntryType = "bug"        // bug
	CommentEntryTypeBugRemark CommentEntryType = "bug_remark" // bug_remark （流转缺陷时候的评论）
	CommentEntryTypeStories   CommentEntryType = "stories"    // stories
	CommentEntryTypeTasks     CommentEntryType = "tasks"      // tasks
	CommentEntryTypeWiki      CommentEntryType = "wiki"       // wiki
	CommentEntryTypeMiniItems CommentEntryType = "mini_items" // mini_items
)

// String CommentEntryType to string
func (t CommentEntryType) String() string {
	return string(t)
}

// Comment 评论
type Comment struct {
	ID          string           `json:"id,omitempty"`           // 评论ID
	Title       string           `json:"title,omitempty"`        // 标题
	Description string           `json:"description,omitempty"`  // 内容
	Author      string           `json:"author,omitempty"`       // 评论人
	EntryType   CommentEntryType `json:"entry_type,omitempty"`   // 评论类型
	EntryID     string           `json:"entry_id,omitempty"`     // 评论所依附的业务对象实体id
	ReplyID     string           `json:"reply_id,omitempty"`     // 评论回复的ID
	RootID      string           `json:"root_id,omitempty"`      // 根评论ID
	Created     string           `json:"created,omitempty"`      // 创建时间
	Modified    string           `json:"modified,omitempty"`     // 最后更改时间
	WorkspaceID string           `json:"workspace_id,omitempty"` // 项目ID
}

// =====================================================================================================================

// CommentService 评论服务
type CommentService struct {
	client *Client
}

// CreateComment 添加评论接口
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/add_comment.html
func (s *CommentService) CreateComment(
	ctx context.Context, request *CreateCommentRequest, opts ...RequestOption,
) (*Comment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "comments", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Comment *Comment `json:"Comment"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Comment, resp, nil
}

type CreateCommentRequest struct {
	Title       *string           `json:"title,omitempty"`        // 标题
	Description *string           `json:"description,omitempty"`  // 内容
	Author      *string           `json:"author,omitempty"`       // 评论人
	EntryType   *CommentEntryType `json:"entry_type,omitempty"`   // 评论类型
	EntryID     *int64            `json:"entry_id,omitempty"`     // 评论所依附的业务对象实体id
	ReplyID     *int64            `json:"reply_id,omitempty"`     // 评论回复的ID
	RootID      *int64            `json:"root_id,omitempty"`      // 根评论ID
	WorkspaceID *int              `json:"workspace_id,omitempty"` // 项目ID
}

// GetComments 获取评论
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments.html
func (s *CommentService) GetComments(
	ctx context.Context, request *GetCommentsRequest, opts ...RequestOption,
) ([]*Comment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "comments", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Comment *Comment `json:"Comment"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	comments := make([]*Comment, 0, len(items))
	for _, item := range items {
		comments = append(comments, item.Comment)
	}

	return comments, resp, nil
}

type GetCommentsRequest struct {
	// 评论ID 支持多ID查询
	ID *Multi[int64] `url:"id,omitempty"`

	// 标题
	Title *string `url:"title,omitempty"`

	// 内容
	Description *string `url:"description,omitempty"`

	// 评论人
	Author *string `url:"author,omitempty"`

	// 评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。多个类型间以竖线隔开） 支持枚举查询
	EntryType *CommentEntryType `url:"entry_type,omitempty"`

	// 评论所依附的业务对象实体id
	EntryID *int64 `url:"entry_id,omitempty"`

	// 创建时间 支持时间查询
	Created *string `url:"created,omitempty"`

	// 最后更改时间 支持时间查询
	Modified *string `url:"modified,omitempty"`

	// 项目ID
	WorkspaceID *int `url:"workspace_id,omitempty"`

	// 根评论ID
	RootID *int64 `url:"root_id,omitempty"`

	// 评论回复的ID
	ReplyID *int64 `url:"reply_id,omitempty"`

	// 设置返回数量限制，默认为30
	Limit *int `url:"limit,omitempty"`

	// 返回当前数量限制下第N页的数据，默认为1（第一页）
	Page *int `url:"page,omitempty"`

	// 排序规则，规则：字段名 ASC或者DESC，然后 urlencode 如按创建时间逆序：order=created%20desc
	Order *Order `url:"order,omitempty"`

	// 设置获取的字段，多个字段间以','逗号隔开
	Fields *Multi[string] `url:"fields,omitempty"`
}

// GetCommentsCount 获取评论数量
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/get_comments_count.html
func (s *CommentService) GetCommentsCount(
	ctx context.Context, request *GetCommentsCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "comments/count", request, opts)
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

type GetCommentsCountRequest struct {
	// 评论ID 支持多ID查询
	ID *Multi[int64] `url:"id,omitempty"`

	// 标题
	Title *string `url:"title,omitempty"`

	// 内容
	Description *string `url:"description,omitempty"`

	// 评论人
	Author *string `url:"author,omitempty"`

	// 评论类型（取值： bug、 bug_remark （流转缺陷时候的评论）、 stories、 tasks 。多个类型间以竖线隔开） 支持枚举查询
	EntryType *CommentEntryType `url:"entry_type,omitempty"`

	// 评论所依附的业务对象实体id
	EntryID *int64 `url:"entry_id,omitempty"`

	// 创建时间 支持时间查询
	Created *string `url:"created,omitempty"`

	// 最后更改时间 支持时间查询
	Modified *string `url:"modified,omitempty"`

	// 项目ID
	WorkspaceID *int `url:"workspace_id,omitempty"`

	// 根评论ID
	RootID *int64 `url:"root_id,omitempty"`

	// 评论回复的ID
	ReplyID *int64 `url:"reply_id,omitempty"`
}

// UpdateComment 更新评论接口
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/comment/update_comment.html
func (s *CommentService) UpdateComment(
	ctx context.Context, request *UpdateCommentRequest, opts ...RequestOption,
) (*Comment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "comments", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Comment *Comment `json:"Comment"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Comment, resp, nil
}

type UpdateCommentRequest struct {
	WorkspaceID   *int    `json:"workspace_id,omitempty"`   // [必须]项目ID
	ID            *int64  `json:"id,omitempty"`             // [必须]评论ID
	Description   *string `json:"description,omitempty"`    // [必须]内容
	ChangeCreator *string `json:"change_creator,omitempty"` // 变更人
}
