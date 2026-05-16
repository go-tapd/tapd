package tapd

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

// TestCaseStatus 测试用例状态
type TestCaseStatus string

const (
	TestCaseStatusNormal   TestCaseStatus = "normal"   // 正常
	TestCaseStatusUpdating TestCaseStatus = "updating" // 待更新
	TestCaseStatusAbandon  TestCaseStatus = "abandon"  // 已废弃
)

func (s TestCaseStatus) String() string {
	return string(s)
}

// TestCaseResultStatus 测试用例执行结果
type TestCaseResultStatus string

const (
	TestCaseResultStatusPass       TestCaseResultStatus = "pass"       // 通过
	TestCaseResultStatusNoPass     TestCaseResultStatus = "no_pass"    // 不通过
	TestCaseResultStatusBlock      TestCaseResultStatus = "block"      // 阻塞
	TestCaseResultStatusUnexecuted TestCaseResultStatus = "unexecuted" // 未执行
)

func (s TestCaseResultStatus) String() string {
	return string(s)
}

// TestCaseFieldsInfoHTMLType 测试用例字段控件类型。
type TestCaseFieldsInfoHTMLType string

const (
	TestCaseFieldsInfoHTMLTypeInput       TestCaseFieldsInfoHTMLType = "input"
	TestCaseFieldsInfoHTMLTypeSelect      TestCaseFieldsInfoHTMLType = "select"
	TestCaseFieldsInfoHTMLTypeQMEditor    TestCaseFieldsInfoHTMLType = "qmeditor"
	TestCaseFieldsInfoHTMLTypeDialog      TestCaseFieldsInfoHTMLType = "dialog"
	TestCaseFieldsInfoHTMLTypeUserChooser TestCaseFieldsInfoHTMLType = "user_chooser"
	TestCaseFieldsInfoHTMLTypeDateInput   TestCaseFieldsInfoHTMLType = "dateinput"
	TestCaseFieldsInfoHTMLTypeMultiSelect TestCaseFieldsInfoHTMLType = "multi_select"
	TestCaseFieldsInfoHTMLTypeText        TestCaseFieldsInfoHTMLType = "text"
)

