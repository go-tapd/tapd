package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBugService_BugSeverity(t *testing.T) {
	tests := []struct {
		severity   BugSeverity
		wantString string
		wantHuman  string
	}{
		{BugSeverityFatal, "fatal", "致命"},
		{BugSeveritySerious, "serious", "严重"},
		{BugSeverityNormal, "normal", "一般"},
		{BugSeverityPrompt, "prompt", "提示"},
		{BugSeverityAdvice, "advice", "建议"},
		{BugSeverity(""), "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.severity.String(), func(t *testing.T) {
			assert.Equal(t, tt.wantString, tt.severity.String())
			assert.Equal(t, tt.wantHuman, tt.severity.Human())
		})
	}
}

func TestBugService_CreateBug(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs", r.URL.Path)

		var req CreateBugRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "API 创建缺陷", *req.Title)
		assert.Equal(t, "缺陷详细描述", *req.Description)
		assert.Equal(t, PriorityLabelHigh, *req.PriorityLabel)
		assert.Equal(t, BugSeverityFatal, *req.Severity)
		assert.Equal(t, "张三", *req.CurrentOwner)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/create_bug.json"))
	}))

	bug, _, err := client.BugService.CreateBug(ctx, &CreateBugRequest{
		WorkspaceID:   Ptr(11112222),
		Title:         Ptr("API 创建缺陷"),
		Description:   Ptr("缺陷详细描述"),
		PriorityLabel: Ptr(PriorityLabelHigh),
		Severity:      Ptr(BugSeverityFatal),
		CurrentOwner:  Ptr("张三"),
	})
	require.NoError(t, err)
	require.NotNil(t, bug)
	assert.Equal(t, "1111122233301037078", bug.ID)
	assert.Equal(t, "API 创建缺陷", bug.Title)
	assert.Equal(t, BugSeverityFatal, bug.Severity)
	assert.Equal(t, "new", bug.Status)
}

