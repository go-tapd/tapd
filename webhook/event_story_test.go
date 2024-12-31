package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoryEvent_StoryCreateEvent(t *testing.T) {
	var event StoryCreateEvent
	loadAndParseWebhookData(t, "story/create.json", &event)

	assert.Equal(t, EventTypeStoryCreate, event.Event)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/11112222/prong/stories/add?workitem_type_id=1111112222001000077", event.Referer)
	assert.Equal(t, "11112222", event.WorkspaceID)
	assert.Equal(t, "current_user", event.CurrentUser)
	assert.Equal(t, "1111112222001071295", event.ID)
	assert.Equal(t, "asdfasdfasdfasdfasdf", event.Name)
	assert.Equal(t, "<p><span style=\"color: #3f4a56;\">adafsasdfasdf</span><br></p>", event.Description)
	assert.Equal(t, "1", event.DescriptionType)
	assert.Equal(t, "", event.Owner)
	assert.Equal(t, "2022-06-13", event.Begin)
	assert.Equal(t, "2022-06-20", event.Due)
	assert.Equal(t, "creator", event.Creator)
	assert.Equal(t, "Middle", event.Priority)
	assert.Equal(t, "1111112222001000077", event.WorkitemTypeID)
	assert.Equal(t, "planning", event.Status)
	assert.Equal(t, "1111112222001000399", event.TemplatedID)
	assert.Equal(t, "Story", event.EntityType)
	assert.Equal(t, "0", event.Remain)
	assert.Equal(t, "", event.ParentID)
	assert.Equal(t, "0", event.AncestorID)
	assert.Equal(t, "", event.ChildrenID)
	assert.Equal(t, "asdfasdfsadfasdf", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "", event.CustomFieldOne)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "2822451111", event.QueueID)
	assert.Equal(t, "1687744222", event.EventID)
	assert.Equal(t, "2024-08-28 19:08:06", event.Created)
}

func TestStoryEvent_StoryUpdateEvent(t *testing.T) {
	var event StoryUpdateEvent
	loadAndParseWebhookData(t, "story/update.json", &event)

	assert.Equal(t, EventTypeStoryUpdate, event.Event)
	assert.Equal(t, "11112222", event.WorkspaceID)
	assert.Equal(t, "web", event.EventFrom)
	assert.Equal(t, "https://www.tapd.cn/tapd_fe/111122222/gantt?dialog_preview_id=story_1111112222001069123", event.Referer)
	assert.Equal(t, "张三", event.CurrentUser)
	assert.Equal(t, "1111112222001069123", event.ID)
	assert.Equal(t, "owner,modified", event.ChangeFields)
	assert.Equal(t, "secret-secret-secret", event.Secret)
	assert.Equal(t, "", event.RioToken)
	assert.Equal(t, "http://websocket-proxy", event.DevProxyHost)
	assert.Equal(t, "281818417", event.QueueID)
	assert.Equal(t, "168601153", event.EventID)
	assert.Equal(t, "2024-08-27 18:07:00", event.Created)
	assert.Equal(t, "1111112222001069123", event.OldID)
	assert.Equal(t, "0", event.OldSecretRootID)
	assert.Equal(t, "106925600000", event.OldSort)
	assert.Equal(t, "1111112222001000008", event.OldWorkitemTypeID)
	assert.Equal(t, "old name", event.OldName)
	assert.Equal(t, "old description", event.OldDescription)
	assert.Equal(t, "", event.OldMarkdownDescription)
	assert.Equal(t, "1", event.OldDescriptionType)
	assert.Equal(t, "old creator", event.OldCreator)
	assert.Equal(t, "2024-08-19 19:18:18", event.OldCreated)
	assert.Equal(t, "2024-08-26 13:02:29", event.OldModified)
	assert.Equal(t, "0", event.OldParentID)
	assert.Equal(t, "|", event.OldChildrenID)
	assert.Equal(t, "1111112222001069123", event.OldAncestorID)
	assert.Equal(t, "1111112222001069123:", event.OldPath)
	assert.Equal(t, "0", event.OldLevel)
	assert.Equal(t, "11112222", event.OldWorkspaceID)
	assert.Equal(t, "audited", event.OldStatus)
	assert.Equal(t, "1", event.OldAppID)
	assert.Equal(t, "old owner", event.OldOwner)
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