type (
	// TestCase 测试用例
	TestCase struct {
		ID                 string          `json:"id,omitempty"`                  // 测试用例ID
		MID                string          `json:"mid,omitempty"`                 // 测试用例MID
		Steps              *string         `json:"steps,omitempty"`               // 用例步骤
		WorkspaceID        string          `json:"workspace_id,omitempty"`        // 项目ID
		CategoryID         string          `json:"category_id,omitempty"`         // 用例目录
		Version            string          `json:"version,omitempty"`             // 版本
		Created            string          `json:"created,omitempty"`             // 创建时间
		Modifier           string          `json:"modifier,omitempty"`            // 最后修改人
		Modified           string          `json:"modified,omitempty"`            // 最后修改时间
		Creator            string          `json:"creator,omitempty"`             // 创建人
		Status             TestCaseStatus  `json:"status,omitempty"`              // 用例状态
		Name               string          `json:"name,omitempty"`                // 用例名称
		Precondition       *string         `json:"precondition,omitempty"`        // 前置条件
		Expectation        *string         `json:"expectation,omitempty"`         // 预期结果
		Sort               string          `json:"sort,omitempty"`                // 排序
		IndexCode          string          `json:"indexcode,omitempty"`           // 编号
		Type               string          `json:"type,omitempty"`                // 用例类型
		Priority           string          `json:"priority,omitempty"`            // 用例等级
		IsAutomated        string          `json:"is_automated,omitempty"`        // 是否自动化
		AutomationType     string          `json:"automation_type,omitempty"`     // 自动化类型
		AutomationPlatform string          `json:"automation_platform,omitempty"` // 自动化平台
		IsServing          string          `json:"is_serving,omitempty"`          // 是否服务中
		TemplateID         string          `json:"template_id,omitempty"`         // 模板ID
		CreatedFrom        string          `json:"created_from,omitempty"`        // 创建来源
		Result             *TestCaseResult `json:"TcaseResult,omitempty"`         // 测试结果
		CustomField1       *string         `json:"custom_field_1,omitempty"`      // 自定义字段
		CustomField2       *string         `json:"custom_field_2,omitempty"`
		CustomField3       *string         `json:"custom_field_3,omitempty"`
		CustomField4       *string         `json:"custom_field_4,omitempty"`
		CustomField5       *string         `json:"custom_field_5,omitempty"`
		CustomField6       *string         `json:"custom_field_6,omitempty"`
		CustomField7       *string         `json:"custom_field_7,omitempty"`
		CustomField8       *string         `json:"custom_field_8,omitempty"`
		CustomField9       *string         `json:"custom_field_9,omitempty"`
		CustomField10      *string         `json:"custom_field_10,omitempty"`
		CustomField11      *string         `json:"custom_field_11,omitempty"`
		CustomField12      *string         `json:"custom_field_12,omitempty"`
		CustomField13      *string         `json:"custom_field_13,omitempty"`
		CustomField14      *string         `json:"custom_field_14,omitempty"`
		CustomField15      *string         `json:"custom_field_15,omitempty"`
		CustomField16      *string         `json:"custom_field_16,omitempty"`
		CustomField17      *string         `json:"custom_field_17,omitempty"`
		CustomField18      *string         `json:"custom_field_18,omitempty"`
		CustomField19      *string         `json:"custom_field_19,omitempty"`
		CustomField20      *string         `json:"custom_field_20,omitempty"`
		CustomField21      *string         `json:"custom_field_21,omitempty"`
		CustomField22      *string         `json:"custom_field_22,omitempty"`
		CustomField23      *string         `json:"custom_field_23,omitempty"`
		CustomField24      *string         `json:"custom_field_24,omitempty"`
		CustomField25      *string         `json:"custom_field_25,omitempty"`
		CustomField26      *string         `json:"custom_field_26,omitempty"`
		CustomField27      *string         `json:"custom_field_27,omitempty"`
		CustomField28      *string         `json:"custom_field_28,omitempty"`
		CustomField29      *string         `json:"custom_field_29,omitempty"`
		CustomField30      *string         `json:"custom_field_30,omitempty"`
		CustomField31      *string         `json:"custom_field_31,omitempty"`
		CustomField32      *string         `json:"custom_field_32,omitempty"`
		CustomField33      *string         `json:"custom_field_33,omitempty"`
		CustomField34      *string         `json:"custom_field_34,omitempty"`
		CustomField35      *string         `json:"custom_field_35,omitempty"`
		CustomField36      *string         `json:"custom_field_36,omitempty"`
		CustomField37      *string         `json:"custom_field_37,omitempty"`
		CustomField38      *string         `json:"custom_field_38,omitempty"`
		CustomField39      *string         `json:"custom_field_39,omitempty"`
		CustomField40      *string         `json:"custom_field_40,omitempty"`
		CustomField41      *string         `json:"custom_field_41,omitempty"`
		CustomField42      *string         `json:"custom_field_42,omitempty"`
		CustomField43      *string         `json:"custom_field_43,omitempty"`
		CustomField44      *string         `json:"custom_field_44,omitempty"`
		CustomField45      *string         `json:"custom_field_45,omitempty"`
		CustomField46      *string         `json:"custom_field_46,omitempty"`
		CustomField47      *string         `json:"custom_field_47,omitempty"`
		CustomField48      *string         `json:"custom_field_48,omitempty"`
		CustomField49      *string         `json:"custom_field_49,omitempty"`
		CustomField50      *string         `json:"custom_field_50,omitempty"`
	}

	CreateTestCaseRequest struct {
		ID            *int64          `json:"id,omitempty"`             // 测试用例ID
		Steps         *string         `json:"steps,omitempty"`          // 用例步骤
		WorkspaceID   *int            `json:"workspace_id,omitempty"`   // [必须]项目ID
		CategoryID    *int64          `json:"category_id,omitempty"`    // 用例目录
		Status        *TestCaseStatus `json:"status,omitempty"`         // 用例状态
		Name          *string         `json:"name,omitempty"`           // [必须]用例名称
		Precondition  *string         `json:"precondition,omitempty"`   // 前置条件
		Expectation   *string         `json:"expectation,omitempty"`    // 预期结果
		Type          *string         `json:"type,omitempty"`           // 用例类型
		Priority      *string         `json:"priority,omitempty"`       // 用例等级
		Creator       *string         `json:"creator,omitempty"`        // 创建人
		CustomField1  *string         `json:"custom_field_1,omitempty"` // 自定义字段
		CustomField2  *string         `json:"custom_field_2,omitempty"`
		CustomField3  *string         `json:"custom_field_3,omitempty"`
		CustomField4  *string         `json:"custom_field_4,omitempty"`
		CustomField5  *string         `json:"custom_field_5,omitempty"`
		CustomField6  *string         `json:"custom_field_6,omitempty"`
		CustomField7  *string         `json:"custom_field_7,omitempty"`
		CustomField8  *string         `json:"custom_field_8,omitempty"`
		CustomField9  *string         `json:"custom_field_9,omitempty"`
		CustomField10 *string         `json:"custom_field_10,omitempty"`
		CustomField11 *string         `json:"custom_field_11,omitempty"`
		CustomField12 *string         `json:"custom_field_12,omitempty"`
		CustomField13 *string         `json:"custom_field_13,omitempty"`
		CustomField14 *string         `json:"custom_field_14,omitempty"`
		CustomField15 *string         `json:"custom_field_15,omitempty"`
		CustomField16 *string         `json:"custom_field_16,omitempty"`
		CustomField17 *string         `json:"custom_field_17,omitempty"`
		CustomField18 *string         `json:"custom_field_18,omitempty"`
		CustomField19 *string         `json:"custom_field_19,omitempty"`
		CustomField20 *string         `json:"custom_field_20,omitempty"`
		CustomField21 *string         `json:"custom_field_21,omitempty"`
		CustomField22 *string         `json:"custom_field_22,omitempty"`
		CustomField23 *string         `json:"custom_field_23,omitempty"`
		CustomField24 *string         `json:"custom_field_24,omitempty"`
		CustomField25 *string         `json:"custom_field_25,omitempty"`
		CustomField26 *string         `json:"custom_field_26,omitempty"`
		CustomField27 *string         `json:"custom_field_27,omitempty"`
		CustomField28 *string         `json:"custom_field_28,omitempty"`
		CustomField29 *string         `json:"custom_field_29,omitempty"`
		CustomField30 *string         `json:"custom_field_30,omitempty"`
		CustomField31 *string         `json:"custom_field_31,omitempty"`
		CustomField32 *string         `json:"custom_field_32,omitempty"`
		CustomField33 *string         `json:"custom_field_33,omitempty"`
		CustomField34 *string         `json:"custom_field_34,omitempty"`
		CustomField35 *string         `json:"custom_field_35,omitempty"`
		CustomField36 *string         `json:"custom_field_36,omitempty"`
		CustomField37 *string         `json:"custom_field_37,omitempty"`
		CustomField38 *string         `json:"custom_field_38,omitempty"`
		CustomField39 *string         `json:"custom_field_39,omitempty"`
		CustomField40 *string         `json:"custom_field_40,omitempty"`
		CustomField41 *string         `json:"custom_field_41,omitempty"`
		CustomField42 *string         `json:"custom_field_42,omitempty"`
		CustomField43 *string         `json:"custom_field_43,omitempty"`
		CustomField44 *string         `json:"custom_field_44,omitempty"`
		CustomField45 *string         `json:"custom_field_45,omitempty"`
		CustomField46 *string         `json:"custom_field_46,omitempty"`
		CustomField47 *string         `json:"custom_field_47,omitempty"`
		CustomField48 *string         `json:"custom_field_48,omitempty"`
		CustomField49 *string         `json:"custom_field_49,omitempty"`
		CustomField50 *string         `json:"custom_field_50,omitempty"`
	}

	BatchCreateTestCasesRequest []*CreateTestCaseRequest

	CreateTestCaseCategoryRequest struct {
		WorkspaceID *int    `json:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string `json:"name,omitempty"`         // [必须]目录名称
		Description *string `json:"description,omitempty"`  // 目录描述
		ParentID    *int64  `json:"parent_id,omitempty"`    // 父目录ID
		Creator     *string `json:"creator,omitempty"`      // 目录创建人
	}

	TestCaseCategory struct {
		ID          string  `json:"id,omitempty"`           // 目录ID
		WorkspaceID string  `json:"workspace_id,omitempty"` // 项目ID
		Name        string  `json:"name,omitempty"`         // 目录名称
		Description *string `json:"description,omitempty"`  // 目录描述
		ParentID    string  `json:"parent_id,omitempty"`    // 父目录ID
		Modified    string  `json:"modified,omitempty"`     // 最后修改时间
		Created     string  `json:"created,omitempty"`      // 创建时间
		Creator     *string `json:"creator,omitempty"`      // 目录创建人
		Modifier    *string `json:"modifier,omitempty"`     // 目录最后修改人
		Sorting     *string `json:"sorting,omitempty"`      // 目录排序序号
	}

	// TestPlan 测试计划
	TestPlan struct {
		ID            string  `json:"id,omitempty"`             // 测试计划ID
		WorkspaceID   string  `json:"workspace_id,omitempty"`   // 项目ID
		Name          string  `json:"name,omitempty"`           // 测试计划标题
		Description   string  `json:"description,omitempty"`    // 测试计划详细描述
		Version       string  `json:"version,omitempty"`        // 版本号
		Owner         string  `json:"owner,omitempty"`          // 测试计划负责人
		Status        string  `json:"status,omitempty"`         // 状态
		Type          string  `json:"type,omitempty"`           // 测试类型
		StartDate     *string `json:"start_date,omitempty"`     // 预计开始
		EndDate       *string `json:"end_date,omitempty"`       // 预计结束
		Creator       string  `json:"creator,omitempty"`        // 创建人
		Created       string  `json:"created,omitempty"`        // 创建时间
		Modified      string  `json:"modified,omitempty"`       // 最后修改时间
		Modifier      string  `json:"modifier,omitempty"`       // 修改人
		IterationID   string  `json:"iteration_id,omitempty"`   // 关联迭代ID
		TemplateID    string  `json:"template_id,omitempty"`    // 模板ID
		CreatedFrom   string  `json:"created_from,omitempty"`   // 创建来源
		CustomField1  *string `json:"custom_field_1,omitempty"` // 自定义字段
		CustomField2  *string `json:"custom_field_2,omitempty"`
		CustomField3  *string `json:"custom_field_3,omitempty"`
		CustomField4  *string `json:"custom_field_4,omitempty"`
		CustomField5  *string `json:"custom_field_5,omitempty"`
		CustomField6  *string `json:"custom_field_6,omitempty"`
		CustomField7  *string `json:"custom_field_7,omitempty"`
		CustomField8  *string `json:"custom_field_8,omitempty"`
		CustomField9  *string `json:"custom_field_9,omitempty"`
		CustomField10 *string `json:"custom_field_10,omitempty"`
	}

	CreateTestPlanRequest struct {
		Name          *string `json:"name,omitempty"`           // [必须]测试计划标题
		Description   *string `json:"description,omitempty"`    // 测试计划详细描述
		WorkspaceID   *int    `json:"workspace_id,omitempty"`   // [必须]项目ID
		Creator       *string `json:"creator,omitempty"`        // 创建人
		Modifier      *string `json:"modifier,omitempty"`       // 修改人
		Owner         *string `json:"owner,omitempty"`          // 测试计划负责人
		StartDate     *string `json:"start_date,omitempty"`     // 预计开始
		EndDate       *string `json:"end_date,omitempty"`       // 预计结束
		IterationID   *int64  `json:"iteration_id,omitempty"`   // 关联迭代ID
		Version       *string `json:"version,omitempty"`        // 版本号
		Type          *string `json:"type,omitempty"`           // 测试类型
		Status        *string `json:"status,omitempty"`         // 状态，默认open
		CustomField1  *string `json:"custom_field_1,omitempty"` // 自定义字段
		CustomField2  *string `json:"custom_field_2,omitempty"`
		CustomField3  *string `json:"custom_field_3,omitempty"`
		CustomField4  *string `json:"custom_field_4,omitempty"`
		CustomField5  *string `json:"custom_field_5,omitempty"`
		CustomField6  *string `json:"custom_field_6,omitempty"`
		CustomField7  *string `json:"custom_field_7,omitempty"`
		CustomField8  *string `json:"custom_field_8,omitempty"`
		CustomField9  *string `json:"custom_field_9,omitempty"`
		CustomField10 *string `json:"custom_field_10,omitempty"`
	}

	AssignTestCaseRequest struct {
		TestPlanID  *int64        `json:"test_plan_id,omitempty"` // [必须]测试计划ID
		TestCaseID  *Multi[int64] `json:"tcase_id,omitempty"`     // 用例ID，多个使用英文逗号分隔
		CategoryID  *int64        `json:"category_id,omitempty"`  // 用例目录ID
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		Executor    *string       `json:"executor,omitempty"`     // 执行人
		Assignee    *string       `json:"assignee,omitempty"`     // 负责人
	}

	CreateTestPlanStoryRelationRequest struct {
		PlanID      *int64        `json:"plan_id,omitempty"`      // [必须]测试计划ID
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		StoryIDs    *Multi[int64] `json:"story_ids,omitempty"`    // [必须]需求ID，多个使用英文逗号分隔
		Creator     *string       `json:"creator,omitempty"`      // [必须]创建人
	}

	CreateTestPlanTestCaseRelationRequest struct {
		TestPlanID  *int64        `json:"test_plan_id,omitempty"` // [必须]测试计划ID
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		TestCaseIDs *Multi[int64] `json:"tcase_ids,omitempty"`    // [必须]测试用例ID，多个使用英文逗号分隔
		Creator     *string       `json:"creator,omitempty"`      // [必须]创建人
	}

	DeleteTestPlanStoryRelationRequest struct {
		PlanID      *int64        `json:"plan_id,omitempty"`      // [必须]测试计划ID
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		StoryIDs    *Multi[int64] `json:"story_ids,omitempty"`    // [必须]需求ID，多个使用英文逗号分隔
		Creator     *string       `json:"creator,omitempty"`      // [必须]操作人
	}

	DeleteTestCaseStoryRelationRequest struct {
		WorkspaceID *int   `json:"workspace_id,omitempty"` // [必须]项目ID
		StoryID     *int64 `json:"story_id,omitempty"`     // [必须]需求ID
		TestCaseID  *int64 `json:"tcase_id,omitempty"`     // [必须]测试用例ID
		TestPlanID  *int64 `json:"test_plan_id,omitempty"` // [必须]测试计划ID
	}

	ExecuteTestCaseRequest struct {
		TestPlanID   *int64                `json:"test_plan_id,omitempty"`  // [必须]测试计划ID
		TestCaseID   *Multi[int64]         `json:"tcase_id,omitempty"`      // [必须]测试用例ID，支持批量执行
		WorkspaceID  *int                  `json:"workspace_id,omitempty"`  // [必须]项目ID
		ResultStatus *TestCaseResultStatus `json:"result_status,omitempty"` // [必须]执行结果
		LastExecutor *string               `json:"last_executor,omitempty"` // [必须]执行人
		ResultRemark *string               `json:"result_remark,omitempty"` // 实际执行结果
	}

	GetTestCaseRelatedStoriesRequest struct {
		WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
		TestCaseIDs *Multi[int64] `url:"tcase_ids,omitempty"`    // [必须]测试用例ID，多个使用英文逗号分隔
	}

	TestCaseRelatedStory struct {
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		TestCaseID  string `json:"tcase_id,omitempty"`     // 测试用例ID
		StoryID     string `json:"story_id,omitempty"`     // 需求ID
		TestPlanID  string `json:"test_plan_id,omitempty"` // 测试计划ID
	}

	GetTestCaseCategoriesRequest struct {
		ID          *Multi[int64]  `url:"id,omitempty"`           // 目录ID，支持多ID查询
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string        `url:"name,omitempty"`         // 目录名称，支持模糊匹配
		Description *string        `url:"description,omitempty"`  // 目录描述
		ParentID    *int64         `url:"parent_id,omitempty"`    // 父目录ID
		Modified    *string        `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		Created     *string        `url:"created,omitempty"`      // 创建时间，支持时间查询
		Creator     *string        `url:"creator,omitempty"`      // 目录创建人
		Modifier    *string        `url:"modifier,omitempty"`     // 目录最后修改人
		Sorting     *int           `url:"sorting,omitempty"`      // 目录排序序号
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetTestCaseCategoriesCountRequest struct {
		ID          *Multi[int64] `url:"id,omitempty"`           // 目录ID，支持多ID查询
		WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
		Name        *string       `url:"name,omitempty"`         // 目录名称，支持模糊匹配
		Description *string       `url:"description,omitempty"`  // 目录描述
		ParentID    *int64        `url:"parent_id,omitempty"`    // 父目录ID
		Modified    *string       `url:"modified,omitempty"`     // 最后修改时间，支持时间查询
		Created     *string       `url:"created,omitempty"`      // 创建时间，支持时间查询
		Creator     *string       `url:"creator,omitempty"`      // 目录创建人
		Modifier    *string       `url:"modifier,omitempty"`     // 目录最后修改人
		Sorting     *int          `url:"sorting,omitempty"`      // 目录排序序号
	}

	GetTestCaseCustomFieldsSettingsRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TestCaseCustomFieldsSetting struct {
		ID              string  `json:"id,omitempty"`           // 自定义字段配置的ID
		WorkspaceID     string  `json:"workspace_id,omitempty"` // 所属项目ID
		AppID           string  `json:"app_id,omitempty"`       // 应用ID
		EntryType       string  `json:"entry_type,omitempty"`   // 所属实体对象
		CustomField     string  `json:"custom_field,omitempty"` // 自定义字段标识
		Type            string  `json:"type,omitempty"`         // 输入类型
		Name            string  `json:"name,omitempty"`         // 自定义字段显示名称
		Options         *string `json:"options,omitempty"`      // 自定义字段可选值
		ExtraConfig     *string `json:"extra_config,omitempty"` // 额外配置
		Enabled         string  `json:"enabled,omitempty"`      // 是否启用
		Freeze          string  `json:"freeze,omitempty"`       // 是否冻结
		Sort            *string `json:"sort,omitempty"`         // 显示时排序系数
		Memo            *string `json:"memo,omitempty"`         // 备注
		OpenExtensionID string  `json:"open_extension_id,omitempty"`
		IsOut           int     `json:"is_out,omitempty"`
		IsUninstall     int     `json:"is_uninstall,omitempty"`
		AppName         string  `json:"app_name,omitempty"`
	}

	GetTestCaseFieldsInfoRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TestCaseFieldsInfoOption struct {
		Key   string `json:"key,omitempty"`   // 英文Key
		Label string `json:"label,omitempty"` // 中文名称
	}

	TestCaseFieldsInfo struct {
		Name     string                     `json:"name,omitempty"`      // 字段名
		HTMLType TestCaseFieldsInfoHTMLType `json:"html_type,omitempty"` // 类型
		Label    string                     `json:"label,omitempty"`     // 中文名称
		Options  []TestCaseFieldsInfoOption `json:"options,omitempty"`   // 候选值
		Default  string                     `json:"default,omitempty"`   // 默认值
	}

	GetTestCaseResultsRequest struct {
		TestPlanID  *int64 `url:"test_plan_id,omitempty"` // [必须]测试计划ID
		TestCaseID  *int64 `url:"tcase_id,omitempty"`     // [必须]测试用例ID
		WorkspaceID *int   `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TestCaseResult struct {
		ID           string               `json:"id,omitempty"`            // 测试结果ID
		TestCaseID   string               `json:"tcase_id,omitempty"`      // 测试用例ID
		Created      string               `json:"created,omitempty"`       // 创建时间
		TestPlanID   string               `json:"test_plan_id,omitempty"`  // 测试计划ID
		ResultStatus TestCaseResultStatus `json:"result_status,omitempty"` // 结果状态
		ResultRemark *string              `json:"result_remark,omitempty"` // 结果备注
		WorkspaceID  string               `json:"workspace_id,omitempty"`  // 项目ID
		Status       string               `json:"status,omitempty"`        // 状态
		Executor     string               `json:"executor,omitempty"`      // 执行人
		ExecutedAt   string               `json:"executed_at,omitempty"`   // 执行时间
		BugID        []string             `json:"bug_id,omitempty"`        // 关联缺陷ID
		Bugs         []*Bug               `json:"Bug,omitempty"`           // 关联缺陷详情
	}

	TestCaseResultItem struct {
		ID     string          `json:"id,omitempty"`     // 测试结果ID
		Result *TestCaseResult `json:"result,omitempty"` // 测试结果
	}

	GetTestCasesRequest struct {
		ID                 *Multi[int64]         `url:"id,omitempty"`                  // 用例ID，支持多ID查询
		Steps              *string               `url:"steps,omitempty"`               // 用例步骤
		WorkspaceID        *int                  `url:"workspace_id,omitempty"`        // [必须]项目ID
		CategoryID         *int64                `url:"category_id,omitempty"`         // 用例目录
		Created            *string               `url:"created,omitempty"`             // 创建时间，支持时间查询
		Modifier           *string               `url:"modifier,omitempty"`            // 最后修改人
		Modified           *string               `url:"modified,omitempty"`            // 最后修改时间，支持时间查询
		Creator            *string               `url:"creator,omitempty"`             // 创建人
		Status             *Enum[TestCaseStatus] `url:"status,omitempty"`              // 用例状态
		Name               *string               `url:"name,omitempty"`                // 用例名称，支持模糊匹配
		Precondition       *string               `url:"precondition,omitempty"`        // 前置条件
		Expectation        *string               `url:"expectation,omitempty"`         // 预期结果
		Type               *string               `url:"type,omitempty"`                // 用例类型
		Priority           *string               `url:"priority,omitempty"`            // 用例等级
		IsAutomated        *string               `url:"is_automated,omitempty"`        // 是否实现自动化
		AutomationType     *string               `url:"automation_type,omitempty"`     // 自动化测试类型
		AutomationPlatform *string               `url:"automation_platform,omitempty"` // 自动化测试平台
		IsServing          *string               `url:"is_serving,omitempty"`          // 是否上架
		Limit              *int                  `url:"limit,omitempty"`               // 设置返回数量限制，默认为30，最大取200
		Page               *int                  `url:"page,omitempty"`                // 返回当前数量限制下第N页的数据，默认为1
		Order              *Order                `url:"order,omitempty"`               // 排序规则
		Fields             *Multi[string]        `url:"fields,omitempty"`              // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetTestCasesCountRequest struct {
		ID                 *Multi[int64]         `url:"id,omitempty"`                  // 用例ID，支持多ID查询
		Steps              *string               `url:"steps,omitempty"`               // 用例步骤
		WorkspaceID        *int                  `url:"workspace_id,omitempty"`        // [必须]项目ID
		CategoryID         *int64                `url:"category_id,omitempty"`         // 用例目录
		Created            *string               `url:"created,omitempty"`             // 创建时间，支持时间查询
		Modifier           *string               `url:"modifier,omitempty"`            // 最后修改人
		Modified           *string               `url:"modified,omitempty"`            // 最后修改时间，支持时间查询
		Creator            *string               `url:"creator,omitempty"`             // 创建人
		Status             *Enum[TestCaseStatus] `url:"status,omitempty"`              // 用例状态
		Name               *string               `url:"name,omitempty"`                // 用例名称，支持模糊匹配
		Precondition       *string               `url:"precondition,omitempty"`        // 前置条件
		Expectation        *string               `url:"expectation,omitempty"`         // 预期结果
		Type               *string               `url:"type,omitempty"`                // 用例类型
		Priority           *string               `url:"priority,omitempty"`            // 用例等级
		IsAutomated        *string               `url:"is_automated,omitempty"`        // 是否实现自动化
		AutomationType     *string               `url:"automation_type,omitempty"`     // 自动化测试类型
		AutomationPlatform *string               `url:"automation_platform,omitempty"` // 自动化测试平台
		IsServing          *string               `url:"is_serving,omitempty"`          // 是否上架
		TestPlanID         *int64                `url:"test_plan_id,omitempty"`        // 测试计划ID
	}

	GetTestPlanRelatedBugsRequest struct {
		ID          *int64         `url:"id,omitempty"`           // [必须]测试计划ID
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	TestPlanRelatedBug struct {
		ID                        string                `json:"id,omitempty"`                       // 测试用例ID
		Name                      string                `json:"name,omitempty"`                     // 测试用例名称
		TestCaseResultRelatedBugs []*TestCaseResultItem `json:"tcase_result_relate_bugs,omitempty"` // 测试结果关联缺陷
	}

	GetIterationTestPlansRequest struct {
		WorkspaceID *int   `url:"workspace_id,omitempty"` // [必须]项目ID
		IterationID *int64 `url:"iteration_id,omitempty"` // [必须]迭代ID
		Limit       *int   `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
		Page        *int   `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order `url:"order,omitempty"`        // 排序规则
	}

	IterationTestPlan struct {
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		IterationID string `json:"iteration_id,omitempty"` // 迭代ID
		TestPlanID  string `json:"test_plan_id,omitempty"` // 测试计划ID
	}

	GetTestPlanResultRequest struct {
		WorkspaceID   *int    `url:"workspace_id,omitempty"`   // [必须]项目ID
		ID            *int64  `url:"id,omitempty"`             // [必须]测试计划ID
		LastExecutor  *string `url:"last_executor,omitempty"`  // 最后执行人
		IncludeRepeat *int    `url:"include_repeat,omitempty"` // 传1可以获取所有数据
	}

	GetTestPlanProgressRequest struct {
		ID          *int64 `url:"id,omitempty"`           // [必须]测试计划ID
		WorkspaceID *int   `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TestPlanProgress struct {
		StoryCount    int                          `json:"story_count,omitempty"`    // 需求数量
		TestCaseCount int                          `json:"tcase_count,omitempty"`    // 测试用例数量
		StatusCounter map[TestCaseResultStatus]int `json:"status_counter,omitempty"` // 各结果状态数量
		ExecutedRate  string                       `json:"executed_rate,omitempty"`  // 执行率
	}

	GetTestPlanTestCaseRelationsRequest struct {
		WorkspaceID *int   `url:"workspace_id,omitempty"` // [必须]项目ID
		TestPlanID  *int64 `url:"test_plan_id,omitempty"` // [必须]测试计划ID
		Limit       *int   `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
		Page        *int   `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order `url:"order,omitempty"`        // 排序规则
	}

	TestPlanTestCaseRelation struct {
		ID          string `json:"id,omitempty"`           // 关系ID
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		TestPlanID  string `json:"test_plan_id,omitempty"` // 测试计划ID
		StoryID     string `json:"story_id,omitempty"`     // 需求ID
		TestCaseID  string `json:"tcase_id,omitempty"`     // 测试用例ID
		Sort        string `json:"sort,omitempty"`         // 排序
		Creator     string `json:"creator,omitempty"`      // 创建人
		Created     string `json:"created,omitempty"`      // 创建时间
	}

	GetTestPlansRequest struct {
		ID          *Multi[int64]  `url:"id,omitempty"`           // 测试计划ID，支持多ID查询
		Name        *string        `url:"name,omitempty"`         // 测试计划标题
		Description *string        `url:"description,omitempty"`  // 测试计划详细描述
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
		Creator     *string        `url:"creator,omitempty"`      // 创建人
		Created     *string        `url:"created,omitempty"`      // 创建时间
		Modifier    *string        `url:"modifier,omitempty"`     // 修改人
		Modified    *string        `url:"modified,omitempty"`     // 最后修改时间
		Owner       *string        `url:"owner,omitempty"`        // 测试计划负责人
		StartDate   *string        `url:"start_date,omitempty"`   // 预计开始
		EndDate     *string        `url:"end_date,omitempty"`     // 预计结束
		IterationID *int64         `url:"iteration_id,omitempty"` // 关联迭代ID
		Version     *string        `url:"version,omitempty"`      // 版本号
		Type        *string        `url:"type,omitempty"`         // 测试类型
		Status      *string        `url:"status,omitempty"`       // 状态，默认open
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30，最大取200
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1
		Order       *Order         `url:"order,omitempty"`        // 排序规则
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}

	GetTestPlansCountRequest struct {
		ID          *Multi[int64] `url:"id,omitempty"`           // 测试计划ID，支持多ID查询
		Name        *string       `url:"name,omitempty"`         // 测试计划标题
		Description *string       `url:"description,omitempty"`  // 测试计划详细描述
		WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
		Creator     *string       `url:"creator,omitempty"`      // 创建人
		Modifier    *string       `url:"modifier,omitempty"`     // 修改人
		Owner       *string       `url:"owner,omitempty"`        // 测试计划负责人
		StartDate   *string       `url:"start_date,omitempty"`   // 预计开始
		EndDate     *string       `url:"end_date,omitempty"`     // 预计结束
		IterationID *int64        `url:"iteration_id,omitempty"` // 关联迭代ID
		Version     *string       `url:"version,omitempty"`      // 版本号
		Type        *string       `url:"type,omitempty"`         // 测试类型
		Status      *string       `url:"status,omitempty"`       // 状态，默认open
	}

	RemoveTestCaseFromTestPlanRequest struct {
		TestPlanID  *int64        `json:"test_plan_id,omitempty"` // [必须]测试计划ID
		WorkspaceID *int          `json:"workspace_id,omitempty"` // [必须]项目ID
		StoryID     *int64        `json:"story_id,omitempty"`     // 需求ID
		TestCaseID  *Multi[int64] `json:"tcase_id,omitempty"`     // [必须]测试用例ID，多个使用英文逗号分隔
	}

	UpdateTestCaseRequest = CreateTestCaseRequest

	UpdateTestPlanRequest struct {
		ID            *int64  `json:"id,omitempty"`             // [必须]测试计划ID
		Name          *string `json:"name,omitempty"`           // 测试计划标题
		Description   *string `json:"description,omitempty"`    // 测试计划详细描述
		WorkspaceID   *int    `json:"workspace_id,omitempty"`   // [必须]项目ID
		Modifier      *string `json:"modifier,omitempty"`       // 修改人
		Owner         *string `json:"owner,omitempty"`          // 测试计划负责人
		StartDate     *string `json:"start_date,omitempty"`     // 预计开始
		EndDate       *string `json:"end_date,omitempty"`       // 预计结束
		Version       *string `json:"version,omitempty"`        // 版本号
		Type          *string `json:"type,omitempty"`           // 测试类型
		Status        *string `json:"status,omitempty"`         // 状态，默认open
		TemplateID    *int64  `json:"template_id,omitempty"`    // 模板ID
		CustomField1  *string `json:"custom_field_1,omitempty"` // 自定义字段
		CustomField2  *string `json:"custom_field_2,omitempty"`
		CustomField3  *string `json:"custom_field_3,omitempty"`
		CustomField4  *string `json:"custom_field_4,omitempty"`
		CustomField5  *string `json:"custom_field_5,omitempty"`
		CustomField6  *string `json:"custom_field_6,omitempty"`
		CustomField7  *string `json:"custom_field_7,omitempty"`
		CustomField8  *string `json:"custom_field_8,omitempty"`
		CustomField9  *string `json:"custom_field_9,omitempty"`
		CustomField10 *string `json:"custom_field_10,omitempty"`
	}

	GetTestPlanFieldsInfoRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
	}

	TestPlanFieldsInfo struct {
		Name     string                     `json:"name,omitempty"`      // 字段名
		HTMLType TestCaseFieldsInfoHTMLType `json:"html_type,omitempty"` // 类型
		Label    string                     `json:"label,omitempty"`     // 中文名称
		Options  []TestCaseFieldsInfoOption `json:"options,omitempty"`   // 候选值
	}
)

// TestService 测试
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/
type TestService interface {
	// CreateTestCase 创建测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase.html
	CreateTestCase(ctx context.Context, request *CreateTestCaseRequest, opts ...RequestOption) (*TestCase, *Response, error)

	// BatchCreateTestCases 批量创建测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/batch_add_tcase.html
	BatchCreateTestCases(
		ctx context.Context, request *BatchCreateTestCasesRequest, opts ...RequestOption,
	) ([]*TestCase, *Response, error)

	// CreateTestCaseCategory 创建测试用例目录
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_tcase_category.html
	CreateTestCaseCategory(
		ctx context.Context, request *CreateTestCaseCategoryRequest, opts ...RequestOption,
	) (*TestCaseCategory, *Response, error)

	// CreateTestPlan 创建测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/add_test_plan.html
	CreateTestPlan(ctx context.Context, request *CreateTestPlanRequest, opts ...RequestOption) (*TestPlan, *Response, error)

	// AssignTestCase 分配测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/assign_tcase_instance.html
	AssignTestCase(ctx context.Context, request *AssignTestCaseRequest, opts ...RequestOption) (bool, *Response, error)

	// CreateTestPlanStoryRelation 创建测试计划和需求关联关系
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/create_story_relation.html
	CreateTestPlanStoryRelation(
		ctx context.Context, request *CreateTestPlanStoryRelationRequest, opts ...RequestOption,
	) (bool, *Response, error)

	// CreateTestPlanTestCaseRelation 创建测试计划和测试用例关联关系
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/create_tcase_relation.html
	CreateTestPlanTestCaseRelation(
		ctx context.Context, request *CreateTestPlanTestCaseRelationRequest, opts ...RequestOption,
	) (bool, *Response, error)

	// DeleteTestPlanStoryRelation 解除测试计划和需求关联关系
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_story_relation.html
	DeleteTestPlanStoryRelation(
		ctx context.Context, request *DeleteTestPlanStoryRelationRequest, opts ...RequestOption,
	) (bool, *Response, error)

	// DeleteTestCaseStoryRelation 解除测试用例关联并移出测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/delete_tcase_story_relation.html
	DeleteTestCaseStoryRelation(
		ctx context.Context, request *DeleteTestCaseStoryRelationRequest, opts ...RequestOption,
	) (bool, *Response, error)

	// ExecuteTestCase 执行测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/execute_tcase_instance.html
	ExecuteTestCase(ctx context.Context, request *ExecuteTestCaseRequest, opts ...RequestOption) (bool, *Response, error)

	// GetTestCaseRelatedStories 获取测试用例关联的需求
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_story_by_tcase_id.html
	GetTestCaseRelatedStories(
		ctx context.Context, request *GetTestCaseRelatedStoriesRequest, opts ...RequestOption,
	) ([]*TestCaseRelatedStory, *Response, error)

	// GetTestCaseCategories 获取测试用例目录
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories.html
	GetTestCaseCategories(
		ctx context.Context, request *GetTestCaseCategoriesRequest, opts ...RequestOption,
	) ([]*TestCaseCategory, *Response, error)

	// GetTestCaseCategoriesCount 获取测试用例目录数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_categories_count.html
	GetTestCaseCategoriesCount(
		ctx context.Context, request *GetTestCaseCategoriesCountRequest, opts ...RequestOption,
	) (int, *Response, error)

	// GetTestCaseCustomFieldsSettings 获取测试用例自定义字段配置
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_custom_fields_settings.html
	GetTestCaseCustomFieldsSettings(
		ctx context.Context, request *GetTestCaseCustomFieldsSettingsRequest, opts ...RequestOption,
	) ([]*TestCaseCustomFieldsSetting, *Response, error)

	// GetTestCaseFieldsInfo 获取测试用例字段所有字段及候选值
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_fields_info.html
	GetTestCaseFieldsInfo(
		ctx context.Context, request *GetTestCaseFieldsInfoRequest, opts ...RequestOption,
	) ([]*TestCaseFieldsInfo, *Response, error)

	// GetTestCaseResults 获取测试用例执行结果
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcase_result.html
	GetTestCaseResults(
		ctx context.Context, request *GetTestCaseResultsRequest, opts ...RequestOption,
	) ([]*TestCaseResultItem, *Response, error)

	// GetTestCases 获取测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases.html
	GetTestCases(ctx context.Context, request *GetTestCasesRequest, opts ...RequestOption) ([]*TestCase, *Response, error)

	// GetTestCasesCount 获取测试用例数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_tcases_count.html
	GetTestCasesCount(ctx context.Context, request *GetTestCasesCountRequest, opts ...RequestOption) (int, *Response, error)

	// GetTestPlanRelatedBugs 获取测试计划关联bug
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_bugs.html
	GetTestPlanRelatedBugs(
		ctx context.Context, request *GetTestPlanRelatedBugsRequest, opts ...RequestOption,
	) ([]*TestPlanRelatedBug, *Response, error)

	// GetIterationTestPlans 获取迭代下测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_by_iteration_id.html
	GetIterationTestPlans(
		ctx context.Context, request *GetIterationTestPlansRequest, opts ...RequestOption,
	) ([]*IterationTestPlan, *Response, error)

	// GetTestPlanResult 获取测试计划测试结果
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_details.html
	GetTestPlanResult(
		ctx context.Context, request *GetTestPlanResultRequest, opts ...RequestOption,
	) ([]*TestCase, *Response, error)

	// GetTestPlanProgress 获取测试计划执行进度
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_progress.html
	GetTestPlanProgress(
		ctx context.Context, request *GetTestPlanProgressRequest, opts ...RequestOption,
	) (*TestPlanProgress, *Response, error)

	// GetTestPlanTestCaseRelations 获取测试计划与测试用例关联关系
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_tcases.html
	GetTestPlanTestCaseRelations(
		ctx context.Context, request *GetTestPlanTestCaseRelationsRequest, opts ...RequestOption,
	) ([]*TestPlanTestCaseRelation, *Response, error)

	// GetTestPlans 获取测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans.html
	GetTestPlans(ctx context.Context, request *GetTestPlansRequest, opts ...RequestOption) ([]*TestPlan, *Response, error)

	// GetTestPlansCount 获取测试计划数量
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plans_count.html
	GetTestPlansCount(ctx context.Context, request *GetTestPlansCountRequest, opts ...RequestOption) (int, *Response, error)

	// RemoveTestCaseFromTestPlan 测试用例移出测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/remove_tcase_instance.html
	RemoveTestCaseFromTestPlan(
		ctx context.Context, request *RemoveTestCaseFromTestPlanRequest, opts ...RequestOption,
	) (bool, *Response, error)

	// UpdateTestCase 更新测试用例
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/update_tcase.html
	UpdateTestCase(ctx context.Context, request *UpdateTestCaseRequest, opts ...RequestOption) (*TestCase, *Response, error)

	// UpdateTestPlan 编辑测试计划
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/update_test_plan.html
	UpdateTestPlan(ctx context.Context, request *UpdateTestPlanRequest, opts ...RequestOption) (*TestPlan, *Response, error)

	// GetTestPlanFieldsInfo 获取测试计划所有字段及候选值
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/tcase/get_test_plan_fields_info.html
	GetTestPlanFieldsInfo(
		ctx context.Context, request *GetTestPlanFieldsInfoRequest, opts ...RequestOption,
	) ([]*TestPlanFieldsInfo, *Response, error)
}

type testService struct {
	client *Client
}

var _ TestService = (*testService)(nil)

func NewTestService(client *Client) TestService {
	return &testService{
		client: client,
	}
}

func (s *testService) CreateTestCase(
	ctx context.Context, request *CreateTestCaseRequest, opts ...RequestOption,
) (*TestCase, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		TestCase *TestCase `json:"Tcase,omitempty"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.TestCase, resp, nil
}

func (s *testService) BatchCreateTestCases(
	ctx context.Context, request *BatchCreateTestCasesRequest, opts ...RequestOption,
) ([]*TestCase, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcases/batch_save", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		TestCase *TestCase `json:"Tcase,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	testCases := make([]*TestCase, 0, len(items))
	for _, item := range items {
		testCases = append(testCases, item.TestCase)
	}

	return testCases, resp, nil
}

func (s *testService) CreateTestCaseCategory(
	ctx context.Context, request *CreateTestCaseCategoryRequest, opts ...RequestOption,
) (*TestCaseCategory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcase_categories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		TestCaseCategory *TestCaseCategory `json:"TcaseCategory,omitempty"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.TestCaseCategory, resp, nil
}

func (s *testService) CreateTestPlan(
	ctx context.Context, request *CreateTestPlanRequest, opts ...RequestOption,
) (*TestPlan, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "test_plans", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		TestPlan *TestPlan `json:"TestPlan,omitempty"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.TestPlan, resp, nil
}

func (s *testService) AssignTestCase(
	ctx context.Context, request *AssignTestCaseRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcase_instance/assign", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) CreateTestPlanStoryRelation(
	ctx context.Context, request *CreateTestPlanStoryRelationRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "test_plans/create_story_relation", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) CreateTestPlanTestCaseRelation(
	ctx context.Context, request *CreateTestPlanTestCaseRelationRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "test_plans/create_tcase_relation", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) DeleteTestPlanStoryRelation(
	ctx context.Context, request *DeleteTestPlanStoryRelationRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "test_plans/delete_story_relation", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) DeleteTestCaseStoryRelation(
	ctx context.Context, request *DeleteTestCaseStoryRelationRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcase_instance/delete_tcase_story_relation", request, opts)
	if err != nil {
		return false, nil, err
	}

	var result bool
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return false, resp, err
	}

	return result, resp, nil
}

