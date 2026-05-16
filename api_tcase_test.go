package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestService_CreateTestCase(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcases", r.URL.Path)

		var req CreateTestCaseRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, "测试浏览器兼容性", *req.Name)
		assert.Equal(t, "第一二三步", *req.Steps)
		assert.Equal(t, "打开浏览器", *req.Precondition)
		assert.Equal(t, "无样式错误", *req.Expectation)
		assert.Equal(t, "其它", *req.Type)
		assert.Equal(t, "高", *req.Priority)
		assert.Equal(t, "tapd", *req.Creator)
		assert.Equal(t, TestCaseStatusUpdating, *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/create_test_case.json"))
	}))

	testCase, _, err := client.TestService.CreateTestCase(ctx, &CreateTestCaseRequest{
		WorkspaceID:  Ptr(10158231),
		Name:         Ptr("测试浏览器兼容性"),
		Steps:        Ptr("第一二三步"),
		Precondition: Ptr("打开浏览器"),
		Expectation:  Ptr("无样式错误"),
		Type:         Ptr("其它"),
		Priority:     Ptr("高"),
		Creator:      Ptr("tapd"),
		Status:       Ptr(TestCaseStatusUpdating),
	})
	assert.NoError(t, err)
	require.NotNil(t, testCase)
	assert.Equal(t, "1010158231077224799", testCase.ID)
	assert.Equal(t, "10158231", testCase.WorkspaceID)
	assert.Equal(t, "-1", testCase.CategoryID)
	assert.Equal(t, "2019-06-26 16:42:59", testCase.Created)
	assert.Equal(t, "tapd", testCase.Modifier)
	assert.Equal(t, "tapd", testCase.Creator)
	assert.Equal(t, "测试浏览器兼容性", testCase.Name)
	require.NotNil(t, testCase.Steps)
	assert.Equal(t, "第一二三步", *testCase.Steps)
	require.NotNil(t, testCase.Precondition)
	assert.Equal(t, "打开浏览器", *testCase.Precondition)
	require.NotNil(t, testCase.Expectation)
	assert.Equal(t, "无样式错误", *testCase.Expectation)
	assert.Equal(t, "其它", testCase.Type)
	assert.Equal(t, "高", testCase.Priority)
}

func TestTestService_BatchCreateTestCases(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcases/batch_save", r.URL.Path)

		var req []CreateTestCaseRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		require.Len(t, req, 2)
		assert.Equal(t, 69992160, *req[0].WorkspaceID)
		assert.Equal(t, "简单用例1", *req[0].Name)
		assert.Equal(t, "XX1", *req[0].Creator)
		assert.Equal(t, 69992160, *req[1].WorkspaceID)
		assert.Equal(t, "简单用例2", *req[1].Name)
		assert.Equal(t, "XX2", *req[1].Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/batch_create_test_cases.json"))
	}))

	testCases, _, err := client.TestService.BatchCreateTestCases(ctx, &BatchCreateTestCasesRequest{
		{
			WorkspaceID: Ptr(69992160),
			Name:        Ptr("简单用例1"),
			Creator:     Ptr("XX1"),
		},
		{
			WorkspaceID: Ptr(69992160),
			Name:        Ptr("简单用例2"),
			Creator:     Ptr("XX2"),
		},
	})
	assert.NoError(t, err)
	require.Len(t, testCases, 2)
	assert.Equal(t, "1069992160077456793", testCases[0].ID)
	assert.Equal(t, "1069992160077456793", testCases[0].MID)
	assert.Equal(t, "69992160", testCases[0].WorkspaceID)
	assert.Equal(t, "简单用例1", testCases[0].Name)
	assert.Equal(t, "XX1", testCases[0].Creator)
	assert.Equal(t, TestCaseStatusNormal, testCases[0].Status)
	assert.Equal(t, "1069992160077456795", testCases[1].ID)
	assert.Equal(t, "简单用例2", testCases[1].Name)
	assert.Equal(t, "XX2", testCases[1].Creator)
}

