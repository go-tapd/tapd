package webhook

// StoryCommentAddEvent story comment add event
type StoryCommentAddEvent struct {
	Event        string `json:"event,omitempty"`
	EventFrom    string `json:"event_from,omitempty"`
	Referer      string `json:"referer,omitempty"`
	WorkspaceID  string `json:"workspace_id,omitempty"`
	CurrentUser  string `json:"current_user,omitempty"`
	ID           string `json:"id,omitempty"`
	Title        string `json:"title,omitempty"`
	Description  string `json:"description,omitempty"`
	Author       string `json:"author,omitempty"`
	EntityID     string `json:"entity_id,omitempty"`
	Secret       string `json:"secret,omitempty"`
	RioToken     string `json:"rio_token,omitempty"`
	DevProxyHost string `json:"devproxy_host,omitempty"`
	QueueID      string `json:"queue_id,omitempty"`
	EventID      string `json:"event_id,omitempty"`
	Created      string `json:"created,omitempty"`
}