func (s *testService) ExecuteTestCase(
	ctx context.Context, request *ExecuteTestCaseRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcase_instance/execute", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) GetTestCaseRelatedStories(
	ctx context.Context, request *GetTestCaseRelatedStoriesRequest, opts ...RequestOption,
) ([]*TestCaseRelatedStory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcases/get_story_by_tcase_id", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var relations []*TestCaseRelatedStory
	resp, err := s.client.Do(req, &relations)
	if err != nil {
		return nil, resp, err
	}

	return relations, resp, nil
}

func (s *testService) GetTestCaseCategories(
	ctx context.Context, request *GetTestCaseCategoriesRequest, opts ...RequestOption,
) ([]*TestCaseCategory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcase_categories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		TestCaseCategory *TestCaseCategory `json:"TcaseCategory,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	categories := make([]*TestCaseCategory, 0, len(items))
	for _, item := range items {
		categories = append(categories, item.TestCaseCategory)
	}

	return categories, resp, nil
}

func (s *testService) GetTestCaseCategoriesCount(
	ctx context.Context, request *GetTestCaseCategoriesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcase_categories/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var count CountResponse
	resp, err := s.client.Do(req, &count)
	if err != nil {
		return 0, resp, err
	}

	return count.Count, resp, nil
}

func (s *testService) GetTestCaseCustomFieldsSettings(
	ctx context.Context, request *GetTestCaseCustomFieldsSettingsRequest, opts ...RequestOption,
) ([]*TestCaseCustomFieldsSetting, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcases/custom_fields_settings", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		CustomFieldConfig *TestCaseCustomFieldsSetting `json:"CustomFieldConfig,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	settings := make([]*TestCaseCustomFieldsSetting, 0, len(items))
	for _, item := range items {
		settings = append(settings, item.CustomFieldConfig)
	}

	return settings, resp, nil
}

