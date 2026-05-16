package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoryService_CreateStoryCategory(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/story_categories", r.URL.Path)

		var req CreateStoryCategoryRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "产品需求", *req.Name)
		assert.Equal(t, "产品需求描述", *req.Description)
		assert.Equal(t, int64(0), *req.ParentID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/create_story_category.json"))
	}))

	category, _, err := client.StoryService.CreateStoryCategory(ctx, &CreateStoryCategoryRequest{
		WorkspaceID: Ptr(11112222),
		Name:        Ptr("产品需求"),
		Description: Ptr("产品需求描述"),
		ParentID:    Ptr[int64](0),
	})
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "1111112222001000056", category.ID)
	assert.Equal(t, "11112222", category.WorkspaceID)
	assert.Equal(t, "产品需求", category.Name)
	assert.Equal(t, "产品需求描述", category.Description)
	assert.Equal(t, "0", category.ParentID)
	assert.Equal(t, "xinweihe", category.Creator)
}

func TestStoryService_CopyStory(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/copy_story", r.URL.Path)

		var req struct {
			WorkspaceID       int         `json:"workspace_id"`
			SrcStoryID        int64       `json:"src_story_id"`
			DstWorkspaceID    int         `json:"dst_workspace_id"`
			SyncFields        string      `json:"sync_fields"`
			DstWorkitemTypeID int64       `json:"dst_workitem_type_id"`
			NewCreator        string      `json:"new_creator"`
			NewStatus         StoryStatus `json:"new_status"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, int64(1111112222001000103), req.SrcStoryID)
		assert.Equal(t, 33334444, req.DstWorkspaceID)
		assert.Equal(t, "name,description,owner", req.SyncFields)
		assert.Equal(t, int64(10001), req.DstWorkitemTypeID)
		assert.Equal(t, "xinweihe", req.NewCreator)
		assert.Equal(t, StoryStatusPlanning, req.NewStatus)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/copy_story.json"))
	}))

	story, _, err := client.StoryService.CopyStory(ctx, &CopyStoryRequest{
		WorkspaceID:       Ptr(11112222),
		SrcStoryID:        Ptr[int64](1111112222001000103),
		DstWorkspaceID:    Ptr(33334444),
		SyncFields:        NewMulti("name", "description", "owner"),
		DstWorkitemTypeID: Ptr[int64](10001),
		NewCreator:        Ptr("xinweihe"),
		NewStatus:         Ptr(StoryStatusPlanning),
	})
	assert.NoError(t, err)
	assert.NotNil(t, story)
	assert.Equal(t, "333344440010000001", story.ID)
	assert.Equal(t, "复制的需求", story.Name)
	assert.Equal(t, "33334444", story.WorkspaceID)
	assert.Equal(t, StoryStatusPlanning, story.Status)
	assert.Equal(t, "xinweihe", story.Creator)
}

func TestStoryService_GetStoryLinkStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_link_stories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_link_stories.json"))
	}))

	relations, _, err := client.StoryService.GetStoryLinkStories(ctx, &GetStoryLinkStoriesRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001000103),
	})
	assert.NoError(t, err)
	assert.Len(t, relations, 2)
	assert.Equal(t, "derivation", relations[0].Type)
	assert.Equal(t, "1111112222001000104", relations[0].ID)
	assert.Equal(t, "1111112222001000103", relations[0].StoryID)
	assert.Equal(t, "11112222", relations[0].WorkspaceID)
	assert.Equal(t, "target", relations[0].ActAs)
	assert.Equal(t, 11112222, relations[0].LinkedWorkspaceID)
	assert.Equal(t, "copy", relations[1].Type)
}

func TestStoryService_GetSecretStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/secret_stories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_secret_stories.json"))
	}))

	stories, _, err := client.StoryService.GetSecretStories(ctx, &GetSecretStoriesRequest{
		WorkspaceID: Ptr(11112222),
		Limit:       Ptr(10),
		Page:        Ptr(2),
		Order:       NewOrder("created", OrderByDesc),
	})
	assert.NoError(t, err)
	assert.Equal(t, []string{
		"1111112222001000103",
		"1111112222001000104",
		"1111112222001000105",
	}, stories)
}

func TestStoryService_GetSecretStoriesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/secret_stories/count", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_secret_stories_count.json"))
	}))

	count, _, err := client.StoryService.GetSecretStoriesCount(ctx, &GetSecretStoriesCountRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}

func TestStoryService_GetStoryCategories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/story_categories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111,1111111111112", r.URL.Query().Get("id"))
		assert.Equal(t, "test name", r.URL.Query().Get("name"))
		assert.Equal(t, "test description", r.URL.Query().Get("description"))
		assert.Equal(t, "1111111111111", r.URL.Query().Get("parent_id"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("modified"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id asc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_categories.json"))
	}))

	categories, _, err := client.StoryService.GetStoryCategories(ctx, &GetStoryCategoriesRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti[int64](1111111111111, 1111111111112),
		Name:        Ptr("test name"),
		Description: Ptr("test description"),
		ParentID:    Ptr(1111111111111),
		Created:     Ptr("2021-01-01"),
		Modified:    Ptr("2021-01-02"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Order:       NewOrder("id", OrderByAsc),
		Fields:      NewMulti("id", "name"),
	})
	assert.NoError(t, err)
	assert.True(t, len(categories) > 0)
	assert.Equal(t, "1111112222001000056", categories[0].ID)
	assert.Equal(t, "11112222", categories[0].WorkspaceID)
	assert.Equal(t, "产品需求", categories[0].Name)
	assert.Equal(t, "产品需求", categories[0].Description)
	assert.Equal(t, "0", categories[0].ParentID)
	assert.Equal(t, "2024-06-20 11:38:37", categories[0].Modified)
	assert.Equal(t, "2018-06-29 15:01:38", categories[0].Created)
	assert.Equal(t, "", categories[0].Creator)
	assert.Equal(t, "张三", categories[0].Modifier)
}

func TestStoryService_UpdateStoryCategory(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/story_categories", r.URL.Path)

		var req UpdateStoryCategoryRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, int64(1111112222001000056), *req.ID)
		assert.Equal(t, "产品需求", *req.Name)
		assert.Equal(t, "产品需求描述", *req.Description)
		assert.Equal(t, int64(0), *req.ParentID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/update_story_category.json"))
	}))

	category, _, err := client.StoryService.UpdateStoryCategory(ctx, &UpdateStoryCategoryRequest{
		WorkspaceID: Ptr(11112222),
		ID:          Ptr[int64](1111112222001000056),
		Name:        Ptr("产品需求"),
		Description: Ptr("产品需求描述"),
		ParentID:    Ptr[int64](0),
	})
	assert.NoError(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, "1111112222001000056", category.ID)
	assert.Equal(t, "11112222", category.WorkspaceID)
	assert.Equal(t, "产品需求", category.Name)
	assert.Nil(t, category.Description)
	assert.Equal(t, "0", category.ParentID)
	assert.Equal(t, "xinweihe", category.Modifier)
}

func TestStoryService_GetStoryCategoriesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/story_categories/count", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111,1111111111112", r.URL.Query().Get("id"))
		assert.Equal(t, "test name", r.URL.Query().Get("name"))
		assert.Equal(t, "test description", r.URL.Query().Get("description"))
		assert.Equal(t, "1111111111111", r.URL.Query().Get("parent_id"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("modified"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_categories_count.json"))
	}))

	count, _, err := client.StoryService.GetStoryCategoriesCount(ctx, &GetStoryCategoriesCountRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti[int64](1111111111111, 1111111111112),
		Name:        Ptr("test name"),
		Description: Ptr("test description"),
		ParentID:    Ptr(1111111111111),
		Created:     Ptr("2021-01-01"),
		Modified:    Ptr("2021-01-02"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 30, count)
}

func TestStoryService_GetStoriesCountByCategories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/count_by_categories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103,1111112222001000108", r.URL.Query().Get("category_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_stories_count_by_categories.json"))
	}))

	counts, _, err := client.StoryService.GetStoriesCountByCategories(ctx, &GetStoriesCountByCategoriesRequest{
		WorkspaceID: Ptr(11112222),
		CategoryID:  NewMulti[int64](1111112222001000103, 1111112222001000108),
	})
	assert.NoError(t, err)
	assert.True(t, len(counts) > 0)
	assert.Contains(t, counts, &StoriesCountByCategory{
		CategoryID: "1111112222001000103",
		Count:      85,
	})
}

func TestStoryService_GetStoryChanges(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/story_changes", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103,1111112222001000108", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_changes.json"))
	}))

	storyChanges, _, err := client.StoryService.GetStoryChanges(ctx, &GetStoryChangesRequest{
		StoryID:     NewMulti[int64](1111112222001000103, 1111112222001000108),
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(storyChanges) > 0)
	assert.Equal(t, "1111112222001275457", storyChanges[0].ID)
	assert.Equal(t, "11112222", storyChanges[0].WorkspaceID)
	assert.Equal(t, "1", storyChanges[0].AppID)
	assert.Equal(t, "0", storyChanges[0].WorkitemTypeID)
	assert.Equal(t, "TAPD", storyChanges[0].Creator)
	assert.Equal(t, "2022-06-10 10:04:12", storyChanges[0].Created)
	assert.Equal(t, "create_story", storyChanges[0].ChangeSummary)
	assert.Nil(t, storyChanges[0].Comment)
	assert.Equal(t, "Story", storyChanges[0].EntityType)
	assert.Equal(t, StoreChangeTypeCreateStory, storyChanges[0].ChangeType)
	assert.Equal(t, "需求创建", storyChanges[0].ChangeTypeText)
	assert.Equal(t, "2024-09-07 23:38:36", storyChanges[0].Updated)
	assert.Equal(t, "1111112222001032850", storyChanges[0].StoryID)
	assert.True(t, len(storyChanges[0].FieldChanges) > 0)
	assert.Equal(t, "name", storyChanges[0].FieldChanges[0].Field)
	assert.Equal(t, "", storyChanges[0].FieldChanges[0].ValueBefore)
	assert.Equal(t, "需求3", storyChanges[0].FieldChanges[0].ValueAfter)
	assert.Equal(t, "--", storyChanges[0].FieldChanges[0].ValueBeforeParsed)
	assert.Equal(t, "需求3", storyChanges[0].FieldChanges[0].ValueAfterParsed)
	assert.Equal(t, "标题", storyChanges[0].FieldChanges[0].FieldLabel)
}

func TestStoryService_GetStoryChangesCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/story_changes/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103,1111112222001000108", r.URL.Query().Get("story_id"))
		assert.Equal(t, "TAPD", r.URL.Query().Get("creator"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_changes_count.json"))
	}))

	count, _, err := client.StoryService.GetStoryChangesCount(ctx, &GetStoryChangesCountRequest{
		StoryID:     NewMulti[int64](1111112222001000103, 1111112222001000108),
		WorkspaceID: Ptr(11112222),
		Creator:     Ptr("TAPD"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 23, count)
}

func TestStoryService_GetStoryCustomFieldsSettings(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/custom_fields_settings", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_custom_fields_settings.json"))
	}))

	settings, _, err := client.StoryService.GetStoryCustomFieldsSettings(ctx, &GetStoryCustomFieldsSettingsRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(settings) > 0)
	assert.Equal(t, "1111112222001000155", settings[0].ID)
	assert.Equal(t, "11112222", settings[0].WorkspaceID)
	assert.Equal(t, "1", settings[0].AppID)
	assert.Equal(t, "story", settings[0].EntryType)
	assert.Equal(t, "custom_field_100", settings[0].CustomField)
	assert.Equal(t, "user_chooser", settings[0].Type)
	assert.Equal(t, "test name", settings[0].Name)
	assert.Nil(t, settings[0].Options)
	assert.Nil(t, settings[0].ExtraConfig)
	assert.Equal(t, "1", settings[0].Enabled)
	assert.Equal(t, "0", settings[0].Freeze)
	assert.Nil(t, settings[0].Sort)
	assert.Nil(t, settings[0].Memo)
	assert.Equal(t, "", settings[0].OpenExtensionID)
	assert.Equal(t, 0, settings[0].IsOut)
	assert.Equal(t, 0, settings[0].IsUninstall)
	assert.Equal(t, "", settings[0].AppName)
}

func TestStoryService_GetStoryTestCaseRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_story_tcase", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_tcase.json"))
	}))

	relations, _, err := client.StoryService.GetStoryTestCaseRelation(ctx, &GetStoryTestCaseRelationRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](33334444),
	})
	assert.NoError(t, err)
	assert.True(t, len(relations) > 0)
	assert.Equal(t, "111111112222001466738", relations[0].ID)
	assert.Equal(t, "1111112222", relations[0].WorkspaceID)
	assert.Equal(t, "0", relations[0].TestPlanID)
	assert.Equal(t, "111111112222001152461", relations[0].StoryID)
	assert.Equal(t, "111111112222001175632", relations[0].TcaseID)
	assert.Equal(t, "0", relations[0].Sort)
	assert.Equal(t, "张三", relations[0].Creator)
	assert.Equal(t, "0000-00-00 00:00:00", relations[0].Created)
}

func TestStoryService_GetStoryTimeRelations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_time_relative_stories", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_time_relations.json"))
	}))

	relations, _, err := client.StoryService.GetStoryTimeRelations(ctx, &GetStoryTimeRelationsRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001000103),
	})
	assert.NoError(t, err)
	assert.Len(t, relations, 2)
	assert.Equal(t, "1210104801000007813", relations[0].ID)
	assert.Equal(t, "11112222", relations[0].WorkspaceID)
	assert.Equal(t, "story", relations[0].WorkitemType)
	assert.Equal(t, "1111112222001000102", relations[0].WorkitemID)
	assert.Equal(t, "begin", relations[0].SrcField)
	assert.Equal(t, "11112222", relations[0].DstWorkspaceID)
	assert.Equal(t, "story", relations[0].DstWorkitemType)
	assert.Equal(t, "1111112222001000103", relations[0].DstWorkitemID)
	assert.Equal(t, "due", relations[0].DstField)
	assert.Equal(t, "after", relations[0].RelationType)
	assert.Equal(t, "1210104801000007815", relations[1].ID)
}

func TestStoryService_SaveStoryTimeRelations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/save_time_relations", r.URL.Path)

		var req SaveStoryTimeRelationsRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "testuser", *req.CurrentUser)
		assert.Len(t, req.Relations, 1)
		assert.Equal(t, int64(1111112222001000102), *req.Relations[0].WorkitemID)
		assert.Equal(t, int64(1111112222001000103), *req.Relations[0].DstWorkitemID)
		assert.Equal(t, "begin", *req.Relations[0].SrcField)
		assert.Equal(t, "due", *req.Relations[0].DstField)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/save_story_time_relations.json"))
	}))

	result, _, err := client.StoryService.SaveStoryTimeRelations(ctx, &SaveStoryTimeRelationsRequest{
		WorkspaceID: Ptr(11112222),
		CurrentUser: Ptr("testuser"),
		Relations: []*SaveStoryTimeRelation{
			{
				WorkitemID:    Ptr[int64](1111112222001000102),
				DstWorkitemID: Ptr[int64](1111112222001000103),
				SrcField:      Ptr("begin"),
				DstField:      Ptr("due"),
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Result)
}

func TestStoryService_DeleteStoryTimeRelations(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/delete_time_relations", r.URL.Path)

		var req DeleteStoryTimeRelationsRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "testuser", *req.CurrentUser)
		assert.Len(t, req.Relations, 1)
		assert.Equal(t, int64(1111112222001000102), *req.Relations[0].WorkitemID)
		assert.Equal(t, int64(1111112222001000103), *req.Relations[0].DstWorkitemID)
		assert.Equal(t, []int64{1210104801000007813}, *req.RelationIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/delete_story_time_relations.json"))
	}))

	result, _, err := client.StoryService.DeleteStoryTimeRelations(ctx, &DeleteStoryTimeRelationsRequest{
		WorkspaceID: Ptr(11112222),
		CurrentUser: Ptr("testuser"),
		Relations: []*DeleteStoryTimeRelation{
			{
				WorkitemID:    Ptr[int64](1111112222001000102),
				DstWorkitemID: Ptr[int64](1111112222001000103),
			},
		},
		RelationIDs: Ptr([]int64{1210104801000007813}),
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Num)
}

func TestStoryService_GetStorySecretInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_secret_info", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_secret_info.json"))
	}))

	info, _, err := client.StoryService.GetStorySecretInfo(ctx, &GetStorySecretInfoRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001000103),
	})
	assert.NoError(t, err)
	assert.NotNil(t, info)
	assert.Equal(t, "xinweihe", info.Creator)
	assert.Equal(t, "xinweihe;1000000000000000002", info.AllowList)
	assert.Equal(t, "1111112222001000103", info.SecretRootID)
	assert.Equal(t, "true", info.AddParticipantFields)
	assert.Equal(t, "secret", info.SecretScope)
}

func TestStoryService_BatchUpdateStorySecretInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/batch_update_secret_info", r.URL.Path)

		var req struct {
			WorkspaceID          int    `json:"workspace_id"`
			StoryIDList          string `json:"story_id_list"`
			SecretScope          string `json:"secret_scope"`
			AllowList            string `json:"allow_list"`
			AddParticipantFields string `json:"add_participant_fields"`
			OperationType        int    `json:"operation_type"`
			CurrentUser          string `json:"current_user"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "1111112222001000103|1111112222001000104", req.StoryIDList)
		assert.Equal(t, "secret", req.SecretScope)
		assert.Equal(t, "xinweihe;1000000000000000002", req.AllowList)
		assert.Equal(t, "false", req.AddParticipantFields)
		assert.Equal(t, 0, req.OperationType)
		assert.Equal(t, "xinweihe", req.CurrentUser)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/batch_update_story_secret_info.json"))
	}))

	result, _, err := client.StoryService.BatchUpdateStorySecretInfo(ctx, &BatchUpdateStorySecretInfoRequest{
		WorkspaceID:          Ptr(11112222),
		StoryIDList:          NewEnum[int64](1111112222001000103, 1111112222001000104),
		SecretScope:          Ptr("secret"),
		AllowList:            Ptr("xinweihe;1000000000000000002"),
		AddParticipantFields: Ptr("false"),
		OperationType:        Ptr(0),
		CurrentUser:          Ptr("xinweihe"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "succeed", result.Code)
	assert.Equal(t, "需求可访问人员设置成功", result.Msg)
}

func TestStoryService_GetStoryWorkitemTypes(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/workitem_types", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103,1111112222001000104", r.URL.Query().Get("id"))
		assert.Equal(t, "需求", r.URL.Query().Get("name"))
		assert.Equal(t, "story", r.URL.Query().Get("entity_type"))
		assert.Equal(t, "custom_story", r.URL.Query().Get("english_name"))
		assert.Equal(t, "1210104801000000001", r.URL.Query().Get("workflow_id"))
		assert.Equal(t, "1", r.URL.Query().Get("status"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "created desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,name,workflow_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_workitem_types.json"))
	}))

	workitemTypes, _, err := client.StoryService.GetStoryWorkitemTypes(ctx, &GetStoryWorkitemTypesRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti[int64](1111112222001000103, 1111112222001000104),
		Name:        Ptr("需求"),
		EntityType:  Ptr("story"),
		EnglishName: Ptr("custom_story"),
		WorkflowID:  Ptr[int64](1210104801000000001),
		Status:      Ptr(1),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Order:       NewOrder("created", OrderByDesc),
		Fields:      NewMulti("id", "name", "workflow_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, workitemTypes, 2)
	assert.Equal(t, "1111112222001000103", workitemTypes[0].ID)
	assert.Equal(t, "11112222", workitemTypes[0].WorkspaceID)
	assert.Equal(t, "用户故事", workitemTypes[0].Name)
	assert.Equal(t, "custom_story", workitemTypes[0].EnglishName)
	assert.Equal(t, "1210104801000000001", workitemTypes[0].WorkflowID)
	assert.Equal(t, "任务", workitemTypes[1].Name)
}

func TestStoryService_BatchUpdateStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/batch_update_story", r.URL.Path)

		var req BatchUpdateStoriesRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Len(t, req.Workitems, 2)
		assert.Equal(t, int64(1111112222001000103), *req.Workitems[0].ID)
		assert.Nil(t, req.Workitems[0].WorkspaceID)
		assert.Equal(t, "first story", *req.Workitems[0].Name)
		assert.Equal(t, "planning", *req.Workitems[0].Status)
		assert.Equal(t, int64(1111112222001000104), *req.Workitems[1].ID)
		assert.Equal(t, "second story", *req.Workitems[1].Name)
		assert.Equal(t, "owner", *req.Workitems[1].Owner)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/batch_update_stories.json"))
	}))

	result, _, err := client.StoryService.BatchUpdateStories(ctx, &BatchUpdateStoriesRequest{
		WorkspaceID: Ptr(11112222),
		Workitems: []*UpdateStoryRequest{
			{
				ID:     Ptr[int64](1111112222001000103),
				Name:   Ptr("first story"),
				Status: Ptr("planning"),
			},
			{
				ID:    Ptr[int64](1111112222001000104),
				Name:  Ptr("second story"),
				Owner: Ptr("owner"),
			},
		},
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "batch update success", result.Msg)
}

