package tapd

import (
	"context"
	"net/http"
)

type (
	// BoardCard 看板工作项
	BoardCard struct {
		ID          string  `json:"id,omitempty"`           // 工作项ID
		Name        string  `json:"name,omitempty"`         // 工作项标题
		Description *string `json:"description,omitempty"`  // 详细描述
		Created     string  `json:"created,omitempty"`      // 创建时间
		Modified    string  `json:"modified,omitempty"`     // 最后修改时间
		WorkspaceID string  `json:"workspace_id,omitempty"` // 项目ID
		Owner       *string `json:"owner,omitempty"`        // 负责人
		Due         *string `json:"due,omitempty"`          // 截止时间
		Label       string  `json:"b_label,omitempty"`      // 标签ID
		Sort        string  `json:"b_sort,omitempty"`       // 排序
		BoardID     string  `json:"b_board_id,omitempty"`   // 看板ID
		ColumnID    string  `json:"b_column_id,omitempty"`  // 板块ID
		Status      string  `json:"status,omitempty"`       // 状态
		CC          *string `json:"cc,omitempty"`           // 参与人
		Begin       *string `json:"begin,omitempty"`        // 开始时间
	}

	CreateBoardCardRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		BoardID     *int64  `json:"b_board_id,omitempty"`   // [必须]看板ID
		ColumnID    *int64  `json:"b_column_id,omitempty"`  // [必须]板块ID
		Name        *string `json:"name,omitempty"`         // [必须]工作项标题
		Owner       *string `json:"owner,omitempty"`        // 负责人
		CC          *string `json:"cc,omitempty"`           // 参与人
		Status      *string `json:"status,omitempty"`       // 状态
		Begin       *string `json:"begin,omitempty"`        // 开始时间
		Due         *string `json:"due,omitempty"`          // 截止时间
		Label       *int64  `json:"b_label,omitempty"`      // 标签ID
		Description *string `json:"description,omitempty"`  // 详细描述
	}

	GetBoardCardsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		ID          *Multi[int64]  `url:"id,omitempty"`           // 工作项ID，支持多ID查询
		BoardID     *int64         `url:"b_board_id,omitempty"`   // 看板ID
		ColumnID    *int64         `url:"b_column_id,omitempty"`  // 板块ID
		Owner       *string        `url:"owner,omitempty"`        // 负责人
		CC          *string        `url:"cc,omitempty"`           // 参与人
		Status      *string        `url:"status,omitempty"`       // 状态
		Name        *string        `url:"name,omitempty"`         // 工作项标题
		Created     *string        `url:"created,omitempty"`      // 创建时间，支持时间查询
		Begin       *string        `url:"begin,omitempty"`        // 开始时间，支持时间查询
		Due         *string        `url:"due,omitempty"`          // 截止时间，支持时间查询
		Label       *int64         `url:"b_label,omitempty"`      // 标签ID
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	UpdateBoardCardRequest struct {
		ID          *int64  `json:"id,omitempty"`           // [必须]工作项ID
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		BoardID     *int64  `json:"b_board_id,omitempty"`   // 看板ID
		ColumnID    *int64  `json:"b_column_id,omitempty"`  // 板块ID
		Name        *string `json:"name,omitempty"`         // 工作项标题
		Owner       *string `json:"owner,omitempty"`        // 负责人
		CC          *string `json:"cc,omitempty"`           // 参与人
		Status      *string `json:"status,omitempty"`       // 状态
		Begin       *string `json:"begin,omitempty"`        // 开始时间
		Due         *string `json:"due,omitempty"`          // 截止时间
		Label       *int64  `json:"b_label,omitempty"`      // 标签ID
		Description *string `json:"description,omitempty"`  // 详细描述
	}

	// BoardColumn 看板板块
	BoardColumn struct {
		ID          string `json:"id,omitempty"`           // 板块ID
		Name        string `json:"name,omitempty"`         // 板块名称
		BoardID     string `json:"board_id,omitempty"`     // 看板ID
		Status      string `json:"status,omitempty"`       // 状态
		Sort        string `json:"sort,omitempty"`         // 排序
		Created     string `json:"created,omitempty"`      // 创建时间
		Creator     string `json:"creator,omitempty"`      // 创建人
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
	}

	GetBoardColumnsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		ID          *Multi[int64]  `url:"id,omitempty"`           // 板块ID，支持多ID查询
		Name        *string        `url:"name,omitempty"`         // 板块名称
		BoardID     *int64         `url:"board_id,omitempty"`     // 看板ID
		Status      *string        `url:"status,omitempty"`       // 状态
		Created     *string        `url:"created,omitempty"`      // 创建时间
		Creator     *string        `url:"creator,omitempty"`      // 创建人
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}
)