type rawTestCaseFieldsInfo map[string]struct {
	Name     string                     `json:"name,omitempty"`      // 字段名
	HTMLType TestCaseFieldsInfoHTMLType `json:"html_type,omitempty"` // 类型
	Label    string                     `json:"label,omitempty"`     // 中文名称
	Options  any                        `json:"options,omitempty"`   // 候选值
	Default  string                     `json:"default,omitempty"`   // 默认值
}

func (s *testService) GetTestCaseFieldsInfo(
	ctx context.Context, request *GetTestCaseFieldsInfoRequest, opts ...RequestOption,
) ([]*TestCaseFieldsInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcases/get_fields_info", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var raw rawTestCaseFieldsInfo
	resp, err := s.client.Do(req, &raw)
	if err != nil {
		return nil, resp, err
	}

	fields := make([]*TestCaseFieldsInfo, 0, len(raw))
	for _, item := range raw {
		fields = append(fields, &TestCaseFieldsInfo{
			Name:     item.Name,
			HTMLType: item.HTMLType,
			Label:    item.Label,
			Options:  newTestCaseFieldsInfoOptions(item.Options),
			Default:  item.Default,
		})
	}

	return fields, resp, nil
}

func (s *testService) GetTestCaseResults(
	ctx context.Context, request *GetTestCaseResultsRequest, opts ...RequestOption,
) ([]*TestCaseResultItem, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcase_instance/result", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var resultMap map[string]*TestCaseResult
	resp, err := s.client.Do(req, &resultMap)
	if err != nil {
		return nil, resp, err
	}

	return newTestCaseResultItems(resultMap), resp, nil
}

