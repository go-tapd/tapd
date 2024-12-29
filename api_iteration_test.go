package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterationService_GetWorkitemTypes(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/iterations/workitem_types", r.URL.Path)
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/iteration/get_workitem_types.json"))
	}))

	workitemTypes, _, err := client.IterationService.GetWorkitemTypes(ctx, &GetWorkitemTypesRequest{
		WorkspaceID: Ptr(111),
	})
	assert.NoError(t, err)
	require.NotNil(t, workitemTypes)
	assert.ElementsMatch(t, []*WorkitemType{
		{
			ID:          "1111110502001000111",
			WorkspaceID: "11112222",
			EntityType:  "iteration",
			Name:        "Tapd Iteration",
			Creator:     "TAPD system",
			Created:     "2024-09-04 15:20:06",
			Modified:    "2024-09-04 15:20:06",
		},
	}, workitemTypes)
}
