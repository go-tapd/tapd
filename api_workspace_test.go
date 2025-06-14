package tapd

import (
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
		RoleID:           []string{"11111122222001000029"},
		Name:             "张三",
		JoinProjectTime:  nil,
		RealJoinTime:     "2018-07-03",
		Status:           "2",
		Allocation:       "100",
		LeaveProjectTime: nil,
	})
	assert.Contains(t, users, &User{
		User:             "李四",
		RoleID:           []string{"11111122222001000028", "11111122222001000143"},
		Name:             "李四",
		JoinProjectTime:  nil,
		RealJoinTime:     "2018-07-09",
		Status:           "1",
		Allocation:       "100",
		LeaveProjectTime: nil,
	})
}
