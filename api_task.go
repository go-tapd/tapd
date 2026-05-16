package tapd

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// TaskStatus 任务状态
type TaskStatus string

const (
	TaskStatusOpen        TaskStatus = "open"        // 未开始
	TaskStatusProgressing TaskStatus = "progressing" // 进行中
	TaskStatusDone        TaskStatus = "done"        // 已完成
)

func (s TaskStatus) String() string {
	return string(s)
}

type TaskFieldsInfoHTMLType string

const (
	TaskFieldsInfoHTMLTypeInput       TaskFieldsInfoHTMLType = "input"
	TaskFieldsInfoHTMLTypeSelect      TaskFieldsInfoHTMLType = "select"
	TaskFieldsInfoHTMLTypeRichEdit    TaskFieldsInfoHTMLType = "rich_edit"
	TaskFieldsInfoHTMLTypeUserChooser TaskFieldsInfoHTMLType = "user_chooser"
	TaskFieldsInfoHTMLTypeDatetime    TaskFieldsInfoHTMLType = "datetime"
	TaskFieldsInfoHTMLTypeFloat       TaskFieldsInfoHTMLType = "float"
	TaskFieldsInfoHTMLTypeMixChooser  TaskFieldsInfoHTMLType = "mix_chooser"
	TaskFieldsInfoHTMLTypeDateInput   TaskFieldsInfoHTMLType = "dateinput"
	TaskFieldsInfoHTMLTypeCheckbox    TaskFieldsInfoHTMLType = "checkbox"
	TaskFieldsInfoHTMLTypeMultiSelect TaskFieldsInfoHTMLType = "multi_select"
)

