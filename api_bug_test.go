package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBugService_GetBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, PriorityLabelHigh.String(), r.URL.Query().Get("priority_label"))

		_, _ = w.Write(loadData(t, ".testdata/api/bug/get_bugs.json"))
	}))

	bugs, _, err := client.BugService.GetBugs(ctx, &GetBugsRequest{
		WorkspaceID:   Ptr(11112222),
		PriorityLabel: Ptr(PriorityLabelHigh),
	})
	require.NoError(t, err)
	require.True(t, len(bugs) > 0)

	bug := bugs[0]
	assert.Equal(t, "11111222333001000268", bug.ID)
	assert.Equal(t, "计算不正确", bug.Title)
	assert.Equal(t, "<strong>前置条件</div><br  />", bug.Description)
	assert.Equal(t, "", bug.Priority)
	assert.Equal(t, "", bug.Severity)
	assert.Equal(t, "", bug.Module)
	assert.Equal(t, "closed", bug.Status)
	assert.Equal(t, "测试人员", bug.Reporter)
	assert.Equal(t, "2018-07-26 17:20:02", bug.Created)
	assert.Equal(t, "项目缺陷", bug.BugType)
	assert.Equal(t, "2018-07-26 18:09:42", bug.Resolved)
	assert.Equal(t, "2018-08-07 10:05:19", bug.Closed)
	assert.Equal(t, "2024-12-23 10:49:16", bug.Modified)
	assert.Equal(t, "李四", bug.LastModify)
	assert.Equal(t, "", bug.Auditer)
	assert.Equal(t, "张三;", bug.De)
	assert.Equal(t, "张三", bug.Fixer)
}
