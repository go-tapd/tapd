package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportService_GetReports(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workspace_reports", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1000000000000000002", r.URL.Query().Get("id"))
		assert.Equal(t, "title", r.URL.Query().Get("title"))
		assert.Equal(t, "author", r.URL.Query().Get("author"))
		assert.Equal(t, "created", r.URL.Query().Get("created"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,title", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/report/get_reports.json"))
	}))

	reports, _, err := client.ReportService.GetReports(ctx, &GetReportsRequest{
		WorkspaceID: Ptr(11112222),
		ID:          Ptr(1000000000000000002),
		Title:       Ptr("title"),
		Author:      Ptr("author"),
		Created:     Ptr("created"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "title"),
	})
	assert.NoError(t, err)
	assert.True(t, len(reports) > 0)
	assert.Equal(t, "10000000000000846", reports[0].ID)
}