func TestTestService_CreateTestCaseCategory(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcase_categories", r.URL.Path)

		var req CreateTestCaseCategoryRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, "用例目录4", *req.Name)
		assert.Equal(t, "回归测试目录", *req.Description)
		assert.Equal(t, int64(0), *req.ParentID)
		assert.Equal(t, "tester", *req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/create_test_case_category.json"))
	}))

	category, _, err := client.TestService.CreateTestCaseCategory(ctx, &CreateTestCaseCategoryRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("用例目录4"),
		Description: Ptr("回归测试目录"),
		ParentID:    Ptr[int64](0),
		Creator:     Ptr("tester"),
	})
	assert.NoError(t, err)
	require.NotNil(t, category)
	assert.Equal(t, "1010158231000082523", category.ID)
	assert.Equal(t, "10158231", category.WorkspaceID)
	assert.Equal(t, "用例目录4", category.Name)
	require.NotNil(t, category.Description)
	assert.Equal(t, "回归测试目录", *category.Description)
	assert.Equal(t, "0", category.ParentID)
	require.NotNil(t, category.Creator)
	assert.Equal(t, "tester", *category.Creator)
}

func TestTestService_CreateTestPlan(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/test_plans", r.URL.Path)

		var req CreateTestPlanRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, "test_plan_12", *req.Name)
		assert.Equal(t, "这不是一个测试", *req.Description)
		assert.Equal(t, "dev", *req.Creator)
		assert.Equal(t, "owner", *req.Owner)
		assert.Equal(t, "2026-05-01", *req.StartDate)
		assert.Equal(t, "2026-05-31", *req.EndDate)
		assert.Equal(t, int64(1010158231000012345), *req.IterationID)
		assert.Equal(t, "123456", *req.Version)
		assert.Equal(t, "open", *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/create_test_plan.json"))
	}))

	plan, _, err := client.TestService.CreateTestPlan(ctx, &CreateTestPlanRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("test_plan_12"),
		Description: Ptr("这不是一个测试"),
		Creator:     Ptr("dev"),
		Owner:       Ptr("owner"),
		StartDate:   Ptr("2026-05-01"),
		EndDate:     Ptr("2026-05-31"),
		IterationID: Ptr[int64](1010158231000012345),
		Version:     Ptr("123456"),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	require.NotNil(t, plan)
	assert.Equal(t, "1000000755000016443", plan.ID)
	assert.Equal(t, "755", plan.WorkspaceID)
	assert.Equal(t, "test_plan_12", plan.Name)
	assert.Equal(t, "这不是一个测试", plan.Description)
	assert.Equal(t, "123456", plan.Version)
	assert.Equal(t, "open", plan.Status)
	assert.Equal(t, "dev", plan.Creator)
	assert.Equal(t, "api", plan.CreatedFrom)
}