// BoardService 看板服务。
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/
type BoardService interface {
	// CreateBoardCard 新建看板工作项
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/add_board_card.html
	CreateBoardCard(ctx context.Context, request *CreateBoardCardRequest, opts ...RequestOption) (*BoardCard, *Response, error)

	// GetBoardCards 获取看板工作项
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/get_board_cards.html
	GetBoardCards(ctx context.Context, request *GetBoardCardsRequest, opts ...RequestOption) ([]*BoardCard, *Response, error)

	// UpdateBoardCard 更新看板工作项
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/update_board_card.html
	UpdateBoardCard(ctx context.Context, request *UpdateBoardCardRequest, opts ...RequestOption) (*BoardCard, *Response, error)

	// GetBoardColumns 获取看板板块
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/board/get_board_columns.html
	GetBoardColumns(ctx context.Context, request *GetBoardColumnsRequest, opts ...RequestOption) ([]*BoardColumn, *Response, error)
}

type boardService struct {
	client *Client
}

var _ BoardService = (*boardService)(nil)

func NewBoardService(client *Client) BoardService {
	return &boardService{
		client: client,
	}
}

func (s *boardService) CreateBoardCard(
	ctx context.Context, request *CreateBoardCardRequest, opts ...RequestOption,
) (*BoardCard, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "board_cards", request, opts)
	if err != nil {
		return nil, nil, err
	}

	cards, resp, err := s.doBoardCards(req)
	if err != nil {
		return nil, resp, err
	}
	if len(cards) == 0 {
		return nil, resp, nil
	}

	return cards[0], resp, nil
}

func (s *boardService) GetBoardCards(
	ctx context.Context, request *GetBoardCardsRequest, opts ...RequestOption,
) ([]*BoardCard, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "board_cards", request, opts)
	if err != nil {
		return nil, nil, err
	}

	return s.doBoardCards(req)
}

func (s *boardService) UpdateBoardCard(
	ctx context.Context, request *UpdateBoardCardRequest, opts ...RequestOption,
) (*BoardCard, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "board_cards", request, opts)
	if err != nil {
		return nil, nil, err
	}

	cards, resp, err := s.doBoardCards(req)
	if err != nil {
		return nil, resp, err
	}
	if len(cards) == 0 {
		return nil, resp, nil
	}

	return cards[0], resp, nil
}

func (s *boardService) GetBoardColumns(
	ctx context.Context, request *GetBoardColumnsRequest, opts ...RequestOption,
) ([]*BoardColumn, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "board_columns", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Column *BoardColumn `json:"Column,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	columns := make([]*BoardColumn, 0, len(items))
	for _, item := range items {
		columns = append(columns, item.Column)
	}

	return columns, resp, nil
}

func (s *boardService) doBoardCards(req *http.Request) ([]*BoardCard, *Response, error) {
	var items []struct {
		BoardCard *BoardCard `json:"BoardCard,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	cards := make([]*BoardCard, 0, len(items))
	for _, item := range items {
		cards = append(cards, item.BoardCard)
	}

	return cards, resp, nil
}