func (s *testService) GetTestCases(
	ctx context.Context, request *GetTestCasesRequest, opts ...RequestOption,
) ([]*TestCase, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		TestCase *TestCase `json:"Tcase,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	testCases := make([]*TestCase, 0, len(items))
	for _, item := range items {
		testCases = append(testCases, item.TestCase)
	}

	return testCases, resp, nil
}

func (s *testService) GetTestCasesCount(
	ctx context.Context, request *GetTestCasesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "tcases/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var count CountResponse
	resp, err := s.client.Do(req, &count)
	if err != nil {
		return 0, resp, err
	}

	return count.Count, resp, nil
}

func (s *testService) GetTestPlanRelatedBugs(
	ctx context.Context, request *GetTestPlanRelatedBugsRequest, opts ...RequestOption,
) ([]*TestPlanRelatedBug, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/result_relation_bugs", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var relations []*TestPlanRelatedBug
	resp, err := s.client.Do(req, &relations)
	if err != nil {
		return nil, resp, err
	}

	return relations, resp, nil
}

func (s *testService) GetIterationTestPlans(
	ctx context.Context, request *GetIterationTestPlansRequest, opts ...RequestOption,
) ([]*IterationTestPlan, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/get_by_iteration_id", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var plans []*IterationTestPlan
	resp, err := s.client.Do(req, &plans)
	if err != nil {
		return nil, resp, err
	}

	return plans, resp, nil
}

func (s *testService) GetTestPlanResult(
	ctx context.Context, request *GetTestPlanResultRequest, opts ...RequestOption,
) ([]*TestCase, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/details", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		TestCase *TestCase `json:"Tcase,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	testCases := make([]*TestCase, 0, len(items))
	for _, item := range items {
		testCases = append(testCases, item.TestCase)
	}

	return testCases, resp, nil
}

