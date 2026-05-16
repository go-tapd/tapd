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

	WikiDrawioData struct {
		ID     string `json:"id,omitempty"`     // drawio 数据ID
		Values string `json:"values,omitempty"` // drawio XML 数据
	}

	GetWikiDrawioDataRequest struct {
		ID          *int64  `url:"id,omitempty"`           // [必须]drawio 数据ID
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		Token       *string `url:"token,omitempty"`        // 验证用 token
	}

	WikiFollower struct {
		ID          string `json:"id,omitempty"`           // id
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		Created     string `json:"created,omitempty"`      // 创建时间
		WikiID      string `json:"wiki_id,omitempty"`      // 关联的 wiki id
		User        string `json:"user,omitempty"`         // 关注者昵称
	}

	GetWikiFollowersRequest struct {
		ID          *int64         `url:"id,omitempty"`           // id
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Created     *string        `url:"created,omitempty"`      // 创建时间，支持时间查询
		WikiID      *int64         `url:"wiki_id,omitempty"`      // 关联的 wiki id
		User        *string        `url:"user,omitempty"`         // 关注者昵称
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetWikiFollowersCountRequest struct {
		ID          *int64  `url:"id,omitempty"`           // id
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		Created     *string `url:"created,omitempty"`      // 创建时间，支持时间查询
		WikiID      *int64  `url:"wiki_id,omitempty"`      // 关联的 wiki id
		User        *string `url:"user,omitempty"`         // 关注者昵称
	}

	WikiEntityPermission struct {
		ID          string `json:"id,omitempty"`           // 记录ID
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		EntryType   string `json:"entry_type,omitempty"`   // 固定值 wiki
		TargetType  string `json:"target_type,omitempty"`  // 可访问的类型
		TargetID    string `json:"target_id,omitempty"`    // 用户昵称或用户组ID
		WikiID      string `json:"wiki_id,omitempty"`      // wiki ID
	}

	GetWikiEntityPermissionsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		WikiID      *int64         `url:"wiki_id,omitempty"`      // [必须]wiki ID
		TargetType  *string        `url:"target_type,omitempty"`  // 可访问的类型
		TargetID    *string        `url:"target_id,omitempty"`    // 用户昵称或用户组ID
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	WikiTag struct {
		Creator string `json:"creator,omitempty"` // 标签创建人 nick
		Created string `json:"created,omitempty"` // 标签创建时间
		WikiID  string `json:"wiki_id,omitempty"` // wiki id
		Tag     string `json:"tag,omitempty"`     // 标签
	}

	GetWikiTagsRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		WikiID      *int64  `url:"wiki_id,omitempty"`      // wiki id，不传取项目下的所有 wiki
		Tag         *string `url:"tag,omitempty"`          // 标签
		Creator     *string `url:"creator,omitempty"`      // 标签创建人 nick
		Created     *string `url:"created,omitempty"`      // 标签创建时间，支持时间查询
		Limit       *int    `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int    `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order  `url:"order,omitempty"`        // 排序规则
	}

	GetWikiTagsCountRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		WikiID      *int64  `url:"wiki_id,omitempty"`      // wiki id，不传取项目下的所有 wiki
		Tag         *string `url:"tag,omitempty"`          // 标签
		Creator     *string `url:"creator,omitempty"`      // 标签创建人 nick
		Created     *string `url:"created,omitempty"`      // 标签创建时间，支持时间查询
	}

	GetWikiAttachmentsCountRequest struct {
		ID          *int64  `url:"id,omitempty"`           // id
		Filename    *string `url:"filename,omitempty"`     // 文件名
		Size        *int    `url:"size,omitempty"`         // 文件大小，字节
		Owner       *string `url:"owner,omitempty"`        // 上传者
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		Created     *string `url:"created,omitempty"`      // 创建时间，支持时间查询
		Modified    *string `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		WikiID      *int64  `url:"wiki_id,omitempty"`      // 关联的 wiki id
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

	// GetWikiDrawioData 获取 wiki drawio 数据
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_drawios.html
	GetWikiDrawioData(ctx context.Context, request *GetWikiDrawioDataRequest, opts ...RequestOption) (*WikiDrawioData, *Response, error)

	// GetWikiFollowers 获取 wiki 关注人数据
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_followers.html
	GetWikiFollowers(ctx context.Context, request *GetWikiFollowersRequest, opts ...RequestOption) ([]*WikiFollower, *Response, error)

	// GetWikiFollowersCount 获取 wiki 关注人数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_followers_count.html
	GetWikiFollowersCount(ctx context.Context, request *GetWikiFollowersCountRequest, opts ...RequestOption) (int, *Response, error)

	// GetWikiEntityPermissions 获取 wiki 可访问范围人员及用户组
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_entity_permissions.html
	GetWikiEntityPermissions(ctx context.Context, request *GetWikiEntityPermissionsRequest, opts ...RequestOption) ([]*WikiEntityPermission, *Response, error)

	// GetWikiTags 获取 wiki 标签信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags.html
	GetWikiTags(ctx context.Context, request *GetWikiTagsRequest, opts ...RequestOption) ([]*WikiTag, *Response, error)

	// GetWikiTagsCount 获取 wiki 标签信息数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_tags_count.html
	GetWikiTagsCount(ctx context.Context, request *GetWikiTagsCountRequest, opts ...RequestOption) (int, *Response, error)

	// GetWikiAttachmentsCount 获取 wiki 附件数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/wiki/get_tapd_wikis_attachments_count.html
	GetWikiAttachmentsCount(ctx context.Context, request *GetWikiAttachmentsCountRequest, opts ...RequestOption) (int, *Response, error)
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

func (s *wikiService) GetWikiDrawioData(
	ctx context.Context, request *GetWikiDrawioDataRequest, opts ...RequestOption,
) (*WikiDrawioData, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_drawios", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		StaticData *WikiDrawioData `json:"StaticData"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.StaticData, resp, nil
}

