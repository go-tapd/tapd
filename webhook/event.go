package webhook

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// EventType represents the type of webhook event.
type EventType string

const (
	// ========================================
	// 需求/任务/缺陷类
	// ========================================

	EventTypeStoryCreate EventType = "story::create"
	EventTypeStoryUpdate EventType = "story::update"
	EventTypeStoryDelete EventType = "story::delete"
	EventTypeTaskCreate  EventType = "task::create"
	EventTypeTaskUpdate  EventType = "task::update"
	EventTypeTaskDelete  EventType = "task::delete"
	EventTypeBugCreate   EventType = "bug::create"
	EventTypeBugUpdate   EventType = "bug::update"
	EventTypeBugDelete   EventType = "bug::delete"

	// ========================================
	// 评论类：需求/任务/缺陷
	// ========================================

	EventTypeStoryCommentAdd    EventType = "story_comment::add"
	EventTypeStoryCommentUpdate EventType = "story_comment::update"
	EventTypeStoryCommentDelete EventType = "story_comment::delete"
	EventTypeTaskCommentAdd     EventType = "task_comment::add"
	EventTypeTaskCommentUpdate  EventType = "task_comment::update"
	EventTypeTaskCommentDelete  EventType = "task_comment::delete"
	EventTypeBugCommentAdd      EventType = "bug_comment::add"
	EventTypeBugCommentUpdate   EventType = "bug_comment::update"
	EventTypeBugCommentDelete   EventType = "bug_comment::delete"

	// ========================================
	// 迭代
	// ========================================

	EventTypeIterationCreate EventType = "iteration::create"
	EventTypeIterationUpdate EventType = "iteration::update"
	EventTypeIterationDelete EventType = "iteration::delete"
)

func (e EventType) String() string {
	return string(e)
}

// ParseWebhookEvent parses the webhook event from the payload.
func ParseWebhookEvent(payload []byte) (EventType, any, error) {
	var raw map[string]any
	if err := json.Unmarshal(payload, &raw); err != nil {
		return "", nil, err
	}

	// get event
	event, ok := raw["event"].(string)
	if !ok {
		return "", nil, errors.New("tapd: webhook event type not found")
	}

	// decode event
	switch EventType(event) {
	case EventTypeStoryCreate:
		return decodeWebhookEvent[StoryCreateEvent](EventTypeStoryCreate, payload)
	case EventTypeStoryUpdate:
		return decodeWebhookEvent[StoryUpdateEvent](EventTypeStoryUpdate, payload)
	case EventTypeStoryDelete:
		return decodeWebhookEvent[StoryDeleteEvent](EventTypeStoryDelete, payload)
	case EventTypeTaskCreate:
		return decodeWebhookEvent[TaskCreateEvent](EventTypeTaskCreate, payload)
	case EventTypeTaskUpdate:
		return decodeWebhookEvent[TaskUpdateEvent](EventTypeTaskUpdate, payload)
	case EventTypeTaskDelete:
		return decodeWebhookEvent[TaskDeleteEvent](EventTypeTaskDelete, payload)
	case EventTypeBugCreate:
		return decodeWebhookEvent[BugCreateEvent](EventTypeBugCreate, payload)
	case EventTypeBugUpdate:
		return decodeWebhookEvent[BugUpdateEvent](EventTypeBugUpdate, payload)
	case EventTypeBugDelete:
		return decodeWebhookEvent[BugDeleteEvent](EventTypeBugDelete, payload)
	case EventTypeStoryCommentAdd:
		return decodeWebhookEvent[StoryCommentAddEvent](EventTypeStoryCommentAdd, payload)
	case EventTypeStoryCommentUpdate:
		return decodeWebhookEvent[StoryCommentUpdateEvent](EventTypeStoryCommentUpdate, payload)
	case EventTypeStoryCommentDelete:
		return decodeWebhookEvent[StoryCommentDeleteEvent](EventTypeStoryCommentDelete, payload)
	case EventTypeTaskCommentAdd:
		return decodeWebhookEvent[TaskCommentAddEvent](EventTypeTaskCommentAdd, payload)
	case EventTypeTaskCommentUpdate:
		return decodeWebhookEvent[TaskCommentUpdateEvent](EventTypeTaskCommentUpdate, payload)
	case EventTypeTaskCommentDelete:
		return decodeWebhookEvent[TaskCommentDeleteEvent](EventTypeTaskCommentDelete, payload)
	case EventTypeBugCommentAdd:
		return decodeWebhookEvent[BugCommentAddEvent](EventTypeBugCommentAdd, payload)
	case EventTypeBugCommentUpdate:
		return decodeWebhookEvent[BugCommentUpdateEvent](EventTypeBugCommentUpdate, payload)
	case EventTypeBugCommentDelete:
		return decodeWebhookEvent[BugCommentDeleteEvent](EventTypeBugCommentDelete, payload)
	case EventTypeIterationCreate:
		return decodeWebhookEvent[IterationCreateEvent](EventTypeIterationCreate, payload)
	case EventTypeIterationUpdate:
		return decodeWebhookEvent[IterationUpdateEvent](EventTypeIterationUpdate, payload)
	case EventTypeIterationDelete:
		return decodeWebhookEvent[IterationDeleteEvent](EventTypeIterationDelete, payload)
	default: // todo: add more event types
		return "", nil, fmt.Errorf("tapd: webhook event type [%s] not supported", event)
	}
}

// decodeWebhookEvent decodes the webhook event from the payload.
func decodeWebhookEvent[T any](eventType EventType, payload []byte) (EventType, *T, error) {
	var event T
	if err := json.Unmarshal(payload, &event); err != nil {
		return eventType, nil, err
	}
	return eventType, &event, nil
}

// ChangeFields represents the changed fields in the webhook event.
//
// Deprecated: todo remove
type ChangeFields []string

var (
	_ json.Marshaler   = (*ChangeFields)(nil)
	_ json.Unmarshaler = (*ChangeFields)(nil)
)

func (f ChangeFields) MarshalJSON() ([]byte, error) {
	if f == nil {
		return json.Marshal(nil)
	}
	return json.Marshal(strings.Join(f, ","))
}

func (f *ChangeFields) UnmarshalJSON(data []byte) error {
	if f == nil {
		return errors.New("tapd: unmarshal nil pointer")
	}

	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*f = strings.Split(raw, ",")
	return nil
}
