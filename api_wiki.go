package tapd

import (
	"context"
	"net/http"
)

type (
	Wiki struct {
		ID                  string `json:"id,omitempty"`                   // ID
		Name                string `json:"name,omitempty"`                 // 标题
		WorkspaceID         string `json:"workspace_id,omitempty"`         // 项目ID
		Description         string `json:"description,omitempty"`          // 富文本
		MarkdownDescription string `json:"markdown_description,omitempty"` // Markdown
		IsRich              string `json:"is_rich,omitempty"`              // 是否富文本
		ParentWikiID        string `json:"parent_wiki_id,omitempty"`       // 父wiki ID
		Author              string `json:"author,omitempty"`               // 修改人
		Creator             string `json:"creator,omitempty"`              // 创建人
		Note                string `json:"note,omitempty"`                 // 备注
		ViewCount           string `json:"view_count,omitempty"`           // 浏览量
		Created             string `json:"created,omitempty"`              // 创建时间
		Modified            string `json:"modified,omitempty"`             // 最后修改时间
		Modifier            string `json:"modifier,omitempty"`             // 最后修改人
	}

	CreateWikiRequest struct {
		Name                *string `json:"name,omitempty"`                 // [必须]标题
		MarkdownDescription *string `json:"markdown_description,omitempty"` // Markdown
		Description         *string `json:"description,omitempty"`          // 富文本
		Creator             *string `json:"creator,omitempty"`              // [必须]创建人
		Note                *string `json:"note,omitempty"`                 // 备注
		WorkspaceID         *int    `json:"workspace_id,omitempty"`         // [必须]项目ID
		ParentWikiID        *string `json:"parent_wiki_id,omitempty"`       // 父wiki ID
	}

	GetWikisRequest struct {
		ID          *int64         `url:"id,omitempty"`           // id
		Name        *string        `url:"name,omitempty"`         // 标题
		Modifier    *string        `url:"modifier,omitempty"`     // 修改人
		Creator     *string        `url:"creator,omitempty"`      // 创建人
		Note        *string        `url:"note,omitempty"`         // 备注
		ViewCount   *string        `url:"view_count,omitempty"`   // 浏览量
		Created     *string        `url:"created,omitempty"`      // 创建时间，支持时间查询
		Modified    *string        `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetWikisCountRequest struct {
		Name        *string `url:"name,omitempty"`         // 标题，支持模糊匹配
		Modifier    *string `url:"modifier,omitempty"`     // 修改人
		Creator     *string `url:"creator,omitempty"`      // 创建人
		Note        *string `url:"note,omitempty"`         // 备注
		ViewCount   *string `url:"view_count,omitempty"`   // 浏览量
		Created     *string `url:"created,omitempty"`      // 创建时间，支持时间查询
		Modified    *string `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	UpdateWikiRequest struct {
		ID                  *int64  `json:"id,omitempty"`                   // [必须]ID
		Name                *string `json:"name,omitempty"`                 // 标题
		MarkdownDescription *string `json:"markdown_description,omitempty"` // Markdown
		Description         *string `json:"description,omitempty"`          // 富文本
		Note                *string `json:"note,omitempty"`                 // 备注
		WorkspaceID         *int    `json:"workspace_id,omitempty"`         // [必须]项目ID
		ParentWikiID        *string `json:"parent_wiki_id,omitempty"`       // 父wiki ID
	}
)

// WikiService Wiki 服务。
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/
type WikiService interface {
	// CreateWiki 创建 wiki
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/add_tapd_wiki.html
	CreateWiki(ctx context.Context, request *CreateWikiRequest, opts ...RequestOption) (*Wiki, *Response, error)

	// GetWikis 获取 wiki
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis.html
	GetWikis(ctx context.Context, request *GetWikisRequest, opts ...RequestOption) ([]*Wiki, *Response, error)

	// GetWikisCount 获取 Wiki 数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_count.html
	GetWikisCount(ctx context.Context, request *GetWikisCountRequest, opts ...RequestOption) (int, *Response, error)

	// UpdateWiki 更新 wiki
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/update_tapd_wiki.html
	UpdateWiki(ctx context.Context, request *UpdateWikiRequest, opts ...RequestOption) (*Wiki, *Response, error)
}

type wikiService struct {
	client *Client
}

var _ WikiService = (*wikiService)(nil)

func NewWikiService(client *Client) WikiService {
	return &wikiService{
		client: client,
	}
}

func (s *wikiService) CreateWiki(
	ctx context.Context, request *CreateWikiRequest, opts ...RequestOption,
) (*Wiki, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tapd_wikis", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Wiki *Wiki `json:"Wiki"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Wiki, resp, nil
}

func (s *wikiService) GetWikis(
	ctx context.Context, request *GetWikisRequest, opts ...RequestOption,
) ([]*Wiki, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Wiki *Wiki `json:"Wiki"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	wikis := make([]*Wiki, 0, len(items))
	for _, item := range items {
		wikis = append(wikis, item.Wiki)
	}

	return wikis, resp, nil
}

func (s *wikiService) GetWikisCount(
	ctx context.Context, request *GetWikisCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis/count", request, opts)
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

func (s *wikiService) UpdateWiki(
	ctx context.Context, request *UpdateWikiRequest, opts ...RequestOption,
) (*Wiki, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tapd_wikis", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Wiki *Wiki `json:"Wiki"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Wiki, resp, nil
}