func (s *wikiService) GetWikiFollowers(
	ctx context.Context, request *GetWikiFollowersRequest, opts ...RequestOption,
) ([]*WikiFollower, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_followers", request, opts)
	if err != nil {
		return nil, nil, err
	}

	return s.doWikiFollowers(req)
}

func (s *wikiService) GetWikiFollowersCount(
	ctx context.Context, request *GetWikiFollowersCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_followers/count", request, opts)
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

func (s *wikiService) GetWikiEntityPermissions(
	ctx context.Context, request *GetWikiEntityPermissionsRequest, opts ...RequestOption,
) ([]*WikiEntityPermission, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_entity_permissions", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		EntityPermission *WikiEntityPermission `json:"EntityPermission"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	permissions := make([]*WikiEntityPermission, 0, len(items))
	for _, item := range items {
		permissions = append(permissions, item.EntityPermission)
	}

	return permissions, resp, nil
}

func (s *wikiService) GetWikiTags(
	ctx context.Context, request *GetWikiTagsRequest, opts ...RequestOption,
) ([]*WikiTag, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_tags", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Tags *WikiTag `json:"Tags"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	tags := make([]*WikiTag, 0, len(items))
	for _, item := range items {
		tags = append(tags, item.Tags)
	}

	return tags, resp, nil
}

func (s *wikiService) GetWikiTagsCount(
	ctx context.Context, request *GetWikiTagsCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_tags/count", request, opts)
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

func (s *wikiService) GetWikiAttachmentsCount(
	ctx context.Context, request *GetWikiAttachmentsCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tapd_wikis_attachments/count", request, opts)
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

func (s *wikiService) doWikiFollowers(req *http.Request) ([]*WikiFollower, *Response, error) {
	var items []map[string]*WikiFollower
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	followers := make([]*WikiFollower, 0, len(items))
	for _, item := range items {
		for _, key := range []string{"WikiFollower", "WikiFollow", "Follower"} {
			if follower := item[key]; follower != nil {
				followers = append(followers, follower)
				goto next
			}
		}
		for _, follower := range item {
			if follower != nil {
				followers = append(followers, follower)
				break
			}
		}
	next:
	}

	return followers, resp, nil
}