func TestStoryService_UpdateStoryWorkitemType(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/change_workitem_type", r.URL.Path)

		var req UpdateStoryWorkitemTypeRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, int64(1111112222001000103), *req.StoryID)
		assert.Equal(t, int64(1111112222001000104), *req.WorkitemTypeID)
		assert.Equal(t, 11112222, *req.WorkspaceID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/update_story_workitem_type.json"))
	}))

	story, _, err := client.StoryService.UpdateStoryWorkitemType(ctx, &UpdateStoryWorkitemTypeRequest{
		StoryID:        Ptr[int64](1111112222001000103),
		WorkitemTypeID: Ptr[int64](1111112222001000104),
		WorkspaceID:    Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.NotNil(t, story)
	assert.Equal(t, "1111112222001000103", story.ID)
	assert.Equal(t, "1111112222001000104", story.WorkitemTypeID)
	assert.Equal(t, "用户故事", story.Name)
	assert.Equal(t, StoryStatusPlanning, story.Status)
}

func TestStoryService_GetStorySteps(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_story_step_list", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111112222001000103", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_steps.json"))
	}))

	steps, _, err := client.StoryService.GetStorySteps(ctx, &GetStoryStepsRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001000103),
	})
	assert.NoError(t, err)
	assert.Len(t, steps, 2)
	assert.Equal(t, "1210104801000007813", steps[0].ID)
	assert.Equal(t, "11112222", steps[0].WorkspaceID)
	assert.Equal(t, "story", steps[0].EntityType)
	assert.Equal(t, "1111112222001000103", steps[0].WorkitemID)
	assert.Equal(t, "step_1", steps[0].Step)
	assert.Equal(t, "0", steps[0].Status)
	assert.Nil(t, steps[0].Begin)
	assert.Nil(t, steps[0].Due)
	assert.Equal(t, "3", steps[0].Effort)
	assert.Equal(t, "2026-01-04 09:38:23", steps[0].CompleteTime)
	assert.Equal(t, "xinweihe", steps[1].Owner)
}

