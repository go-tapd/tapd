package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		ID:          NewMulti(1111111111111, 1111111111112),
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
		CategoryID:  NewMulti(1111112222001000103, 1111112222001000108),
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
		StoryID:     NewMulti(33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.True(t, len(relatedBugs) > 0)
	assert.Equal(t, 11112222, relatedBugs[0].WorkspaceID)
	assert.Equal(t, "1111112222001063941", relatedBugs[0].StoryID)
	assert.Equal(t, "1111112222001035927", relatedBugs[0].BugID)
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
		TemplateID:  Ptr(1111111111111),
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
		StoryIDs:    NewMulti(33334444, 55556666),
	})
	assert.NoError(t, err)
	assert.Equal(t, "11111111111", response.QueryToken)
	assert.Contains(t, response.Href, "11111111111")
}
