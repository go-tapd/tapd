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

func TestWorkspaceService_GetCustomWorkCalendar(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workspaces/get_custom_work_calendar", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "2025", r.URL.Query().Get("year"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/workspace/get_custom_work_calendar.json"))
	}))

	calendar, _, err := client.WorkspaceService.GetCustomWorkCalendar(ctx, &GetCustomWorkCalendarRequest{
		WorkspaceID: Ptr(11112222),
		Year:        Ptr("2025"),
	})
	require.NoError(t, err)
	assert.Equal(t, &CustomWorkCalendar{
		Weekdays: []string{"1", "2", "3", "4", "5", "6", "7"},
		Holidays: []string{"2025-01-01"},
		Workdays: []string{"2025-01-02", "2025-01-03", "2025-01-04"},
	}, calendar)
}
