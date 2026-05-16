package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseService_CreateRelease(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/new_releases", r.URL.Path)

		var req CreateReleaseRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, "test2", *req.Name)
		assert.Equal(t, "发布计划描述", *req.Description)
		assert.Equal(t, "2020-10-20", *req.StartDate)
		assert.Equal(t, "2020-11-20", *req.EndDate)
		assert.Equal(t, "dev", *req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/create_release.json"))
	}))

	release, _, err := client.ReleaseService.CreateRelease(ctx, &CreateReleaseRequest{
		WorkspaceID: Ptr(10104801),
		Name:        Ptr("test2"),
		Description: Ptr("发布计划描述"),
		StartDate:   Ptr("2020-10-20"),
		EndDate:     Ptr("2020-11-20"),
		Creator:     Ptr("dev"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801100003081", release.ID)
	assert.Equal(t, "10104801", release.WorkspaceID)
	assert.Equal(t, "test2", release.Name)
	assert.Nil(t, release.Description)
	assert.Equal(t, "2020-10-20", release.StartDate)
	assert.Equal(t, "2020-11-20", release.EndDate)
	assert.Equal(t, "dev", release.Creator)
	assert.Equal(t, "2020-10-27 11:09:14", release.Created)
	assert.Equal(t, "2020-10-27 11:09:14", release.Modified)
	assert.Equal(t, "open", release.Status)
}

func TestReleaseService_GetReleases(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/releases", r.URL.Path)
		assert.Equal(t, "1010158231100000905,1010158231100000906", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "发布计划", r.URL.Query().Get("name"))
		assert.Equal(t, "敏捷研发", r.URL.Query().Get("description"))
		assert.Equal(t, "2017-06-12", r.URL.Query().Get("startdate"))
		assert.Equal(t, "2017-07-07", r.URL.Query().Get("enddate"))
		assert.Equal(t, "anyechen", r.URL.Query().Get("creator"))
		assert.Equal(t, "2017-06-20", r.URL.Query().Get("created"))
		assert.Equal(t, "2017-06-21", r.URL.Query().Get("modified"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_releases.json"))
	}))

	releases, _, err := client.ReleaseService.GetReleases(ctx, &GetReleasesRequest{
		ID:          NewMulti[int64](1010158231100000905, 1010158231100000906),
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("发布计划"),
		Description: Ptr("敏捷研发"),
		StartDate:   Ptr("2017-06-12"),
		EndDate:     Ptr("2017-07-07"),
		Creator:     Ptr("anyechen"),
		Created:     Ptr("2017-06-20"),
		Modified:    Ptr("2017-06-21"),
		Status:      Ptr("open"),
		Limit:       Ptr(10),
		Page:        Ptr(2),
		Order:       NewOrder("created", OrderByDesc),
		Fields:      NewMulti("id", "name", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, releases, 1)
	assert.Equal(t, "1010158231100000905", releases[0].ID)
	assert.Equal(t, "10158231", releases[0].WorkspaceID)
	assert.Equal(t, "发布计划1", releases[0].Name)
	assert.Equal(t, "熟悉敏捷研发全生命周期", *releases[0].Description)
	assert.Equal(t, "2017-06-12", releases[0].StartDate)
	assert.Equal(t, "2017-07-07", releases[0].EndDate)
	assert.Equal(t, "anyechen", releases[0].Creator)
	assert.Equal(t, "2017-06-20 16:49:01", releases[0].Created)
	assert.Equal(t, "2017-06-20 16:49:01", releases[0].Modified)
	assert.Equal(t, "open", releases[0].Status)
}

func TestReleaseService_GetReleasesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/releases/count", r.URL.Path)
		assert.Equal(t, "1010158231100000905,1010158231100000906", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "发布计划", r.URL.Query().Get("name"))
		assert.Equal(t, "敏捷研发", r.URL.Query().Get("description"))
		assert.Equal(t, "2017-06-12", r.URL.Query().Get("startdate"))
		assert.Equal(t, "2017-07-07", r.URL.Query().Get("enddate"))
		assert.Equal(t, "anyechen", r.URL.Query().Get("creator"))
		assert.Equal(t, "2017-06-20", r.URL.Query().Get("created"))
		assert.Equal(t, "2017-06-21", r.URL.Query().Get("modified"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_releases_count.json"))
	}))

	count, _, err := client.ReleaseService.GetReleasesCount(ctx, &GetReleasesCountRequest{
		ID:          NewMulti[int64](1010158231100000905, 1010158231100000906),
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("发布计划"),
		Description: Ptr("敏捷研发"),
		StartDate:   Ptr("2017-06-12"),
		EndDate:     Ptr("2017-07-07"),
		Creator:     Ptr("anyechen"),
		Created:     Ptr("2017-06-20"),
		Modified:    Ptr("2017-06-21"),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestReleaseService_UpdateRelease(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/new_releases", r.URL.Path)

		var req UpdateReleaseRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, int64(1010104801100003081), *req.ID)
		assert.Equal(t, "test2", *req.Name)
		assert.Equal(t, "内容被更新", *req.Description)
		assert.Equal(t, "2020-10-20", *req.StartDate)
		assert.Equal(t, "2020-11-20", *req.EndDate)
		assert.Equal(t, "open", *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/update_release.json"))
	}))

	release, _, err := client.ReleaseService.UpdateRelease(ctx, &UpdateReleaseRequest{
		WorkspaceID: Ptr(10104801),
		ID:          Ptr[int64](1010104801100003081),
		Name:        Ptr("test2"),
		Description: Ptr("内容被更新"),
		StartDate:   Ptr("2020-10-20"),
		EndDate:     Ptr("2020-11-20"),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801100003081", release.ID)
	assert.Equal(t, "10104801", release.WorkspaceID)
	assert.Equal(t, "test2", release.Name)
	assert.Equal(t, "内容被更新", *release.Description)
	assert.Equal(t, "2020-10-27 11:24:48", release.Modified)
	assert.Equal(t, "open", release.Status)
}

func TestReleaseService_GetLaunchAccessories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_accessories", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010104801000402051", r.URL.Query().Get("form_id"))
		assert.Equal(t, "1010104801000253485", r.URL.Query().Get("id"))
		assert.Equal(t, "v_xuanfang", r.URL.Query().Get("created_by"))
		assert.Equal(t, "2020-06-11", r.URL.Query().Get("created"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_accessories.json"))
	}))

	accessories, _, err := client.ReleaseService.GetLaunchAccessories(ctx, &GetLaunchAccessoriesRequest{
		WorkspaceID: Ptr(10104801),
		FormID:      Ptr[int64](1010104801000402051),
		ID:          Ptr[int64](1010104801000253485),
		CreatedBy:   Ptr("v_xuanfang"),
		Created:     Ptr("2020-06-11"),
	})
	assert.NoError(t, err)
	assert.Len(t, accessories, 2)
	assert.Equal(t, "1010104801000253485", accessories[0].ID)
	assert.Equal(t, "1010104801000402051", accessories[0].FormID)
	assert.Equal(t, "10104801", accessories[0].WorkspaceID)
	assert.Equal(t, "launch_tasks_list", accessories[0].Type)
	assert.Equal(t, "任务列表", accessories[0].Title)
	assert.Equal(t, "1010104801500601739", accessories[0].Content)
	assert.Equal(t, "task", *accessories[0].ContentType)
	assert.Equal(t, "v_xuanfang", accessories[0].CreatedBy)
	assert.Equal(t, "2020-06-11 16:17:56", accessories[0].Created)
}

