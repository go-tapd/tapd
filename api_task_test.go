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

func TestTaskService_GetTaskCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/custom_fields_settings", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_task_custom_fields_settings.json"))
	}))

	settings, _, err := client.TaskService.GetTaskCustomFieldsSettings(ctx, &GetTaskCustomFieldsSettingsRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.Len(t, settings, 1)
	assert.Equal(t, "1111112222001000155", settings[0].ID)
	assert.Equal(t, "11112222", settings[0].WorkspaceID)
	assert.Equal(t, "task", settings[0].EntryType)
	assert.Equal(t, "custom_field_100", settings[0].CustomField)
	assert.Equal(t, "user_chooser", settings[0].Type)
	assert.Equal(t, "任务处理人", settings[0].Name)
	assert.Nil(t, settings[0].Options)
	assert.Equal(t, "1", settings[0].Enabled)
	assert.Equal(t, 0, settings[0].IsOut)
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

func TestTaskService_BatchUpdateTasks(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tasks/batch_update_task", r.URL.Path)

		var req BatchUpdateTasksRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Len(t, req.Workitems, 2)
		assert.Equal(t, int64(1111112222001138994), *req.Workitems[0].ID)
		assert.Equal(t, "first task", *req.Workitems[0].Name)
		assert.Equal(t, TaskStatusProgressing, *req.Workitems[0].Status)
		assert.Equal(t, int64(1111112222001138995), *req.Workitems[1].ID)
		assert.Equal(t, "second task", *req.Workitems[1].Name)
		assert.Equal(t, "owner", *req.Workitems[1].Owner)

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/batch_update_task.json"))
	}))

	result, _, err := client.TaskService.BatchUpdateTasks(ctx, &BatchUpdateTasksRequest{
		WorkspaceID: Ptr(11112222),
		Workitems: []*UpdateTaskRequest{
			{
				ID:     Ptr[int64](1111112222001138994),
				Name:   Ptr("first task"),
				Status: Ptr(TaskStatusProgressing),
			},
			{
				ID:    Ptr[int64](1111112222001138995),
				Name:  Ptr("second task"),
				Owner: Ptr("owner"),
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "success", result.Msg)
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

func TestTaskService_GetRemovedTasks(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/get_removed_tasks", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111,1111111111112", r.URL.Query().Get("id"))
		assert.Equal(t, "creator", r.URL.Query().Get("creator"))
		assert.Equal(t, "1", r.URL.Query().Get("is_archived"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("deleted"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_removed_tasks.json"))
	}))

	tasks, _, err := client.TaskService.GetRemovedTasks(ctx, &GetRemovedTasksRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti(1111111111111, 1111111111112),
		Creator:     Ptr("creator"),
		IsArchived:  Ptr(1),
		Created:     Ptr("2021-01-01"),
		Deleted:     Ptr("2021-01-02"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
	})
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "1111112222001138994", tasks[0].ID)
	assert.Equal(t, "已删除任务", tasks[0].Name)
	assert.Equal(t, "张三", tasks[0].Creator)
	assert.Equal(t, "张三", tasks[0].OperationUser)
	assert.Equal(t, "2024-08-20 11:28:23", tasks[0].Deleted)
	assert.Equal(t, "0", tasks[0].IsArchived)
}

func TestTaskService_GetTasksByViewConfID(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tasks/get_tasks_by_view_conf_id", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000001", r.URL.Query().Get("view_conf_id"))
		assert.Equal(t, "testuser", r.URL.Query().Get("current_user"))
		assert.Equal(t, "20", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,name,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/task/get_tasks_by_view_conf_id.json"))
	}))

	tasks, _, err := client.TaskService.GetTasksByViewConfID(ctx, &GetTasksByViewConfIDRequest{
		WorkspaceID: Ptr(11112222),
		ViewConfID:  Ptr[int64](1111112222001000001),
		CurrentUser: Ptr("testuser"),
		Limit:       Ptr(20),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "name", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, "1111112222001138994", tasks[0].ID)
	assert.Equal(t, "视图任务", tasks[0].Name)
	assert.Equal(t, "11112222", tasks[0].WorkspaceID)
	assert.Equal(t, TaskStatusOpen, tasks[0].Status)
	assert.Equal(t, "owner", tasks[0].Owner)
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
