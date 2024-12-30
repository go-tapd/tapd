package webhook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookEvent_EventType(t *testing.T) {
	tests := []struct {
		name string
		want EventType
	}{
		{"story::create", EventTypeStoryCreate},
		{"story::update", EventTypeStoryUpdate},
		{"task::update", EventTypeTaskUpdate},
		{"story_comment::add", EventTypeStoryCommentAdd},
		{"bug::create", EventTypeBugCreate},
		{"bug::update", EventTypeBugUpdate},
		{"bug_comment::update", EventTypeBugCommentUpdate},
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
		{"story_create_event.json", EventTypeStoryCreate, &StoryCreateEvent{}},
		{"story_update_event.json", EventTypeStoryUpdate, &StoryUpdateEvent{}},
		{"bug_create_event.json", EventTypeBugCreate, &BugCreateEvent{}},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			payload := loadData(t, "../.testdata/webhook/"+tt.filename)
			eventType, event, err := ParseWebhookEvent(payload)
			assert.NoError(t, err)
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
