package webhook

type BugCreateEvent struct {
	Event            EventType `json:"event,omitempty"`
	EventFrom        string    `json:"event_from,omitempty"`
	Referer          string    `json:"referer,omitempty"`
	WorkspaceID      string    `json:"workspace_id,omitempty"`
	CurrentUser      string    `json:"current_user,omitempty"`
	ID               string    `json:"id,omitempty"`
	Title            string    `json:"title,omitempty"`
	IssueID          string    `json:"issue_id,omitempty"`
	IsNewStatus      string    `json:"is_new_status,omitempty"`
	IsReplicate      string    `json:"is_replicate,omitempty"`
	CreateLink       string    `json:"create_link,omitempty"`
	IsJenkins        string    `json:"is_jenkins,omitempty"`
	TemplateID       string    `json:"template_id,omitempty"`
	Description      string    `json:"description,omitempty"`
	IterationID      string    `json:"iteration_id,omitempty"`
	CustomFieldThree string    `json:"custom_field_three,omitempty"`
	Severity         string    `json:"severity,omitempty"`
	Priority         string    `json:"priority,omitempty"`
	CustomFieldFour  string    `json:"custom_field_four,omitempty"`
	CurrentOwner     string    `json:"current_owner,omitempty"`
	Cc               string    `json:"cc,omitempty"`
	De               string    `json:"de,omitempty"`
	Te               string    `json:"te,omitempty"`
	CustomField6     string    `json:"custom_field_6,omitempty"`
	Platform         string    `json:"platform,omitempty"`
	BugType          string    `json:"bugtype,omitempty"`
	OriginPhase      string    `json:"originphase,omitempty"`
	Source           string    `json:"source,omitempty"`
	CustomFieldOne   string    `json:"custom_field_one,omitempty"`
	DescriptionType  string    `json:"description_type,omitempty"`
	ProjectID        string    `json:"project_id,omitempty"`
	IsDraft          string    `json:"is_draft,omitempty"`
	Begin            string    `json:"begin,omitempty"`
	Due              string    `json:"due,omitempty"`
	Status           string    `json:"status,omitempty"`
	Reporter         string    `json:"reporter,omitempty"`
	Flows            string    `json:"flows,omitempty"`
	Resolution       string    `json:"resolution,omitempty"`
	Resolved         string    `json:"resolved,omitempty"`
	Closed           string    `json:"closed,omitempty"`
	InProgressTime   string    `json:"in_progress_time,omitempty"`
	VerifyTime       string    `json:"verify_time,omitempty"`
	RejectTime       string    `json:"reject_time,omitempty"`
	AuditTime        string    `json:"audit_time,omitempty"`
	SuspendTime      string    `json:"suspend_time,omitempty"`
	Secret           string    `json:"secret,omitempty"`
	RioToken         string    `json:"rio_token,omitempty"`
	DevProxyHost     string    `json:"devproxy_host,omitempty"`
	QueueID          string    `json:"queue_id,omitempty"`
	EventID          string    `json:"event_id,omitempty"`
	Created          string    `json:"created,omitempty"`
}

type BugUpdateEvent struct {
	Event EventType `json:"event,omitempty"`
}

type BugDeleteEvent struct {
	Event EventType `json:"event,omitempty"`
}