func (s *testService) GetTestPlanProgress(
	ctx context.Context, request *GetTestPlanProgressRequest, opts ...RequestOption,
) (*TestPlanProgress, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/progress", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var progress TestPlanProgress
	resp, err := s.client.Do(req, &progress)
	if err != nil {
		return nil, resp, err
	}

	return &progress, resp, nil
}

func (s *testService) GetTestPlanTestCaseRelations(
	ctx context.Context, request *GetTestPlanTestCaseRelationsRequest, opts ...RequestOption,
) ([]*TestPlanTestCaseRelation, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/get_test_plan_tcase", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Relation *TestPlanTestCaseRelation `json:"TestPlanStoryTcaseRelation,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	relations := make([]*TestPlanTestCaseRelation, 0, len(items))
	for _, item := range items {
		relations = append(relations, item.Relation)
	}

	return relations, resp, nil
}

func (s *testService) GetTestPlans(
	ctx context.Context, request *GetTestPlansRequest, opts ...RequestOption,
) ([]*TestPlan, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		TestPlan *TestPlan `json:"TestPlan,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	plans := make([]*TestPlan, 0, len(items))
	for _, item := range items {
		plans = append(plans, item.TestPlan)
	}

	return plans, resp, nil
}

func (s *testService) GetTestPlansCount(
	ctx context.Context, request *GetTestPlansCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var count CountResponse
	resp, err := s.client.Do(req, &count)
	if err != nil {
		return 0, resp, err
	}

	return count.Count, resp, nil
}

func (s *testService) RemoveTestCaseFromTestPlan(
	ctx context.Context, request *RemoveTestCaseFromTestPlanRequest, opts ...RequestOption,
) (bool, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcase_instance/remove_tcase", request, opts)
	if err != nil {
		return false, nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return false, resp, err
	}

	return true, resp, nil
}

func (s *testService) UpdateTestCase(
	ctx context.Context, request *UpdateTestCaseRequest, opts ...RequestOption,
) (*TestCase, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "tcases", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		TestCase *TestCase `json:"Tcase,omitempty"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.TestCase, resp, nil
}

