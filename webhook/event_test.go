package webhook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhookEvent_EventType(t *testing.T) {
	tests := []struct {
		name string
		want EventType
	}{
		// 需求/任务/缺陷类
		{"story::create", EventTypeStoryCreate},
		{"story::update", EventTypeStoryUpdate},
		{"story::delete", EventTypeStoryDelete},
		{"task::create", EventTypeTaskCreate},
		{"task::update", EventTypeTaskUpdate},
		{"task::delete", EventTypeTaskDelete},
		{"bug::create", EventTypeBugCreate},
		{"bug::update", EventTypeBugUpdate},
		{"bug::delete", EventTypeBugDelete},
		// 评论类：需求/任务/缺陷
		{"story_comment::add", EventTypeStoryCommentAdd},
		{"story_comment::update", EventTypeStoryCommentUpdate},
		{"story_comment::delete", EventTypeStoryCommentDelete},
		{"task_comment::add", EventTypeTaskCommentAdd},
		{"task_comment::update", EventTypeTaskCommentUpdate},
		{"task_comment::delete", EventTypeTaskCommentDelete},
		{"bug_comment::add", EventTypeBugCommentAdd},
		{"bug_comment::update", EventTypeBugCommentUpdate},
		{"bug_comment::delete", EventTypeBugCommentDelete},
		// 迭代
		{"iteration::create", EventTypeIterationCreate},
		{"iteration::update", EventTypeIterationUpdate},
		{"iteration::delete", EventTypeIterationDelete},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, EventType(tt.name))
			assert.Equal(t, tt.name, tt.want.String())
		})
	}
}

func TestWebhookEvent_ParseWebhookEvent(t *testing.T) {
	tests := []struct {
		filename  string
		eventType EventType
		event     any
	}{
		// 需求/任务/缺陷类
		{"story/create.json", EventTypeStoryCreate, &StoryCreateEvent{}},
		{"story/update.json", EventTypeStoryUpdate, &StoryUpdateEvent{}},
		{"story/delete.json", EventTypeStoryDelete, &StoryDeleteEvent{}},
		{"task/create.json", EventTypeTaskCreate, &TaskCreateEvent{}},
		{"task/update.json", EventTypeTaskUpdate, &TaskUpdateEvent{}},
		{"task/delete.json", EventTypeTaskDelete, &TaskDeleteEvent{}},
		{"bug/create.json", EventTypeBugCreate, &BugCreateEvent{}},
		{"bug/update.json", EventTypeBugUpdate, &BugUpdateEvent{}},
		{"bug/delete.json", EventTypeBugDelete, &BugDeleteEvent{}},
		// 评论类：需求/任务/缺陷
		{"story_comment/add.json", EventTypeStoryCommentAdd, &StoryCommentAddEvent{}},
		{"story_comment/update.json", EventTypeStoryCommentUpdate, &StoryCommentUpdateEvent{}},
		{"story_comment/delete.json", EventTypeStoryCommentDelete, &StoryCommentDeleteEvent{}},
		{"task_comment/add.json", EventTypeTaskCommentAdd, &TaskCommentAddEvent{}},
		{"task_comment/update.json", EventTypeTaskCommentUpdate, &TaskCommentUpdateEvent{}},
		{"task_comment/delete.json", EventTypeTaskCommentDelete, &TaskCommentDeleteEvent{}},
		{"bug_comment/add.json", EventTypeBugCommentAdd, &BugCommentAddEvent{}},
		{"bug_comment/update.json", EventTypeBugCommentUpdate, &BugCommentUpdateEvent{}},
		{"bug_comment/delete.json", EventTypeBugCommentDelete, &BugCommentDeleteEvent{}},
		// 迭代
		{"iteration/create.json", EventTypeIterationCreate, &IterationCreateEvent{}},
		{"iteration/update.json", EventTypeIterationUpdate, &IterationUpdateEvent{}},
		{"iteration/delete.json", EventTypeIterationDelete, &IterationDeleteEvent{}},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			payload := loadWebhookData(t, tt.filename)
			eventType, event, err := ParseWebhookEvent(payload)
			require.NoError(t, err)
			assert.Equal(t, tt.eventType, eventType)
			assert.IsType(t, tt.event, event)
		})
	}
}

func TestWebhookEvent_EventChangeFields(t *testing.T) {
	fields := EventChangeFields{"field1", "field2"}
	bytes, err := json.Marshal(fields)
	assert.NoError(t, err)
	assert.Equal(t, `"field1,field2"`, string(bytes))

	var fields2 EventChangeFields
	assert.NoError(t, json.Unmarshal(bytes, &fields2))
	assert.Equal(t, fields, fields2)

	t.Log(fields, fields2)
}

func TestWebhookEvent_EventChangeFields_Extends(t *testing.T) {
	type Extends struct {
		Name   string            `json:"name"`
		Fields EventChangeFields `json:"fields,omitempty"`
	}

	extends := Extends{
		Name:   "extends",
		Fields: EventChangeFields{"field1", "field2"},
	}

	bytes, err := json.Marshal(extends)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"extends","fields":"field1,field2"}`, string(bytes))

	var extends2 Extends
	assert.NoError(t, json.Unmarshal(bytes, &extends2))
	assert.Equal(t, extends, extends2)

	t.Log(extends, extends2)
}
