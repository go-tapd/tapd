package webhook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventBug_BugCreateEvent(t *testing.T) {
	var event BugCreateEvent
	assert.NoError(t, json.Unmarshal(loadData(t, "../.testdata/webhook/bug_create_event.json"), &event))
	assert.Equal(t, EventTypeBugCreate, event.Event)
	assert.Equal(t, "11112222", event.WorkspaceID)
}
