package webhook

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoryEvent_StoryUpdateEvent(t *testing.T) {
	var event StoryUpdateEvent
	loadAndParseWebhookData(t, "story/update.json", &event)

	assert.Equal(t, EventTypeStoryUpdate, event.Event)
	assert.Equal(t, "11112222", event.WorkspaceID)
	assert.Equal(t, "1111112222001069123", event.ID)
	assert.Equal(t, "1111112222001069123", *event.StoryUpdateEventOldFields.ID)
	assert.Len(t, event.ChangeFields, 2)
	assert.Equal(t, "owner", event.ChangeFields[0])

	bytes, err := json.MarshalIndent(event, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(bytes), "owner,modified")
}

func TestStoryEvent_StoryDeleteEvent(t *testing.T) {
	var event StoryDeleteEvent
	loadAndParseWebhookData(t, "story/delete.json", &event)

	assert.Equal(t, EventTypeStoryDelete, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/tapd_fe/111222333/story/list?useScene=storyList&groupType=&conf_id=11111222333001035191", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001078517", event.ID)
	assert.Equal(t, "delete", event.OpType)
	assert.Equal(t, "0", event.OldIterationID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318959744", event.QueueID)
	assert.Equal(t, "183735666", event.EventID)
	assert.Equal(t, "2024-12-30 18:22:57", event.Created)
}