func TestReleaseService_GetLaunchForms(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_forms", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010104801079697767", r.URL.Query().Get("id"))
		assert.Equal(t, "v_xuanfang", r.URL.Query().Get("creator"))
		assert.Equal(t, "2021-01-15", r.URL.Query().Get("created"))
		assert.Equal(t, "发布评审", r.URL.Query().Get("title"))
		assert.Equal(t, "LAUNCHFORM_STATUS_INITIAL", r.URL.Query().Get("status"))
		assert.Equal(t, "version", r.URL.Query().Get("version_type"))
		assert.Equal(t, "baseline", r.URL.Query().Get("baseline"))
		assert.Equal(t, "module", r.URL.Query().Get("release_model"))
		assert.Equal(t, "roadmap", r.URL.Query().Get("roadmap_version"))
		assert.Equal(t, "正常发布", r.URL.Query().Get("release_type"))
		assert.Equal(t, "normal", r.URL.Query().Get("change_type"))
		assert.Equal(t, "signer", r.URL.Query().Get("signed_by"))
		assert.Equal(t, "archiver", r.URL.Query().Get("archived_by"))
		assert.Equal(t, "cc", r.URL.Query().Get("cc"))
		assert.Equal(t, "notifier", r.URL.Query().Get("change_notifier"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,title,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_forms.json"))
	}))

	forms, _, err := client.ReleaseService.GetLaunchForms(ctx, &GetLaunchFormsRequest{
		WorkspaceID:    Ptr(10104801),
		ID:             Ptr[int64](1010104801079697767),
		Creator:        Ptr("v_xuanfang"),
		Created:        Ptr("2021-01-15"),
		Title:          Ptr("发布评审"),
		Status:         Ptr("LAUNCHFORM_STATUS_INITIAL"),
		VersionType:    Ptr("version"),
		Baseline:       Ptr("baseline"),
		ReleaseModel:   Ptr("module"),
		RoadmapVersion: Ptr("roadmap"),
		ReleaseType:    Ptr("正常发布"),
		ChangeType:     Ptr("normal"),
		SignedBy:       Ptr("signer"),
		ArchivedBy:     Ptr("archiver"),
		CC:             Ptr("cc"),
		ChangeNotifier: Ptr("notifier"),
		Limit:          Ptr(10),
		Page:           Ptr(1),
		Fields:         NewMulti("id", "title", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, forms, 1)
	assert.Equal(t, "1010104801079697767", forms[0].ID)
	assert.Nil(t, forms[0].Title)
	assert.Equal(t, "202101150008", forms[0].Name)
	assert.Equal(t, "v_xuanfang", forms[0].Creator)
	assert.Equal(t, "10104801", forms[0].WorkspaceID)
	assert.Equal(t, "LAUNCHFORM_STATUS_INITIAL", forms[0].Status)
	assert.Equal(t, "正常发布", *forms[0].ReleaseType)
	assert.Equal(t, ";v_xuanfang;", *forms[0].Participator)
	assert.Equal(t, "1010104801065798331", forms[0].TemplateID)
	assert.Equal(t, "LAUNCHFORM_STATUS_INITIAL", forms[0].Flows)
}

func TestReleaseService_CreateLaunchForm(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/launch_forms", r.URL.Path)

		var req CreateLaunchFormRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, "tapd_api", *req.Creator)
		assert.Equal(t, "1010104801065798331", *req.TemplateID)
		assert.Equal(t, "发布评审", *req.Title)
		assert.Equal(t, "version", *req.VersionType)
		assert.Equal(t, "baseline", *req.Baseline)
		assert.Equal(t, "module", *req.ReleaseModel)
		assert.Equal(t, "roadmap", *req.RoadmapVersion)
		assert.Equal(t, "正常发布", *req.ReleaseType)
		assert.Equal(t, "signer", *req.SignedBy)
		assert.Equal(t, "archiver", *req.ArchivedBy)
		assert.Equal(t, "cc", *req.CC)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/create_launch_form.json"))
	}))

	form, _, err := client.ReleaseService.CreateLaunchForm(ctx, &CreateLaunchFormRequest{
		WorkspaceID:    Ptr(10104801),
		Creator:        Ptr("tapd_api"),
		TemplateID:     Ptr("1010104801065798331"),
		Title:          Ptr("发布评审"),
		VersionType:    Ptr("version"),
		Baseline:       Ptr("baseline"),
		ReleaseModel:   Ptr("module"),
		RoadmapVersion: Ptr("roadmap"),
		ReleaseType:    Ptr("正常发布"),
		SignedBy:       Ptr("signer"),
		ArchivedBy:     Ptr("archiver"),
		CC:             Ptr("cc"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801079724013", form.ID)
	assert.Nil(t, form.Title)
	assert.Equal(t, "202107210009", form.Name)
	assert.Equal(t, "v_xuanfang", form.Creator)
	assert.Equal(t, "10104801", form.WorkspaceID)
	assert.Equal(t, "initial", form.Status)
	assert.Equal(t, "正常发布", *form.ReleaseType)
	assert.Equal(t, ";v_xuanfang;", *form.Participator)
	assert.Equal(t, "1010104801065798331", form.TemplateID)
}

func TestReleaseService_CreateLaunchAccessory(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/launch_accessories", r.URL.Path)

		var req CreateLaunchAccessoryRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, *req.WorkspaceID)
		assert.Equal(t, int64(1010104801079533889), *req.FormID)
		assert.Equal(t, "launch_url", *req.Type)
		assert.Equal(t, "https://www.tapd.cn/", *req.Content)

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/create_launch_accessory.json"))
	}))

	accessory, _, err := client.ReleaseService.CreateLaunchAccessory(ctx, &CreateLaunchAccessoryRequest{
		WorkspaceID: Ptr(10104801),
		FormID:      Ptr[int64](1010104801079533889),
		Type:        Ptr("launch_url"),
		Content:     Ptr("https://www.tapd.cn/"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1010104801000254035", accessory.ID)
	assert.Equal(t, "1010104801079533889", accessory.FormID)
	assert.Equal(t, "10104801", accessory.WorkspaceID)
	assert.Equal(t, "launch_url", accessory.Type)
	assert.Equal(t, "URL", accessory.Title)
	assert.Equal(t, "https://www.tapd.cn/", accessory.Content)
	assert.Equal(t, "tapd", accessory.CreatedBy)
	assert.Equal(t, "2022-09-08 16:45:30", accessory.Created)
}

func TestReleaseService_GetLaunchFormsCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_forms/count", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010104801079697767", r.URL.Query().Get("id"))
		assert.Equal(t, "v_xuanfang", r.URL.Query().Get("creator"))
		assert.Equal(t, "2021-01-15", r.URL.Query().Get("created"))
		assert.Equal(t, "发布评审", r.URL.Query().Get("title"))
		assert.Equal(t, "LAUNCHFORM_STATUS_INITIAL", r.URL.Query().Get("status"))
		assert.Equal(t, "version", r.URL.Query().Get("version_type"))
		assert.Equal(t, "baseline", r.URL.Query().Get("baseline"))
		assert.Equal(t, "module", r.URL.Query().Get("release_model"))
		assert.Equal(t, "roadmap", r.URL.Query().Get("roadmap_version"))
		assert.Equal(t, "正常发布", r.URL.Query().Get("release_type"))
		assert.Equal(t, "normal", r.URL.Query().Get("change_type"))
		assert.Equal(t, "signer", r.URL.Query().Get("signed_by"))
		assert.Equal(t, "archiver", r.URL.Query().Get("archived_by"))
		assert.Equal(t, "cc", r.URL.Query().Get("cc"))
		assert.Equal(t, "notifier", r.URL.Query().Get("change_notifier"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,title", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_forms_count.json"))
	}))

	count, _, err := client.ReleaseService.GetLaunchFormsCount(ctx, &GetLaunchFormsCountRequest{
		WorkspaceID:    Ptr(10104801),
		ID:             Ptr[int64](1010104801079697767),
		Creator:        Ptr("v_xuanfang"),
		Created:        Ptr("2021-01-15"),
		Title:          Ptr("发布评审"),
		Status:         Ptr("LAUNCHFORM_STATUS_INITIAL"),
		VersionType:    Ptr("version"),
		Baseline:       Ptr("baseline"),
		ReleaseModel:   Ptr("module"),
		RoadmapVersion: Ptr("roadmap"),
		ReleaseType:    Ptr("正常发布"),
		ChangeType:     Ptr("normal"),
		SignedBy:       Ptr("signer"),
		ArchivedBy:     Ptr("archiver"),
		CC:             Ptr("cc"),
		ChangeNotifier: Ptr("notifier"),
		Limit:          Ptr(10),
		Page:           Ptr(1),
		Fields:         NewMulti("id", "title"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestReleaseService_GetLaunchFormCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_forms/custom_fields_settings", r.URL.Path)
		assert.Equal(t, "20003271", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_form_custom_fields_settings.json"))
	}))

	settings, _, err := client.ReleaseService.GetLaunchFormCustomFieldsSettings(
		ctx,
		&GetLaunchFormCustomFieldsSettingsRequest{
			WorkspaceID: Ptr(20003271),
		},
	)
	assert.NoError(t, err)
	assert.Len(t, settings, 1)
	assert.Equal(t, "1120003271001000004", settings[0].ID)
	assert.Equal(t, "20003271", settings[0].WorkspaceID)
	assert.Equal(t, "launchform", settings[0].EntryType)
	assert.Equal(t, "custom_field_one", settings[0].CustomField)
	assert.Equal(t, "textarea", settings[0].Type)
	assert.Equal(t, "DB变更", settings[0].Name)
	assert.Nil(t, settings[0].Options)
	assert.Equal(t, "1", settings[0].Enabled)
	assert.Nil(t, settings[0].Sort)
	assert.Nil(t, settings[0].Memo)
}

