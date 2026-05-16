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
