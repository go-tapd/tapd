package webhook

import (
	"testing"

	"github.com/go-tapd/tapd"
	"github.com/stretchr/testify/assert"
)

func TestTaskEvent_TaskCreateEvent(t *testing.T) {
	var event TaskCreateEvent
	loadAndParseWebhookData(t, "task/create.json", &event)

	assert.Equal(t, EventTypeTaskCreate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/11122233/prong/iterations/card_view/1111122233001002079?q=d41112314b0bf2a49649d1677bc2ac4", event.Referer)
	assert.Equal(t, "11122233", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "1111122233001116469", event.ID)
	assert.Equal(t, "任务合并-数据开发", event.Name)
	assert.Equal(t, tapd.TaskStatusOpen, event.Status)
	assert.Equal(t, "1111122233001116468", event.StoryID)
	assert.Equal(t, "1111122233001002079", event.IterationID)
	assert.Equal(t, "张三", event.Creator)
	assert.Equal(t, "2024-12-30  17:58:03", event.Created)
	assert.Equal(t, "", event.Description)
	assert.Equal(t, "", event.MarkdownDescription)
	assert.Equal(t, "1", event.DescriptionType)
	assert.Equal(t, "", event.Effort)
	assert.Equal(t, "0", event.ReleaseID)
	assert.Equal(t, "张三;", event.Owner)
	assert.Equal(t, "Task", event.EntityType)
	assert.Equal(t, "0", event.CustomPlanField1)
	assert.Equal(t, "0", event.CustomPlanField2)
	assert.Equal(t, "0", event.CustomPlanField3)
	assert.Equal(t, "0", event.CustomPlanField4)
	assert.Equal(t, "0", event.CustomPlanField5)
	assert.Equal(t, "0", event.CustomPlanField6)
	assert.Equal(t, "0", event.CustomPlanField7)
	assert.Equal(t, "0", event.CustomPlanField8)
	assert.Equal(t, "0", event.CustomPlanField9)
	assert.Equal(t, "0", event.CustomPlanField10)
	assert.Equal(t, "0", event.ParentID)
	assert.Equal(t, "0", event.AncestorID)
	assert.Equal(t, "", event.ChildrenID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318939988", event.QueueID)
	assert.Equal(t, "183727534", event.EventID)
}

func TestTaskEvent_TaskUpdateEvent(t *testing.T) {
	var event TaskUpdateEvent
	loadAndParseWebhookData(t, "task/update.json", &event)

	assert.Equal(t, EventTypeTaskUpdate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/tasks?conf_id=11111222333001035219&dialog_preview_id=task_11111222333001116478", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001116478", event.ID)
	assert.Equal(t, "11111222333001116478", event.OldID)
	assert.Equal(t, "test story", event.OldName)
	assert.Equal(t, "", event.OldDescription)
	assert.Equal(t, "test story13", event.NewName)
	assert.Equal(t, "2024-12-30 18:23:57", event.NewModified)
	assert.Equal(t, "", event.NewFollower)
	assert.Equal(t, "name,modified", event.ChangeFields)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318960449", event.QueueID)
	assert.Equal(t, "183735958", event.EventID)
	assert.Equal(t, "2024-12-30 18:23:57", event.Created)
}

func TestTaskEvent_TaskDeleteEvent(t *testing.T) {
	var event TaskDeleteEvent
	loadAndParseWebhookData(t, "task/delete.json", &event)

	assert.Equal(t, EventTypeTaskDelete, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/tasks?conf_id=11111222333001035219&dialog_preview_id=task_11111222333001116470", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001116470", event.ID)
	assert.Equal(t, "delete", event.OpType)
	assert.Equal(t, "0", event.OldIterationID)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318944628", event.QueueID)
	assert.Equal(t, "183729448", event.EventID)
	assert.Equal(t, "2024-12-30 18:04:51", event.Created)
}