func TestBugService_CopyBug(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/copy_bug", r.URL.Path)

		var req struct {
			WorkspaceID    int    `json:"workspace_id"`
			SourceBugID    int64  `json:"src_bug_id"`
			DstWorkspaceID int    `json:"dst_workspace_id"`
			SyncFields     string `json:"sync_fields"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, int64(1111122233301037078), req.SourceBugID)
		assert.Equal(t, 33334444, req.DstWorkspaceID)
		assert.Equal(t, "title,description,status", req.SyncFields)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/copy_bug.json"))
	}))

	bug, _, err := client.BugService.CopyBug(ctx, &CopyBugRequest{
		WorkspaceID:    Ptr(11112222),
		SourceBugID:    Ptr[int64](1111122233301037078),
		DstWorkspaceID: Ptr(33334444),
		SyncFields:     NewMulti("title", "description", "status"),
	})
	require.NoError(t, err)
	require.NotNil(t, bug)
	assert.Equal(t, "333344440010000001", bug.ID)
	assert.Equal(t, "API 创建缺陷", bug.Title)
	assert.Equal(t, "33334444", bug.WorkspaceID)
	assert.Equal(t, "new", bug.Status)
}

func TestBugService_GetBugChanges(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bug_changes", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233301037078,1111122233301037079", r.URL.Query().Get("bug_id"))
		assert.Equal(t, "severity", r.URL.Query().Get("field"))
		assert.Equal(t, "1", r.URL.Query().Get("include_add_bug"))
		assert.Equal(t, "20", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_changes.json"))
	}))

	changes, _, err := client.BugService.GetBugChanges(ctx, &GetBugChangesRequest{
		WorkspaceID:   Ptr(11112222),
		BugID:         NewMulti[int64](1111122233301037078, 1111122233301037079),
		Field:         Ptr("severity"),
		IncludeAddBug: Ptr(1),
		Limit:         Ptr(20),
		Page:          Ptr(1),
		Order:         NewOrder("created", OrderByDesc),
	})
	require.NoError(t, err)
	require.Len(t, changes, 2)
	assert.Equal(t, "1111122233300015913", changes[0].ID)
	assert.Equal(t, "1111122233301037078", changes[0].BugID)
	assert.Equal(t, "severity", changes[0].Field)
	assert.Equal(t, "normal", changes[0].OldValue)
	assert.Equal(t, "fatal", changes[0].NewValue)
	assert.Nil(t, changes[0].Memo)
	assert.Equal(t, "11112222", changes[0].WorkspaceID)
}

func TestBugService_GetBugChangesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bug_changes/count", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233301037078", r.URL.Query().Get("bug_id"))
		assert.Equal(t, "severity", r.URL.Query().Get("field"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_changes_count.json"))
	}))

	count, _, err := client.BugService.GetBugChangesCount(ctx, &GetBugChangesCountRequest{
		WorkspaceID: Ptr(11112222),
		BugID:       NewMulti[int64](1111122233301037078),
		Field:       Ptr("severity"),
	})
	require.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestBugService_GetBugCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/custom_fields_settings", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_custom_fields_settings.json"))
	}))

	settings, _, err := client.BugService.GetBugCustomFieldsSettings(ctx, &GetBugCustomFieldsSettingsRequest{
		WorkspaceID: Ptr(11112222),
	})
	require.NoError(t, err)
	require.NotEmpty(t, settings)
	assert.Equal(t, "11111222333077902981", settings[0].ID)
	assert.Equal(t, "11112222", settings[0].WorkspaceID)
	assert.Equal(t, "bug", settings[0].EntryType)
	assert.Equal(t, "custom_field_one", settings[0].CustomField)
	assert.Equal(t, "radio", settings[0].Type)
	assert.Equal(t, "安全漏洞类型", settings[0].Name)
	assert.Equal(t, "XSS注入|SQL注入|越权", *settings[0].Options)
	assert.Equal(t, "1", settings[0].Enabled)
	assert.Equal(t, "1", *settings[0].Sort)
}

func TestBugService_GetBugLinkBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_link_bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233301037078", r.URL.Query().Get("bug_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_link_bugs.json"))
	}))

	relations, _, err := client.BugService.GetBugLinkBugs(ctx, &GetBugLinkBugsRequest{
		WorkspaceID: Ptr(11112222),
		BugID:       Ptr[int64](1111122233301037078),
	})
	require.NoError(t, err)
	require.Len(t, relations, 2)
	assert.Equal(t, "repeat", relations[0].Type)
	assert.Equal(t, "1111122233301037079", relations[0].ID)
	assert.Equal(t, "11112222", relations[0].WorkspaceID)
	assert.Equal(t, "target", relations[0].ActAs)
	assert.Equal(t, 11112222, relations[0].LinkedWorkspaceID)
	assert.Equal(t, "1162187798001000534", relations[0].LinkID)
}

func TestBugService_GetBugTemplates(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/template_list", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_templates.json"))
	}))

	templates, _, err := client.BugService.GetBugTemplates(ctx, &GetBugTemplatesRequest{
		WorkspaceID: Ptr(11112222),
	})
	require.NoError(t, err)
	require.NotEmpty(t, templates)
	assert.Equal(t, "1111222233300068639", templates[0].ID)
	assert.Equal(t, "创建模板", templates[0].Name)
	assert.Equal(t, "AA", templates[0].Description)
	assert.Equal(t, "1", templates[0].Sort)
	assert.Equal(t, "0", templates[0].Default)
	assert.Equal(t, "v_xuanfang", templates[0].Creator)
	assert.Equal(t, "1", templates[0].EditorType)
}

func TestBugService_GetBugTemplateFields(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_default_bug_template", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111222233300068639", r.URL.Query().Get("template_id"))
		assert.Equal(t, "1", r.URL.Query().Get("use_priority_label"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_template_fields.json"))
	}))

	fields, _, err := client.BugService.GetBugTemplateFields(ctx, &GetBugTemplateFieldsRequest{
		WorkspaceID:      Ptr(11112222),
		TemplateID:       Ptr[int64](1111222233300068639),
		UsePriorityLabel: Ptr(1),
	})
	require.NoError(t, err)
	require.NotEmpty(t, fields)
	assert.Equal(t, "1111222233300778831", fields[0].ID)
	assert.Equal(t, "11112222", fields[0].WorkspaceID)
	assert.Equal(t, "bug", fields[0].Type)
	assert.Equal(t, "1111222233300068639", fields[0].TemplateID)
	assert.Equal(t, "title", fields[0].Field)
	assert.Equal(t, "", fields[0].Value)
	assert.Equal(t, "1", fields[0].Required)
	assert.Equal(t, "0", fields[0].Sort)
}

func TestBugService_GetBugsByViewConfID(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_bugs_by_view_conf_id", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233301000001", r.URL.Query().Get("view_conf_id"))
		assert.Equal(t, "xinweihe", r.URL.Query().Get("current_user"))
		assert.Equal(t, "new", r.URL.Query().Get("status"))
		assert.Equal(t, "20", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,title,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bugs_by_view_conf_id.json"))
	}))

	bugs, _, err := client.BugService.GetBugsByViewConfID(ctx, &GetBugsByViewConfIDRequest{
		ViewConfID:  Ptr[int64](1111122233301000001),
		CurrentUser: Ptr("xinweihe"),
		GetBugsRequest: GetBugsRequest{
			WorkspaceID: Ptr(11112222),
			Status:      NewEnum("new"),
			Limit:       Ptr(20),
			Page:        Ptr(1),
			Fields:      NewMulti("id", "title", "status"),
		},
	})
	require.NoError(t, err)
	require.Len(t, bugs, 2)
	assert.Equal(t, "11111222333084955735", bugs[0].ID)
	assert.Equal(t, "视图缺陷一", bugs[0].Title)
	assert.Equal(t, "new", bugs[0].Status)
	assert.Equal(t, "11111222333083011055", bugs[1].ID)
}

func TestBugService_GetBugFieldsInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_fields_info", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1", r.URL.Query().Get("all_options"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_fields_info.json"))
	}))

	fields, _, err := client.BugService.GetBugFieldsInfo(ctx, &GetBugFieldsInfoRequest{
		WorkspaceID: Ptr(11112222),
		AllOptions:  Ptr(1),
	})
	require.NoError(t, err)
	require.NotEmpty(t, fields)

	var foundID, foundStatus bool
	for _, field := range fields {
		if field.Name == "id" {
			foundID = true
			assert.Equal(t, "ID", field.Label)
			assert.Equal(t, BugFieldsInfoHTMLTypeInput, field.HTMLType)
		}
		if field.Name == "status" {
			foundStatus = true
			assert.Equal(t, "状态", field.Label)
			assert.Equal(t, BugFieldsInfoHTMLTypeSelect, field.HTMLType)
			assert.Contains(t, field.Options, BugFieldsInfoOption{
				Value: "new",
				Label: "新",
			})
			assert.Contains(t, field.PureOptions, BugFieldsInfoPureOption{
				ParentID:    "0",
				WorkspaceID: "11112222",
				Sort:        "1",
				Value:       "new",
				Label:       "新",
				Panel:       0,
			})
		}
	}
	assert.True(t, foundID)
	assert.True(t, foundStatus)
}

func TestBugService_GetBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, PriorityLabelHigh.String(), r.URL.Query().Get("priority_label"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bugs.json"))
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
	assert.Equal(t, BugSeverityNormal, bug.Severity)
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

func TestBugService_UpdateBug(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs", r.URL.Path)

		var req struct {
			WorkspaceID   int           `json:"workspace_id"`
			ID            int           `json:"id"`
			PriorityLabel PriorityLabel `json:"priority_label"`
			Severity      string        `json:"severity"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, 11111222330268, req.ID)
		assert.Equal(t, PriorityLabelHigh, req.PriorityLabel)
		assert.Equal(t, "fatal|serious", req.Severity)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/update_bug.json"))
	}))

	bug, _, err := client.BugService.UpdateBug(ctx, &UpdateBugRequest{
		WorkspaceID:   Ptr(11112222),
		ID:            Ptr(int64(11111222330268)),
		PriorityLabel: Ptr(PriorityLabelHigh),
		Severity:      NewEnum(BugSeverityFatal, BugSeveritySerious),
	})
	require.NoError(t, err)

	assert.Equal(t, "11111222333001037077", bug.ID)
	assert.Equal(t, "计算不正确222", bug.Title)
	assert.Equal(t, "", bug.Description)
	assert.Equal(t, "", bug.Priority)
	assert.Equal(t, BugSeverityNormal, bug.Severity)
}

