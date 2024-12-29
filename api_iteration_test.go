package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterationService_GetIterations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/iterations", r.URL.Path)
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/iteration/get_iterations.json"))
	}))

	iterations, _, err := client.IterationService.GetIterations(ctx, &GetIterationsRequest{
		WorkspaceID: Ptr(111),
	})
	assert.NoError(t, err)
	require.NotNil(t, iterations)
	require.Len(t, iterations, 8)

	iteration := iterations[0]
	assert.Equal(t, "11111222001002235", iteration.ID)
	assert.Equal(t, "2025 年 M1-迭代", iteration.Name)
	assert.Equal(t, "111222", iteration.WorkspaceID)
	assert.Equal(t, "2025-01-01", iteration.StartDate)
	assert.Equal(t, "2025-01-31", iteration.EndDate)
	assert.Equal(t, "open", iteration.Status)
	assert.Equal(t, "creator name", iteration.Creator)
	assert.Equal(t, "2024-12-27 17:04:43", iteration.Created)
	assert.Equal(t, "2024-12-27 17:04:43", iteration.Modified)
	assert.Equal(t, "iteration", iteration.EntityType)
	assert.Equal(t, "0", iteration.ParentID)
	assert.Equal(t, "11111222001002235", iteration.AncestorID)
	assert.Equal(t, "11111222001002235:", iteration.Path)
	assert.Equal(t, "11111222001000098", iteration.WorkitemTypeID)
	assert.Equal(t, "11111222001000218", iteration.TemplatedID)
}

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

func TestIterationService_GetTemplateList(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/iterations/template_list", r.URL.Path)
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/iteration/get_template_list.json"))
	}))

	templates, _, err := client.IterationService.GetTemplateList(ctx, &GetTemplateListRequest{
		WorkspaceID: Ptr(111),
	})
	assert.NoError(t, err)
	require.NotNil(t, templates)
	assert.ElementsMatch(t, []*WorkitemTemplate{
		{
			ID:          "1111110502001000111",
			WorkspaceID: "11112222",
			Type:        "iteration",
			Name:        "Tapd Template",
			Creator:     "Tapd System",
			Created:     "2022-06-10 10:04:08",
			Modified:    "2022-06-10 10:04:08",
		},
	}, templates)
}