func TestStoryService_GetStoryFieldsLabel(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_fields_lable", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_fields_lable.json"))
	}))

	labels, _, err := client.StoryService.GetStoryFieldsLabel(ctx, &GetStoryFieldsLabelRequest{
		WorkspaceID: Ptr(11112222),
	})
	assert.NoError(t, err)
	assert.True(t, len(labels) > 0)

	// Verify some of the field labels
	labelMap := make(map[string]string)
	for _, label := range labels {
		labelMap[label.EN] = label.CN
	}

	assert.Equal(t, "ID", labelMap["id"])
	assert.Equal(t, "标题", labelMap["name"])
	assert.Equal(t, "详细描述", labelMap["description"])
	assert.Equal(t, "项目ID", labelMap["workspace_id"])
	assert.Equal(t, "创建人", labelMap["creator"])
}

func TestStoryService_GetStoryRelatedBugs(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_related_bugs", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444,55556666", r.URL.Query().Get("story_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_related_bugs.json"))
	}))

	relatedBugs, _, err := client.StoryService.GetStoryRelatedBugs(ctx, &GetStoryRelatedBugsRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     NewMulti[int64](33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.True(t, len(relatedBugs) > 0)
	assert.Equal(t, 11112222, relatedBugs[0].WorkspaceID)
	assert.Equal(t, "1111112222001063941", relatedBugs[0].StoryID)
	assert.Equal(t, "1111112222001035927", relatedBugs[0].BugID)
}

func TestStoryService_RemoveStoryBugRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/remove_story_bug_raletions", r.URL.Path)

		var req RemoveStoryBugRelationRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, int64(1111112222001063941), *req.StoryID)
		assert.Equal(t, int64(1111112222001035927), *req.BugID)
		assert.Equal(t, "xinweihe", *req.CurrentUser)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/remove_story_bug_relation.json"))
	}))

	result, _, err := client.StoryService.RemoveStoryBugRelation(ctx, &RemoveStoryBugRelationRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001063941),
		BugID:       Ptr[int64](1111112222001035927),
		CurrentUser: Ptr("xinweihe"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
}

