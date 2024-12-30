package webhook

type TaskCreateEvent struct {
	Event EventType `json:"event,omitempty"`
}

type TaskUpdateEvent struct {
	Event EventType `json:"event,omitempty"`
}

type TaskDeleteEvent struct {
	Event EventType `json:"event,omitempty"`
}
