package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkspaceService_GetUsers(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workspaces/users", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "张三,李四", r.URL.Query().Get("user"))
		assert.Equal(t, "id,name", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/users.json"))
	}))

	users, _, err := client.WorkspaceService.GetUsers(ctx, &GetUsersRequest{
		WorkspaceID: Ptr(11112222),
		User:        NewMulti("张三", "李四"),
		Fields:      NewMulti("id", "name"),
	})
	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Contains(t, users, &User{
		User:             "张三",
		UserID:           "",
		RoleID:           []string{"11111122222001000029"},
		Name:             "张三",
		Email:            "",
		JoinProjectTime:  nil,
		RealJoinTime:     "2018-07-03",
		Status:           "2",
		Allocation:       "100",
		LeaveProjectTime: nil,
	})
	assert.Contains(t, users, &User{
		User:             "李四",
		UserID:           "",
		RoleID:           []string{"11111122222001000028", "11111122222001000143"},
		Name:             "李四",
		Email:            "",
		JoinProjectTime:  nil,
		RealJoinTime:     "2018-07-09",
		Status:           "1",
		Allocation:       "100",
		LeaveProjectTime: nil,
	})
}

func TestWorkspaceService_GetUsersList(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workspaces/users", r.URL.Path)
		assert.Equal(t, "20003271", r.URL.Query().Get("workspace_id"))
		assert.Empty(t, r.URL.Query().Get("user"))
		assert.Equal(t, "user,user_id,role_id,name,email,real_join_time", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/users_list.json"))
	}))

	users, _, err := client.WorkspaceService.GetUsers(ctx, &GetUsersRequest{
		WorkspaceID: Ptr(20003271),
		Fields:      NewMulti("user", "user_id", "role_id", "name", "email", "real_join_time"),
	})
	require.NoError(t, err)
	require.Len(t, users, 1)
	assert.Equal(t, "davidning", users[0].User)
	assert.Equal(t, "123456", users[0].UserID)
	assert.Equal(t, []string{"1000000000000000010", "1000000000000000015"}, users[0].RoleID)
	assert.Equal(t, "David", users[0].Name)
	assert.Equal(t, "david@example.com", users[0].Email)
	assert.Equal(t, "2025-04-17", users[0].RealJoinTime)
}

func TestWorkspaceService_AddWorkspaceMember(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/workspaces/add_workspace_member", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			Nick        string `json:"nick"`
			CompanyID   int    `json:"company_id"`
			RoleIDs     string `json:"role_ids"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 10104801, req.WorkspaceID)
		assert.Equal(t, "davidning", req.Nick)
		assert.Equal(t, 20003271, req.CompanyID)
		assert.Equal(t, "1000000000000000010,1000000000000000015", req.RoleIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/add_workspace_member.json"))
	}))

	result, _, err := client.WorkspaceService.AddWorkspaceMember(ctx, &AddWorkspaceMemberRequest{
		WorkspaceID: Ptr(10104801),
		Nick:        Ptr("davidning"),
		CompanyID:   Ptr(20003271),
		RoleIDs:     NewMulti[int64](1000000000000000010, 1000000000000000015),
	})
	require.NoError(t, err)
	assert.True(t, result.Success)
}

func TestWorkspaceService_GetWorkspaceRoles(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/roles", r.URL.Path)
		assert.Equal(t, "10104801", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/roles.json"))
	}))

	roles, _, err := client.WorkspaceService.GetWorkspaceRoles(ctx, &GetWorkspaceRolesRequest{
		WorkspaceID: Ptr(10104801),
	})
	require.NoError(t, err)
	assert.Contains(t, roles, &WorkspaceRole{ID: "1000000000000000002", Name: "管理员"})
	assert.Contains(t, roles, &WorkspaceRole{ID: "1000000000000000010", Name: "测试人员"})
}

func TestWorkspaceService_UpdateWorkspaceInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/workspaces/update_workspace_info", r.URL.Path)

		var req UpdateWorkspaceInfoRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 69999237, *req.WorkspaceID)
		assert.Equal(t, "end_date", *req.Field)
		assert.Equal(t, "2025-03-03", *req.Value)

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/update_workspace_info.json"))
	}))

	result, _, err := client.WorkspaceService.UpdateWorkspaceInfo(ctx, &UpdateWorkspaceInfoRequest{
		WorkspaceID: Ptr(69999237),
		Field:       Ptr("end_date"),
		Value:       Ptr("2025-03-03"),
	})
	require.NoError(t, err)
	assert.Equal(t, "update workspace success", result)
}

func TestWorkspaceService_GetWorkspaceInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workspaces/get_workspace_info", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/get_workspace_info.json"))
	}))

	workspace, _, err := client.WorkspaceService.GetWorkspaceInfo(ctx, &GetWorkspaceInfoRequest{
		WorkspaceID: Ptr(11112222),
	})
	require.NoError(t, err)
	assert.Equal(t, "1112222", workspace.ID)
	assert.Equal(t, "T8", workspace.Name)
	assert.Equal(t, "1112222", workspace.PrettyName)
	assert.Equal(t, "project", workspace.Category)
	assert.Equal(t, "normal", workspace.Status)
	assert.Equal(t, "描述信息", workspace.Description)
	assert.Equal(t, "张三", workspace.Creator)
	assert.Equal(t, "2018-06-29 15:01:33", workspace.Created)
}
