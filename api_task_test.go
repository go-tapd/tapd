package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskService_CreateTask(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tasks", r.URL.Path)

		var req CreateTaskRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))

		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "Test Task", *req.Name)
		assert.Equal(t, "This is a test task", *req.Description)
		assert.Equal(t, "testuser", *req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/create_task.json"))
	}))

	task, _, err := client.TaskService.CreateTask(ctx, &CreateTaskRequest{
		WorkspaceID: Ptr(11112222),
		Name:        Ptr("Test Task"),
		Description: Ptr("This is a test task"),
		Creator:     Ptr("testuser"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, task)

	assert.Equal(t, "1111112222001138994", task.ID)
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, "This is a test task", task.Description)
	assert.Equal(t, "11112222", task.WorkspaceID)
	assert.Equal(t, "testuser", task.Creator)
	assert.Equal(t, "2025-06-26 21:49:02", task.Created)
	assert.Equal(t, "2025-06-26 21:49:02", task.Modified)
	assert.Equal(t, TaskStatusOpen, task.Status)
}

func TestTaskService_GetTasks(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks", r.URL.Path)

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_tasks.json"))
	}))

	tasks, _, err := client.TaskService.GetTasks(ctx, &GetTasksRequest{
		WorkspaceID: Ptr(11112222),
		Status:      NewEnum(TaskStatusOpen, TaskStatusDone),
		Fields:      NewMulti("id", "workspace_id"),
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

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_tasks_count.json"))
	}))

	count, _, err := client.TaskService.GetTasksCount(ctx, &GetTasksCountRequest{
		WorkspaceID: Ptr(11112222),
		Status:      NewEnum(TaskStatusOpen, TaskStatusDone),
	})
	assert.NoError(t, err)
	assert.Equal(t, 36, count)
}

func TestTaskService_UpdateTask(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tasks", r.URL.Path)

		var req struct {
			ID                 int64      `json:"id"`
			WorkspaceID        int        `json:"workspace_id"`
			CurrentUser        string     `json:"current_user"`
			Name               string     `json:"name"`
			Description        string     `json:"description"`
			Status             TaskStatus `json:"status"`
			Owner              string     `json:"owner"`
			Begin              string     `json:"begin"`
			Due                string     `json:"due"`
			StoryID            int64      `json:"story_id"`
			IterationID        int64      `json:"iteration_id"`
			PriorityLabel      string     `json:"priority_label"`
			Label              string     `json:"label"`
			Progress           int        `json:"progress"`
			Effort             string     `json:"effort"`
			AutoCompleteEffort int        `json:"auto_complete_effort"`
			CustomFieldOne     string     `json:"custom_field_one"`
			CustomPlanField1   string     `json:"custom_plan_field_1"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1111112222001138994), req.ID)
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "testuser", req.CurrentUser)
		assert.Equal(t, "Updated Task", req.Name)
		assert.Equal(t, "This is an updated task", req.Description)
		assert.Equal(t, TaskStatusProgressing, req.Status)
		assert.Equal(t, "owner", req.Owner)
		assert.Equal(t, "2025-06-27", req.Begin)
		assert.Equal(t, "2025-06-30", req.Due)
		assert.Equal(t, int64(1111112222001047639), req.StoryID)
		assert.Equal(t, int64(1111112222001001779), req.IterationID)
		assert.Equal(t, "High", req.PriorityLabel)
		assert.Equal(t, "frontend|urgent", req.Label)
		assert.Equal(t, 50, req.Progress)
		assert.Equal(t, "8", req.Effort)
		assert.Equal(t, 1, req.AutoCompleteEffort)
		assert.Equal(t, "custom value", req.CustomFieldOne)
		assert.Equal(t, "plan value", req.CustomPlanField1)

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/update_task.json"))
	}))

	task, _, err := client.TaskService.UpdateTask(ctx, &UpdateTaskRequest{
		ID:                 Ptr[int64](1111112222001138994),
		WorkspaceID:        Ptr(11112222),
		CurrentUser:        Ptr("testuser"),
		Name:               Ptr("Updated Task"),
		Description:        Ptr("This is an updated task"),
		Status:             Ptr(TaskStatusProgressing),
		Owner:              Ptr("owner"),
		Begin:              Ptr("2025-06-27"),
		Due:                Ptr("2025-06-30"),
		StoryID:            Ptr[int64](1111112222001047639),
		IterationID:        Ptr[int64](1111112222001001779),
		PriorityLabel:      Ptr(PriorityLabelHigh),
		Label:              NewEnum("frontend", "urgent"),
		Progress:           Ptr(50),
		Effort:             Ptr("8"),
		AutoCompleteEffort: Ptr(1),
		CustomFieldOne:     Ptr("custom value"),
		CustomPlanField1:   Ptr("plan value"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, task)
	assert.Equal(t, "1111112222001138994", task.ID)
	assert.Equal(t, "Updated Task", task.Name)
	assert.Equal(t, "This is an updated task", task.Description)
	assert.Equal(t, "11112222", task.WorkspaceID)
	assert.Equal(t, TaskStatusProgressing, task.Status)
	assert.Equal(t, "owner", task.Owner)
	assert.Equal(t, "2025-06-27", task.Begin)
	assert.Equal(t, "2025-06-30", task.Due)
	assert.Equal(t, "1111112222001047639", task.StoryID)
	assert.Equal(t, "1111112222001001779", task.IterationID)
	assert.Equal(t, "50", task.Progress)
	assert.Equal(t, "8", task.Effort)
	assert.Equal(t, PriorityLabelHigh, task.PriorityLabel)
	assert.Equal(t, "custom value", task.CustomFieldOne)
	assert.Equal(t, "plan value", task.CustomPlanField1)
}

func TestTaskService_GetTaskChanges(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/task_changes", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_task_changes.json"))
	}))

	changes, _, err := client.TaskService.GetTaskChanges(t.Context(), &GetTaskChangesRequest{
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

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_task_changes_count.json"))
	}))

	count, _, err := client.TaskService.GetTaskChangesCount(t.Context(), &GetTaskChangesCountRequest{
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

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_task_fields_info.json"))
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
