package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWikiService_CreateWiki(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tapd_wikis", r.URL.Path)

		var req CreateWikiRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, "test111", *req.Name)
		assert.Equal(t, "## markdown", *req.MarkdownDescription)
		assert.Equal(t, "xxxxxxx", *req.Description)
		assert.Equal(t, "v_xuanfang", *req.Creator)
		assert.Equal(t, "note", *req.Note)
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, "0", *req.ParentWikiID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/wiki/create_wiki.json"))
	}))

	wiki, _, err := client.WikiService.CreateWiki(ctx, &CreateWikiRequest{
		Name:                Ptr("test111"),
		MarkdownDescription: Ptr("## markdown"),
		Description:         Ptr("xxxxxxx"),
		Creator:             Ptr("v_xuanfang"),
		Note:                Ptr("note"),
		WorkspaceID:         Ptr(10104801),
		ParentWikiID:        Ptr("0"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1210104801000043897", wiki.ID)
	assert.Equal(t, "test111", wiki.Name)
	assert.Equal(t, "10104801", wiki.WorkspaceID)
	assert.Equal(t, "xxxxxxx", wiki.Description)
	assert.Equal(t, "## markdown", wiki.MarkdownDescription)
	assert.Equal(t, "1", wiki.IsRich)
	assert.Equal(t, "0", wiki.ParentWikiID)
	assert.Equal(t, "note", wiki.Note)
	assert.Equal(t, "0", wiki.ViewCount)
	assert.Equal(t, "2020-08-26 10:15:28", wiki.Created)
	assert.Equal(t, "v_xuanfang", wiki.Creator)
	assert.Equal(t, "2020-08-26 10:15:28", wiki.Modified)
	assert.Equal(t, "v_xuanfang", wiki.Modifier)
}

func TestWikiService_GetWikis(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tapd_wikis", r.URL.Path)
		assert.Equal(t, "1210104801000043827", r.URL.Query().Get("id"))
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "dev", r.URL.Query().Get("modifier"))
		assert.Equal(t, "dev", r.URL.Query().Get("creator"))
		assert.Equal(t, "note", r.URL.Query().Get("note"))
		assert.Equal(t, "4", r.URL.Query().Get("view_count"))
		assert.Equal(t, "2020-08-25", r.URL.Query().Get("created"))
		assert.Equal(t, "2020-08-26", r.URL.Query().Get("modified"))
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/wiki/get_wikis.json"))
	}))

	wikis, _, err := client.WikiService.GetWikis(ctx, &GetWikisRequest{
		ID:          Ptr[int64](1210104801000043827),
		Name:        Ptr("test"),
		Modifier:    Ptr("dev"),
		Creator:     Ptr("dev"),
		Note:        Ptr("note"),
		ViewCount:   Ptr("4"),
		Created:     Ptr("2020-08-25"),
		Modified:    Ptr("2020-08-26"),
		WorkspaceID: Ptr(10104801),
		Limit:       Ptr(10),
		Page:        Ptr(2),
		Order:       NewOrder("created", OrderByDesc),
		Fields:      NewMulti("id", "name", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, wikis, 2)
	assert.Equal(t, "1210104801000043827", wikis[0].ID)
	assert.Equal(t, "test888", wikis[0].Name)
	assert.Equal(t, "10104801", wikis[0].WorkspaceID)
	assert.Equal(t, "", wikis[0].Description)
	assert.Equal(t, "", wikis[0].MarkdownDescription)
	assert.Equal(t, "0", wikis[0].IsRich)
	assert.Equal(t, "0", wikis[0].ParentWikiID)
	assert.Equal(t, "0", wikis[0].ViewCount)
	assert.Equal(t, "2020-08-25 11:24:44", wikis[0].Created)
	assert.Equal(t, "dev", wikis[0].Creator)
	assert.Equal(t, "2020-08-25 11:24:44", wikis[0].Modified)
	assert.Equal(t, "dev", wikis[0].Modifier)
}

func TestWikiService_GetWikisCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tapd_wikis/count", r.URL.Path)
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "dev", r.URL.Query().Get("modifier"))
		assert.Equal(t, "dev", r.URL.Query().Get("creator"))
		assert.Equal(t, "note", r.URL.Query().Get("note"))
		assert.Equal(t, "4", r.URL.Query().Get("view_count"))
		assert.Equal(t, "2020-08-25", r.URL.Query().Get("created"))
		assert.Equal(t, "2020-08-26", r.URL.Query().Get("modified"))
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/wiki/get_wikis_count.json"))
	}))

	count, _, err := client.WikiService.GetWikisCount(ctx, &GetWikisCountRequest{
		Name:        Ptr("test"),
		Modifier:    Ptr("dev"),
		Creator:     Ptr("dev"),
		Note:        Ptr("note"),
		ViewCount:   Ptr("4"),
		Created:     Ptr("2020-08-25"),
		Modified:    Ptr("2020-08-26"),
		WorkspaceID: Ptr(10104801),
	})
	assert.NoError(t, err)
	assert.Equal(t, 23, count)
}

func TestWikiService_UpdateWiki(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tapd_wikis", r.URL.Path)

		var req UpdateWikiRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1210104801000043897), *req.ID)
		assert.Equal(t, "test111", *req.Name)
		assert.Equal(t, "## updated", *req.MarkdownDescription)
		assert.Equal(t, "内容被更新", *req.Description)
		assert.Equal(t, "note updated", *req.Note)
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, "0", *req.ParentWikiID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/wiki/update_wiki.json"))
	}))

	wiki, _, err := client.WikiService.UpdateWiki(ctx, &UpdateWikiRequest{
		ID:                  Ptr[int64](1210104801000043897),
		Name:                Ptr("test111"),
		MarkdownDescription: Ptr("## updated"),
		Description:         Ptr("内容被更新"),
		Note:                Ptr("note updated"),
		WorkspaceID:         Ptr(10104801),
		ParentWikiID:        Ptr("0"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1210104801000043897", wiki.ID)
	assert.Equal(t, "test111", wiki.Name)
	assert.Equal(t, "10104801", wiki.WorkspaceID)
	assert.Equal(t, "内容被更新", wiki.Description)
	assert.Equal(t, "## updated", wiki.MarkdownDescription)
	assert.Equal(t, "1", wiki.IsRich)
	assert.Equal(t, "0", wiki.ParentWikiID)
	assert.Equal(t, "1", wiki.ViewCount)
	assert.Equal(t, "2020-08-26 10:30:11", wiki.Modified)
	assert.Equal(t, "dev", wiki.Modifier)
}
