package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoardService_CreateBoardCard(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/board_cards", r.URL.Path)

		var req CreateBoardCardRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 20355782, *req.WorkspaceID)
		assert.Equal(t, int64(1020355782000010725), *req.BoardID)
		assert.Equal(t, int64(1020355782000045509), *req.ColumnID)
		assert.Equal(t, "test1", *req.Name)
		assert.Equal(t, "open", *req.Status)
		assert.Equal(t, "tester", *req.Owner)
		assert.Equal(t, "dev", *req.CC)
		assert.Equal(t, "2019-09-11", *req.Begin)
		assert.Equal(t, "2019-09-19", *req.Due)
		assert.Equal(t, int64(0), *req.Label)
		assert.Equal(t, "看板工作项描述", *req.Description)

		_, _ = w.Write(loadData(t, "internal/testdata/api/board/create_board_card.json"))
	}))

	card, _, err := client.BoardService.CreateBoardCard(ctx, &CreateBoardCardRequest{
		WorkspaceID: Ptr(20355782),
		BoardID:     Ptr[int64](1020355782000010725),
		ColumnID:    Ptr[int64](1020355782000045509),
		Name:        Ptr("test1"),
		Status:      Ptr("open"),
		Owner:       Ptr("tester"),
		CC:          Ptr("dev"),
		Begin:       Ptr("2019-09-11"),
		Due:         Ptr("2019-09-19"),
		Label:       Ptr[int64](0),
		Description: Ptr("看板工作项描述"),
	})
	assert.NoError(t, err)
	require.NotNil(t, card)
	assert.Equal(t, "1020355782500624255", card.ID)
	assert.Equal(t, "test1", card.Name)
	assert.Equal(t, "20355782", card.WorkspaceID)
	assert.Equal(t, "1020355782000010725", card.BoardID)
	assert.Equal(t, "1020355782000045509", card.ColumnID)
	assert.Equal(t, "open", card.Status)
}

func TestBoardService_GetBoardCards(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/board_cards", r.URL.Path)
		assert.Equal(t, "20355782", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1020355782500624163,1020355782500624255", r.URL.Query().Get("id"))
		assert.Equal(t, "1020355782000010725", r.URL.Query().Get("b_board_id"))
		assert.Equal(t, "1020355782000045509", r.URL.Query().Get("b_column_id"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "30", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,name,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/board/get_board_cards.json"))
	}))

	cards, _, err := client.BoardService.GetBoardCards(ctx, &GetBoardCardsRequest{
		WorkspaceID: Ptr(20355782),
		ID:          NewMulti[int64](1020355782500624163, 1020355782500624255),
		BoardID:     Ptr[int64](1020355782000010725),
		ColumnID:    Ptr[int64](1020355782000045509),
		Status:      Ptr("open"),
		Name:        Ptr("test"),
		Limit:       Ptr(30),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "name", "status"),
	})
	assert.NoError(t, err)
	require.Len(t, cards, 2)
	assert.Equal(t, "1020355782500624163", cards[0].ID)
	assert.Equal(t, "('测试2')", cards[0].Name)
	assert.Equal(t, "1020355782000045509", cards[0].ColumnID)
	assert.Equal(t, "1020355782500624255", cards[1].ID)
	assert.Equal(t, "test1", cards[1].Name)
}

func TestBoardService_UpdateBoardCard(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/board_cards", r.URL.Path)

		var req UpdateBoardCardRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1020355782500624255), *req.ID)
		assert.Equal(t, 20355782, *req.WorkspaceID)
		assert.Equal(t, int64(1020355782000045510), *req.ColumnID)
		assert.Equal(t, "test1 updated", *req.Name)
		assert.Equal(t, "done", *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/board/update_board_card.json"))
	}))

	card, _, err := client.BoardService.UpdateBoardCard(ctx, &UpdateBoardCardRequest{
		ID:          Ptr[int64](1020355782500624255),
		WorkspaceID: Ptr(20355782),
		ColumnID:    Ptr[int64](1020355782000045510),
		Name:        Ptr("test1 updated"),
		Status:      Ptr("done"),
	})
	assert.NoError(t, err)
	require.NotNil(t, card)
	assert.Equal(t, "1020355782500624255", card.ID)
	assert.Equal(t, "test1 updated", card.Name)
	assert.Equal(t, "1020355782000045510", card.ColumnID)
	assert.Equal(t, "done", card.Status)
}

func TestBoardService_GetBoardColumns(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/board_columns", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010104801000032321,1010104801000032319", r.URL.Query().Get("id"))
		assert.Equal(t, "1010104801000007781", r.URL.Query().Get("board_id"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))
		assert.Equal(t, "30", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/board/get_board_columns.json"))
	}))

	columns, _, err := client.BoardService.GetBoardColumns(ctx, &GetBoardColumnsRequest{
		WorkspaceID: Ptr(10104801),
		ID:          NewMulti[int64](1010104801000032321, 1010104801000032319),
		BoardID:     Ptr[int64](1010104801000007781),
		Status:      Ptr("open"),
		Limit:       Ptr(30),
		Page:        Ptr(1),
		Order:       NewOrder("created", OrderByDesc),
		Fields:      NewMulti("id", "name", "status"),
	})
	assert.NoError(t, err)
	require.Len(t, columns, 2)
	assert.Equal(t, "1010104801000032321", columns[0].ID)
	assert.Equal(t, "已上线", columns[0].Name)
	assert.Equal(t, "1010104801000007781", columns[0].BoardID)
	assert.Equal(t, "open", columns[0].Status)
	assert.Equal(t, "7", columns[0].Sort)
	assert.Equal(t, "1010104801000032319", columns[1].ID)
	assert.Equal(t, "开发中", columns[1].Name)
}
