package webhook

type TaskCommentAddEvent struct {
	Event EventType `json:"event,omitempty"`
}

type TaskCommentUpdateEvent struct {
	Event EventType `json:"event,omitempty"`
}

type TaskCommentDeleteEvent struct {
	Event EventType `json:"event,omitempty"`
}