func (s *testService) UpdateTestPlan(
	ctx context.Context, request *UpdateTestPlanRequest, opts ...RequestOption,
) (*TestPlan, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "test_plans", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		TestPlan *TestPlan `json:"TestPlan,omitempty"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.TestPlan, resp, nil
}

type rawTestPlanFieldsInfo map[string]struct {
	HTMLType TestCaseFieldsInfoHTMLType `json:"html_type,omitempty"` // 类型
	Label    string                     `json:"label,omitempty"`     // 中文名称
	Options  any                        `json:"options,omitempty"`   // 候选值
}

func (s *testService) GetTestPlanFieldsInfo(
	ctx context.Context, request *GetTestPlanFieldsInfoRequest, opts ...RequestOption,
) ([]*TestPlanFieldsInfo, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "test_plans/get_fields_info", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var raw rawTestPlanFieldsInfo
	resp, err := s.client.Do(req, &raw)
	if err != nil {
		return nil, resp, err
	}

	fields := make([]*TestPlanFieldsInfo, 0, len(raw))
	for name, item := range raw {
		fields = append(fields, &TestPlanFieldsInfo{
			Name:     name,
			HTMLType: item.HTMLType,
			Label:    item.Label,
			Options:  newTestCaseFieldsInfoOptions(item.Options),
		})
	}

	return fields, resp, nil
}

func (r *TestCaseResult) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID           json.RawMessage      `json:"id,omitempty"`
		TestCaseID   json.RawMessage      `json:"tcase_id,omitempty"`
		Created      string               `json:"created,omitempty"`
		TestPlanID   json.RawMessage      `json:"test_plan_id,omitempty"`
		ResultStatus TestCaseResultStatus `json:"result_status,omitempty"`
		ResultRemark *string              `json:"result_remark,omitempty"`
		WorkspaceID  json.RawMessage      `json:"workspace_id,omitempty"`
		Status       string               `json:"status,omitempty"`
		Executor     string               `json:"executor,omitempty"`
		ExecutedAt   string               `json:"executed_at,omitempty"`
		BugID        []string             `json:"bug_id,omitempty"`
		Bugs         []*Bug               `json:"Bug,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*r = TestCaseResult{
		ID:           stringifyJSONRaw(raw.ID),
		TestCaseID:   stringifyJSONRaw(raw.TestCaseID),
		Created:      raw.Created,
		TestPlanID:   stringifyJSONRaw(raw.TestPlanID),
		ResultStatus: raw.ResultStatus,
		ResultRemark: raw.ResultRemark,
		WorkspaceID:  stringifyJSONRaw(raw.WorkspaceID),
		Status:       raw.Status,
		Executor:     raw.Executor,
		ExecutedAt:   raw.ExecutedAt,
		BugID:        raw.BugID,
		Bugs:         raw.Bugs,
	}

	return nil
}

