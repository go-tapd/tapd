package tapd

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_GetTasks(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks", r.URL.Path)

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_tasks.json"))
	}))

	tasks, _, err := client.TaskService.GetTasks(ctx, &GetTasksRequest{
		WorkspaceID: Ptr(11112222),
		Status:      Enum(TaskStatusOpen, TaskStatusDone),
		Fields:      Multi("id", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.True(t, len(tasks) > 0)
}

func TestTaskService_GetTasksCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "open|done", r.URL.Query().Get("status"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_tasks_count.json"))
	}))

	count, _, err := client.TaskService.GetTasksCount(ctx, &GetTasksCountRequest{
		WorkspaceID: Ptr(11112222),
		Status:      Enum(TaskStatusOpen, TaskStatusDone),
	})
	assert.NoError(t, err)
	assert.Equal(t, 36, count)
}

func TestTaskService_GetTaskChanges(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/task_changes", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_task_changes.json"))
	}))

	changes, _, err := client.TaskService.GetTaskChanges(context.Background(), &GetTaskChangesRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(changes) > 0)

	var flag1, flag2 bool
	for _, change := range changes {
		if change.ID == "1111112222001019140" {
			for _, fieldChange := range change.FieldChanges {
				if fieldChange.Field == "remain" {
					assert.Equal(t, "1", fieldChange.ValueBefore)
					assert.Equal(t, "0", fieldChange.ValueAfter)
					assert.Equal(t, "剩余工时", fieldChange.FieldLabel)
					flag1 = true
				}
				if fieldChange.Field == "effort_completed" {
					assert.Equal(t, "0", fieldChange.ValueBefore)
					assert.Equal(t, "1", fieldChange.ValueAfter)
					assert.Equal(t, "完成工时", fieldChange.FieldLabel)
					flag2 = true
				}
			}
		}
	}
	assert.True(t, flag1)
	assert.True(t, flag2)
}

func TestTaskService_GetTaskChangesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/task_changes/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_task_changes_count.json"))
	}))

	count, _, err := client.TaskService.GetTaskChangesCount(context.Background(), &GetTaskChangesCountRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.Equal(t, 189, count)
}

func TestTaskService_GetTaskFieldsInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/get_fields_info", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, ".testdata/api/task/get_task_fields_info.json"))
	}))

	fields, _, err := client.TaskService.GetTaskFieldsInfo(ctx, &GetTaskFieldsInfoRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(fields) > 0)

	var flag1, flag2 bool

	// slice with id name
	for _, field := range fields {
		if field.Name == "id" {
			assert.Equal(t, TaskFieldsInfoHTMLTypeInput, field.HTMLType)
			assert.Equal(t, "ID", field.Label)
			flag1 = true
		}

		if field.Name == "iteration_id" {
			flag2 = true
			assert.Equal(t, "迭代", field.Label)
			assert.Equal(t, TaskFieldsInfoHTMLTypeSelect, field.HTMLType)
			assert.Contains(t, field.Options, TaskFieldsInfoOption{
				Value: "1111112222001001246",
				Label: "迭代2",
			})
			assert.Contains(t, field.PureOptions, TaskFieldsInfoPureOption{
				ParentID:    "0",
				WorkspaceID: "11112222",
				Sort:        "100124600000",
				Value:       "1111112222001001246",
				Label:       "迭代2",
				Panel:       0,
			})
		}
	}
	assert.True(t, flag1)
	assert.True(t, flag2)
}