func TestReleaseService_GetLaunchFormTemplates(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_forms/templates", r.URL.Path)
		assert.Equal(t, "20042301", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_form_templates.json"))
	}))

	templates, _, err := client.ReleaseService.GetLaunchFormTemplates(ctx, &GetLaunchFormTemplatesRequest{
		WorkspaceID: Ptr(20042301),
	})
	assert.NoError(t, err)
	assert.Len(t, templates, 3)
	assert.Equal(t, "1120042301001000009", templates[0].ID)
	assert.Equal(t, "系统默认流程", templates[0].Name)
	assert.Equal(t, "1120042301001000079", templates[2].ID)
	assert.Equal(t, "迭代", templates[2].Name)
}

func TestReleaseService_GetLaunchFormActivityLogs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/launch_forms/get_activity_logs", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010104801079777231", r.URL.Query().Get("form_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/release/get_launch_form_activity_logs.json"))
	}))

	logs, _, err := client.ReleaseService.GetLaunchFormActivityLogs(ctx, &GetLaunchFormActivityLogsRequest{
		WorkspaceID: Ptr(10104801),
		FormID:      Ptr[int64](1010104801079777231),
	})
	assert.NoError(t, err)
	assert.Len(t, logs, 2)
	assert.Equal(t, "1010104801083448610", logs[0].ID)
	assert.Equal(t, "10104801", logs[0].WorkspaceID)
	assert.Equal(t, "audit", logs[0].Type)
	assert.Equal(t, "1010104801079777231", logs[0].FormID)
	assert.Equal(t, "audit_result", logs[0].Field)
	assert.JSONEq(t, `{"result":"pass","audited_by":"v_xuanfang"}`, string(logs[0].NewValue))
	assert.JSONEq(t, `[]`, string(logs[0].FactorResult))
	assert.Equal(t, "v_xuanfang", logs[0].CreatedBy)
	assert.Equal(t, "2023-07-13 10:41:23", logs[0].Created)
	assert.Equal(t, "audit_absolutely", logs[0].Operation)
	assert.Equal(t, "initialization", logs[1].Field)
	assert.JSONEq(t, `null`, string(logs[1].OldValue))
	assert.JSONEq(t, `null`, string(logs[1].NewValue))
}