func TestBugService_UpdateBugSystemSelectFieldOptions(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/update_system_select_field_options", r.URL.Path)

		var req UpdateBugSystemSelectFieldOptionsRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "bugtype", *req.Field)
		require.Len(t, req.Options, 2)
		assert.Equal(t, "test", *req.Options[0].Value)
		assert.Equal(t, "test111", *req.Options[1].Value)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/update_bug_system_select_field_options.json"))
	}))

	result, _, err := client.BugService.UpdateBugSystemSelectFieldOptions(ctx, &UpdateBugSystemSelectFieldOptionsRequest{
		WorkspaceID: Ptr(11112222),
		Field:       Ptr("bugtype"),
		Options: []*BugSystemSelectFieldOption{
			{Value: Ptr("test")},
			{Value: Ptr("test111")},
		},
	})
	require.NoError(t, err)
	assert.True(t, result)
}

func TestBugService_BatchUpdateBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/batch_update_bug", r.URL.Path)

		var req struct {
			ProjectID int `json:"project_id"`
			Workitems []struct {
				ID           int64  `json:"id"`
				Title        string `json:"title"`
				Status       string `json:"status"`
				CurrentOwner string `json:"current_owner"`
				WorkspaceID  *int   `json:"workspace_id"`
			} `json:"workitems"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.ProjectID)
		require.Len(t, req.Workitems, 2)
		assert.Equal(t, int64(1111122233300103707), req.Workitems[0].ID)
		assert.Nil(t, req.Workitems[0].WorkspaceID)
		assert.Equal(t, "first bug", req.Workitems[0].Title)
		assert.Equal(t, "new", req.Workitems[0].Status)
		assert.Equal(t, int64(1111122233300103708), req.Workitems[1].ID)
		assert.Equal(t, "second bug", req.Workitems[1].Title)
		assert.Equal(t, "owner", req.Workitems[1].CurrentOwner)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/batch_update_bugs.json"))
	}))

	result, _, err := client.BugService.BatchUpdateBugs(ctx, &BatchUpdateBugsRequest{
		ProjectID: Ptr(11112222),
		Workitems: []*UpdateBugRequest{
			{
				ID:     Ptr[int64](1111122233300103707),
				Title:  Ptr("first bug"),
				Status: NewEnum("new"),
			},
			{
				ID:           Ptr[int64](1111122233300103708),
				Title:        Ptr("second bug"),
				CurrentOwner: Ptr("owner"),
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, "batch update success", result.Msg)
}

func TestBugService_GetBugsCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "anyechen", r.URL.Query().Get("current_owner"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bugs_count.json"))
	}))

	count, _, err := client.BugService.GetBugsCount(ctx, &GetBugsCountRequest{
		WorkspaceID:  Ptr(11112222),
		CurrentOwner: Ptr("anyechen"),
	})
	require.NoError(t, err)
	assert.Equal(t, 2, count)
}

func TestBugService_GetRemovedBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_removed_bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233300103707,1111122233300103708", r.URL.Query().Get("id"))
		assert.Equal(t, "creator", r.URL.Query().Get("creator"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("modified"))
		assert.Equal(t, "1", r.URL.Query().Get("include_all"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_removed_bugs.json"))
	}))

	bugs, _, err := client.BugService.GetRemovedBugs(ctx, &GetRemovedBugsRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti[int64](1111122233300103707, 1111122233300103708),
		Creator:     Ptr("creator"),
		Created:     Ptr("2021-01-01"),
		Modified:    Ptr("2021-01-02"),
		IncludeAll:  Ptr(1),
		Limit:       Ptr(10),
		Page:        Ptr(1),
	})
	require.NoError(t, err)
	require.Len(t, bugs, 2)
	assert.Equal(t, "1111122233300103707", bugs[0].ID)
	assert.Equal(t, "回收站缺陷一", bugs[0].Title)
	assert.Equal(t, "creator", bugs[0].Reporter)
	assert.Equal(t, "delete", bugs[0].Type)
	assert.Equal(t, "{\"action\":\"delete\"}", bugs[0].RemovedComment)
	assert.Equal(t, "http://tapd.example.com/bugs/view?bug_id=1111122233300103707", bugs[1].NewBugURL)
}

func TestBugService_GetBugRelatedStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_related_stories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111122233301037078,1111122233301037079", r.URL.Query().Get("bug_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_related_stories.json"))
	}))

	stories, _, err := client.BugService.GetBugRelatedStories(ctx, &GetBugRelatedStoriesRequest{
		WorkspaceID: Ptr(11112222),
		BugID:       NewMulti[int64](1111122233301037078, 1111122233301037079),
	})
	require.NoError(t, err)
	require.Len(t, stories, 2)
	assert.Equal(t, "11112222", stories[0].WorkspaceID)
	assert.Equal(t, "1111122233301037078", stories[0].BugID)
	assert.Equal(t, "1111112222001063941", stories[0].StoryID)
}

func TestBugService_LinkBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/link_bugs", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			BugID       int64  `json:"bug_id"`
			RelateBugs  string `json:"relate_bugs"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, int64(1111122233301037078), req.BugID)
		assert.Equal(t, "1111122233301037079,1111122233301037080", req.RelateBugs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/link_bugs.json"))
	}))

	result, _, err := client.BugService.LinkBugs(ctx, &LinkBugsRequest{
		WorkspaceID: Ptr(11112222),
		BugID:       Ptr[int64](1111122233301037078),
		RelateBugs:  NewMulti[int64](1111122233301037079, 1111122233301037080),
	})
	require.NoError(t, err)
	assert.True(t, result)
}