func TestStoryService_UpdateStoryParent(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/update_story_parent", r.URL.Path)

		var req UpdateStoryParentRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, int64(1111112222001063941), *req.StoryID)
		assert.Equal(t, int64(1111112222001060000), *req.ParentID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/update_story_parent.json"))
	}))

	story, _, err := client.StoryService.UpdateStoryParent(ctx, &UpdateStoryParentRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001063941),
		ParentID:    Ptr[int64](1111112222001060000),
	})
	assert.NoError(t, err)
	assert.NotNil(t, story)
	assert.Equal(t, "1111112222001063941", story.ID)
	assert.Equal(t, "1111112222001060000", story.ParentID)
	assert.Equal(t, "11112222", story.WorkspaceID)
	assert.Equal(t, "1111112222001060000:1111112222001063941:", story.Path)
}

func TestStoryService_CreateStoryBugRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/relations", r.URL.Path)

		var req CreateStoryBugRelationRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, *req.WorkspaceID)
		assert.Equal(t, "story", *req.SourceType)
		assert.Equal(t, int64(1111112222001063941), *req.SourceID)
		assert.Equal(t, "bug", *req.TargetType)
		assert.Equal(t, int64(1111112222001035927), *req.TargetID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/create_story_bug_relation.json"))
	}))

	relation, _, err := client.StoryService.CreateStoryBugRelation(ctx, &CreateStoryBugRelationRequest{
		WorkspaceID: Ptr(11112222),
		SourceType:  Ptr("story"),
		SourceID:    Ptr[int64](1111112222001063941),
		TargetType:  Ptr("bug"),
		TargetID:    Ptr[int64](1111112222001035927),
	})
	assert.NoError(t, err)
	assert.NotNil(t, relation)
	assert.Equal(t, "22265547", relation.ID)
	assert.Equal(t, "11112222", relation.WorkspaceID)
	assert.Equal(t, "story", relation.SourceType)
	assert.Equal(t, "1111112222001063941", relation.SourceID)
	assert.Equal(t, "bug", relation.TargetType)
	assert.Equal(t, "1111112222001035927", relation.TargetID)
}

