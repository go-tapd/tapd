package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIterationService_CreateIteration(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/iterations", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			Name        string `json:"name"`
			StartDate   string `json:"startdate"`
			EndDate     string `json:"enddate"`
			Creator     string `json:"creator"`
			Label       string `json:"label"`
		}

		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 111, req.WorkspaceID)
		assert.Equal(t, "测试迭代1", req.Name)
		assert.Equal(t, "2025-01-01", req.StartDate)
		assert.Equal(t, "2025-01-31", req.EndDate)
		assert.Equal(t, "creator name", req.Creator)
		assert.Equal(t, "label1|label2", req.Label)

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/create_iteration.json"))
	}))

	iteration, _, err := client.IterationService.CreateIteration(ctx, &CreateIterationRequest{
		WorkspaceID: Ptr(111),
		Name:        Ptr("测试迭代1"),
		StartDate:   Ptr("2025-01-01"),
		EndDate:     Ptr("2025-01-31"),
		Creator:     Ptr("creator name"),
		Label:       NewEnum("label1", "label2"),
	})
	assert.NoError(t, err)
	require.NotNil(t, iteration)

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

func TestIterationService_GetIterations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/iterations", r.URL.Path)
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/get_iterations.json"))
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

func TestIterationService_GetIterationsCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/iterations/count", r.URL.Path)
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/get_iterations_count.json"))
	}))

	count, _, err := client.IterationService.GetIterationsCount(ctx, &GetIterationsCountRequest{
		WorkspaceID: Ptr(111),
	})
	assert.NoError(t, err)
	assert.Equal(t, 106, count)
}

func TestIterationService_UpdateIteration(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/iterations", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			Name        string `json:"name"`
			StartDate   string `json:"startdate"`
			EndDate     string `json:"enddate"`
			Creator     string `json:"creator"`
			Label       string `json:"label"`
		}

		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 111, req.WorkspaceID)
		assert.Equal(t, "测试迭代1", req.Name)
		assert.Equal(t, "2025-01-01", req.StartDate)
		assert.Equal(t, "2025-01-31", req.EndDate)
		assert.Equal(t, "creator name", req.Creator)
		assert.Equal(t, "label1|label2", req.Label)

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/update_iteration.json"))
	}))

	iteration, _, err := client.IterationService.UpdateIteration(ctx, &UpdateIterationRequest{
		WorkspaceID: Ptr(111),
		ID:          Ptr(int64(11111222001002235)),
		CurrentUser: Ptr("current user"),
		Name:        Ptr("测试迭代1"),
		StartDate:   Ptr("2025-01-01"),
		EndDate:     Ptr("2025-01-31"),
		Creator:     Ptr("creator name"),
		Label:       NewEnum("label1", "label2"),
	})
	assert.NoError(t, err)
	require.NotNil(t, iteration)

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

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/get_workitem_types.json"))
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

		_, _ = w.Write(loadData(t, "internal/testdata/api/iteration/get_template_list.json"))
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
