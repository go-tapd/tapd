package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestTimesheetService_CreateTimesheet(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/timesheets", r.URL.Path)

		var req struct {
			EntityType  EntityType `json:"entity_type"`
			EntityID    int        `json:"entity_id"`
			Timespent   string     `json:"timespent"`
			Timeremain  string     `json:"timeremain"`
			Spentdate   string     `json:"spentdate"`
			Owner       string     `json:"owner"`
			WorkspaceID int        `json:"workspace_id"`
			Memo        string     `json:"memo"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, EntityTypeStory, req.EntityType)
		assert.Equal(t, 11223344, req.EntityID)
		assert.Equal(t, "2", req.Timespent)
		assert.Equal(t, "0", req.Timeremain)
		assert.Equal(t, "2024-08-22", req.Spentdate)
		assert.Equal(t, "1", req.Owner)
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "1", req.Memo)

		_, _ = w.Write(loadData(t, "internal/testdata/api/timesheet/create_timesheet.json"))
	}))

	timesheet, _, err := client.TimesheetService.CreateTimesheet(ctx, &CreateTimesheetRequest{
		EntityType:  new(EntityTypeStory),
		EntityID:    new(int64(11223344)),
		Timespent:   new("2"),
		Timeremain:  new("0"),
		Spentdate:   new("2024-08-22"),
		Owner:       new("1"),
		WorkspaceID: new(11112222),
		Memo:        new("1"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1134190502001044767", timesheet.ID)
	assert.Equal(t, EntityTypeStory, timesheet.EntityType)
	assert.Equal(t, "1134190502001057318", timesheet.EntityID)
	assert.Equal(t, "2", timesheet.Timespent)
	assert.Equal(t, "2024-08-22", timesheet.Spentdate)
	assert.Equal(t, "1", timesheet.Owner)
	assert.Equal(t, "2024-08-27 08:55:16", timesheet.Created)
	assert.Equal(t, "2024-08-27 08:55:16", timesheet.Modified)
	assert.Equal(t, "11112222", timesheet.WorkspaceID)
	assert.Equal(t, "1", timesheet.Memo)
	assert.Equal(t, "0", timesheet.IsDelete)
}

func TestTimesheetService_GetTimesheets(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/timesheets", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "story", r.URL.Query().Get("entity_type"))
		assert.Equal(t, "111111222222", r.URL.Query().Get("entity_id"))
		assert.Equal(t, "2", r.URL.Query().Get("timespent"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("spentdate"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("modified"))
		assert.Equal(t, "1", r.URL.Query().Get("owner"))
		assert.Equal(t, "1", r.URL.Query().Get("include_parent_story_timesheet"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("created"))
		assert.Equal(t, "1", r.URL.Query().Get("memo"))
		assert.Equal(t, "0", r.URL.Query().Get("is_delete"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,workspace_id", r.URL.Query().Get("fields"))
		assert.Equal(t, "id desc", r.URL.Query().Get("order"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/timesheet/get_timesheets.json"))
	}))

	timesheets, _, err := client.TimesheetService.GetTimesheets(ctx, &GetTimesheetsRequest{
		WorkspaceID:                 new(11112222),
		EntityType:                  new(EntityTypeStory),
		EntityID:                    new(int64(111111222222)),
		Timespent:                   new("2"),
		Spentdate:                   new("2024-08-22"),
		Modified:                    new("2024-08-22"),
		Owner:                       new("1"),
		IncludeParentStoryTimesheet: new(1),
		Created:                     new("2024-08-22"),
		Memo:                        new("1"),
		IsDelete:                    new(0),
		Limit:                       new(10),
		Page:                        new(1),
		Order:                       NewOrder("id", OrderByDesc),
		Fields:                      NewMulti("id", "workspace_id"),
	})
	assert.NoError(t, err)
	spew.Dump(timesheets)
	assert.Len(t, timesheets, 6)
	assert.Equal(t, "111111111", timesheets[0].ID)
	assert.Equal(t, EntityTypeStory, timesheets[0].EntityType)
	assert.Equal(t, "1134190502001057318", timesheets[0].EntityID)
	assert.Equal(t, "2", timesheets[0].Timespent)
	assert.Equal(t, "2024-08-22", timesheets[0].Spentdate)
	assert.Equal(t, "1", timesheets[0].Owner)
	assert.Equal(t, "2024-08-27 08:55:16", timesheets[0].Created)
	assert.Equal(t, "2024-08-27 08:55:16", timesheets[0].Modified)
	assert.Equal(t, "11112222", timesheets[0].WorkspaceID)
	assert.Equal(t, "1", timesheets[0].Memo)
	assert.Equal(t, "0", timesheets[0].IsDelete)
}

func TestTimesheetService_GetTimesheetsCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/timesheets/count", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "story", r.URL.Query().Get("entity_type"))
		assert.Equal(t, "111111222222", r.URL.Query().Get("entity_id"))
		assert.Equal(t, "2", r.URL.Query().Get("timespent"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("spentdate"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("modified"))
		assert.Equal(t, "1", r.URL.Query().Get("owner"))
		assert.Equal(t, "1", r.URL.Query().Get("include_parent_story_timesheet"))
		assert.Equal(t, "2024-08-22", r.URL.Query().Get("created"))
		assert.Equal(t, "1", r.URL.Query().Get("memo"))
		assert.Equal(t, "0", r.URL.Query().Get("is_delete"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/timesheet/get_timesheets_count.json"))
	}))

	count, _, err := client.TimesheetService.GetTimesheetsCount(ctx, &GetTimesheetsCountRequest{
		WorkspaceID:                 new(11112222),
		EntityType:                  new(EntityTypeStory),
		EntityID:                    new(int64(111111222222)),
		Timespent:                   new("2"),
		Spentdate:                   new("2024-08-22"),
		Modified:                    new("2024-08-22"),
		Owner:                       new("1"),
		IncludeParentStoryTimesheet: new(1),
		Created:                     new("2024-08-22"),
		Memo:                        new("1"),
		IsDelete:                    new(0),
	})
	assert.NoError(t, err)
	assert.Equal(t, 6, count)
}

func TestTimesheetService_UpdateTimesheet(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/timesheets", r.URL.Path)

		var req struct {
			ID          int    `json:"id"`
			Timespent   string `json:"timespent,omitempty"`
			Timeremain  string `json:"timeremain,omitempty"`
			WorkspaceID int    `json:"workspace_id,omitempty"`
			Memo        string `json:"memo,omitempty"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 1134190502001044767, req.ID)
		assert.Equal(t, "2", req.Timespent)
		assert.Equal(t, "0", req.Timeremain)
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "1", req.Memo)

		_, _ = w.Write(loadData(t, "internal/testdata/api/timesheet/update_timesheet.json"))
	}))

	timesheet, _, err := client.TimesheetService.UpdateTimesheet(ctx, &UpdateTimesheetRequest{
		ID:          new(int64(1134190502001044767)),
		Timespent:   new("2"),
		Timeremain:  new("0"),
		WorkspaceID: new(11112222),
		Memo:        new("1"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1134190502001044767", timesheet.ID)
	assert.Equal(t, EntityTypeStory, timesheet.EntityType)
	assert.Equal(t, "1134190502001057318", timesheet.EntityID)
	assert.Equal(t, "2", timesheet.Timespent)
	assert.Equal(t, "2024-08-22", timesheet.Spentdate)
	assert.Equal(t, "1", timesheet.Owner)
	assert.Equal(t, "2024-08-27 08:55:16", timesheet.Created)
	assert.Equal(t, "2024-08-27 08:55:16", timesheet.Modified)
	assert.Equal(t, "11112222", timesheet.WorkspaceID)
	assert.Equal(t, "1", timesheet.Memo)
	assert.Equal(t, "0", timesheet.IsDelete)
}

func TestTimesheetService_DeleteTimesheets(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/timesheets/delete_timesheets", r.URL.Path)

		var req struct {
			EntityType  EntityType `json:"entity_type"`
			EntityID    int64      `json:"entity_id"`
			WorkspaceID int        `json:"workspace_id"`
			CostIDs     []int64    `json:"cost_ids"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, EntityTypeStory, req.EntityType)
		assert.Equal(t, int64(1134190502001057318), req.EntityID)
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, []int64{1134190502001044767, 1134190502001044768}, req.CostIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/timesheet/delete_timesheets.json"))
	}))

	response, _, err := client.TimesheetService.DeleteTimesheets(ctx, &DeleteTimesheetsRequest{
		EntityType:  new(EntityTypeStory),
		EntityID:    new(int64(1134190502001057318)),
		WorkspaceID: new(11112222),
		CostIDs:     new([]int64{1134190502001044767, 1134190502001044768}),
	})
	assert.NoError(t, err)
	assert.Equal(t, "删除成功", response.Msg)
	assert.Equal(t, []string{"1134190502001044767"}, response.Data.Success.CostIDs)
	assert.Equal(t, "删除成功", response.Data.Success.Msg)
	assert.Len(t, response.Data.Failed, 1)
	assert.Equal(t, []string{"1134190502001044768"}, response.Data.Failed[0].CostIDs)
	assert.Equal(t, "工时花费不存在", response.Data.Failed[0].Msg)
}
