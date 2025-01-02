package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommentService_CommentEntryType(t *testing.T) {
	assert.Equal(t, "bug", CommentEntryTypeBug.String())
	assert.Equal(t, "bug_remark", CommentEntryTypeBugRemark.String())
	assert.Equal(t, "stories", CommentEntryTypeStories.String())
	assert.Equal(t, "tasks", CommentEntryTypeTasks.String())
	assert.Equal(t, "wiki", CommentEntryTypeWiki.String())
	assert.Equal(t, "mini_items", CommentEntryTypeMiniItems.String())
}

func TestCommentService_CreateComment(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/comments", r.URL.Path)

		var req struct {
			Title       string           `json:"title"`
			Description string           `json:"description"`
			Author      string           `json:"author"`
			EntryType   CommentEntryType `json:"entry_type"`
			EntryID     int              `json:"entry_id"`
			ReplyID     int              `json:"reply_id"`
			RootID      int              `json:"root_id"`
			WorkspaceID int              `json:"workspace_id"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, "title", req.Title)
		assert.Equal(t, "description", req.Description)
		assert.Equal(t, "author", req.Author)
		assert.Equal(t, CommentEntryTypeStories, req.EntryType)
		assert.Equal(t, 123, req.EntryID)
		assert.Equal(t, 0, req.ReplyID)
		assert.Equal(t, 111, req.RootID)
		assert.Equal(t, 111, req.WorkspaceID)

		_, _ = w.Write(loadData(t, "internal/testdata/api/comment/create_comment.json"))
	}))

	comment, _, err := client.CommentService.CreateComment(ctx, &CreateCommentRequest{
		Title:       Ptr("title"),
		Description: Ptr("description"),
		Author:      Ptr("author"),
		EntryType:   Ptr(CommentEntryTypeStories),
		EntryID:     Ptr(123),
		ReplyID:     Ptr(0),
		RootID:      Ptr(111),
		WorkspaceID: Ptr(111),
	})
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, "1111112222001033109", comment.ID)
	assert.Equal(t, "test title", comment.Title)
	assert.Equal(t, "test description", comment.Description)
	assert.Equal(t, "go-tapd", comment.Author)
	assert.Equal(t, CommentEntryTypeStories, comment.EntryType)
	assert.Equal(t, "1111112222001071295", comment.EntryID)
	assert.Equal(t, "0", comment.ReplyID)
	assert.Equal(t, "1111112222001033105", comment.RootID)
	assert.Equal(t, "2024-08-28 23:03:06", comment.Created)
	assert.Equal(t, "2024-08-28 23:03:06", comment.Modified)
	assert.Equal(t, "11112222", comment.WorkspaceID)
}

func TestCommentService_GetComments(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/comments", r.URL.Path)
		assert.Equal(t, "111,222", r.URL.Query().Get("id"))
		assert.Equal(t, "title", r.URL.Query().Get("title"))
		assert.Equal(t, "description", r.URL.Query().Get("description"))
		assert.Equal(t, "author", r.URL.Query().Get("author"))
		assert.Equal(t, CommentEntryTypeStories.String(), r.URL.Query().Get("entry_type"))
		assert.Equal(t, "123", r.URL.Query().Get("entry_id"))
		assert.Equal(t, "created", r.URL.Query().Get("created"))
		assert.Equal(t, "modified", r.URL.Query().Get("modified"))
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "222", r.URL.Query().Get("root_id"))
		assert.Equal(t, "333", r.URL.Query().Get("reply_id"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id desc", r.URL.Query().Get("order"))
		assert.Equal(t, "id,title", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/comment/get_comments.json"))
	}))

	comments, _, err := client.CommentService.GetComments(ctx, &GetCommentsRequest{
		ID:          NewMulti(111, 222),
		Title:       Ptr("title"),
		Description: Ptr("description"),
		Author:      Ptr("author"),
		EntryType:   Ptr(CommentEntryTypeStories),
		EntryID:     Ptr(123),
		Created:     Ptr("created"),
		Modified:    Ptr("modified"),
		WorkspaceID: Ptr(111),
		RootID:      Ptr(222),
		ReplyID:     Ptr(333),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Order:       NewOrder("id", OrderByDesc),
		Fields:      NewMulti("id", "title"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, comments)
	assert.True(t, len(comments) > 0)
	assert.Equal(t, "1123402991001033061", comments[0].ID)
	assert.Equal(t, "title1", comments[0].Title)
	assert.Equal(t, "description1", comments[0].Description)
	assert.Equal(t, "author1", comments[0].Author)
	assert.Equal(t, CommentEntryTypeBug, comments[0].EntryType)
	assert.Equal(t, "1123402991001037223", comments[0].EntryID)
	assert.Equal(t, "0", comments[0].ReplyID)
	assert.Equal(t, "0", comments[0].RootID)
	assert.Equal(t, "2024-08-28 10:06:24", comments[0].Created)
	assert.Equal(t, "2024-08-28 10:06:24", comments[0].Modified)
	assert.Equal(t, "11112222", comments[0].WorkspaceID)
}

func TestCommentService_GetCommentsCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/comments/count", r.URL.Path)
		assert.Equal(t, "111,222", r.URL.Query().Get("id"))
		assert.Equal(t, "test title", r.URL.Query().Get("title"))
		assert.Equal(t, "test description", r.URL.Query().Get("description"))
		assert.Equal(t, "test author", r.URL.Query().Get("author"))
		assert.Equal(t, CommentEntryTypeStories.String(), r.URL.Query().Get("entry_type"))
		assert.Equal(t, "123", r.URL.Query().Get("entry_id"))
		assert.Equal(t, "2024-08-28", r.URL.Query().Get("created"))
		assert.Equal(t, "2024-08-28", r.URL.Query().Get("modified"))
		assert.Equal(t, "111", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "222", r.URL.Query().Get("root_id"))
		assert.Equal(t, "333", r.URL.Query().Get("reply_id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/comment/get_comments_count.json"))
	}))

	count, _, err := client.CommentService.GetCommentsCount(ctx, &GetCommentsCountRequest{
		ID:          NewMulti(111, 222),
		Title:       Ptr("test title"),
		Description: Ptr("test description"),
		Author:      Ptr("test author"),
		EntryType:   Ptr(CommentEntryTypeStories),
		EntryID:     Ptr(123),
		Created:     Ptr("2024-08-28"),
		Modified:    Ptr("2024-08-28"),
		WorkspaceID: Ptr(111),
		RootID:      Ptr(222),
		ReplyID:     Ptr(333),
	})
	assert.NoError(t, err)
	assert.NotNil(t, count)
	assert.Equal(t, 2284, count)
}

func TestCommentService_UpdateComment(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/comments", r.URL.Path)

		var req struct {
			WorkspaceID   int    `json:"workspace_id"`
			ID            int    `json:"id"`
			Description   string `json:"description"`
			ChangeCreator string `json:"change_creator"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 111, req.WorkspaceID)
		assert.Equal(t, 111, req.ID)
		assert.Equal(t, "test description 2", req.Description)
		assert.Equal(t, "test creator", req.ChangeCreator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/comment/update_comment.json"))
	}))

	comment, _, err := client.CommentService.UpdateComment(ctx, &UpdateCommentRequest{
		WorkspaceID:   Ptr(111),
		ID:            Ptr(111),
		Description:   Ptr("test description 2"),
		ChangeCreator: Ptr("test creator"),
	})
	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, "1111112222001033109", comment.ID)
	assert.Equal(t, "test title", comment.Title)
	assert.Equal(t, "test description", comment.Description)
	assert.Equal(t, "go-tapd", comment.Author)
	assert.Equal(t, CommentEntryTypeStories, comment.EntryType)
	assert.Equal(t, "1111112222001071295", comment.EntryID)
	assert.Equal(t, "0", comment.ReplyID)
	assert.Equal(t, "1111112222001033105", comment.RootID)
	assert.Equal(t, "2024-08-28 23:03:06", comment.Created)
	assert.Equal(t, "2024-08-28 23:03:06", comment.Modified)
	assert.Equal(t, "11112222", comment.WorkspaceID)
}