type (
	// Task 任务
	Task struct {
		ID                string        `json:"id,omitempty"`               // 任务ID
		Name              string        `json:"name,omitempty"`             // 任务标题
		Description       string        `json:"description,omitempty"`      // 任务详细描述
		WorkspaceID       string        `json:"workspace_id,omitempty"`     // 项目ID
		Creator           string        `json:"creator,omitempty"`          // 创建人
		Created           string        `json:"created,omitempty"`          // 创建时间
		Modified          string        `json:"modified,omitempty"`         // 最后修改时间
		Status            TaskStatus    `json:"status,omitempty"`           // 状态
		Owner             string        `json:"owner,omitempty"`            // 任务当前处理人
		CC                string        `json:"cc,omitempty"`               // 抄送人
		Begin             string        `json:"begin,omitempty"`            // 预计开始
		Due               string        `json:"due,omitempty"`              // 预计结束
		StoryID           string        `json:"story_id,omitempty"`         // 关联需求的ID
		IterationID       string        `json:"iteration_id,omitempty"`     // 所属迭代的ID
		Priority          string        `json:"priority,omitempty"`         // 优先级
		Progress          string        `json:"progress,omitempty"`         // 进度
		Completed         string        `json:"completed,omitempty"`        // 完成时间
		EffortCompleted   string        `json:"effort_completed,omitempty"` // 完成工时
		Exceed            string        `json:"exceed,omitempty"`           // 超出工时
		Remain            string        `json:"remain,omitempty"`           // 剩余工时
		Effort            string        `json:"effort,omitempty"`           // 预估工时
		HasAttachment     string        `json:"has_attachment,omitempty"`   // 是否有附件
		ReleaseID         string        `json:"release_id,omitempty"`       // 发布计划ID
		Label             string        `json:"label,omitempty"`            // 标签
		CustomFieldOne    string        `json:"custom_field_one,omitempty"` // 自定义字段
		CustomFieldTwo    string        `json:"custom_field_two,omitempty"`
		CustomFieldThree  string        `json:"custom_field_three,omitempty"`
		CustomFieldFour   string        `json:"custom_field_four,omitempty"`
		CustomFieldFive   string        `json:"custom_field_five,omitempty"`
		CustomFieldSix    string        `json:"custom_field_six,omitempty"`
		CustomFieldSeven  string        `json:"custom_field_seven,omitempty"`
		CustomFieldEight  string        `json:"custom_field_eight,omitempty"`
		CustomField9      string        `json:"custom_field_9,omitempty"`
		CustomField10     string        `json:"custom_field_10,omitempty"`
		CustomField11     string        `json:"custom_field_11,omitempty"`
		CustomField12     string        `json:"custom_field_12,omitempty"`
		CustomField13     string        `json:"custom_field_13,omitempty"`
		CustomField14     string        `json:"custom_field_14,omitempty"`
		CustomField15     string        `json:"custom_field_15,omitempty"`
		CustomField16     string        `json:"custom_field_16,omitempty"`
		CustomField17     string        `json:"custom_field_17,omitempty"`
		CustomField18     string        `json:"custom_field_18,omitempty"`
		CustomField19     string        `json:"custom_field_19,omitempty"`
		CustomField20     string        `json:"custom_field_20,omitempty"`
		CustomField21     string        `json:"custom_field_21,omitempty"`
		CustomField22     string        `json:"custom_field_22,omitempty"`
		CustomField23     string        `json:"custom_field_23,omitempty"`
		CustomField24     string        `json:"custom_field_24,omitempty"`
		CustomField25     string        `json:"custom_field_25,omitempty"`
		CustomField26     string        `json:"custom_field_26,omitempty"`
		CustomField27     string        `json:"custom_field_27,omitempty"`
		CustomField28     string        `json:"custom_field_28,omitempty"`
		CustomField29     string        `json:"custom_field_29,omitempty"`
		CustomField30     string        `json:"custom_field_30,omitempty"`
		CustomField31     string        `json:"custom_field_31,omitempty"`
		CustomField32     string        `json:"custom_field_32,omitempty"`
		CustomField33     string        `json:"custom_field_33,omitempty"`
		CustomField34     string        `json:"custom_field_34,omitempty"`
		CustomField35     string        `json:"custom_field_35,omitempty"`
		CustomField36     string        `json:"custom_field_36,omitempty"`
		CustomField37     string        `json:"custom_field_37,omitempty"`
		CustomField38     string        `json:"custom_field_38,omitempty"`
		CustomField39     string        `json:"custom_field_39,omitempty"`
		CustomField40     string        `json:"custom_field_40,omitempty"`
		CustomField41     string        `json:"custom_field_41,omitempty"`
		CustomField42     string        `json:"custom_field_42,omitempty"`
		CustomField43     string        `json:"custom_field_43,omitempty"`
		CustomField44     string        `json:"custom_field_44,omitempty"`
		CustomField45     string        `json:"custom_field_45,omitempty"`
		CustomField46     string        `json:"custom_field_46,omitempty"`
		CustomField47     string        `json:"custom_field_47,omitempty"`
		CustomField48     string        `json:"custom_field_48,omitempty"`
		CustomField49     string        `json:"custom_field_49,omitempty"`
		CustomField50     string        `json:"custom_field_50,omitempty"`
		CustomPlanField1  string        `json:"custom_plan_field_1,omitempty"`
		CustomPlanField2  string        `json:"custom_plan_field_2,omitempty"`
		CustomPlanField3  string        `json:"custom_plan_field_3,omitempty"`
		CustomPlanField4  string        `json:"custom_plan_field_4,omitempty"`
		CustomPlanField5  string        `json:"custom_plan_field_5,omitempty"`
		CustomPlanField6  string        `json:"custom_plan_field_6,omitempty"`
		CustomPlanField7  string        `json:"custom_plan_field_7,omitempty"`
		CustomPlanField8  string        `json:"custom_plan_field_8,omitempty"`
		CustomPlanField9  string        `json:"custom_plan_field_9,omitempty"`
		CustomPlanField10 string        `json:"custom_plan_field_10,omitempty"`
		PriorityLabel     PriorityLabel `json:"priority_label,omitempty"` // 优先级
	}

	CreateTaskRequest struct {
		Name             *string           `json:"name,omitempty"`             // 任务标题	支持模糊匹配
		Description      *string           `json:"description,omitempty"`      // 任务详细描述
		WorkspaceID      *int              `json:"workspace_id,omitempty"`     // [必须]项目ID
		Creator          *string           `json:"creator,omitempty"`          // 创建人	支持多人员查询
		Created          *string           `json:"created,omitempty"`          // 创建时间	支持时间查询
		Modified         *string           `json:"modified,omitempty"`         // 最后修改时间	支持时间查询
		Status           *Enum[TaskStatus] `json:"status,omitempty"`           // 状态	支持枚举查询
		Label            *Enum[string]     `json:"label,omitempty"`            // 标签查询	支持枚举查询
		Owner            *string           `json:"owner,omitempty"`            // 任务当前处理人	支持模糊匹配
		CC               *string           `json:"cc,omitempty"`               // 抄送人
		Begin            *string           `json:"begin,omitempty"`            // 预计开始	支持时间查询
		Due              *string           `json:"due,omitempty"`              // 预计结束	支持时间查询
		StoryID          *Multi[int64]     `json:"story_id,omitempty"`         // 关联需求的ID	支持多ID查询
		IterationID      *Enum[int64]      `json:"iteration_id,omitempty"`     // 所属迭代的ID	支持枚举查询
		Priority         *string           `json:"priority,omitempty"`         //nolint:lll // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
		PriorityLabel    *PriorityLabel    `json:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
		Progress         *int              `json:"progress,omitempty"`         // 进度
		Completed        *string           `json:"completed,omitempty"`        // 完成时间	支持时间查询
		EffortCompleted  *string           `json:"effort_completed,omitempty"` // 完成工时
		Exceed           *float64          `json:"exceed,omitempty"`           // 超出工时
		Remain           *float64          `json:"remain,omitempty"`           // 剩余工时
		Effort           *string           `json:"effort,omitempty"`           // 预估工时
		CustomFieldOne   *string           `json:"custom_field_one,omitempty"`
		CustomFieldTwo   *string           `json:"custom_field_two,omitempty"`
		CustomFieldThree *string           `json:"custom_field_three,omitempty"`
		CustomFieldFour  *string           `json:"custom_field_four,omitempty"`
		CustomFieldFive  *string           `json:"custom_field_five,omitempty"`
		CustomFieldSix   *string           `json:"custom_field_six,omitempty"`
		CustomFieldSeven *string           `json:"custom_field_seven,omitempty"`
		CustomFieldEight *string           `json:"custom_field_eight,omitempty"`
		CustomField9     *string           `json:"custom_field_9,omitempty"`
		CustomField10    *string           `json:"custom_field_10,omitempty"`
		CustomField11    *string           `json:"custom_field_11,omitempty"`
		CustomField12    *string           `json:"custom_field_12,omitempty"`
		CustomField13    *string           `json:"custom_field_13,omitempty"`
		CustomField14    *string           `json:"custom_field_14,omitempty"`
		CustomField15    *string           `json:"custom_field_15,omitempty"`
		CustomField16    *string           `json:"custom_field_16,omitempty"`
		CustomField17    *string           `json:"custom_field_17,omitempty"`
		CustomField18    *string           `json:"custom_field_18,omitempty"`
		CustomField19    *string           `json:"custom_field_19,omitempty"`
		CustomField20    *string           `json:"custom_field_20,omitempty"`
		CustomField21    *string           `json:"custom_field_21,omitempty"`
		CustomField22    *string           `json:"custom_field_22,omitempty"`
		CustomField23    *string           `json:"custom_field_23,omitempty"`
		CustomField24    *string           `json:"custom_field_24,omitempty"`
		CustomField25    *string           `json:"custom_field_25,omitempty"`
		CustomField26    *string           `json:"custom_field_26,omitempty"`
		CustomField27    *string           `json:"custom_field_27,omitempty"`
		CustomField28    *string           `json:"custom_field_28,omitempty"`
		CustomField29    *string           `json:"custom_field_29,omitempty"`
		CustomField30    *string           `json:"custom_field_30,omitempty"`
		CustomField31    *string           `json:"custom_field_31,omitempty"`
		CustomField32    *string           `json:"custom_field_32,omitempty"`
		CustomField33    *string           `json:"custom_field_33,omitempty"`
		CustomField34    *string           `json:"custom_field_34,omitempty"`
		CustomField35    *string           `json:"custom_field_35,omitempty"`
		CustomField36    *string           `json:"custom_field_36,omitempty"`
		CustomField37    *string           `json:"custom_field_37,omitempty"`
		CustomField38    *string           `json:"custom_field_38,omitempty"`
		CustomField39    *string           `json:"custom_field_39,omitempty"`
		CustomField40    *string           `json:"custom_field_40,omitempty"`
		CustomField41    *string           `json:"custom_field_41,omitempty"`
		CustomField42    *string           `json:"custom_field_42,omitempty"`
		CustomField43    *string           `json:"custom_field_43,omitempty"`
		CustomField44    *string           `json:"custom_field_44,omitempty"`
		CustomField45    *string           `json:"custom_field_45,omitempty"`
		CustomField46    *string           `json:"custom_field_46,omitempty"`
		CustomField47    *string           `json:"custom_field_47,omitempty"`
		CustomField48    *string           `json:"custom_field_48,omitempty"`
		CustomField49    *string           `json:"custom_field_49,omitempty"`
		CustomField50    *string           `json:"custom_field_50,omitempty"`
	}

	UpdateTaskRequest struct {
		ID                 *int64         `json:"id,omitempty"`                   // [必须]任务ID
		WorkspaceID        *int           `json:"workspace_id,omitempty"`         // [必须]项目ID
		CurrentUser        *string        `json:"current_user,omitempty"`         // 当前操作人
		Name               *string        `json:"name,omitempty"`                 // 任务标题
		Description        *string        `json:"description,omitempty"`          // 任务详细描述
		Creator            *string        `json:"creator,omitempty"`              // 创建人
		Status             *TaskStatus    `json:"status,omitempty"`               // 状态
		Label              *Enum[string]  `json:"label,omitempty"`                // 标签
		Owner              *string        `json:"owner,omitempty"`                // 任务当前处理人
		CC                 *string        `json:"cc,omitempty"`                   // 抄送人
		Begin              *string        `json:"begin,omitempty"`                // 预计开始
		Due                *string        `json:"due,omitempty"`                  // 预计结束
		StoryID            *int64         `json:"story_id,omitempty"`             // 关联需求的ID
		IterationID        *int64         `json:"iteration_id,omitempty"`         // 所属迭代的ID
		Priority           *string        `json:"priority,omitempty"`             // 优先级。为了兼容自定义优先级，请使用 priority_label 字段
		PriorityLabel      *PriorityLabel `json:"priority_label,omitempty"`       // 优先级。推荐使用这个字段
		Progress           *int           `json:"progress,omitempty"`             // 进度
		Effort             *string        `json:"effort,omitempty"`               // 预估工时
		AutoCompleteEffort *int           `json:"auto_complete_effort,omitempty"` // 完成任务时是否自动填写完成工时
		CustomFieldOne     *string        `json:"custom_field_one,omitempty"`
		CustomFieldTwo     *string        `json:"custom_field_two,omitempty"`
		CustomFieldThree   *string        `json:"custom_field_three,omitempty"`
		CustomFieldFour    *string        `json:"custom_field_four,omitempty"`
		CustomFieldFive    *string        `json:"custom_field_five,omitempty"`
		CustomFieldSix     *string        `json:"custom_field_six,omitempty"`
		CustomFieldSeven   *string        `json:"custom_field_seven,omitempty"`
		CustomFieldEight   *string        `json:"custom_field_eight,omitempty"`
		CustomField9       *string        `json:"custom_field_9,omitempty"`
		CustomField10      *string        `json:"custom_field_10,omitempty"`
		CustomField11      *string        `json:"custom_field_11,omitempty"`
		CustomField12      *string        `json:"custom_field_12,omitempty"`
		CustomField13      *string        `json:"custom_field_13,omitempty"`
		CustomField14      *string        `json:"custom_field_14,omitempty"`
		CustomField15      *string        `json:"custom_field_15,omitempty"`
		CustomField16      *string        `json:"custom_field_16,omitempty"`
		CustomField17      *string        `json:"custom_field_17,omitempty"`
		CustomField18      *string        `json:"custom_field_18,omitempty"`
		CustomField19      *string        `json:"custom_field_19,omitempty"`
		CustomField20      *string        `json:"custom_field_20,omitempty"`
		CustomField21      *string        `json:"custom_field_21,omitempty"`
		CustomField22      *string        `json:"custom_field_22,omitempty"`
		CustomField23      *string        `json:"custom_field_23,omitempty"`
		CustomField24      *string        `json:"custom_field_24,omitempty"`
		CustomField25      *string        `json:"custom_field_25,omitempty"`
		CustomField26      *string        `json:"custom_field_26,omitempty"`
		CustomField27      *string        `json:"custom_field_27,omitempty"`
		CustomField28      *string        `json:"custom_field_28,omitempty"`
		CustomField29      *string        `json:"custom_field_29,omitempty"`
		CustomField30      *string        `json:"custom_field_30,omitempty"`
		CustomField31      *string        `json:"custom_field_31,omitempty"`
		CustomField32      *string        `json:"custom_field_32,omitempty"`
		CustomField33      *string        `json:"custom_field_33,omitempty"`
		CustomField34      *string        `json:"custom_field_34,omitempty"`
		CustomField35      *string        `json:"custom_field_35,omitempty"`
		CustomField36      *string        `json:"custom_field_36,omitempty"`
		CustomField37      *string        `json:"custom_field_37,omitempty"`
		CustomField38      *string        `json:"custom_field_38,omitempty"`
		CustomField39      *string        `json:"custom_field_39,omitempty"`
		CustomField40      *string        `json:"custom_field_40,omitempty"`
		CustomField41      *string        `json:"custom_field_41,omitempty"`
		CustomField42      *string        `json:"custom_field_42,omitempty"`
		CustomField43      *string        `json:"custom_field_43,omitempty"`
		CustomField44      *string        `json:"custom_field_44,omitempty"`
		CustomField45      *string        `json:"custom_field_45,omitempty"`
		CustomField46      *string        `json:"custom_field_46,omitempty"`
		CustomField47      *string        `json:"custom_field_47,omitempty"`
		CustomField48      *string        `json:"custom_field_48,omitempty"`
		CustomField49      *string        `json:"custom_field_49,omitempty"`
		CustomField50      *string        `json:"custom_field_50,omitempty"`
		CustomPlanField1   *string        `json:"custom_plan_field_1,omitempty"` // 自定义计划应用参数
		CustomPlanField2   *string        `json:"custom_plan_field_2,omitempty"`
		CustomPlanField3   *string        `json:"custom_plan_field_3,omitempty"`
		CustomPlanField4   *string        `json:"custom_plan_field_4,omitempty"`
		CustomPlanField5   *string        `json:"custom_plan_field_5,omitempty"`
		CustomPlanField6   *string        `json:"custom_plan_field_6,omitempty"`
		CustomPlanField7   *string        `json:"custom_plan_field_7,omitempty"`
		CustomPlanField8   *string        `json:"custom_plan_field_8,omitempty"`
		CustomPlanField9   *string        `json:"custom_plan_field_9,omitempty"`
		CustomPlanField10  *string        `json:"custom_plan_field_10,omitempty"`
	}

	TaskChange struct {
		ID             string                  `json:"id,omitempty"`
		WorkspaceID    string                  `json:"workspace_id,omitempty"`
		AppID          string                  `json:"app_id,omitempty"`
		WorkitemTypeID string                  `json:"workitem_type_id,omitempty"`
		Creator        string                  `json:"creator,omitempty"`
		Created        string                  `json:"created,omitempty"`
		ChangeSummary  string                  `json:"change_summary,omitempty"`
		Comment        string                  `json:"comment,omitempty"`
		Changes        string                  `json:"changes,omitempty"`
		EntityType     string                  `json:"entity_type,omitempty"`
		ChangeType     string                  `json:"change_type,omitempty"`
		ChangeTypeText string                  `json:"change_type_text,omitempty"`
		FieldChanges   []TaskChangeFieldChange `json:"field_changes,omitempty"`
		TaskID         string                  `json:"task_id,omitempty"`
	}

	TaskChangeFieldChange struct {
		Field             string `json:"field,omitempty"`
		ValueBefore       string `json:"value_before,omitempty"`
		ValueAfter        string `json:"value_after,omitempty"`
		ValueBeforeParsed string `json:"value_before_parsed,omitempty"`
		ValueAfterParsed  string `json:"value_after_parsed,omitempty"`
		FieldLabel        string `json:"field_label,omitempty"`
	}

	GetTaskChangesRequest struct {
		ID               *Multi[int64]  `url:"id,omitempty"`                 // 支持多ID查询
		WorkspaceID      *int           `url:"workspace_id,omitempty"`       // [必须]项目ID
		TaskID           *int64         `url:"task_id,omitempty"`            // 任务ID
		Creator          *string        `url:"creator,omitempty"`            // 创建人（操作人）
		Created          *string        `url:"created,omitempty"`            // 创建时间（变更时间）	支持时间查询
		ChangeSummary    *string        `url:"change_summary,omitempty"`     // 需求变更描述
		Comment          *string        `url:"comment,omitempty"`            // 评论
		Changes          *string        `url:"changes,omitempty"`            // 变更详细记录
		EntityType       *string        `url:"entity_type,omitempty"`        // 变更的对象类型
		NeedParseChanges *int           `url:"need_parse_changes,omitempty"` // 设置field_changes字段是否返回（默认取 1。取 0 则不返回）
		Limit            *int           `url:"limit,omitempty"`              // 设置返回数量限制，默认为30
		Page             *int           `url:"page,omitempty"`               // 返回当前数量限制下第N页的数据，默认为1（第一页）
		Order            *Order         `url:"order,omitempty"`              //nolint:lll // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode	如按创建时间逆序：order=created%20desc
		Fields           *Multi[string] `url:"fields,omitempty"`             // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetTaskChangesCountRequest struct {
		ID            *Multi[int64] `url:"id,omitempty"`             // 支持多ID查询
		WorkspaceID   *int          `url:"workspace_id,omitempty"`   // [必须]项目ID
		TaskID        *int64        `url:"task_id,omitempty"`        // 任务ID
		Creator       *string       `url:"creator,omitempty"`        // 创建人（操作人）
		Created       *string       `url:"created,omitempty"`        // 创建时间（变更时间）	支持时间查询
		ChangeSummary *string       `url:"change_summary,omitempty"` // 需求变更描述
		Comment       *string       `url:"comment,omitempty"`        // 评论
		Changes       *string       `url:"changes,omitempty"`        // 变更详细记录
		EntityType    *string       `url:"entity_type,omitempty"`    // 变更的对象类型
	}

	GetTasksRequest struct {
		ID               *Multi[int64]     `url:"id,omitempty"`               // 支持多ID查询、模糊匹配
		Name             *string           `url:"name,omitempty"`             // 任务标题	支持模糊匹配
		Description      *string           `url:"description,omitempty"`      // 任务详细描述
		WorkspaceID      *int              `url:"workspace_id,omitempty"`     // [必须]项目ID
		Creator          *string           `url:"creator,omitempty"`          // 创建人	支持多人员查询
		Created          *string           `url:"created,omitempty"`          // 创建时间	支持时间查询
		Modified         *string           `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
		Status           *Enum[TaskStatus] `url:"status,omitempty"`           // 状态	支持枚举查询
		Label            *Enum[string]     `url:"label,omitempty"`            // 标签查询	支持枚举查询
		Owner            *string           `url:"owner,omitempty"`            // 任务当前处理人	支持模糊匹配
		CC               *string           `url:"cc,omitempty"`               // 抄送人
		Begin            *string           `url:"begin,omitempty"`            // 预计开始	支持时间查询
		Due              *string           `url:"due,omitempty"`              // 预计结束	支持时间查询
		StoryID          *Multi[int64]     `url:"story_id,omitempty"`         // 关联需求的ID	支持多ID查询
		IterationID      *Enum[int64]      `url:"iteration_id,omitempty"`     // 所属迭代的ID	支持枚举查询
		Priority         *string           `url:"priority,omitempty"`         //nolint:lll // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
		PriorityLabel    *PriorityLabel    `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
		Progress         *int              `url:"progress,omitempty"`         // 进度
		Completed        *string           `url:"completed,omitempty"`        // 完成时间	支持时间查询
		EffortCompleted  *string           `url:"effort_completed,omitempty"` // 完成工时
		Exceed           *float64          `url:"exceed,omitempty"`           // 超出工时
		Remain           *float64          `url:"remain,omitempty"`           // 剩余工时
		Effort           *string           `url:"effort,omitempty"`           // 预估工时
		CustomFieldOne   *string           `url:"custom_field_one,omitempty"`
		CustomFieldTwo   *string           `url:"custom_field_two,omitempty"`
		CustomFieldThree *string           `url:"custom_field_three,omitempty"`
		CustomFieldFour  *string           `url:"custom_field_four,omitempty"`
		CustomFieldFive  *string           `url:"custom_field_five,omitempty"`
		CustomFieldSix   *string           `url:"custom_field_six,omitempty"`
		CustomFieldSeven *string           `url:"custom_field_seven,omitempty"`
		CustomFieldEight *string           `url:"custom_field_eight,omitempty"`
		CustomField9     *string           `url:"custom_field_9,omitempty"`
		CustomField10    *string           `url:"custom_field_10,omitempty"`
		CustomField11    *string           `url:"custom_field_11,omitempty"`
		CustomField12    *string           `url:"custom_field_12,omitempty"`
		CustomField13    *string           `url:"custom_field_13,omitempty"`
		CustomField14    *string           `url:"custom_field_14,omitempty"`
		CustomField15    *string           `url:"custom_field_15,omitempty"`
		CustomField16    *string           `url:"custom_field_16,omitempty"`
		CustomField17    *string           `url:"custom_field_17,omitempty"`
		CustomField18    *string           `url:"custom_field_18,omitempty"`
		CustomField19    *string           `url:"custom_field_19,omitempty"`
		CustomField20    *string           `url:"custom_field_20,omitempty"`
		CustomField21    *string           `url:"custom_field_21,omitempty"`
		CustomField22    *string           `url:"custom_field_22,omitempty"`
		CustomField23    *string           `url:"custom_field_23,omitempty"`
		CustomField24    *string           `url:"custom_field_24,omitempty"`
		CustomField25    *string           `url:"custom_field_25,omitempty"`
		CustomField26    *string           `url:"custom_field_26,omitempty"`
		CustomField27    *string           `url:"custom_field_27,omitempty"`
		CustomField28    *string           `url:"custom_field_28,omitempty"`
		CustomField29    *string           `url:"custom_field_29,omitempty"`
		CustomField30    *string           `url:"custom_field_30,omitempty"`
		CustomField31    *string           `url:"custom_field_31,omitempty"`
		CustomField32    *string           `url:"custom_field_32,omitempty"`
		CustomField33    *string           `url:"custom_field_33,omitempty"`
		CustomField34    *string           `url:"custom_field_34,omitempty"`
		CustomField35    *string           `url:"custom_field_35,omitempty"`
		CustomField36    *string           `url:"custom_field_36,omitempty"`
		CustomField37    *string           `url:"custom_field_37,omitempty"`
		CustomField38    *string           `url:"custom_field_38,omitempty"`
		CustomField39    *string           `url:"custom_field_39,omitempty"`
		CustomField40    *string           `url:"custom_field_40,omitempty"`
		CustomField41    *string           `url:"custom_field_41,omitempty"`
		CustomField42    *string           `url:"custom_field_42,omitempty"`
		CustomField43    *string           `url:"custom_field_43,omitempty"`
		CustomField44    *string           `url:"custom_field_44,omitempty"`
		CustomField45    *string           `url:"custom_field_45,omitempty"`
		CustomField46    *string           `url:"custom_field_46,omitempty"`
		CustomField47    *string           `url:"custom_field_47,omitempty"`
		CustomField48    *string           `url:"custom_field_48,omitempty"`
		CustomField49    *string           `url:"custom_field_49,omitempty"`
		CustomField50    *string           `url:"custom_field_50,omitempty"`
		Limit            *int              `url:"limit,omitempty"`  // 设置返回数量限制，默认为30
		Page             *int              `url:"page,omitempty"`   // 返回当前数量限制下第N页的数据，默认为1（第一页）
		Order            *Order            `url:"order,omitempty"`  //nolint:lll // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode	如按创建时间逆序：order=created%20desc
		Fields           *Multi[string]    `url:"fields,omitempty"` // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetTasksCountRequest struct {
		ID               *Multi[int64]     `url:"id,omitempty"`               // 支持多ID查询、模糊匹配
		Name             *string           `url:"name,omitempty"`             // 任务标题	支持模糊匹配
		Description      *string           `url:"description,omitempty"`      // 任务详细描述
		WorkspaceID      *int              `url:"workspace_id,omitempty"`     // [必须]项目ID
		Creator          *string           `url:"creator,omitempty"`          // 创建人	支持多人员查询
		Created          *string           `url:"created,omitempty"`          // 创建时间	支持时间查询
		Modified         *string           `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
		Status           *Enum[TaskStatus] `url:"status,omitempty"`           // 状态	支持枚举查询
		Label            *Enum[string]     `url:"label,omitempty"`            // 标签查询	支持枚举查询
		Owner            *string           `url:"owner,omitempty"`            // 任务当前处理人	支持模糊匹配
		CC               *string           `url:"cc,omitempty"`               // 抄送人
		Begin            *string           `url:"begin,omitempty"`            // 预计开始	支持时间查询
		Due              *string           `url:"due,omitempty"`              // 预计结束	支持时间查询
		StoryID          *Multi[int64]     `url:"story_id,omitempty"`         // 关联需求的ID	支持多ID查询
		IterationID      *Enum[int64]      `url:"iteration_id,omitempty"`     // 所属迭代的ID	支持枚举查询
		Priority         *string           `url:"priority,omitempty"`         //nolint:lll // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
		PriorityLabel    *PriorityLabel    `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
		Progress         *int              `url:"progress,omitempty"`         // 进度
		Completed        *string           `url:"completed,omitempty"`        // 完成时间	支持时间查询
		EffortCompleted  *string           `url:"effort_completed,omitempty"` // 完成工时
		Exceed           *float64          `url:"exceed,omitempty"`           // 超出工时
		Remain           *float64          `url:"remain,omitempty"`           // 剩余工时
		Effort           *string           `url:"effort,omitempty"`           // 预估工时
		CustomFieldOne   *string           `url:"custom_field_one,omitempty"`
		CustomFieldTwo   *string           `url:"custom_field_two,omitempty"`
		CustomFieldThree *string           `url:"custom_field_three,omitempty"`
		CustomFieldFour  *string           `url:"custom_field_four,omitempty"`
		CustomFieldFive  *string           `url:"custom_field_five,omitempty"`
		CustomFieldSix   *string           `url:"custom_field_six,omitempty"`
		CustomFieldSeven *string           `url:"custom_field_seven,omitempty"`
		CustomFieldEight *string           `url:"custom_field_eight,omitempty"`
		CustomField9     *string           `url:"custom_field_9,omitempty"`
		CustomField10    *string           `url:"custom_field_10,omitempty"`
		CustomField11    *string           `url:"custom_field_11,omitempty"`
		CustomField12    *string           `url:"custom_field_12,omitempty"`
		CustomField13    *string           `url:"custom_field_13,omitempty"`
		CustomField14    *string           `url:"custom_field_14,omitempty"`
		CustomField15    *string           `url:"custom_field_15,omitempty"`
		CustomField16    *string           `url:"custom_field_16,omitempty"`
		CustomField17    *string           `url:"custom_field_17,omitempty"`
		CustomField18    *string           `url:"custom_field_18,omitempty"`
		CustomField19    *string           `url:"custom_field_19,omitempty"`
		CustomField20    *string           `url:"custom_field_20,omitempty"`
		CustomField21    *string           `url:"custom_field_21,omitempty"`
		CustomField22    *string           `url:"custom_field_22,omitempty"`
		CustomField23    *string           `url:"custom_field_23,omitempty"`
		CustomField24    *string           `url:"custom_field_24,omitempty"`
		CustomField25    *string           `url:"custom_field_25,omitempty"`
		CustomField26    *string           `url:"custom_field_26,omitempty"`
		CustomField27    *string           `url:"custom_field_27,omitempty"`
		CustomField28    *string           `url:"custom_field_28,omitempty"`
		CustomField29    *string           `url:"custom_field_29,omitempty"`
		CustomField30    *string           `url:"custom_field_30,omitempty"`
		CustomField31    *string           `url:"custom_field_31,omitempty"`
		CustomField32    *string           `url:"custom_field_32,omitempty"`
		CustomField33    *string           `url:"custom_field_33,omitempty"`
		CustomField34    *string           `url:"custom_field_34,omitempty"`
		CustomField35    *string           `url:"custom_field_35,omitempty"`
		CustomField36    *string           `url:"custom_field_36,omitempty"`
		CustomField37    *string           `url:"custom_field_37,omitempty"`
		CustomField38    *string           `url:"custom_field_38,omitempty"`
		CustomField39    *string           `url:"custom_field_39,omitempty"`
		CustomField40    *string           `url:"custom_field_40,omitempty"`
		CustomField41    *string           `url:"custom_field_41,omitempty"`
		CustomField42    *string           `url:"custom_field_42,omitempty"`
		CustomField43    *string           `url:"custom_field_43,omitempty"`
		CustomField44    *string           `url:"custom_field_44,omitempty"`
		CustomField45    *string           `url:"custom_field_45,omitempty"`
		CustomField46    *string           `url:"custom_field_46,omitempty"`
		CustomField47    *string           `url:"custom_field_47,omitempty"`
		CustomField48    *string           `url:"custom_field_48,omitempty"`
		CustomField49    *string           `url:"custom_field_49,omitempty"`
		CustomField50    *string           `url:"custom_field_50,omitempty"`
	}

	GetTaskFieldsInfoRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TaskFieldsInfoOption struct {
		Value string `json:"key,omitempty"`   // 值
		Label string `json:"label,omitempty"` // 中文名称
	}

	TaskFieldsInfoColorOption struct {
		Value string `json:"value,omitempty"`
		Label string `json:"label,omitempty"`
		Color string `json:"color,omitempty"`
	}

	TaskFieldsInfoPureOption struct {
		ParentID    string `json:"parent_id,omitempty"`
		WorkspaceID string `json:"workspace_id,omitempty"`
		Sort        string `json:"sort,omitempty"`
		OriginName  string `json:"origin_name,omitempty"`
		Value       string `json:"value,omitempty"`
		Label       string `json:"label,omitempty"`
		Panel       int    `json:"panel,omitempty"`
	}

	TaskFieldsInfo struct {
		Name         string                      `json:"name,omitempty"`      // name
		HTMLType     TaskFieldsInfoHTMLType      `json:"html_type,omitempty"` // 类型
		Label        string                      `json:"label,omitempty"`     // 中文名称
		Options      []TaskFieldsInfoOption      `json:"options,omitempty"`   // 候选值
		ColorOptions []TaskFieldsInfoColorOption `json:"color_options,omitempty"`
		PureOptions  []TaskFieldsInfoPureOption  `json:"pure_options,omitempty"`
	}
)

// TaskService 任务服务
type TaskService interface {
	// CreateTask 创建任务
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/add_task.html
	CreateTask(ctx context.Context, request *CreateTaskRequest, opts ...RequestOption) (*Task, *Response, error)

	// GetTaskChanges 获取任务变更历史
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes.html
	GetTaskChanges(ctx context.Context, request *GetTaskChangesRequest, opts ...RequestOption) ([]*TaskChange, *Response, error)

	// GetTaskChangesCount 获取任务变更次数
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_changes_count.html
	GetTaskChangesCount(ctx context.Context, request *GetTaskChangesCountRequest, opts ...RequestOption) (int, *Response, error)

	// 获取任务自定义字段配置

	// GetTasks 获取任务
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks.html
	GetTasks(ctx context.Context, request *GetTasksRequest, opts ...RequestOption) ([]*Task, *Response, error)

	// GetTasksCount 获取任务数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_tasks_count.html
	GetTasksCount(ctx context.Context, request *GetTasksCountRequest, opts ...RequestOption) (int, *Response, error)

	// UpdateTask 更新任务
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/update_task.html
	UpdateTask(ctx context.Context, request *UpdateTaskRequest, opts ...RequestOption) (*Task, *Response, error)

	// 获取回收站下的任务
	// 获取视图对应的任务列表

	// GetTaskFieldsInfo 获取任务字段信息
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/task/get_task_fields_info.html
	GetTaskFieldsInfo(ctx context.Context, request *GetTaskFieldsInfoRequest, opts ...RequestOption) ([]*TaskFieldsInfo, *Response, error)
}

type taskService struct {
	client *Client
}

var _ TaskService = (*taskService)(nil)

func NewTaskService(client *Client) TaskService {
	return &taskService{
		client: client,
	}
}

func (s *taskService) CreateTask(ctx context.Context, request *CreateTaskRequest, opts ...RequestOption) (*Task, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tasks", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Task *Task `json:"task"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Task, resp, nil
}

// ↓↓↓↓ 这段代码是为了解决 Tapd API 返回的不同数据类型问题，官方的 API 写的非常好 🙂🙂----开始
type rawTaskChange struct {
	TaskChange
	FieldChanges []struct {
		TaskChangeFieldChange
		ValueBefore any `json:"value_before"` // 为了兼容自定义字段，value_before 和 value_after 为 any 类型
		ValueAfter  any `json:"value_after"`  // 为了兼容自定义字段，value_before 和 value_after 为 any 类型
	} `json:"field_changes,omitempty"`
}

func parseRawTaskChange(raw *rawTaskChange) (*TaskChange, error) {
	fieldChanges := make([]TaskChangeFieldChange, 0, len(raw.TaskChange.FieldChanges))

	for _, rawFieldChange := range raw.FieldChanges {
		fieldChange := rawFieldChange.TaskChangeFieldChange

		// value_before 和 value_after 为 any 类型，需要根据实际类型解析
		valueBefore, err := decodeGetTaskChangesFieldChangesValue(rawFieldChange.ValueBefore)
		if err != nil {
			return nil, err
		}
		fieldChange.ValueBefore = valueBefore

		valueAfter, err := decodeGetTaskChangesFieldChangesValue(rawFieldChange.ValueAfter)
		if err != nil {
			return nil, err
		}
		fieldChange.ValueAfter = valueAfter

		fieldChanges = append(fieldChanges, fieldChange)
	}

	change := raw.TaskChange
	change.FieldChanges = fieldChanges
	return &change, nil
}

func decodeGetTaskChangesFieldChangesValue(v any) (string, error) {
	switch v := v.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("unexpected type %T", v)
	}
}

