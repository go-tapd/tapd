package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkflowService_GetAllLastSteps(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workflows/all_last_steps", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "story", r.URL.Query().Get("system"))

		_, _ = w.Write(loadData(t, ".testdata/api/workflow/get_all_last_steps.json"))
	}))

	steps, _, err := client.WorkflowService.GetAllLastSteps(ctx, &GetAllLastStepsRequest{
		WorkspaceID: Ptr(11112222),
		System:      Ptr("story"),
	})
	assert.NoError(t, err)
	require.Len(t, steps, 2)
}
