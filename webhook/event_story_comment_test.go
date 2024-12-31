package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoryCommentEvent_StoryCommentAddEvent(t *testing.T) {
	var event StoryCommentAddEvent
	loadAndParseWebhookData(t, "story_comment/add.json", &event)

	assert.Equal(t, EventTypeStoryCommentAdd, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/stories/view/11111222333001116468?from_iteration_id=11111222333001002079", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001036003", event.ID)
	assert.Equal(t, "在状态 [规划中] 添加", event.Title)
	assert.Equal(t, "<p>hello world</p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222333001116468", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318940853", event.QueueID)
	assert.Equal(t, "183727896", event.EventID)
	assert.Equal(t, "2024-12-30 17:59:31", event.Created)
}

func TestStoryCommentEvent_StoryCommentUpdateEvent(t *testing.T) {
	var event StoryCommentUpdateEvent
	loadAndParseWebhookData(t, "story_comment/update.json", &event)

	assert.Equal(t, EventTypeStoryCommentUpdate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/tapd_fe/111222/story/list?useScene=storyList&groupType=&conf_id=11111222001035191&dialog_preview_id=story_11111222001112724", event.Referer)
	assert.Equal(t, "111222", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222001036005", event.ID)
	assert.Equal(t, "<p>123123123123</p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222001112724", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318946074", event.QueueID)
	assert.Equal(t, "183729987", event.EventID)
	assert.Equal(t, "2024-12-30 18:07:03", event.Created)
}

func TestStoryCommentEvent_StoryCommentDeleteEvent(t *testing.T) {
	var event StoryCommentDeleteEvent
	loadAndParseWebhookData(t, "story_comment/delete.json", &event)

	assert.Equal(t, EventTypeStoryCommentDelete, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/tapd_fe/111222/story/list?useScene=storyList&groupType=&conf_id=11111222001035191&dialog_preview_id=story_11111222001112724", event.Referer)
	assert.Equal(t, "111222", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222001036005", event.ID)
	assert.Equal(t, "<p>123123123123</p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222001112724", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318947018", event.QueueID)
	assert.Equal(t, "183730338", event.EventID)
	assert.Equal(t, "2024-12-30 18:08:15", event.Created)
}
