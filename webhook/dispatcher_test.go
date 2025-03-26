package webhook

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func loadData(t *testing.T, filepath string) []byte {
	content, err := os.ReadFile(filepath)
	require.NoError(t, err)
	return content
}

func loadWebhookData(t *testing.T, filename string) []byte {
	return loadData(t, "../internal/testdata/webhook/"+filename)
}

func loadAndParseWebhookData(t *testing.T, filename string, v any) {
	require.NoError(t, json.Unmarshal(loadWebhookData(t, filename), v))
}

func TestDispatcher_Dispatch(t *testing.T) {
	dispatcher := NewDispatcher(
		WithRegisters(&testListener{t: t}),
	)

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/webhook", r.URL.Path)
		assert.NoError(t,
			dispatcher.DispatchRequest(r, WithDispatchRequestContext(
				newDispatcherContext(r.Context()),
			)),
		)

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{"status":"ok"}`))
		assert.NoError(t, err)
	}

	tests := []struct {
		name     string
		filename string
	}{
		{"story create", "story/create.json"},
		{"story update", "story/update.json"},
		{"story delete", "story/delete.json"},
		{"task create", "task/create.json"},
		{"task update", "task/update.json"},
		{"task delete", "task/delete.json"},
		{"bug create", "bug/create.json"},
		{"bug update", "bug/update.json"},
		{"bug delete", "bug/delete.json"},
		{"story comment add", "story_comment/add.json"},
		{"story comment update", "story_comment/update.json"},
		{"story comment delete", "story_comment/delete.json"},
		{"task comment add", "task_comment/add.json"},
		{"task comment update", "task_comment/update.json"},
		{"task comment delete", "task_comment/delete.json"},
		{"bug comment add", "bug_comment/add.json"},
		{"bug comment update", "bug_comment/update.json"},
		{"bug comment delete", "bug_comment/delete.json"},
		{"iteration create", "iteration/create.json"},
		{"iteration update", "iteration/update.json"},
		{"iteration delete", "iteration/delete.json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewReader(loadWebhookData(t, tt.filename)))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			handler(w, req)

			resp := w.Result()
			defer resp.Body.Close() // nolint:errcheck

			assert.Equal(t, http.StatusOK, resp.StatusCode)
			buf, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			assert.Equal(t, `{"status":"ok"}`, string(buf))
		})
	}
}

type dispatcherContextKey struct{}

func newDispatcherContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, dispatcherContextKey{}, "DispatcherContext")
}

func testDispatcherContext(ctx context.Context, t *testing.T) {
	v, ok := ctx.Value(dispatcherContextKey{}).(string)
	require.True(t, ok)
	assert.Equal(t, "DispatcherContext", v)
}

type testListener struct {
	t *testing.T
}

var (
	_ StoryCreateListener        = (*testListener)(nil)
	_ StoryUpdateListener        = (*testListener)(nil)
	_ StoryDeleteListener        = (*testListener)(nil)
	_ TaskCreateListener         = (*testListener)(nil)
	_ TaskUpdateListener         = (*testListener)(nil)
	_ TaskDeleteListener         = (*testListener)(nil)
	_ BugCreateListener          = (*testListener)(nil)
	_ BugUpdateListener          = (*testListener)(nil)
	_ BugDeleteListener          = (*testListener)(nil)
	_ StoryCommentAddListener    = (*testListener)(nil)
	_ StoryCommentUpdateListener = (*testListener)(nil)
	_ StoryCommentDeleteListener = (*testListener)(nil)
	_ TaskCommentAddListener     = (*testListener)(nil)
	_ TaskCommentUpdateListener  = (*testListener)(nil)
	_ TaskCommentDeleteListener  = (*testListener)(nil)
	_ BugCommentAddListener      = (*testListener)(nil)
	_ BugCommentUpdateListener   = (*testListener)(nil)
	_ BugCommentDeleteListener   = (*testListener)(nil)
	_ IterationCreateListener    = (*testListener)(nil)
	_ IterationUpdateListener    = (*testListener)(nil)
	_ IterationDeleteListener    = (*testListener)(nil)
)

func (t testListener) OnStoryCreate(ctx context.Context, event *StoryCreateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryCreate, event.Event)
	return nil
}

func (t testListener) OnStoryUpdate(ctx context.Context, event *StoryUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryUpdate, event.Event)
	return nil
}

func (t testListener) OnStoryDelete(ctx context.Context, event *StoryDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryDelete, event.Event)
	return nil
}

func (t testListener) OnTaskCreate(ctx context.Context, event *TaskCreateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskCreate, event.Event)
	return nil
}

func (t testListener) OnTaskUpdate(ctx context.Context, event *TaskUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskUpdate, event.Event)
	return nil
}

func (t testListener) OnTaskDelete(ctx context.Context, event *TaskDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskDelete, event.Event)
	return nil
}

func (t testListener) OnBugCreate(ctx context.Context, event *BugCreateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugCreate, event.Event)
	return nil
}

func (t testListener) OnBugUpdate(ctx context.Context, event *BugUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugUpdate, event.Event)
	return nil
}

func (t testListener) OnBugDelete(ctx context.Context, event *BugDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugDelete, event.Event)
	return nil
}

func (t testListener) OnStoryCommentAdd(ctx context.Context, event *StoryCommentAddEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryCommentAdd, event.Event)
	return nil
}

func (t testListener) OnStoryCommentUpdate(ctx context.Context, event *StoryCommentUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryCommentUpdate, event.Event)
	return nil
}

func (t testListener) OnStoryCommentDelete(ctx context.Context, event *StoryCommentDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeStoryCommentDelete, event.Event)
	return nil
}

func (t testListener) OnTaskCommentAdd(ctx context.Context, event *TaskCommentAddEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskCommentAdd, event.Event)
	return nil
}

func (t testListener) OnTaskCommentUpdate(ctx context.Context, event *TaskCommentUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskCommentUpdate, event.Event)
	return nil
}

func (t testListener) OnTaskCommentDelete(ctx context.Context, event *TaskCommentDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeTaskCommentDelete, event.Event)
	return nil
}

func (t testListener) OnBugCommentAdd(ctx context.Context, event *BugCommentAddEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugCommentAdd, event.Event)
	return nil
}

func (t testListener) OnBugCommentUpdate(ctx context.Context, event *BugCommentUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugCommentUpdate, event.Event)
	return nil
}

func (t testListener) OnBugCommentDelete(ctx context.Context, event *BugCommentDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeBugCommentDelete, event.Event)
	return nil
}

func (t testListener) OnIterationCreate(ctx context.Context, event *IterationCreateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeIterationCreate, event.Event)
	return nil
}

func (t testListener) OnIterationUpdate(ctx context.Context, event *IterationUpdateEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeIterationUpdate, event.Event)
	return nil
}

func (t testListener) OnIterationDelete(ctx context.Context, event *IterationDeleteEvent) error {
	testDispatcherContext(ctx, t.t)
	assert.Equal(t.t, EventTypeIterationDelete, event.Event)
	return nil
}
