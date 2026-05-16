package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseService_CreateRelease(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/new_releases", r.URL.Path)

		var req CreateReleaseRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, "test2", *req.Name)
		assert.Equal(t, "发布计划描述", *req.Description)
		assert.Equal(t, "2020-10-20", *req.StartDate)
		assert.Equal(t, "2020-11-20", *req.EndDate)
		assert.Equal(t, "dev", *req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/create_release.json"))
	}))

	release, _, err := client.ReleaseService.CreateRelease(ctx, &CreateReleaseRequest{
		WorkspaceID: Ptr(10104801),
		Name:        Ptr("test2"),
		Description: Ptr("发布计划描述"),
		StartDate:   Ptr("2020-10-20"),
		EndDate:     Ptr("2020-11-20"),
		Creator:     Ptr("dev"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801100003081", release.ID)
	assert.Equal(t, "10104801", release.WorkspaceID)
	assert.Equal(t, "test2", release.Name)
	assert.Nil(t, release.Description)
	assert.Equal(t, "2020-10-20", release.StartDate)
	assert.Equal(t, "2020-11-20", release.EndDate)
	assert.Equal(t, "dev", release.Creator)
	assert.Equal(t, "2020-10-27 11:09:14", release.Created)
	assert.Equal(t, "2020-10-27 11:09:14", release.Modified)
	assert.Equal(t, "open", release.Status)
}

func TestReleaseService_GetReleases(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/releases", r.URL.Path)
		assert.Equal(t, "1010158231100000905,1010158231100000906", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "发布计划", r.URL.Query().Get("name"))
		assert.Equal(t, "敏捷研发", r.URL.Query().Get("description"))
		assert.Equal(t, "2017-06-12", r.URL.Query().Get("startdate"))
		assert.Equal(t, "2017-07-07", r.URL.Query().Get("enddate"))
		assert.Equal(t, "anyechen", r.URL.Query().Get("creator"))
		assert.Equal(t, "2017-06-20", r.URL.Query().Get("created"))
		assert.Equal(t, "2017-06-21", r.URL.Query().Get("modified"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_releases.json"))
	}))

	releases, _, err := client.ReleaseService.GetReleases(ctx, &GetReleasesRequest{
		ID:          NewMulti[int64](1010158231100000905, 1010158231100000906),
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("发布计划"),
		Description: Ptr("敏捷研发"),
		StartDate:   Ptr("2017-06-12"),
		EndDate:     Ptr("2017-07-07"),
		Creator:     Ptr("anyechen"),
		Created:     Ptr("2017-06-20"),
		Modified:    Ptr("2017-06-21"),
		Status:      Ptr("open"),
		Limit:       Ptr(10),
		Page:        Ptr(2),
		Order:       NewOrder("created", OrderByDesc),
		Fields:      NewMulti("id", "name", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, releases, 1)
	assert.Equal(t, "1010158231100000905", releases[0].ID)
	assert.Equal(t, "10158231", releases[0].WorkspaceID)
	assert.Equal(t, "发布计划1", releases[0].Name)
	assert.Equal(t, "熟悉敏捷研发全生命周期", *releases[0].Description)
	assert.Equal(t, "2017-06-12", releases[0].StartDate)
	assert.Equal(t, "2017-07-07", releases[0].EndDate)
	assert.Equal(t, "anyechen", releases[0].Creator)
	assert.Equal(t, "2017-06-20 16:49:01", releases[0].Created)
	assert.Equal(t, "2017-06-20 16:49:01", releases[0].Modified)
	assert.Equal(t, "open", releases[0].Status)
}

func TestReleaseService_GetReleasesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/releases/count", r.URL.Path)
		assert.Equal(t, "1010158231100000905,1010158231100000906", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "发布计划", r.URL.Query().Get("name"))
		assert.Equal(t, "敏捷研发", r.URL.Query().Get("description"))
		assert.Equal(t, "2017-06-12", r.URL.Query().Get("startdate"))
		assert.Equal(t, "2017-07-07", r.URL.Query().Get("enddate"))
		assert.Equal(t, "anyechen", r.URL.Query().Get("creator"))
		assert.Equal(t, "2017-06-20", r.URL.Query().Get("created"))
		assert.Equal(t, "2017-06-21", r.URL.Query().Get("modified"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_releases_count.json"))
	}))

	count, _, err := client.ReleaseService.GetReleasesCount(ctx, &GetReleasesCountRequest{
		ID:          NewMulti[int64](1010158231100000905, 1010158231100000906),
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("发布计划"),
		Description: Ptr("敏捷研发"),
		StartDate:   Ptr("2017-06-12"),
		EndDate:     Ptr("2017-07-07"),
		Creator:     Ptr("anyechen"),
		Created:     Ptr("2017-06-20"),
		Modified:    Ptr("2017-06-21"),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestReleaseService_UpdateRelease(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/new_releases", r.URL.Path)

		var req UpdateReleaseRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, int64(1010104801100003081), *req.ID)
		assert.Equal(t, "test2", *req.Name)
		assert.Equal(t, "内容被更新", *req.Description)
		assert.Equal(t, "2020-10-20", *req.StartDate)
		assert.Equal(t, "2020-11-20", *req.EndDate)
		assert.Equal(t, "open", *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/update_release.json"))
	}))

	release, _, err := client.ReleaseService.UpdateRelease(ctx, &UpdateReleaseRequest{
		WorkspaceID: Ptr(10104801),
		ID:          Ptr[int64](1010104801100003081),
		Name:        Ptr("test2"),
		Description: Ptr("内容被更新"),
		StartDate:   Ptr("2020-10-20"),
		EndDate:     Ptr("2020-11-20"),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801100003081", release.ID)
	assert.Equal(t, "10104801", release.WorkspaceID)
	assert.Equal(t, "test2", release.Name)
	assert.Equal(t, "内容被更新", *release.Description)
	assert.Equal(t, "2020-10-27 11:24:48", release.Modified)
	assert.Equal(t, "open", release.Status)
}