func TestBugService_DeleteLinkBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/delete_link_bugs", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			BugID       int64  `json:"bug_id"`
			LinkIDs     string `json:"link_ids"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, int64(1111122233301037078), req.BugID)
		assert.Equal(t, "1162187798001000534,1162187798001000535", req.LinkIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/delete_link_bugs.json"))
	}))

	result, _, err := client.BugService.DeleteLinkBugs(ctx, &DeleteLinkBugsRequest{
		WorkspaceID: Ptr(11112222),
		BugID:       Ptr[int64](1111122233301037078),
		LinkIDs:     NewMulti[int64](1162187798001000534, 1162187798001000535),
	})
	require.NoError(t, err)
	assert.True(t, result)
}

func TestBugService_GetConvertBugIDsToQueryToken(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/bugs/ids_to_query_token", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			BugIDs      string `json:"ids"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "1111122233301037078,1111122233301037079", req.BugIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_convert_bug_ids_to_query_token.json"))
	}))

	response, _, err := client.BugService.GetConvertBugIDsToQueryToken(ctx, &GetConvertBugIDsToQueryTokenRequest{
		WorkspaceID: Ptr(11112222),
		BugIDs:      NewMulti[int64](1111122233301037078, 1111122233301037079),
	})
	require.NoError(t, err)
	assert.Equal(t, "71ab88eeb45d084d8fbc85686a0d2399", response.QueryToken)
	assert.Contains(t, response.Href, "71ab88eeb45d084d8fbc85686a0d2399")
}

func TestBugService_GetBugFieldsLabel(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/bugs/get_fields_lable", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/bug/get_bug_fields_lable.json"))
	}))

	labels, _, err := client.BugService.GetBugFieldsLabel(ctx, &GetBugFieldsLabelRequest{
		WorkspaceID: Ptr(11112222),
	})
	require.NoError(t, err)
	require.NotEmpty(t, labels)

	labelMap := make(map[string]string, len(labels))
	for _, label := range labels {
		labelMap[label.EN] = label.CN
	}

	assert.Equal(t, "ID", labelMap["id"])
	assert.Equal(t, "标题", labelMap["title"])
	assert.Equal(t, "详细描述", labelMap["description"])
	assert.Equal(t, "项目ID", labelMap["workspace_id"])
	assert.Equal(t, "严重程度", labelMap["severity"])
	assert.Equal(t, "处理人", labelMap["current_owner"])
}
