package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTaskCommentEvent_TaskCommentAddEvent(t *testing.T) {
	var event TaskCommentAddEvent
	loadAndParseWebhookData(t, "task_comment/add.json", &event)

	assert.Equal(t, EventTypeTaskCommentAdd, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222/comments/add_workitem_commment?keepThis=true&TB_iframe=true&height=750&width=1120&entity_type=task&entity_id=11111222001116478&rand=1735553344339", event.Referer)
	assert.Equal(t, "111222", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222001036008", event.ID)
	assert.Equal(t, "在状态 [未开始] 添加", event.Title)
	assert.Equal(t, "<p><span style=\"color: #3f4a56;\">123</span><br></p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222001116478", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318947664", event.QueueID)
	assert.Equal(t, "183730585", event.EventID)
	assert.Equal(t, "2024-12-30 18:09:06", event.Created)
}

func TestTaskCommentEvent_TaskCommentUpdateEvent(t *testing.T) {
	var event TaskCommentUpdateEvent
	loadAndParseWebhookData(t, "task_comment/update.json", &event)

	assert.Equal(t, EventTypeTaskCommentUpdate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/comments/edit_workitem_comment/11111222333001036008?keepThis=true&TB_iframe=true&height=750&width=1120&entity_type=task&rand=1735554507185", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001036008", event.ID)
	assert.Equal(t, "<p><span style=\"color: #3f4a56;\">123222</span><br></p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222333001116478", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318963265", event.QueueID)
	assert.Equal(t, "183737120", event.EventID)
	assert.Equal(t, "2024-12-30 18:28:29", event.Created)
}

func TestTaskCommentEvent_TaskCommentDeleteEvent(t *testing.T) {
	var event TaskCommentDeleteEvent
	loadAndParseWebhookData(t, "task_comment/delete.json", &event)

	assert.Equal(t, EventTypeTaskCommentDelete, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/tasks/view/11111222333001116478", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001036008", event.ID)
	assert.Equal(t, "<p><span style=\"color: #3f4a56;\">123222</span><br  /></p>", event.Description)
	assert.Equal(t, "张三", event.Author)
	assert.Equal(t, "11111222333001116478", event.EntityID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318963459", event.QueueID)
	assert.Equal(t, "183737210", event.EventID)
	assert.Equal(t, "2024-12-30 18:28:58", event.Created)
}
