package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBug_BugCreateEvent(t *testing.T) {
	var event BugCreateEvent
	loadAndParseWebhookData(t, "bug/create.json", &event)

	assert.Equal(t, EventTypeBugCreate, event.Event)
	assert.Equal(t, "111222333", event.WorkspaceID)
}
