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
