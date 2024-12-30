package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterationEvent_IterationCreateEvent(t *testing.T) {
	var event IterationCreateEvent
	loadAndParseWebhookData(t, "iteration/create.json", &event)

	assert.Equal(t, EventTypeIterationCreate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/iterations/card_view/11111222333001002236?q=06aa20b2e5eeee207f0f98f266651024", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001002244", event.ID)
	assert.Equal(t, "产品简称+计划发布版本号43", event.Name)
	assert.Equal(t, "<p>简要输入计划发布版本的主要内容。</p>", event.Description)
	assert.Equal(t, "<p>简要输入计划发布版本的主要内容。</p>", event.MarkdownDescription)
	assert.Equal(t, "1", event.DescriptionType)
	assert.Equal(t, "iteration", event.EntityType)
	assert.Equal(t, "11111222333001000128", event.WorkitemTypeID)
	assert.Equal(t, "1936319866181128277", event.AttachmentToken)
	assert.Equal(t, "[{\"key\":\"StartDate\",\"name\":\"开始时间\",\"period\":\"\",\"value\":\"2024-12-04\"},{\"key\":\"EndDate\",\"name\":\"结束时间\",\"period\":\"\",\"value\":\"2024-12-05\"}]", event.CrucialMoment)
	assert.Equal(t, "2024-12-04", event.StartDate)
	assert.Equal(t, "2024-12-05", event.EndDate)
	assert.Equal(t, "11111222333001000401", event.TemplatedID)
	assert.Equal(t, "张三", event.Creator)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318950442", event.QueueID)
	assert.Equal(t, "183731786", event.EventID)
	assert.Equal(t, "2024-12-30 18:11:44", event.Created)
}

func TestIterationEvent_IterationUpdateEvent(t *testing.T) {
	var event IterationUpdateEvent
	loadAndParseWebhookData(t, "iteration/update.json", &event)

	assert.Equal(t, EventTypeIterationUpdate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/1112223/prong/iterations/card_view/111112223001002236?q=06aa20b2e5eeee207f0f98f266651024", event.Referer)
	assert.Equal(t, "1112223", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "111112223001002244", event.ID)
	assert.Equal(t, "0", event.OldSort)
	assert.Equal(t, "111112223001000128", event.OldWorkitemTypeID)
	assert.Equal(t, "iteration", event.OldEntityType)
	assert.Nil(t, event.NewCustomField47)
	assert.Equal(t, "sort,ancestor_id,path,modified", event.ChangeFields)
	assert.Equal(t, "318950437", event.QueueID)
}

func TestIterationEvent_IterationDeleteEvent(t *testing.T) {
	var event IterationDeleteEvent
	loadAndParseWebhookData(t, "iteration/delete.json", &event)

	assert.Equal(t, EventTypeIterationDelete, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/111222333/prong/iterations/card_view/11111222333001002244?q=7e09c11b30332326c641664c4fce56c8", event.Referer)
	assert.Equal(t, "111222333", event.WorkspaceID)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "11111222333001002244", event.ID)
	assert.Equal(t, "delete", event.OpType)
	assert.Equal(t, "", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "318953897", event.QueueID)
	assert.Equal(t, "183733251", event.EventID)
	assert.Equal(t, "2024-12-30 18:15:14", event.Created)
}