// ↑↑↑↑ 这段代码是为了解决 Tapd API 返回的不同数据类型问题，官方的 API 写的非常好 🙂🙂----结束(再次👏）

func (s *taskService) GetTaskChanges(
	ctx context.Context, request *GetTaskChangesRequest, opts ...RequestOption,
) ([]*TaskChange, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "task_changes", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var rawItems []struct {
		WorkitemChange *rawTaskChange `json:"WorkitemChange"`
	}
	resp, err := s.client.Do(req, &rawItems)
	if err != nil {
		return nil, resp, err
	}

	changes := make([]*TaskChange, 0, len(rawItems))
	for _, rawItem := range rawItems {
		change, err := parseRawTaskChange(rawItem.WorkitemChange)
		if err != nil {
			return nil, resp, err
		}
		changes = append(changes, change)
	}

	return changes, resp, nil
}

func (s *taskService) GetTaskChangesCount(
	ctx context.Context, request *GetTaskChangesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "task_changes/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response CountResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

func (s *taskService) GetTasks(
	ctx context.Context, request *GetTasksRequest, opts ...RequestOption,
) ([]*Task, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Task *Task `json:"Task"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	tasks := make([]*Task, 0, len(items))
	for _, item := range items {
		tasks = append(tasks, item.Task)
	}

	return tasks, resp, nil
}

func (s *taskService) GetTasksCount(
	ctx context.Context, request *GetTasksCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response CountResponse
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

func (s *taskService) UpdateTask(
	ctx context.Context, request *UpdateTaskRequest, opts ...RequestOption,
) (*Task, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tasks", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		Task *Task `json:"Task"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.Task, resp, nil
}

type rawTaskFieldsInfo map[string]struct {
	Name         string                      `json:"name,omitempty"`      // name
	HTMLType     TaskFieldsInfoHTMLType      `json:"html_type,omitempty"` // 类型
	Label        string                      `json:"label,omitempty"`     // 中文名称
	Options      any                         `json:"options,omitempty"`   // 候选值
	ColorOptions []TaskFieldsInfoColorOption `json:"color_options,omitempty"`
	PureOptions  []TaskFieldsInfoPureOption  `json:"pure_options,omitempty"`
}

func (s *taskService) GetTaskFieldsInfo(
	ctx context.Context, request *GetTaskFieldsInfoRequest, opts ...RequestOption,
) ([]*TaskFieldsInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tasks/get_fields_info", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var raw rawTaskFieldsInfo
	resp, err := s.client.Do(req, &raw)
	if err != nil {
		return nil, resp, err
	}

	fields := make([]*TaskFieldsInfo, 0, len(raw))
	for _, item := range raw {
		options := make([]TaskFieldsInfoOption, 0)

		if item.Options != nil {
			if os, ok := item.Options.(map[string]any); ok {
				options = make([]TaskFieldsInfoOption, 0, len(os))
				for key, value := range os {
					if v, ok2 := value.(string); ok2 {
						options = append(options, TaskFieldsInfoOption{
							Value: key,
							Label: v,
						})
					}
				}
			}
		}

		fields = append(fields, &TaskFieldsInfo{
			Name:         item.Name,
			HTMLType:     item.HTMLType,
			Label:        item.Label,
			Options:      options,
			ColorOptions: item.ColorOptions,
			PureOptions:  item.PureOptions,
		})
	}

	return fields, resp, nil
}