func TestTestService_AssignTestCase(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcase_instance/assign", r.URL.Path)

		var req struct {
			TestPlanID  int64  `json:"test_plan_id"`
			TestCaseID  string `json:"tcase_id"`
			WorkspaceID int    `json:"workspace_id"`
			Executor    string `json:"executor"`
			Assignee    string `json:"assignee"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.TestPlanID)
		assert.Equal(t, "1020357849077231381,1020357849077231382", req.TestCaseID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, "peter", req.Executor)
		assert.Equal(t, "tester", req.Assignee)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/assign_test_case.json"))
	}))

	ok, _, err := client.TestService.AssignTestCase(ctx, &AssignTestCaseRequest{
		TestPlanID:  Ptr[int64](1010158231077224799),
		TestCaseID:  NewMulti[int64](1020357849077231381, 1020357849077231382),
		WorkspaceID: Ptr(10158231),
		Executor:    Ptr("peter"),
		Assignee:    Ptr("tester"),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_CreateTestPlanStoryRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/test_plans/create_story_relation", r.URL.Path)

		var req struct {
			PlanID      int64  `json:"plan_id"`
			WorkspaceID int    `json:"workspace_id"`
			StoryIDs    string `json:"story_ids"`
			Creator     string `json:"creator"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.PlanID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, "123123123,123123124", req.StoryIDs)
		assert.Equal(t, "peter", req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/create_test_plan_story_relation.json"))
	}))

	ok, _, err := client.TestService.CreateTestPlanStoryRelation(ctx, &CreateTestPlanStoryRelationRequest{
		PlanID:      Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
		StoryIDs:    NewMulti[int64](123123123, 123123124),
		Creator:     Ptr("peter"),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_CreateTestPlanTestCaseRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/test_plans/create_tcase_relation", r.URL.Path)

		var req struct {
			TestPlanID  int64  `json:"test_plan_id"`
			WorkspaceID int    `json:"workspace_id"`
			TestCaseIDs string `json:"tcase_ids"`
			Creator     string `json:"creator"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.TestPlanID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, "1020357849077231603,1020357849077231393", req.TestCaseIDs)
		assert.Equal(t, "peter", req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/create_test_plan_test_case_relation.json"))
	}))

	ok, _, err := client.TestService.CreateTestPlanTestCaseRelation(ctx, &CreateTestPlanTestCaseRelationRequest{
		TestPlanID:  Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
		TestCaseIDs: NewMulti[int64](1020357849077231603, 1020357849077231393),
		Creator:     Ptr("peter"),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_DeleteTestPlanStoryRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/test_plans/delete_story_relation", r.URL.Path)

		var req struct {
			PlanID      int64  `json:"plan_id"`
			WorkspaceID int    `json:"workspace_id"`
			StoryIDs    string `json:"story_ids"`
			Creator     string `json:"creator"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.PlanID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, "123123123,123123124", req.StoryIDs)
		assert.Equal(t, "peter", req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/delete_test_plan_story_relation.json"))
	}))

	ok, _, err := client.TestService.DeleteTestPlanStoryRelation(ctx, &DeleteTestPlanStoryRelationRequest{
		PlanID:      Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
		StoryIDs:    NewMulti[int64](123123123, 123123124),
		Creator:     Ptr("peter"),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_DeleteTestCaseStoryRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcase_instance/delete_tcase_story_relation", r.URL.Path)

		var req DeleteTestCaseStoryRelationRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, int64(1020357849500705291), *req.StoryID)
		assert.Equal(t, int64(1020357849077231363), *req.TestCaseID)
		assert.Equal(t, int64(1020357849000015397), *req.TestPlanID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/delete_test_case_story_relation.json"))
	}))

	ok, _, err := client.TestService.DeleteTestCaseStoryRelation(ctx, &DeleteTestCaseStoryRelationRequest{
		WorkspaceID: Ptr(10158231),
		StoryID:     Ptr[int64](1020357849500705291),
		TestCaseID:  Ptr[int64](1020357849077231363),
		TestPlanID:  Ptr[int64](1020357849000015397),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_ExecuteTestCase(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcase_instance/execute", r.URL.Path)

		var req struct {
			TestPlanID   int64  `json:"test_plan_id"`
			TestCaseID   string `json:"tcase_id"`
			WorkspaceID  int    `json:"workspace_id"`
			ResultStatus string `json:"result_status"`
			LastExecutor string `json:"last_executor"`
			ResultRemark string `json:"result_remark"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.TestPlanID)
		assert.Equal(t, "1020357849077231381,1020357849077231382", req.TestCaseID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, "pass", req.ResultStatus)
		assert.Equal(t, "peter", req.LastExecutor)
		assert.Equal(t, "执行通过", req.ResultRemark)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/execute_test_case.json"))
	}))

	ok, _, err := client.TestService.ExecuteTestCase(ctx, &ExecuteTestCaseRequest{
		TestPlanID:   Ptr[int64](1010158231077224799),
		TestCaseID:   NewMulti[int64](1020357849077231381, 1020357849077231382),
		WorkspaceID:  Ptr(10158231),
		ResultStatus: Ptr(TestCaseResultStatusPass),
		LastExecutor: Ptr("peter"),
		ResultRemark: Ptr("执行通过"),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_GetTestCaseRelatedStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcases/get_story_by_tcase_id", r.URL.Path)
		assert.Equal(t, "20358306", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1020358306077237055,1020358306077237053", r.URL.Query().Get("tcase_ids"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_related_stories.json"))
	}))

	relations, _, err := client.TestService.GetTestCaseRelatedStories(ctx, &GetTestCaseRelatedStoriesRequest{
		WorkspaceID: Ptr(20358306),
		TestCaseIDs: NewMulti[int64](1020358306077237055, 1020358306077237053),
	})
	assert.NoError(t, err)
	require.Len(t, relations, 2)
	assert.Equal(t, "20358306", relations[0].WorkspaceID)
	assert.Equal(t, "1020358306077237053", relations[0].TestCaseID)
	assert.Equal(t, "1020358306854812395", relations[0].StoryID)
	assert.Equal(t, "0", relations[0].TestPlanID)
	assert.Equal(t, "1020358306077237055", relations[1].TestCaseID)
}

func TestTestService_GetTestCaseCategories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcase_categories", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "用例目录", r.URL.Query().Get("name"))
		assert.Equal(t, "30", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,name,parent_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_categories.json"))
	}))

	categories, _, err := client.TestService.GetTestCaseCategories(ctx, &GetTestCaseCategoriesRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("用例目录"),
		Limit:       Ptr(30),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "name", "parent_id"),
	})
	assert.NoError(t, err)
	require.Len(t, categories, 2)
	assert.Equal(t, "1010158231075917759", categories[0].ID)
	assert.Equal(t, "10158231", categories[0].WorkspaceID)
	assert.Equal(t, "None Category", categories[0].Name)
	require.NotNil(t, categories[0].Description)
	assert.Equal(t, "未规划目录", *categories[0].Description)
	assert.Equal(t, "0", categories[0].ParentID)
	assert.Nil(t, categories[0].Creator)
	assert.Equal(t, "1010158231000082521", categories[1].ID)
	assert.Equal(t, "用例目录3", categories[1].Name)
	require.NotNil(t, categories[1].Creator)
	assert.Equal(t, "system", *categories[1].Creator)
}

func TestTestService_GetTestCaseCategoriesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcase_categories/count", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "用例目录", r.URL.Query().Get("name"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_categories_count.json"))
	}))

	count, _, err := client.TestService.GetTestCaseCategoriesCount(ctx, &GetTestCaseCategoriesCountRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("用例目录"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, count)
}

func TestTestService_GetTestCaseCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcases/custom_fields_settings", r.URL.Path)
		assert.Equal(t, "755", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_custom_fields_settings.json"))
	}))

	settings, _, err := client.TestService.GetTestCaseCustomFieldsSettings(
		ctx,
		&GetTestCaseCustomFieldsSettingsRequest{
			WorkspaceID: Ptr(755),
		},
	)
	assert.NoError(t, err)
	require.Len(t, settings, 1)
	assert.Equal(t, "1000000755214854654", settings[0].ID)
	assert.Equal(t, "755", settings[0].WorkspaceID)
	assert.Equal(t, "tcase", settings[0].EntryType)
	assert.Equal(t, "custom_field_30", settings[0].CustomField)
	assert.Equal(t, "AT已实现？", settings[0].Name)
	require.NotNil(t, settings[0].Options)
	assert.Equal(t, `{"1":"已实现","2":"未实现"}`, *settings[0].Options)
}

func TestTestService_GetTestCaseFieldsInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcases/get_fields_info", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_fields_info.json"))
	}))

	fields, _, err := client.TestService.GetTestCaseFieldsInfo(ctx, &GetTestCaseFieldsInfoRequest{
		WorkspaceID: Ptr(10104801),
	})
	assert.NoError(t, err)
	require.Len(t, fields, 3)

	var statusField *TestCaseFieldsInfo
	for _, field := range fields {
		if field.Name == "status" {
			statusField = field
			break
		}
	}
	require.NotNil(t, statusField)
	assert.Equal(t, TestCaseFieldsInfoHTMLTypeSelect, statusField.HTMLType)
	assert.Equal(t, "用例状态", statusField.Label)
	assert.Equal(t, "normal", statusField.Default)
	assert.Contains(t, statusField.Options, TestCaseFieldsInfoOption{
		Key:   "normal",
		Label: "正常",
	})
}

func TestTestService_GetTestCaseResults(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcase_instance/result", r.URL.Path)
		assert.Equal(t, "1010158231077224799", r.URL.Query().Get("test_plan_id"))
		assert.Equal(t, "1020357849077231381", r.URL.Query().Get("tcase_id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_case_results.json"))
	}))

	results, _, err := client.TestService.GetTestCaseResults(ctx, &GetTestCaseResultsRequest{
		TestPlanID:  Ptr[int64](1010158231077224799),
		TestCaseID:  Ptr[int64](1020357849077231381),
		WorkspaceID: Ptr(10158231),
	})
	assert.NoError(t, err)
	require.Len(t, results, 2)
	passed := findTestCaseResultItem(results, "1020357849000703565")
	require.NotNil(t, passed)
	require.NotNil(t, passed.Result)
	assert.Equal(t, TestCaseResultStatusPass, passed.Result.ResultStatus)
	assert.Equal(t, "jeffjffang", passed.Result.Executor)
	failed := findTestCaseResultItem(results, "1020357849000703483")
	require.NotNil(t, failed)
	require.NotNil(t, failed.Result)
	assert.Equal(t, []string{"1020357849500655643"}, failed.Result.BugID)
	require.Len(t, failed.Result.Bugs, 1)
	assert.Equal(t, "用例失败bug关联", failed.Result.Bugs[0].Title)
}

func TestTestService_GetTestCases(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcases", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "测试浏览器", r.URL.Query().Get("name"))
		assert.Equal(t, "normal", r.URL.Query().Get("status"))
		assert.Equal(t, "20", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,name,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_cases.json"))
	}))

	testCases, _, err := client.TestService.GetTestCases(ctx, &GetTestCasesRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("测试浏览器"),
		Status:      NewEnum(TestCaseStatusNormal),
		Limit:       Ptr(20),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "name", "status"),
	})
	assert.NoError(t, err)
	require.Len(t, testCases, 2)
	assert.Equal(t, "1120003271001000049", testCases[0].ID)
	assert.Equal(t, TestCaseStatusAbandon, testCases[0].Status)
	assert.Equal(t, "1010158231077224799", testCases[1].ID)
	assert.Equal(t, "测试浏览器兼容性", testCases[1].Name)
}

func TestTestService_GetTestCasesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/tcases/count", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1020357849000015397", r.URL.Query().Get("test_plan_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_cases_count.json"))
	}))

	count, _, err := client.TestService.GetTestCasesCount(ctx, &GetTestCasesCountRequest{
		WorkspaceID: Ptr(10158231),
		TestPlanID:  Ptr[int64](1020357849000015397),
	})
	assert.NoError(t, err)
	assert.Equal(t, 10, count)
}

func TestTestService_GetTestPlanRelatedBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/result_relation_bugs", r.URL.Path)
		assert.Equal(t, "1010158231077224799", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plan_related_bugs.json"))
	}))

	relations, _, err := client.TestService.GetTestPlanRelatedBugs(ctx, &GetTestPlanRelatedBugsRequest{
		ID:          Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
	})
	assert.NoError(t, err)
	require.Len(t, relations, 1)
	assert.Equal(t, "1020357849077231363", relations[0].ID)
	assert.Equal(t, "用例1", relations[0].Name)
	result := findTestCaseResultItem(relations[0].TestCaseResultRelatedBugs, "1020357849000703483")
	require.NotNil(t, result)
	require.NotNil(t, result.Result)
	assert.Equal(t, TestCaseResultStatusPass, result.Result.ResultStatus)
	require.Len(t, result.Result.Bugs, 1)
	assert.Equal(t, "1231", result.Result.Bugs[0].Title)
}

func findTestCaseResultItem(items []*TestCaseResultItem, id string) *TestCaseResultItem {
	for _, item := range items {
		if item.ID == id {
			return item
		}
	}

	return nil
}

func TestTestService_GetIterationTestPlans(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/get_by_iteration_id", r.URL.Path)
		assert.Equal(t, "51650666", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1151650666001000111", r.URL.Query().Get("iteration_id"))
		assert.Equal(t, "30", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_iteration_test_plans.json"))
	}))

	plans, _, err := client.TestService.GetIterationTestPlans(ctx, &GetIterationTestPlansRequest{
		WorkspaceID: Ptr(51650666),
		IterationID: Ptr[int64](1151650666001000111),
		Limit:       Ptr(30),
		Page:        Ptr(1),
	})
	assert.NoError(t, err)
	require.Len(t, plans, 2)
	assert.Equal(t, "51650666", plans[0].WorkspaceID)
	assert.Equal(t, "1151650666001000111", plans[0].IterationID)
	assert.Equal(t, "1151650666001000019", plans[0].TestPlanID)
}

func TestTestService_GetTestPlanResult(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/details", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1010158231000005241", r.URL.Query().Get("id"))
		assert.Equal(t, "1", r.URL.Query().Get("include_repeat"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plan_result.json"))
	}))

	testCases, _, err := client.TestService.GetTestPlanResult(ctx, &GetTestPlanResultRequest{
		WorkspaceID:   Ptr(10158231),
		ID:            Ptr[int64](1010158231000005241),
		IncludeRepeat: Ptr(1),
	})
	assert.NoError(t, err)
	require.Len(t, testCases, 2)
	assert.Equal(t, "1010158231075919347", testCases[0].ID)
	require.NotNil(t, testCases[0].Result)
	assert.Equal(t, "0", testCases[0].Result.ID)
	assert.Equal(t, TestCaseResultStatusUnexecuted, testCases[0].Result.ResultStatus)
	assert.Equal(t, "1010158231075919345", testCases[1].ID)
	require.NotNil(t, testCases[1].Result)
	assert.Equal(t, "1010158231000703323", testCases[1].Result.ID)
	assert.Equal(t, TestCaseResultStatusBlock, testCases[1].Result.ResultStatus)
}

func TestTestService_GetTestPlanProgress(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/progress", r.URL.Path)
		assert.Equal(t, "1010158231077224799", r.URL.Query().Get("id"))
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plan_progress.json"))
	}))

	progress, _, err := client.TestService.GetTestPlanProgress(ctx, &GetTestPlanProgressRequest{
		ID:          Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
	})
	assert.NoError(t, err)
	require.NotNil(t, progress)
	assert.Equal(t, 1, progress.StoryCount)
	assert.Equal(t, 10, progress.TestCaseCount)
	assert.Equal(t, 5, progress.StatusCounter[TestCaseResultStatusPass])
	assert.Equal(t, 5, progress.StatusCounter[TestCaseResultStatusUnexecuted])
	assert.Equal(t, "50%", progress.ExecutedRate)
}

func TestTestService_GetTestPlanTestCaseRelations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/get_test_plan_tcase", r.URL.Path)
		assert.Equal(t, "755", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1000000755077233617", r.URL.Query().Get("test_plan_id"))
		assert.Equal(t, "30", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plan_test_case_relations.json"))
	}))

	relations, _, err := client.TestService.GetTestPlanTestCaseRelations(ctx, &GetTestPlanTestCaseRelationsRequest{
		WorkspaceID: Ptr(755),
		TestPlanID:  Ptr[int64](1000000755077233617),
		Limit:       Ptr(30),
		Page:        Ptr(1),
	})
	assert.NoError(t, err)
	require.Len(t, relations, 2)
	assert.Equal(t, "1000000755002248699", relations[0].ID)
	assert.Equal(t, "755", relations[0].WorkspaceID)
	assert.Equal(t, "1000000755077233617", relations[0].TestPlanID)
	assert.Equal(t, "1000000755000026804", relations[0].TestCaseID)
}

func TestTestService_GetTestPlans(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans", r.URL.Path)
		assert.Equal(t, "10158231", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "test_plan", r.URL.Query().Get("name"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))
		assert.Equal(t, "20", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,name,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plans.json"))
	}))

	plans, _, err := client.TestService.GetTestPlans(ctx, &GetTestPlansRequest{
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("test_plan"),
		Status:      Ptr("open"),
		Limit:       Ptr(20),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "name", "status"),
	})
	assert.NoError(t, err)
	require.Len(t, plans, 2)
	assert.Equal(t, "1000000755000016443", plans[0].ID)
	assert.Equal(t, "test_plan_12", plans[0].Name)
	assert.Equal(t, "open", plans[0].Status)
	assert.Equal(t, "1000000755000016444", plans[1].ID)
	assert.Equal(t, "close", plans[1].Status)
}

func TestTestService_GetTestPlansCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/count", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "open", r.URL.Query().Get("status"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plans_count.json"))
	}))

	count, _, err := client.TestService.GetTestPlansCount(ctx, &GetTestPlansCountRequest{
		WorkspaceID: Ptr(10104801),
		Status:      Ptr("open"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 4, count)
}

func TestTestService_RemoveTestCaseFromTestPlan(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcase_instance/remove_tcase", r.URL.Path)

		var req struct {
			TestPlanID  int64  `json:"test_plan_id"`
			WorkspaceID int    `json:"workspace_id"`
			StoryID     int64  `json:"story_id"`
			TestCaseID  string `json:"tcase_id"`
		}
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), req.TestPlanID)
		assert.Equal(t, 10158231, req.WorkspaceID)
		assert.Equal(t, int64(1020357849500705291), req.StoryID)
		assert.Equal(t, "1020357849077231363,1020357849077231364", req.TestCaseID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/remove_test_case_from_test_plan.json"))
	}))

	ok, _, err := client.TestService.RemoveTestCaseFromTestPlan(ctx, &RemoveTestCaseFromTestPlanRequest{
		TestPlanID:  Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
		StoryID:     Ptr[int64](1020357849500705291),
		TestCaseID:  NewMulti[int64](1020357849077231363, 1020357849077231364),
	})
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestTestService_UpdateTestCase(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/tcases", r.URL.Path)

		var req UpdateTestCaseRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1010158231077224799), *req.ID)
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, TestCaseStatusAbandon, *req.Status)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/update_test_case.json"))
	}))

	testCase, _, err := client.TestService.UpdateTestCase(ctx, &UpdateTestCaseRequest{
		ID:          Ptr[int64](1010158231077224799),
		WorkspaceID: Ptr(10158231),
		Status:      Ptr(TestCaseStatusAbandon),
	})
	assert.NoError(t, err)
	require.NotNil(t, testCase)
	assert.Equal(t, "1010158231077224799", testCase.ID)
	assert.Equal(t, TestCaseStatusAbandon, testCase.Status)
	assert.Equal(t, "api_doc_oauth", testCase.Modifier)
}

func TestTestService_UpdateTestPlan(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/test_plans", r.URL.Path)

		var req UpdateTestPlanRequest
		require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1000000755000016443), *req.ID)
		assert.Equal(t, 10158231, *req.WorkspaceID)
		assert.Equal(t, "test", *req.Name)
		assert.Equal(t, "tapd", *req.Modifier)

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/update_test_plan.json"))
	}))

	plan, _, err := client.TestService.UpdateTestPlan(ctx, &UpdateTestPlanRequest{
		ID:          Ptr[int64](1000000755000016443),
		WorkspaceID: Ptr(10158231),
		Name:        Ptr("test"),
		Modifier:    Ptr("tapd"),
	})
	assert.NoError(t, err)
	require.NotNil(t, plan)
	assert.Equal(t, "1000000755000016443", plan.ID)
	assert.Equal(t, "test", plan.Name)
	assert.Equal(t, "tapd", plan.Modifier)
}

func TestTestService_GetTestPlanFieldsInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/test_plans/get_fields_info", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/tcase/get_test_plan_fields_info.json"))
	}))

	fields, _, err := client.TestService.GetTestPlanFieldsInfo(ctx, &GetTestPlanFieldsInfoRequest{
		WorkspaceID: Ptr(10104801),
	})
	assert.NoError(t, err)
	require.Len(t, fields, 4)
	statusField := findTestPlanFieldsInfo(fields, "status")
	require.NotNil(t, statusField)
	assert.Equal(t, TestCaseFieldsInfoHTMLTypeSelect, statusField.HTMLType)
	assert.Equal(t, "状态", statusField.Label)
	assert.Contains(t, statusField.Options, TestCaseFieldsInfoOption{
		Key:   "open",
		Label: "开启",
	})
}

func findTestPlanFieldsInfo(fields []*TestPlanFieldsInfo, name string) *TestPlanFieldsInfo {
	for _, field := range fields {
		if field.Name == name {
			return field
		}
	}

	return nil
}