func (p *TestPlanProgress) UnmarshalJSON(data []byte) error {
	var raw struct {
		StoryCount    int                                      `json:"story_count,omitempty"`
		TestCaseCount int                                      `json:"tcase_count,omitempty"`
		StatusCounter map[TestCaseResultStatus]json.RawMessage `json:"status_counter,omitempty"`
		ExecutedRate  string                                   `json:"executed_rate,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	counter := make(map[TestCaseResultStatus]int, len(raw.StatusCounter))
	for status, value := range raw.StatusCounter {
		count, err := parseJSONInt(value)
		if err != nil {
			return err
		}
		counter[status] = count
	}

	*p = TestPlanProgress{
		StoryCount:    raw.StoryCount,
		TestCaseCount: raw.TestCaseCount,
		StatusCounter: counter,
		ExecutedRate:  raw.ExecutedRate,
	}

	return nil
}

func (r *TestPlanRelatedBug) UnmarshalJSON(data []byte) error {
	var raw struct {
		ID                        json.RawMessage            `json:"id,omitempty"`
		Name                      string                     `json:"name,omitempty"`
		TestCaseResultRelatedBugs map[string]*TestCaseResult `json:"tcase_result_relate_bugs,omitempty"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*r = TestPlanRelatedBug{
		ID:                        stringifyJSONRaw(raw.ID),
		Name:                      raw.Name,
		TestCaseResultRelatedBugs: newTestCaseResultItems(raw.TestCaseResultRelatedBugs),
	}

	return nil
}

func newTestCaseResultItems(results map[string]*TestCaseResult) []*TestCaseResultItem {
	items := make([]*TestCaseResultItem, 0, len(results))
	for id, result := range results {
		items = append(items, &TestCaseResultItem{
			ID:     id,
			Result: result,
		})
	}

	return items
}

func newTestCaseFieldsInfoOptions(options any) []TestCaseFieldsInfoOption {
	os, ok := options.(map[string]any)
	if !ok {
		return nil
	}

	items := make([]TestCaseFieldsInfoOption, 0, len(os))
	for key, value := range os {
		if v, ok2 := value.(string); ok2 {
			items = append(items, TestCaseFieldsInfoOption{
				Key:   key,
				Label: v,
			})
		}
	}

	return items
}

func parseJSONInt(raw json.RawMessage) (int, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return 0, nil
	}
	if raw[0] == '"' {
		var value string
		if err := json.Unmarshal(raw, &value); err != nil {
			return 0, err
		}

		return strconv.Atoi(value)
	}

	var value int
	if err := json.Unmarshal(raw, &value); err != nil {
		return 0, err
	}

	return value, nil
}

func stringifyJSONRaw(raw json.RawMessage) string {
	if len(raw) == 0 || string(raw) == "null" {
		return ""
	}
	if raw[0] != '"' {
		return string(raw)
	}

	var value string
	_ = json.Unmarshal(raw, &value)
	return value
}
