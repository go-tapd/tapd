package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBugCommentEvent_BugCommentAddEvent(t *testing.T) {
	var event BugCommentAddEvent
	loadAndParseWebhookData(t, "bug_comment/add.json", &event)

	assert.Equal(t, event, BugCommentAddEvent{
		Event:        EventTypeBugCommentAdd,
		EventFrom:    "web",
		Referer:      "https://www.tapd.cn/tapd_fe/111222/bug/list?confId=11111222001048467&dialog_preview_id=bug_11111222001037212",
		WorkspaceID:  "111222",
		CurrentUser:  "张三",
		ID:           "11111222001036001",
		Title:        "在状态 [新] 添加",
		Description:  "<p>123123</p>",
		Author:       "张三",
		EntityID:     "11111222001037212",
		Secret:       "",
		RioToken:     "",
		DevProxyHost: "http://websocket-proxy",
		QueueID:      "318928285",
		EventID:      "183722797",
		Created:      "2024-12-30 17:41:02",
	})
}

func TestBugCommentEvent_BugCommentUpdateEvent(t *testing.T) {
	var event BugCommentUpdateEvent
	loadAndParseWebhookData(t, "bug_comment/update.json", &event)

	assert.Equal(t, event, BugCommentUpdateEvent{
		Event:        EventTypeBugCommentUpdate,
		EventFrom:    "web",
		Referer:      "https://www.tapd.cn/tapd_fe/111222333/bug/list?confId=11111222333001048467&dialog_preview_id=bug_11111222333001037212",
		WorkspaceID:  "111222333",
		CurrentUser:  "张三",
		ID:           "11111222333001036001",
		Description:  "<p>12312323</p>",
		Author:       "张三",
		EntityID:     "11111222333001037212",
		Secret:       "",
		RioToken:     "",
		DevProxyHost: "http://websocket-proxy",
		QueueID:      "318963935",
		EventID:      "183737413",
		Created:      "2024-12-30 18:30:02",
	})
}

func TestBugCommentEvent_BugCommentDeleteEvent(t *testing.T) {
	var event BugCommentDeleteEvent
	loadAndParseWebhookData(t, "bug_comment/delete.json", &event)

	assert.Equal(t, event, BugCommentDeleteEvent{
		Event:        EventTypeBugCommentDelete,
		EventFrom:    "web",
		Referer:      "https://www.tapd.cn/111111/comments/delete_comment/",
		WorkspaceID:  "111111",
		CurrentUser:  "李四",
		ID:           "11111111001035977",
		Description:  "<p>這個可能是需求</p>",
		Author:       "李四",
		EntityID:     "11111111001039882",
		Secret:       "",
		RioToken:     "",
		DevProxyHost: "http://websocket-proxy",
		QueueID:      "318938077",
		EventID:      "183726824",
		Created:      "2024-12-30 17:55:19",
	})
}