func TestStoryService_CreateStoryTestCaseRelation(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/add_story_tcase", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			StoryID     int64  `json:"story_id"`
			TestCaseID  string `json:"tcase_id"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, int64(1111112222001063941), req.StoryID)
		assert.Equal(t, "1111112222001077291,1111112222001077292", req.TestCaseID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/create_story_test_case_relation.json"))
	}))

	result, _, err := client.StoryService.CreateStoryTestCaseRelation(ctx, &CreateStoryTestCaseRelationRequest{
		WorkspaceID: Ptr(11112222),
		StoryID:     Ptr[int64](1111112222001063941),
		TestCaseID:  NewMulti[int64](1111112222001077291, 1111112222001077292),
	})
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, []string{"1111112222001077291", "1111112222001077292"}, result.SuccessID)
}

func TestStoryService_GetStoryTemplates(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/template_list", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1", r.URL.Query().Get("workitem_type_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_templates.json"))
	}))

	templates, _, err := client.StoryService.GetStoryTemplates(ctx, &GetStoryTemplatesRequest{
		WorkspaceID:    Ptr(11112222),
		WorkitemTypeID: Ptr(1),
	})
	assert.NoError(t, err)
	assert.True(t, len(templates) > 0)
	assert.Equal(t, "1111112222001000015", templates[0].ID)
	assert.Equal(t, "System default template", templates[0].Name)
	assert.Equal(t, "Auto created by the system", templates[0].Description)
	assert.Equal(t, "0", templates[0].Sort)
	assert.Equal(t, "0", templates[0].Default)
	assert.Equal(t, "SYSTEM", templates[0].Creator)
	assert.Equal(t, "1", templates[0].EditorType)
}

func TestStoryService_GetStoryTemplateFields(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_default_story_template", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111", r.URL.Query().Get("template_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_story_template_fields.json"))
	}))

	fields, _, err := client.StoryService.GetStoryTemplateFields(ctx, &GetStoryTemplateFieldsRequest{
		WorkspaceID: Ptr(11112222),
		TemplateID:  Ptr(int64(1111111111111)),
	})
	assert.NoError(t, err)
	assert.True(t, len(fields) > 0)
	assert.Equal(t, "1111112222001000113", fields[0].ID)
	assert.Equal(t, "11112222", fields[0].WorkspaceID)
	assert.Equal(t, "story", fields[0].Type)
	assert.Equal(t, "1111112222001000015", fields[0].TemplateID)
	assert.Equal(t, "name", fields[0].Field)
	assert.Equal(t, "", fields[0].Value)
	assert.Equal(t, "1", fields[0].Required)
	assert.Equal(t, "0", fields[0].Sort)
	assert.Equal(t, "", fields[0].LinkageRules)
}

func TestStoryService_GetRemovedStories(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/stories/get_removed_stories", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "1111111111111,1111111111112", r.URL.Query().Get("id"))
		assert.Equal(t, "creator", r.URL.Query().Get("creator"))
		assert.Equal(t, "1", r.URL.Query().Get("is_archived"))
		assert.Equal(t, "2021-01-01", r.URL.Query().Get("created"))
		assert.Equal(t, "2021-01-02", r.URL.Query().Get("deleted"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_removed_stories.json"))
	}))

	stories, _, err := client.StoryService.GetRemovedStories(ctx, &GetRemovedStoriesRequest{
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
	assert.True(t, len(stories) > 0)
	assert.Equal(t, "1111112222001069791", stories[0].ID)
	assert.Equal(t, "測試測試", stories[0].Name)
	assert.Equal(t, "张三", stories[0].Creator)
	assert.Equal(t, "2024-08-20 11:22:49", stories[0].Created)
	assert.Equal(t, "张三", stories[0].OperationUser)
	assert.Equal(t, "2024-08-20 11:28:23", stories[0].Deleted)
	assert.Equal(t, "0", stories[0].IsArchived)
}

func TestStoryService_GetConvertStoryIDsToQueryToken(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/stories/ids_to_query_token", r.URL.Path)

		var req struct {
			WorkspaceID int    `json:"workspace_id"`
			StoryIDs    string `json:"ids"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "33334444,55556666", req.StoryIDs)

		_, _ = w.Write(loadData(t, "internal/testdata/api/story/get_convert_story_ids_to_query_token.json"))
	}))

	response, _, err := client.StoryService.GetConvertStoryIDsToQueryToken(ctx, &GetConvertStoryIDsToQueryTokenRequest{
		WorkspaceID: Ptr(11112222),
		StoryIDs:    NewMulti[int64](33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.Equal(t, "11111111111", response.QueryToken)
	assert.Contains(t, response.Href, "11111111111")
}
