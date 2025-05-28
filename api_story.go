package tapd

import (
	"context"
	"net/http"
)

type StoryStatus string

const (
	StoryStatusAudited    StoryStatus = "audited"     // 已评审
	StoryStatusDeveloping StoryStatus = "developing"  // 开发中
	StoryStatusForTest    StoryStatus = "for_test"    // 待测试
	StoryStatusInProgress StoryStatus = "in_progress" // 处理中
	StoryStatusPlanning   StoryStatus = "planning"    // 规划中
	StoryStatusRejected   StoryStatus = "rejected"    // 已拒绝
	StoryStatusResolved   StoryStatus = "resolved"    // 已处理
	StoryStatusTesting    StoryStatus = "testing"     // 测试中
)

func (s StoryStatus) String() string {
	return string(s)
}

type StoryService struct {
	client *Client
}

// CreateStory 创建需求
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/add_story.html
func (s *StoryService) CreateStory(
	ctx context.Context, request *CreateStoryRequest, opts ...RequestOption,
) (*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Story, resp, nil
}

type CreateStoryRequest struct {
	WorkspaceID     *int           `json:"workspace_id,omitempty"`     // [必须]项目ID
	Name            *string        `json:"name,omitempty"`             // [必须]标题
	Priority        *string        `json:"priority,omitempty"`         // 优先级
	PriorityLabel   *PriorityLabel `json:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue   *int           `json:"business_value,omitempty"`   // 业务价值
	Version         *string        `json:"version,omitempty"`          // 版本
	Module          *string        `json:"module,omitempty"`           // 模块
	TestFocus       *string        `json:"test_focus,omitempty"`       // 测试重点
	Size            *int           `json:"size,omitempty"`             // 规模
	Owner           *string        `json:"owner,omitempty"`            // 处理人
	CC              *string        `json:"cc,omitempty"`               // 抄送人
	Creator         *string        `json:"creator,omitempty"`          // 创建人
	Developer       *string        `json:"developer,omitempty"`        // 开发人员
	Begin           *string        `json:"begin,omitempty"`            // 预计开始
	Due             *string        `json:"due,omitempty"`              // 预计结束
	IterationID     *string        `json:"iteration_id,omitempty"`     // 迭代ID
	TemplatedID     *int           `json:"templated_id,omitempty"`     // 模板ID
	ParentID        *int           `json:"parent_id,omitempty"`        // 父需求ID
	Effort          *string        `json:"effort,omitempty"`           // 预估工时
	EffortCompleted *string        `json:"effort_completed,omitempty"` // 完成工时
	Remain          *float64       `json:"remain,omitempty"`           // 剩余工时
	Exceed          *float64       `json:"exceed,omitempty"`           // 超出工时
	CategoryID      *int           `json:"category_id,omitempty"`      // 需求分类
	WorkitemTypeID  *int           `json:"workitem_type_id,omitempty"` // 需求类别
	ReleaseID       *int           `json:"release_id,omitempty"`       // 发布计划
	Source          *string        `json:"source,omitempty"`           // 来源
	Type            *string        `json:"type,omitempty"`             // 类型
	Description     *string        `json:"description,omitempty"`      // 详细描述
	Label           *string        `json:"label,omitempty"`            // 标签，标签不存在时将自动创建，多个以英文坚线分格
}

// 创建需求分类
// 复制需求
// 获取需求与其它需求的所有关联关系

type Story struct {
	ID                string        `json:"id,omitempty"`
	WorkitemTypeID    string        `json:"workitem_type_id,omitempty"`
	Name              string        `json:"name,omitempty"`
	Description       string        `json:"description,omitempty"`
	WorkspaceID       string        `json:"workspace_id,omitempty"`
	Creator           string        `json:"creator,omitempty"`
	Created           string        `json:"created,omitempty"`
	Modified          string        `json:"modified,omitempty"`
	Status            string        `json:"status,omitempty"`
	Step              string        `json:"step,omitempty"`
	Owner             string        `json:"owner,omitempty"`
	Cc                string        `json:"cc,omitempty"`
	Begin             *string       `json:"begin,omitempty"`
	Due               *string       `json:"due,omitempty"`
	Size              *string       `json:"size,omitempty"`
	Priority          string        `json:"priority,omitempty"`
	Developer         string        `json:"developer,omitempty"`
	IterationID       string        `json:"iteration_id,omitempty"`
	TestFocus         string        `json:"test_focus,omitempty"`
	Type              string        `json:"type,omitempty"`
	Source            string        `json:"source,omitempty"`
	Module            string        `json:"module,omitempty"`
	Version           string        `json:"version,omitempty"`
	Completed         *string       `json:"completed,omitempty"`
	CategoryID        string        `json:"category_id,omitempty"`
	Path              string        `json:"path,omitempty"`
	ParentID          string        `json:"parent_id,omitempty"`
	ChildrenID        string        `json:"children_id,omitempty"`
	AncestorID        string        `json:"ancestor_id,omitempty"`
	Level             string        `json:"level,omitempty"`
	BusinessValue     *string       `json:"business_value,omitempty"`
	Effort            *string       `json:"effort,omitempty"`
	EffortCompleted   string        `json:"effort_completed,omitempty"`
	Exceed            string        `json:"exceed,omitempty"`
	Remain            string        `json:"remain,omitempty"`
	ReleaseID         string        `json:"release_id,omitempty"`
	BugID             string        `json:"bug_id,omitempty"`
	TemplatedID       string        `json:"templated_id,omitempty"`
	CreatedFrom       string        `json:"created_from,omitempty"`
	Feature           string        `json:"feature,omitempty"`
	Label             string        `json:"label,omitempty"`
	Progress          string        `json:"progress,omitempty"`
	IsArchived        string        `json:"is_archived,omitempty"`
	TechRisk          *string       `json:"tech_risk,omitempty"`
	Flows             *string       `json:"flows,omitempty"`
	SecretRootID      string        `json:"secret_root_id,omitempty"`
	PriorityLabel     PriorityLabel `json:"priority_label,omitempty"`
	CustomFieldOne    string        `json:"custom_field_one,omitempty"`
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
	CustomField51     string        `json:"custom_field_51,omitempty"`
	CustomField52     string        `json:"custom_field_52,omitempty"`
	CustomField53     string        `json:"custom_field_53,omitempty"`
	CustomField54     string        `json:"custom_field_54,omitempty"`
	CustomField55     string        `json:"custom_field_55,omitempty"`
	CustomField56     string        `json:"custom_field_56,omitempty"`
	CustomField57     string        `json:"custom_field_57,omitempty"`
	CustomField58     string        `json:"custom_field_58,omitempty"`
	CustomField59     string        `json:"custom_field_59,omitempty"`
	CustomField60     string        `json:"custom_field_60,omitempty"`
	CustomField61     string        `json:"custom_field_61,omitempty"`
	CustomField62     string        `json:"custom_field_62,omitempty"`
	CustomField63     string        `json:"custom_field_63,omitempty"`
	CustomField64     string        `json:"custom_field_64,omitempty"`
	CustomField65     string        `json:"custom_field_65,omitempty"`
	CustomField66     string        `json:"custom_field_66,omitempty"`
	CustomField67     string        `json:"custom_field_67,omitempty"`
	CustomField68     string        `json:"custom_field_68,omitempty"`
	CustomField69     string        `json:"custom_field_69,omitempty"`
	CustomField70     string        `json:"custom_field_70,omitempty"`
	CustomField71     string        `json:"custom_field_71,omitempty"`
	CustomField72     string        `json:"custom_field_72,omitempty"`
	CustomField73     string        `json:"custom_field_73,omitempty"`
	CustomField74     string        `json:"custom_field_74,omitempty"`
	CustomField75     string        `json:"custom_field_75,omitempty"`
	CustomField76     string        `json:"custom_field_76,omitempty"`
	CustomField77     string        `json:"custom_field_77,omitempty"`
	CustomField78     string        `json:"custom_field_78,omitempty"`
	CustomField79     string        `json:"custom_field_79,omitempty"`
	CustomField80     string        `json:"custom_field_80,omitempty"`
	CustomField81     string        `json:"custom_field_81,omitempty"`
	CustomField82     string        `json:"custom_field_82,omitempty"`
	CustomField83     string        `json:"custom_field_83,omitempty"`
	CustomField84     string        `json:"custom_field_84,omitempty"`
	CustomField85     string        `json:"custom_field_85,omitempty"`
	CustomField86     string        `json:"custom_field_86,omitempty"`
	CustomField87     string        `json:"custom_field_87,omitempty"`
	CustomField88     string        `json:"custom_field_88,omitempty"`
	CustomField89     string        `json:"custom_field_89,omitempty"`
	CustomField90     string        `json:"custom_field_90,omitempty"`
	CustomField91     string        `json:"custom_field_91,omitempty"`
	CustomField92     string        `json:"custom_field_92,omitempty"`
	CustomField93     string        `json:"custom_field_93,omitempty"`
	CustomField94     string        `json:"custom_field_94,omitempty"`
	CustomField95     string        `json:"custom_field_95,omitempty"`
	CustomField96     string        `json:"custom_field_96,omitempty"`
	CustomField97     string        `json:"custom_field_97,omitempty"`
	CustomField98     string        `json:"custom_field_98,omitempty"`
	CustomField99     string        `json:"custom_field_99,omitempty"`
	CustomField100    string        `json:"custom_field_100,omitempty"`
	CustomField101    string        `json:"custom_field_101,omitempty"`
	CustomField102    string        `json:"custom_field_102,omitempty"`
	CustomField103    string        `json:"custom_field_103,omitempty"`
	CustomField104    string        `json:"custom_field_104,omitempty"`
	CustomField105    string        `json:"custom_field_105,omitempty"`
	CustomField106    string        `json:"custom_field_106,omitempty"`
	CustomField107    string        `json:"custom_field_107,omitempty"`
	CustomField108    string        `json:"custom_field_108,omitempty"`
	CustomField109    string        `json:"custom_field_109,omitempty"`
	CustomField110    string        `json:"custom_field_110,omitempty"`
	CustomField111    string        `json:"custom_field_111,omitempty"`
	CustomField112    string        `json:"custom_field_112,omitempty"`
	CustomField113    string        `json:"custom_field_113,omitempty"`
	CustomField114    string        `json:"custom_field_114,omitempty"`
	CustomField115    string        `json:"custom_field_115,omitempty"`
	CustomField116    string        `json:"custom_field_116,omitempty"`
	CustomField117    string        `json:"custom_field_117,omitempty"`
	CustomField118    string        `json:"custom_field_118,omitempty"`
	CustomField119    string        `json:"custom_field_119,omitempty"`
	CustomField120    string        `json:"custom_field_120,omitempty"`
	CustomField121    string        `json:"custom_field_121,omitempty"`
	CustomField122    string        `json:"custom_field_122,omitempty"`
	CustomField123    string        `json:"custom_field_123,omitempty"`
	CustomField124    string        `json:"custom_field_124,omitempty"`
	CustomField125    string        `json:"custom_field_125,omitempty"`
	CustomField126    string        `json:"custom_field_126,omitempty"`
	CustomField127    string        `json:"custom_field_127,omitempty"`
	CustomField128    string        `json:"custom_field_128,omitempty"`
	CustomField129    string        `json:"custom_field_129,omitempty"`
	CustomField130    string        `json:"custom_field_130,omitempty"`
	CustomField131    string        `json:"custom_field_131,omitempty"`
	CustomField132    string        `json:"custom_field_132,omitempty"`
	CustomField133    string        `json:"custom_field_133,omitempty"`
	CustomField134    string        `json:"custom_field_134,omitempty"`
	CustomField135    string        `json:"custom_field_135,omitempty"`
	CustomField136    string        `json:"custom_field_136,omitempty"`
	CustomField137    string        `json:"custom_field_137,omitempty"`
	CustomField138    string        `json:"custom_field_138,omitempty"`
	CustomField139    string        `json:"custom_field_139,omitempty"`
	CustomField140    string        `json:"custom_field_140,omitempty"`
	CustomField141    string        `json:"custom_field_141,omitempty"`
	CustomField142    string        `json:"custom_field_142,omitempty"`
	CustomField143    string        `json:"custom_field_143,omitempty"`
	CustomField144    string        `json:"custom_field_144,omitempty"`
	CustomField145    string        `json:"custom_field_145,omitempty"`
	CustomField146    string        `json:"custom_field_146,omitempty"`
	CustomField147    string        `json:"custom_field_147,omitempty"`
	CustomField148    string        `json:"custom_field_148,omitempty"`
	CustomField149    string        `json:"custom_field_149,omitempty"`
	CustomField150    string        `json:"custom_field_150,omitempty"`
	CustomField151    string        `json:"custom_field_151,omitempty"`
	CustomField152    string        `json:"custom_field_152,omitempty"`
	CustomField153    string        `json:"custom_field_153,omitempty"`
	CustomField154    string        `json:"custom_field_154,omitempty"`
	CustomField155    string        `json:"custom_field_155,omitempty"`
	CustomField156    string        `json:"custom_field_156,omitempty"`
	CustomField157    string        `json:"custom_field_157,omitempty"`
	CustomField158    string        `json:"custom_field_158,omitempty"`
	CustomField159    string        `json:"custom_field_159,omitempty"`
	CustomField160    string        `json:"custom_field_160,omitempty"`
	CustomField161    string        `json:"custom_field_161,omitempty"`
	CustomField162    string        `json:"custom_field_162,omitempty"`
	CustomField163    string        `json:"custom_field_163,omitempty"`
	CustomField164    string        `json:"custom_field_164,omitempty"`
	CustomField165    string        `json:"custom_field_165,omitempty"`
	CustomField166    string        `json:"custom_field_166,omitempty"`
	CustomField167    string        `json:"custom_field_167,omitempty"`
	CustomField168    string        `json:"custom_field_168,omitempty"`
	CustomField169    string        `json:"custom_field_169,omitempty"`
	CustomField170    string        `json:"custom_field_170,omitempty"`
	CustomField171    string        `json:"custom_field_171,omitempty"`
	CustomField172    string        `json:"custom_field_172,omitempty"`
	CustomField173    string        `json:"custom_field_173,omitempty"`
	CustomField174    string        `json:"custom_field_174,omitempty"`
	CustomField175    string        `json:"custom_field_175,omitempty"`
	CustomField176    string        `json:"custom_field_176,omitempty"`
	CustomField177    string        `json:"custom_field_177,omitempty"`
	CustomField178    string        `json:"custom_field_178,omitempty"`
	CustomField179    string        `json:"custom_field_179,omitempty"`
	CustomField180    string        `json:"custom_field_180,omitempty"`
	CustomField181    string        `json:"custom_field_181,omitempty"`
	CustomField182    string        `json:"custom_field_182,omitempty"`
	CustomField183    string        `json:"custom_field_183,omitempty"`
	CustomField184    string        `json:"custom_field_184,omitempty"`
	CustomField185    string        `json:"custom_field_185,omitempty"`
	CustomField186    string        `json:"custom_field_186,omitempty"`
	CustomField187    string        `json:"custom_field_187,omitempty"`
	CustomField188    string        `json:"custom_field_188,omitempty"`
	CustomField189    string        `json:"custom_field_189,omitempty"`
	CustomField190    string        `json:"custom_field_190,omitempty"`
	CustomField191    string        `json:"custom_field_191,omitempty"`
	CustomField192    string        `json:"custom_field_192,omitempty"`
	CustomField193    string        `json:"custom_field_193,omitempty"`
	CustomField194    string        `json:"custom_field_194,omitempty"`
	CustomField195    string        `json:"custom_field_195,omitempty"`
	CustomField196    string        `json:"custom_field_196,omitempty"`
	CustomField197    string        `json:"custom_field_197,omitempty"`
	CustomField198    string        `json:"custom_field_198,omitempty"`
	CustomField199    string        `json:"custom_field_199,omitempty"`
	CustomField200    string        `json:"custom_field_200,omitempty"`
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
}

// GetStories 获取需求
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories.html
func (s *StoryService) GetStories(
	ctx context.Context, request *GetStoriesRequest, opts ...RequestOption,
) ([]*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	stories := make([]*Story, 0, len(items))
	for _, item := range items {
		stories = append(stories, item.Story)
	}

	return stories, resp, nil
}

type GetStoriesRequest struct {
	ID                *Multi[int64]      `url:"id,omitempty"`               // ID	支持多ID查询,多个ID用逗号分隔
	Name              *string            `url:"name,omitempty"`             // 标题	支持模糊匹配
	Priority          *string            `url:"priority,omitempty"`         // 优先级
	PriorityLabel     *PriorityLabel     `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue     *int               `url:"business_value,omitempty"`   // 业务价值
	Status            *Enum[StoryStatus] `url:"status,omitempty"`           // 状态	支持枚举查询
	VStatus           *string            `url:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	WithVStatus       *string            `url:"with_v_status,omitempty"`    // 值=1可以返回中文状态
	Label             *string            `url:"label,omitempty"`            // 标签查询	支持枚举查询
	WorkitemTypeID    *string            `url:"workitem_type_id,omitempty"` // 需求类别ID	支持枚举查询
	Version           *string            `url:"version,omitempty"`          // 版本
	Module            *string            `url:"module,omitempty"`           // 模块
	Feature           *string            `url:"feature,omitempty"`          // 特性
	TestFocus         *string            `url:"test_focus,omitempty"`       // 测试重点
	Size              *int               `url:"size,omitempty"`             // 规模
	Owner             *string            `url:"owner,omitempty"`            // 处理人	支持模糊匹配
	CC                *string            `url:"cc,omitempty"`               // 抄送人	支持模糊匹配
	Creator           *string            `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Developer         *string            `url:"developer,omitempty"`        // 开发人员
	Begin             *string            `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due               *string            `url:"due,omitempty"`              // 预计结束	支持时间查询
	Created           *string            `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified          *string            `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Completed         *string            `url:"completed,omitempty"`        // 完成时间	支持时间查询
	IterationID       *string            `url:"iteration_id,omitempty"`     // 迭代ID	支持不等于查询
	Effort            *string            `url:"effort,omitempty"`           // 预估工时
	EffortCompleted   *string            `url:"effort_completed,omitempty"` // 完成工时
	Remain            *float64           `url:"remain,omitempty"`           // 剩余工时
	Exceed            *float64           `url:"exceed,omitempty"`           // 超出工时
	CategoryID        *string            `url:"category_id,omitempty"`      // 需求分类	支持枚举查询
	ReleaseID         *string            `url:"release_id,omitempty"`       // 发布计划
	Source            *string            `url:"source,omitempty"`           // 需求来源
	Type              *string            `url:"type,omitempty"`             // 需求类型
	ParentID          *string            `url:"parent_id,omitempty"`        // 父需求
	ChildrenID        *string            `url:"children_id,omitempty"`      // 子需求	为空查询传：丨
	Description       *string            `url:"description,omitempty"`      // 详细描述	支持模糊匹配
	WorkspaceID       *int               `url:"workspace_id,omitempty"`     // 项目ID
	Limit             *int               `url:"limit,omitempty"`            // 设置返回数量限制，默认为30
	Page              *int               `url:"page,omitempty"`             // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order             *Order             `url:"order,omitempty"`            // 排序规则，规则：字段名 ASC或者DESC
	Fields            *Multi[string]     `url:"fields,omitempty"`           // 设置获取的字段，多个字段间以','逗号隔开
	CustomFieldOne    *string            `url:"custom_field_one,omitempty"`
	CustomFieldTwo    *string            `url:"custom_field_two,omitempty"`
	CustomFieldThree  *string            `url:"custom_field_three,omitempty"`
	CustomFieldFour   *string            `url:"custom_field_four,omitempty"`
	CustomFieldFive   *string            `url:"custom_field_five,omitempty"`
	CustomFieldSix    *string            `url:"custom_field_six,omitempty"`
	CustomFieldSeven  *string            `url:"custom_field_seven,omitempty"`
	CustomFieldEight  *string            `url:"custom_field_eight,omitempty"`
	CustomField9      *string            `url:"custom_field_9,omitempty"`
	CustomField10     *string            `url:"custom_field_10,omitempty"`
	CustomField11     *string            `url:"custom_field_11,omitempty"`
	CustomField12     *string            `url:"custom_field_12,omitempty"`
	CustomField13     *string            `url:"custom_field_13,omitempty"`
	CustomField14     *string            `url:"custom_field_14,omitempty"`
	CustomField15     *string            `url:"custom_field_15,omitempty"`
	CustomField16     *string            `url:"custom_field_16,omitempty"`
	CustomField17     *string            `url:"custom_field_17,omitempty"`
	CustomField18     *string            `url:"custom_field_18,omitempty"`
	CustomField19     *string            `url:"custom_field_19,omitempty"`
	CustomField20     *string            `url:"custom_field_20,omitempty"`
	CustomField21     *string            `url:"custom_field_21,omitempty"`
	CustomField22     *string            `url:"custom_field_22,omitempty"`
	CustomField23     *string            `url:"custom_field_23,omitempty"`
	CustomField24     *string            `url:"custom_field_24,omitempty"`
	CustomField25     *string            `url:"custom_field_25,omitempty"`
	CustomField26     *string            `url:"custom_field_26,omitempty"`
	CustomField27     *string            `url:"custom_field_27,omitempty"`
	CustomField28     *string            `url:"custom_field_28,omitempty"`
	CustomField29     *string            `url:"custom_field_29,omitempty"`
	CustomField30     *string            `url:"custom_field_30,omitempty"`
	CustomField31     *string            `url:"custom_field_31,omitempty"`
	CustomField32     *string            `url:"custom_field_32,omitempty"`
	CustomField33     *string            `url:"custom_field_33,omitempty"`
	CustomField34     *string            `url:"custom_field_34,omitempty"`
	CustomField35     *string            `url:"custom_field_35,omitempty"`
	CustomField36     *string            `url:"custom_field_36,omitempty"`
	CustomField37     *string            `url:"custom_field_37,omitempty"`
	CustomField38     *string            `url:"custom_field_38,omitempty"`
	CustomField39     *string            `url:"custom_field_39,omitempty"`
	CustomField40     *string            `url:"custom_field_40,omitempty"`
	CustomField41     *string            `url:"custom_field_41,omitempty"`
	CustomField42     *string            `url:"custom_field_42,omitempty"`
	CustomField43     *string            `url:"custom_field_43,omitempty"`
	CustomField44     *string            `url:"custom_field_44,omitempty"`
	CustomField45     *string            `url:"custom_field_45,omitempty"`
	CustomField46     *string            `url:"custom_field_46,omitempty"`
	CustomField47     *string            `url:"custom_field_47,omitempty"`
	CustomField48     *string            `url:"custom_field_48,omitempty"`
	CustomField49     *string            `url:"custom_field_49,omitempty"`
	CustomField50     *string            `url:"custom_field_50,omitempty"`
	CustomField51     *string            `url:"custom_field_51,omitempty"`
	CustomField52     *string            `url:"custom_field_52,omitempty"`
	CustomField53     *string            `url:"custom_field_53,omitempty"`
	CustomField54     *string            `url:"custom_field_54,omitempty"`
	CustomField55     *string            `url:"custom_field_55,omitempty"`
	CustomField56     *string            `url:"custom_field_56,omitempty"`
	CustomField57     *string            `url:"custom_field_57,omitempty"`
	CustomField58     *string            `url:"custom_field_58,omitempty"`
	CustomField59     *string            `url:"custom_field_59,omitempty"`
	CustomField60     *string            `url:"custom_field_60,omitempty"`
	CustomField61     *string            `url:"custom_field_61,omitempty"`
	CustomField62     *string            `url:"custom_field_62,omitempty"`
	CustomField63     *string            `url:"custom_field_63,omitempty"`
	CustomField64     *string            `url:"custom_field_64,omitempty"`
	CustomField65     *string            `url:"custom_field_65,omitempty"`
	CustomField66     *string            `url:"custom_field_66,omitempty"`
	CustomField67     *string            `url:"custom_field_67,omitempty"`
	CustomField68     *string            `url:"custom_field_68,omitempty"`
	CustomField69     *string            `url:"custom_field_69,omitempty"`
	CustomField70     *string            `url:"custom_field_70,omitempty"`
	CustomField71     *string            `url:"custom_field_71,omitempty"`
	CustomField72     *string            `url:"custom_field_72,omitempty"`
	CustomField73     *string            `url:"custom_field_73,omitempty"`
	CustomField74     *string            `url:"custom_field_74,omitempty"`
	CustomField75     *string            `url:"custom_field_75,omitempty"`
	CustomField76     *string            `url:"custom_field_76,omitempty"`
	CustomField77     *string            `url:"custom_field_77,omitempty"`
	CustomField78     *string            `url:"custom_field_78,omitempty"`
	CustomField79     *string            `url:"custom_field_79,omitempty"`
	CustomField80     *string            `url:"custom_field_80,omitempty"`
	CustomField81     *string            `url:"custom_field_81,omitempty"`
	CustomField82     *string            `url:"custom_field_82,omitempty"`
	CustomField83     *string            `url:"custom_field_83,omitempty"`
	CustomField84     *string            `url:"custom_field_84,omitempty"`
	CustomField85     *string            `url:"custom_field_85,omitempty"`
	CustomField86     *string            `url:"custom_field_86,omitempty"`
	CustomField87     *string            `url:"custom_field_87,omitempty"`
	CustomField88     *string            `url:"custom_field_88,omitempty"`
	CustomField89     *string            `url:"custom_field_89,omitempty"`
	CustomField90     *string            `url:"custom_field_90,omitempty"`
	CustomField91     *string            `url:"custom_field_91,omitempty"`
	CustomField92     *string            `url:"custom_field_92,omitempty"`
	CustomField93     *string            `url:"custom_field_93,omitempty"`
	CustomField94     *string            `url:"custom_field_94,omitempty"`
	CustomField95     *string            `url:"custom_field_95,omitempty"`
	CustomField96     *string            `url:"custom_field_96,omitempty"`
	CustomField97     *string            `url:"custom_field_97,omitempty"`
	CustomField98     *string            `url:"custom_field_98,omitempty"`
	CustomField99     *string            `url:"custom_field_99,omitempty"`
	CustomField100    *string            `url:"custom_field_100,omitempty"`
	CustomField101    *string            `url:"custom_field_101,omitempty"`
	CustomField102    *string            `url:"custom_field_102,omitempty"`
	CustomField103    *string            `url:"custom_field_103,omitempty"`
	CustomField104    *string            `url:"custom_field_104,omitempty"`
	CustomField105    *string            `url:"custom_field_105,omitempty"`
	CustomField106    *string            `url:"custom_field_106,omitempty"`
	CustomField107    *string            `url:"custom_field_107,omitempty"`
	CustomField108    *string            `url:"custom_field_108,omitempty"`
	CustomField109    *string            `url:"custom_field_109,omitempty"`
	CustomField110    *string            `url:"custom_field_110,omitempty"`
	CustomField111    *string            `url:"custom_field_111,omitempty"`
	CustomField112    *string            `url:"custom_field_112,omitempty"`
	CustomField113    *string            `url:"custom_field_113,omitempty"`
	CustomField114    *string            `url:"custom_field_114,omitempty"`
	CustomField115    *string            `url:"custom_field_115,omitempty"`
	CustomField116    *string            `url:"custom_field_116,omitempty"`
	CustomField117    *string            `url:"custom_field_117,omitempty"`
	CustomField118    *string            `url:"custom_field_118,omitempty"`
	CustomField119    *string            `url:"custom_field_119,omitempty"`
	CustomField120    *string            `url:"custom_field_120,omitempty"`
	CustomField121    *string            `url:"custom_field_121,omitempty"`
	CustomField122    *string            `url:"custom_field_122,omitempty"`
	CustomField123    *string            `url:"custom_field_123,omitempty"`
	CustomField124    *string            `url:"custom_field_124,omitempty"`
	CustomField125    *string            `url:"custom_field_125,omitempty"`
	CustomField126    *string            `url:"custom_field_126,omitempty"`
	CustomField127    *string            `url:"custom_field_127,omitempty"`
	CustomField128    *string            `url:"custom_field_128,omitempty"`
	CustomField129    *string            `url:"custom_field_129,omitempty"`
	CustomField130    *string            `url:"custom_field_130,omitempty"`
	CustomField131    *string            `url:"custom_field_131,omitempty"`
	CustomField132    *string            `url:"custom_field_132,omitempty"`
	CustomField133    *string            `url:"custom_field_133,omitempty"`
	CustomField134    *string            `url:"custom_field_134,omitempty"`
	CustomField135    *string            `url:"custom_field_135,omitempty"`
	CustomField136    *string            `url:"custom_field_136,omitempty"`
	CustomField137    *string            `url:"custom_field_137,omitempty"`
	CustomField138    *string            `url:"custom_field_138,omitempty"`
	CustomField139    *string            `url:"custom_field_139,omitempty"`
	CustomField140    *string            `url:"custom_field_140,omitempty"`
	CustomField141    *string            `url:"custom_field_141,omitempty"`
	CustomField142    *string            `url:"custom_field_142,omitempty"`
	CustomField143    *string            `url:"custom_field_143,omitempty"`
	CustomField144    *string            `url:"custom_field_144,omitempty"`
	CustomField145    *string            `url:"custom_field_145,omitempty"`
	CustomField146    *string            `url:"custom_field_146,omitempty"`
	CustomField147    *string            `url:"custom_field_147,omitempty"`
	CustomField148    *string            `url:"custom_field_148,omitempty"`
	CustomField149    *string            `url:"custom_field_149,omitempty"`
	CustomField150    *string            `url:"custom_field_150,omitempty"`
	CustomField151    *string            `url:"custom_field_151,omitempty"`
	CustomField152    *string            `url:"custom_field_152,omitempty"`
	CustomField153    *string            `url:"custom_field_153,omitempty"`
	CustomField154    *string            `url:"custom_field_154,omitempty"`
	CustomField155    *string            `url:"custom_field_155,omitempty"`
	CustomField156    *string            `url:"custom_field_156,omitempty"`
	CustomField157    *string            `url:"custom_field_157,omitempty"`
	CustomField158    *string            `url:"custom_field_158,omitempty"`
	CustomField159    *string            `url:"custom_field_159,omitempty"`
	CustomField160    *string            `url:"custom_field_160,omitempty"`
	CustomField161    *string            `url:"custom_field_161,omitempty"`
	CustomField162    *string            `url:"custom_field_162,omitempty"`
	CustomField163    *string            `url:"custom_field_163,omitempty"`
	CustomField164    *string            `url:"custom_field_164,omitempty"`
	CustomField165    *string            `url:"custom_field_165,omitempty"`
	CustomField166    *string            `url:"custom_field_166,omitempty"`
	CustomField167    *string            `url:"custom_field_167,omitempty"`
	CustomField168    *string            `url:"custom_field_168,omitempty"`
	CustomField169    *string            `url:"custom_field_169,omitempty"`
	CustomField170    *string            `url:"custom_field_170,omitempty"`
	CustomField171    *string            `url:"custom_field_171,omitempty"`
	CustomField172    *string            `url:"custom_field_172,omitempty"`
	CustomField173    *string            `url:"custom_field_173,omitempty"`
	CustomField174    *string            `url:"custom_field_174,omitempty"`
	CustomField175    *string            `url:"custom_field_175,omitempty"`
	CustomField176    *string            `url:"custom_field_176,omitempty"`
	CustomField177    *string            `url:"custom_field_177,omitempty"`
	CustomField178    *string            `url:"custom_field_178,omitempty"`
	CustomField179    *string            `url:"custom_field_179,omitempty"`
	CustomField180    *string            `url:"custom_field_180,omitempty"`
	CustomField181    *string            `url:"custom_field_181,omitempty"`
	CustomField182    *string            `url:"custom_field_182,omitempty"`
	CustomField183    *string            `url:"custom_field_183,omitempty"`
	CustomField184    *string            `url:"custom_field_184,omitempty"`
	CustomField185    *string            `url:"custom_field_185,omitempty"`
	CustomField186    *string            `url:"custom_field_186,omitempty"`
	CustomField187    *string            `url:"custom_field_187,omitempty"`
	CustomField188    *string            `url:"custom_field_188,omitempty"`
	CustomField189    *string            `url:"custom_field_189,omitempty"`
	CustomField190    *string            `url:"custom_field_190,omitempty"`
	CustomField191    *string            `url:"custom_field_191,omitempty"`
	CustomField192    *string            `url:"custom_field_192,omitempty"`
	CustomField193    *string            `url:"custom_field_193,omitempty"`
	CustomField194    *string            `url:"custom_field_194,omitempty"`
	CustomField195    *string            `url:"custom_field_195,omitempty"`
	CustomField196    *string            `url:"custom_field_196,omitempty"`
	CustomField197    *string            `url:"custom_field_197,omitempty"`
	CustomField198    *string            `url:"custom_field_198,omitempty"`
	CustomField199    *string            `url:"custom_field_199,omitempty"`
	CustomField200    *string            `url:"custom_field_200,omitempty"`
	CustomPlanField1  *string            `url:"custom_plan_field_1,omitempty"`
	CustomPlanField2  *string            `url:"custom_plan_field_2,omitempty"`
	CustomPlanField3  *string            `url:"custom_plan_field_3,omitempty"`
	CustomPlanField4  *string            `url:"custom_plan_field_4,omitempty"`
	CustomPlanField5  *string            `url:"custom_plan_field_5,omitempty"`
	CustomPlanField6  *string            `url:"custom_plan_field_6,omitempty"`
	CustomPlanField7  *string            `url:"custom_plan_field_7,omitempty"`
	CustomPlanField8  *string            `url:"custom_plan_field_8,omitempty"`
	CustomPlanField9  *string            `url:"custom_plan_field_9,omitempty"`
	CustomPlanField10 *string            `url:"custom_plan_field_10,omitempty"`
}

// GetStoriesCount 获取需求数量
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_stories_count.html
func (s *StoryService) GetStoriesCount(
	ctx context.Context, request *GetStoriesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	var response struct {
		Count int `json:"count,omitempty"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

type GetStoriesCountRequest struct {
	ID                *Multi[int64]  `url:"id,omitempty"`               // ID	支持多ID查询,多个ID用逗号分隔
	Name              *string        `url:"name,omitempty"`             // 标题	支持模糊匹配
	Priority          *string        `url:"priority,omitempty"`         // 优先级。
	PriorityLabel     *PriorityLabel `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	BusinessValue     *int           `url:"business_value,omitempty"`   // 业务价值
	Status            *string        `url:"status,omitempty"`           // 状态	支持枚举查询
	VStatus           *string        `url:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	WithVStatus       *string        `url:"with_v_status,omitempty"`    // 值=1可以返回中文状态
	Label             *string        `url:"label,omitempty"`            // 标签查询	支持枚举查询
	WorkitemTypeID    *string        `url:"workitem_type_id,omitempty"` // 需求类别ID	支持枚举查询
	Version           *string        `url:"version,omitempty"`          // 版本
	Module            *string        `url:"module,omitempty"`           // 模块
	Feature           *string        `url:"feature,omitempty"`          // 特性
	TestFocus         *string        `url:"test_focus,omitempty"`       // 测试重点
	Size              *int           `url:"size,omitempty"`             // 规模
	Owner             *string        `url:"owner,omitempty"`            // 处理人	支持模糊匹配
	CC                *string        `url:"cc,omitempty"`               // 抄送人	支持模糊匹配
	Creator           *string        `url:"creator,omitempty"`          // 创建人	支持多人员查询
	Developer         *string        `url:"developer,omitempty"`        // 开发人员
	Begin             *string        `url:"begin,omitempty"`            // 预计开始	支持时间查询
	Due               *string        `url:"due,omitempty"`              // 预计结束	支持时间查询
	Created           *string        `url:"created,omitempty"`          // 创建时间	支持时间查询
	Modified          *string        `url:"modified,omitempty"`         // 最后修改时间	支持时间查询
	Completed         *string        `url:"completed,omitempty"`        // 完成时间	支持时间查询
	IterationID       *string        `url:"iteration_id,omitempty"`     // 迭代ID	支持不等于查询
	Effort            *string        `url:"effort,omitempty"`           // 预估工时
	EffortCompleted   *string        `url:"effort_completed,omitempty"` // 完成工时
	Remain            *float64       `url:"remain,omitempty"`           // 剩余工时
	Exceed            *float64       `url:"exceed,omitempty"`           // 超出工时
	CategoryID        *string        `url:"category_id,omitempty"`      // 需求分类	支持枚举查询
	ReleaseID         *string        `url:"release_id,omitempty"`       // 发布计划
	Source            *string        `url:"source,omitempty"`           // 需求来源
	Type              *string        `url:"type,omitempty"`             // 需求类型
	ParentID          *string        `url:"parent_id,omitempty"`        // 父需求
	ChildrenID        *string        `url:"children_id,omitempty"`      // 子需求	为空查询传：丨
	Description       *string        `url:"description,omitempty"`      // 详细描述	支持模糊匹配
	WorkspaceID       *int           `url:"workspace_id,omitempty"`     // 项目ID
	CustomFieldOne    *string        `url:"custom_field_one,omitempty"`
	CustomFieldTwo    *string        `url:"custom_field_two,omitempty"`
	CustomFieldThree  *string        `url:"custom_field_three,omitempty"`
	CustomFieldFour   *string        `url:"custom_field_four,omitempty"`
	CustomFieldFive   *string        `url:"custom_field_five,omitempty"`
	CustomFieldSix    *string        `url:"custom_field_six,omitempty"`
	CustomFieldSeven  *string        `url:"custom_field_seven,omitempty"`
	CustomFieldEight  *string        `url:"custom_field_eight,omitempty"`
	CustomField9      *string        `url:"custom_field_9,omitempty"`
	CustomField10     *string        `url:"custom_field_10,omitempty"`
	CustomField11     *string        `url:"custom_field_11,omitempty"`
	CustomField12     *string        `url:"custom_field_12,omitempty"`
	CustomField13     *string        `url:"custom_field_13,omitempty"`
	CustomField14     *string        `url:"custom_field_14,omitempty"`
	CustomField15     *string        `url:"custom_field_15,omitempty"`
	CustomField16     *string        `url:"custom_field_16,omitempty"`
	CustomField17     *string        `url:"custom_field_17,omitempty"`
	CustomField18     *string        `url:"custom_field_18,omitempty"`
	CustomField19     *string        `url:"custom_field_19,omitempty"`
	CustomField20     *string        `url:"custom_field_20,omitempty"`
	CustomField21     *string        `url:"custom_field_21,omitempty"`
	CustomField22     *string        `url:"custom_field_22,omitempty"`
	CustomField23     *string        `url:"custom_field_23,omitempty"`
	CustomField24     *string        `url:"custom_field_24,omitempty"`
	CustomField25     *string        `url:"custom_field_25,omitempty"`
	CustomField26     *string        `url:"custom_field_26,omitempty"`
	CustomField27     *string        `url:"custom_field_27,omitempty"`
	CustomField28     *string        `url:"custom_field_28,omitempty"`
	CustomField29     *string        `url:"custom_field_29,omitempty"`
	CustomField30     *string        `url:"custom_field_30,omitempty"`
	CustomField31     *string        `url:"custom_field_31,omitempty"`
	CustomField32     *string        `url:"custom_field_32,omitempty"`
	CustomField33     *string        `url:"custom_field_33,omitempty"`
	CustomField34     *string        `url:"custom_field_34,omitempty"`
	CustomField35     *string        `url:"custom_field_35,omitempty"`
	CustomField36     *string        `url:"custom_field_36,omitempty"`
	CustomField37     *string        `url:"custom_field_37,omitempty"`
	CustomField38     *string        `url:"custom_field_38,omitempty"`
	CustomField39     *string        `url:"custom_field_39,omitempty"`
	CustomField40     *string        `url:"custom_field_40,omitempty"`
	CustomField41     *string        `url:"custom_field_41,omitempty"`
	CustomField42     *string        `url:"custom_field_42,omitempty"`
	CustomField43     *string        `url:"custom_field_43,omitempty"`
	CustomField44     *string        `url:"custom_field_44,omitempty"`
	CustomField45     *string        `url:"custom_field_45,omitempty"`
	CustomField46     *string        `url:"custom_field_46,omitempty"`
	CustomField47     *string        `url:"custom_field_47,omitempty"`
	CustomField48     *string        `url:"custom_field_48,omitempty"`
	CustomField49     *string        `url:"custom_field_49,omitempty"`
	CustomField50     *string        `url:"custom_field_50,omitempty"`
	CustomField51     *string        `url:"custom_field_51,omitempty"`
	CustomField52     *string        `url:"custom_field_52,omitempty"`
	CustomField53     *string        `url:"custom_field_53,omitempty"`
	CustomField54     *string        `url:"custom_field_54,omitempty"`
	CustomField55     *string        `url:"custom_field_55,omitempty"`
	CustomField56     *string        `url:"custom_field_56,omitempty"`
	CustomField57     *string        `url:"custom_field_57,omitempty"`
	CustomField58     *string        `url:"custom_field_58,omitempty"`
	CustomField59     *string        `url:"custom_field_59,omitempty"`
	CustomField60     *string        `url:"custom_field_60,omitempty"`
	CustomField61     *string        `url:"custom_field_61,omitempty"`
	CustomField62     *string        `url:"custom_field_62,omitempty"`
	CustomField63     *string        `url:"custom_field_63,omitempty"`
	CustomField64     *string        `url:"custom_field_64,omitempty"`
	CustomField65     *string        `url:"custom_field_65,omitempty"`
	CustomField66     *string        `url:"custom_field_66,omitempty"`
	CustomField67     *string        `url:"custom_field_67,omitempty"`
	CustomField68     *string        `url:"custom_field_68,omitempty"`
	CustomField69     *string        `url:"custom_field_69,omitempty"`
	CustomField70     *string        `url:"custom_field_70,omitempty"`
	CustomField71     *string        `url:"custom_field_71,omitempty"`
	CustomField72     *string        `url:"custom_field_72,omitempty"`
	CustomField73     *string        `url:"custom_field_73,omitempty"`
	CustomField74     *string        `url:"custom_field_74,omitempty"`
	CustomField75     *string        `url:"custom_field_75,omitempty"`
	CustomField76     *string        `url:"custom_field_76,omitempty"`
	CustomField77     *string        `url:"custom_field_77,omitempty"`
	CustomField78     *string        `url:"custom_field_78,omitempty"`
	CustomField79     *string        `url:"custom_field_79,omitempty"`
	CustomField80     *string        `url:"custom_field_80,omitempty"`
	CustomField81     *string        `url:"custom_field_81,omitempty"`
	CustomField82     *string        `url:"custom_field_82,omitempty"`
	CustomField83     *string        `url:"custom_field_83,omitempty"`
	CustomField84     *string        `url:"custom_field_84,omitempty"`
	CustomField85     *string        `url:"custom_field_85,omitempty"`
	CustomField86     *string        `url:"custom_field_86,omitempty"`
	CustomField87     *string        `url:"custom_field_87,omitempty"`
	CustomField88     *string        `url:"custom_field_88,omitempty"`
	CustomField89     *string        `url:"custom_field_89,omitempty"`
	CustomField90     *string        `url:"custom_field_90,omitempty"`
	CustomField91     *string        `url:"custom_field_91,omitempty"`
	CustomField92     *string        `url:"custom_field_92,omitempty"`
	CustomField93     *string        `url:"custom_field_93,omitempty"`
	CustomField94     *string        `url:"custom_field_94,omitempty"`
	CustomField95     *string        `url:"custom_field_95,omitempty"`
	CustomField96     *string        `url:"custom_field_96,omitempty"`
	CustomField97     *string        `url:"custom_field_97,omitempty"`
	CustomField98     *string        `url:"custom_field_98,omitempty"`
	CustomField99     *string        `url:"custom_field_99,omitempty"`
	CustomField100    *string        `url:"custom_field_100,omitempty"`
	CustomField101    *string        `url:"custom_field_101,omitempty"`
	CustomField102    *string        `url:"custom_field_102,omitempty"`
	CustomField103    *string        `url:"custom_field_103,omitempty"`
	CustomField104    *string        `url:"custom_field_104,omitempty"`
	CustomField105    *string        `url:"custom_field_105,omitempty"`
	CustomField106    *string        `url:"custom_field_106,omitempty"`
	CustomField107    *string        `url:"custom_field_107,omitempty"`
	CustomField108    *string        `url:"custom_field_108,omitempty"`
	CustomField109    *string        `url:"custom_field_109,omitempty"`
	CustomField110    *string        `url:"custom_field_110,omitempty"`
	CustomField111    *string        `url:"custom_field_111,omitempty"`
	CustomField112    *string        `url:"custom_field_112,omitempty"`
	CustomField113    *string        `url:"custom_field_113,omitempty"`
	CustomField114    *string        `url:"custom_field_114,omitempty"`
	CustomField115    *string        `url:"custom_field_115,omitempty"`
	CustomField116    *string        `url:"custom_field_116,omitempty"`
	CustomField117    *string        `url:"custom_field_117,omitempty"`
	CustomField118    *string        `url:"custom_field_118,omitempty"`
	CustomField119    *string        `url:"custom_field_119,omitempty"`
	CustomField120    *string        `url:"custom_field_120,omitempty"`
	CustomField121    *string        `url:"custom_field_121,omitempty"`
	CustomField122    *string        `url:"custom_field_122,omitempty"`
	CustomField123    *string        `url:"custom_field_123,omitempty"`
	CustomField124    *string        `url:"custom_field_124,omitempty"`
	CustomField125    *string        `url:"custom_field_125,omitempty"`
	CustomField126    *string        `url:"custom_field_126,omitempty"`
	CustomField127    *string        `url:"custom_field_127,omitempty"`
	CustomField128    *string        `url:"custom_field_128,omitempty"`
	CustomField129    *string        `url:"custom_field_129,omitempty"`
	CustomField130    *string        `url:"custom_field_130,omitempty"`
	CustomField131    *string        `url:"custom_field_131,omitempty"`
	CustomField132    *string        `url:"custom_field_132,omitempty"`
	CustomField133    *string        `url:"custom_field_133,omitempty"`
	CustomField134    *string        `url:"custom_field_134,omitempty"`
	CustomField135    *string        `url:"custom_field_135,omitempty"`
	CustomField136    *string        `url:"custom_field_136,omitempty"`
	CustomField137    *string        `url:"custom_field_137,omitempty"`
	CustomField138    *string        `url:"custom_field_138,omitempty"`
	CustomField139    *string        `url:"custom_field_139,omitempty"`
	CustomField140    *string        `url:"custom_field_140,omitempty"`
	CustomField141    *string        `url:"custom_field_141,omitempty"`
	CustomField142    *string        `url:"custom_field_142,omitempty"`
	CustomField143    *string        `url:"custom_field_143,omitempty"`
	CustomField144    *string        `url:"custom_field_144,omitempty"`
	CustomField145    *string        `url:"custom_field_145,omitempty"`
	CustomField146    *string        `url:"custom_field_146,omitempty"`
	CustomField147    *string        `url:"custom_field_147,omitempty"`
	CustomField148    *string        `url:"custom_field_148,omitempty"`
	CustomField149    *string        `url:"custom_field_149,omitempty"`
	CustomField150    *string        `url:"custom_field_150,omitempty"`
	CustomField151    *string        `url:"custom_field_151,omitempty"`
	CustomField152    *string        `url:"custom_field_152,omitempty"`
	CustomField153    *string        `url:"custom_field_153,omitempty"`
	CustomField154    *string        `url:"custom_field_154,omitempty"`
	CustomField155    *string        `url:"custom_field_155,omitempty"`
	CustomField156    *string        `url:"custom_field_156,omitempty"`
	CustomField157    *string        `url:"custom_field_157,omitempty"`
	CustomField158    *string        `url:"custom_field_158,omitempty"`
	CustomField159    *string        `url:"custom_field_159,omitempty"`
	CustomField160    *string        `url:"custom_field_160,omitempty"`
	CustomField161    *string        `url:"custom_field_161,omitempty"`
	CustomField162    *string        `url:"custom_field_162,omitempty"`
	CustomField163    *string        `url:"custom_field_163,omitempty"`
	CustomField164    *string        `url:"custom_field_164,omitempty"`
	CustomField165    *string        `url:"custom_field_165,omitempty"`
	CustomField166    *string        `url:"custom_field_166,omitempty"`
	CustomField167    *string        `url:"custom_field_167,omitempty"`
	CustomField168    *string        `url:"custom_field_168,omitempty"`
	CustomField169    *string        `url:"custom_field_169,omitempty"`
	CustomField170    *string        `url:"custom_field_170,omitempty"`
	CustomField171    *string        `url:"custom_field_171,omitempty"`
	CustomField172    *string        `url:"custom_field_172,omitempty"`
	CustomField173    *string        `url:"custom_field_173,omitempty"`
	CustomField174    *string        `url:"custom_field_174,omitempty"`
	CustomField175    *string        `url:"custom_field_175,omitempty"`
	CustomField176    *string        `url:"custom_field_176,omitempty"`
	CustomField177    *string        `url:"custom_field_177,omitempty"`
	CustomField178    *string        `url:"custom_field_178,omitempty"`
	CustomField179    *string        `url:"custom_field_179,omitempty"`
	CustomField180    *string        `url:"custom_field_180,omitempty"`
	CustomField181    *string        `url:"custom_field_181,omitempty"`
	CustomField182    *string        `url:"custom_field_182,omitempty"`
	CustomField183    *string        `url:"custom_field_183,omitempty"`
	CustomField184    *string        `url:"custom_field_184,omitempty"`
	CustomField185    *string        `url:"custom_field_185,omitempty"`
	CustomField186    *string        `url:"custom_field_186,omitempty"`
	CustomField187    *string        `url:"custom_field_187,omitempty"`
	CustomField188    *string        `url:"custom_field_188,omitempty"`
	CustomField189    *string        `url:"custom_field_189,omitempty"`
	CustomField190    *string        `url:"custom_field_190,omitempty"`
	CustomField191    *string        `url:"custom_field_191,omitempty"`
	CustomField192    *string        `url:"custom_field_192,omitempty"`
	CustomField193    *string        `url:"custom_field_193,omitempty"`
	CustomField194    *string        `url:"custom_field_194,omitempty"`
	CustomField195    *string        `url:"custom_field_195,omitempty"`
	CustomField196    *string        `url:"custom_field_196,omitempty"`
	CustomField197    *string        `url:"custom_field_197,omitempty"`
	CustomField198    *string        `url:"custom_field_198,omitempty"`
	CustomField199    *string        `url:"custom_field_199,omitempty"`
	CustomField200    *string        `url:"custom_field_200,omitempty"`
	CustomPlanField1  *string        `url:"custom_plan_field_1,omitempty"`
	CustomPlanField2  *string        `url:"custom_plan_field_2,omitempty"`
	CustomPlanField3  *string        `url:"custom_plan_field_3,omitempty"`
	CustomPlanField4  *string        `url:"custom_plan_field_4,omitempty"`
	CustomPlanField5  *string        `url:"custom_plan_field_5,omitempty"`
	CustomPlanField6  *string        `url:"custom_plan_field_6,omitempty"`
	CustomPlanField7  *string        `url:"custom_plan_field_7,omitempty"`
	CustomPlanField8  *string        `url:"custom_plan_field_8,omitempty"`
	CustomPlanField9  *string        `url:"custom_plan_field_9,omitempty"`
	CustomPlanField10 *string        `url:"custom_plan_field_10,omitempty"`
}

// 获取保密需求
// 获取保密需求数量

// GetStoryCategories 获取需求分类
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_categories.html
func (s *StoryService) GetStoryCategories(
	ctx context.Context, request *GetStoryCategoriesRequest, opts ...RequestOption,
) ([]*StoryCategory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_categories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Category *StoryCategory
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	categories := make([]*StoryCategory, 0, len(items))
	for _, item := range items {
		categories = append(categories, item.Category)
	}

	return categories, resp, nil
}

type GetStoryCategoriesRequest struct {
	WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
	ID          *Multi[int64]  `url:"id,omitempty"`           // ID 支持多ID查询，多个ID用逗号分隔
	Name        *string        `url:"name,omitempty"`         // 需求分类名称	支持模糊匹配
	Description *string        `url:"description,omitempty"`  // 需求分类描述
	ParentID    *int           `url:"parent_id,omitempty"`    // 父分类ID
	Created     *string        `url:"created,omitempty"`      // 创建时间	支持时间查询
	Modified    *string        `url:"modified,omitempty"`     // 最后修改时间	支持时间查询
	Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
	Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order       *Order         `url:"order,omitempty"`        //nolint:lll // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode	如按创建时间逆序：order=created%20desc
	Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
}

type StoryCategory struct {
	ID          string `json:"id,omitempty"`           // ID
	WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
	Name        string `json:"name,omitempty"`         // 需求分类名称
	Description string `json:"description,omitempty"`  // 需求分类描述
	ParentID    string `json:"parent_id,omitempty"`    // 父分类ID
	Created     string `json:"created,omitempty"`      // 创建时间
	Modified    string `json:"modified,omitempty"`     // 最后修改时间
	Creator     string `json:"creator,omitempty"`      // 创建人
	Modifier    string `json:"modifier,omitempty"`     // 最后修改人
}

// -----------------------------------------------------------------------------
// 获取需求分类数量
// -----------------------------------------------------------------------------

type GetStoryCategoriesCountRequest struct {
	WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
	ID          *Multi[int64] `url:"id,omitempty"`           // ID 支持多ID查询，多个ID用逗号分隔
	Name        *string       `url:"name,omitempty"`         // 需求分类名称	支持模糊匹配
	Description *string       `url:"description,omitempty"`  // 需求分类描述
	ParentID    *int          `url:"parent_id,omitempty"`    // 父分类ID
	Created     *string       `url:"created,omitempty"`      // 创建时间	支持时间查询
	Modified    *string       `url:"modified,omitempty"`     // 最后修改时间	支持时间查询
}

// GetStoryCategoriesCount 获取需求分类数量
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_categories_count.html
func (s *StoryService) GetStoryCategoriesCount(
	ctx context.Context, request *GetStoryCategoriesCountRequest, opts ...RequestOption,
) (int, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_categories/count", request, opts)
	if err != nil {
		return 0, nil, err
	}

	response := new(CountResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return 0, resp, err
	}

	return response.Count, resp, nil
}

// -----------------------------------------------------------------------------
// 获取指定分类需求数量
// -----------------------------------------------------------------------------

type GetStoriesCountByCategoriesRequest struct {
	WorkspaceID *int          `url:"workspace_id,omitempty"` // [必须]项目ID
	CategoryID  *Multi[int64] `url:"category_id,omitempty"`  // 需求分类 支持多ID。比如 id1,id2,id3
}

type StoriesCountByCategory struct {
	CategoryID string `json:"category_id,omitempty"`
	Count      int    `json:"count,omitempty"`
}

// GetStoriesCountByCategories 获取指定分类下需求数量
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/count_by_categories.html
func (s *StoryService) GetStoriesCountByCategories(
	ctx context.Context, request *GetStoriesCountByCategoriesRequest, opts ...RequestOption,
) ([]*StoriesCountByCategory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/count_by_categories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items map[string]int
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	counts := make([]*StoriesCountByCategory, 0, len(items))
	for categoryID, count := range items {
		counts = append(counts, &StoriesCountByCategory{
			CategoryID: categoryID,
			Count:      count,
		})
	}

	return counts, resp, nil
}

// GetStoryChanges 获取需求变更历史
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html
func (s *StoryService) GetStoryChanges(
	ctx context.Context, request *GetStoryChangesRequest, opts ...RequestOption,
) ([]*StoryChange, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "story_changes", request, opts)
	if err != nil {
		return nil, nil, err
	}

	items := make([]struct {
		WorkitemChange *StoryChange `json:"WorkitemChange"`
	}, 0)
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	changes := make([]*StoryChange, 0, len(items))
	for _, item := range items {
		changes = append(changes, item.WorkitemChange)
	}

	return changes, resp, nil
}

// StoreChangeType 变更类型
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_changes.html
type StoreChangeType string

const (
	StoreChangeTypeSyncCopy            StoreChangeType = "sync_copy"              // 同步复制联动
	StoreChangeTypeStoryStatusRelation StoreChangeType = "story_status_relation"  // 父子需求联动
	StoreChangeTypeStoryTaskRelation   StoreChangeType = "story_task_relation"    // 需求任务联动
	StoreChangeTypeAPI                 StoreChangeType = "api"                    // API变更
	StoreChangeTypeSmartCommit         StoreChangeType = "smart_commit"           // Smart Commit触发
	StoreChangeTypeAutoTask            StoreChangeType = "auto_task"              // 自动化任务触发
	StoreChangeTypeAutoWorkflow        StoreChangeType = "auto_workflow"          // 自动化工作流触发
	StoreChangeTypeManualUpdate        StoreChangeType = "manual_update"          // 手动变更
	StoreChangeTypeImportUpdate        StoreChangeType = "import_update"          // 导入更新
	StoreChangeTypeCodeChange          StoreChangeType = "code_change"            // 代码变更
	StoreChangeTypeStatusDelete        StoreChangeType = "status_delete"          // 状态删除
	StoreChangeTypeExitWorkspace       StoreChangeType = "exit_workspace"         // 退出项目触发
	StoreChangeTypeLinkUpdate          StoreChangeType = "link_update"            // 更新关联
	StoreChangeTypeLinkCreate          StoreChangeType = "link_create"            // 创建关联
	StoreChangeTypeLinkDelete          StoreChangeType = "link_delete"            // 删除关联
	StoreChangeTypeCreateStoryFromCopy StoreChangeType = "create_story_from_copy" // 复制创建
	StoreChangeTypeCreateStory         StoreChangeType = "create_story"           // 创建需求
)

type GetStoryChangesRequest struct {
	ID               *Multi[int64]    `url:"id,omitempty"`
	StoryID          *Multi[int64]    `url:"story_id,omitempty"`           // 需求id	支持多ID查询
	WorkspaceID      *int             `url:"workspace_id,omitempty"`       // [必须]项目ID
	Creator          *string          `url:"creator,omitempty"`            // 创建人（操作人）
	Created          *string          `url:"created,omitempty"`            // 创建时间（变更时间）	支持时间查询
	ChangeType       *StoreChangeType `url:"change_type,omitempty"`        // 变更类型
	ChangeSummary    *string          `url:"change_summary,omitempty"`     // 需求变更描述
	Comment          *string          `url:"comment,omitempty"`            // 评论
	EntityType       *string          `url:"entity_type,omitempty"`        // 变更的对象类型
	ChangeField      *string          `url:"change_field,omitempty"`       // 设置获取变更字段如（status）
	NeedParseChanges *int             `url:"need_parse_changes,omitempty"` // 设置field_changes字段是否返回（默认取 1。取 0 则不返回）
	Limit            *int             `url:"limit,omitempty"`              // 设置返回数量限制，默认为30，最大取 100
	Page             *int             `url:"page,omitempty"`               // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order            *Order           `url:"order,omitempty"`              // 排序规则，规则：字段名 ASC或者DESC
	Fields           *Multi[string]   `url:"fields,omitempty"`             // 设置获取的字段，多个字段间以','逗号隔开
}

type StoryChange struct {
	ID             string          `json:"id,omitempty"`
	WorkspaceID    string          `json:"workspace_id,omitempty"`
	AppID          string          `json:"app_id,omitempty"`
	WorkitemTypeID string          `json:"workitem_type_id,omitempty"`
	Creator        string          `json:"creator,omitempty"`
	Created        string          `json:"created,omitempty"`
	ChangeSummary  string          `json:"change_summary,omitempty"`
	Comment        *string         `json:"comment,omitempty"`
	Changes        string          `json:"changes,omitempty"`
	EntityType     string          `json:"entity_type,omitempty"`
	ChangeType     StoreChangeType `json:"change_type,omitempty"`
	ChangeTypeText string          `json:"change_type_text,omitempty"`
	Updated        string          `json:"updated,omitempty"`
	FieldChanges   []struct {
		Field             string `json:"field,omitempty"`
		ValueBefore       any    `json:"value_before,omitempty"` // todo: any to string
		ValueAfter        any    `json:"value_after,omitempty"`  // todo: any to string
		ValueBeforeParsed string `json:"value_before_parsed,omitempty"`
		ValueAfterParsed  string `json:"value_after_parsed,omitempty"`
		FieldLabel        string `json:"field_label,omitempty"`
	} `json:"field_changes,omitempty"`
	StoryID string `json:"story_id,omitempty"`
}

// -----------------------------------------------------------------------------
// 获取需求变更次数
// -----------------------------------------------------------------------------

// GetStoryCustomFieldsSettings 获取需求自定义字段配置
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_custom_fields_settings.html
func (s *StoryService) GetStoryCustomFieldsSettings(
	ctx context.Context, request *GetStoryCustomFieldsSettingsRequest, opts ...RequestOption,
) ([]*StoryCustomFieldsSetting, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/custom_fields_settings", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := make([]struct {
		CustomFieldConfig *StoryCustomFieldsSetting `json:"CustomFieldConfig,omitempty"`
	}, 0)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	settings := make([]*StoryCustomFieldsSetting, 0, len(response))
	for _, item := range response {
		settings = append(settings, item.CustomFieldConfig)
	}

	return settings, resp, nil
}

type GetStoryCustomFieldsSettingsRequest struct {
	WorkspaceID *int `url:"workspace_id,omitempty"` // 项目ID
}

type StoryCustomFieldsSetting struct {
	ID              string  `json:"id,omitempty"`           // 自定义字段配置的ID
	WorkspaceID     string  `json:"workspace_id,omitempty"` // 所属项目ID
	AppID           string  `json:"app_id,omitempty"`
	EntryType       string  `json:"entry_type,omitempty"`   // 所属实体对象
	CustomField     string  `json:"custom_field,omitempty"` // 自定义字段标识（英文名）
	Type            string  `json:"type,omitempty"`         // 输入类型
	Name            string  `json:"name,omitempty"`         // 自定义字段显示名称
	Options         *string `json:"options,omitempty"`      // 自定义字段可选值
	ExtraConfig     *string `json:"extra_config,omitempty"` // 额外配置
	Enabled         string  `json:"enabled,omitempty"`      // 是否启用
	Freeze          string  `json:"freeze,omitempty"`
	Sort            *string `json:"sort,omitempty"` // 显示时排序系数
	Memo            *string `json:"memo,omitempty"`
	OpenExtensionID string  `json:"open_extension_id,omitempty"`
	IsOut           int     `json:"is_out,omitempty"`
	IsUninstall     int     `json:"is_uninstall,omitempty"`
	AppName         string  `json:"app_name,omitempty"`
}

// -----------------------------------------------------------------------------
// 获取需求与测试用例关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量新增或修改需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量删除需求前后置关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求保密信息
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 批量修改保密信息
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求类别
// -----------------------------------------------------------------------------

// UpdateStory 更新需求
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/update_story.html
func (s *StoryService) UpdateStory(
	ctx context.Context, request *UpdateStoryRequest, opts ...RequestOption,
) (*Story, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Story *Story `json:"story"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Story, resp, nil
}

type UpdateStoryRequest struct {
	ID                *int64         `json:"id"`                           // 必须
	WorkspaceID       *int           `json:"workspace_id"`                 // 必须
	Name              *string        `json:"name,omitempty"`               // 标题
	Priority          *string        `json:"priority,omitempty"`           // 优先级。
	PriorityLabel     *PriorityLabel `json:"priority_label,omitempty"`     // 优先级。推荐使用这个字段
	BusinessValue     *int           `json:"business_value,omitempty"`     // 业务价值
	Status            *string        `json:"status,omitempty"`             // 状态
	VStatus           *string        `json:"v_status,omitempty"`           // 中文状态名称
	Version           *string        `json:"version,omitempty"`            // 版本
	Module            *string        `json:"module,omitempty"`             // 模块
	TestFocus         *string        `json:"test_focus,omitempty"`         // 测试重点
	Size              *int           `json:"size,omitempty"`               // 规模
	Owner             *string        `json:"owner,omitempty"`              // 处理人
	CurrentUser       *string        `json:"current_user,omitempty"`       // 变更人
	CC                *string        `json:"cc,omitempty"`                 // 抄送人
	Developer         *string        `json:"developer,omitempty"`          // 开发人员
	Begin             *string        `json:"begin,omitempty"`              // 预计开始
	Due               *string        `json:"due,omitempty"`                // 预计结束
	IterationID       *string        `json:"iteration_id,omitempty"`       // 迭代ID
	Effort            *string        `json:"effort,omitempty"`             // 预估工时
	EffortCompleted   *string        `json:"effort_completed,omitempty"`   // 完成工时
	Remain            *float64       `json:"remain,omitempty"`             // 剩余工时
	Exceed            *float64       `json:"exceed,omitempty"`             // 超出工时
	CategoryID        *int64         `json:"category_id,omitempty"`        // 需求分类ID
	ReleaseID         *int           `json:"release_id,omitempty"`         // 发布计划ID
	Source            *string        `json:"source,omitempty"`             // 来源
	Type              *string        `json:"type,omitempty"`               // 类型
	Description       *string        `json:"description,omitempty"`        // 详细描述
	IsAutoCloseTask   *int           `json:"is_auto_close_task,omitempty"` // 自动关闭关联任务
	Label             *string        `json:"label,omitempty"`              // 标签
	CustomFieldOne    *string        `json:"custom_field_one,omitempty"`
	CustomFieldTwo    *string        `json:"custom_field_two,omitempty"`
	CustomFieldThree  *string        `json:"custom_field_three,omitempty"`
	CustomFieldFour   *string        `json:"custom_field_four,omitempty"`
	CustomFieldFive   *string        `json:"custom_field_five,omitempty"`
	CustomFieldSix    *string        `json:"custom_field_six,omitempty"`
	CustomFieldSeven  *string        `json:"custom_field_seven,omitempty"`
	CustomFieldEight  *string        `json:"custom_field_eight,omitempty"`
	CustomField9      *string        `json:"custom_field_9,omitempty"`
	CustomField10     *string        `json:"custom_field_10,omitempty"`
	CustomField11     *string        `json:"custom_field_11,omitempty"`
	CustomField12     *string        `json:"custom_field_12,omitempty"`
	CustomField13     *string        `json:"custom_field_13,omitempty"`
	CustomField14     *string        `json:"custom_field_14,omitempty"`
	CustomField15     *string        `json:"custom_field_15,omitempty"`
	CustomField16     *string        `json:"custom_field_16,omitempty"`
	CustomField17     *string        `json:"custom_field_17,omitempty"`
	CustomField18     *string        `json:"custom_field_18,omitempty"`
	CustomField19     *string        `json:"custom_field_19,omitempty"`
	CustomField20     *string        `json:"custom_field_20,omitempty"`
	CustomField21     *string        `json:"custom_field_21,omitempty"`
	CustomField22     *string        `json:"custom_field_22,omitempty"`
	CustomField23     *string        `json:"custom_field_23,omitempty"`
	CustomField24     *string        `json:"custom_field_24,omitempty"`
	CustomField25     *string        `json:"custom_field_25,omitempty"`
	CustomField26     *string        `json:"custom_field_26,omitempty"`
	CustomField27     *string        `json:"custom_field_27,omitempty"`
	CustomField28     *string        `json:"custom_field_28,omitempty"`
	CustomField29     *string        `json:"custom_field_29,omitempty"`
	CustomField30     *string        `json:"custom_field_30,omitempty"`
	CustomField31     *string        `json:"custom_field_31,omitempty"`
	CustomField32     *string        `json:"custom_field_32,omitempty"`
	CustomField33     *string        `json:"custom_field_33,omitempty"`
	CustomField34     *string        `json:"custom_field_34,omitempty"`
	CustomField35     *string        `json:"custom_field_35,omitempty"`
	CustomField36     *string        `json:"custom_field_36,omitempty"`
	CustomField37     *string        `json:"custom_field_37,omitempty"`
	CustomField38     *string        `json:"custom_field_38,omitempty"`
	CustomField39     *string        `json:"custom_field_39,omitempty"`
	CustomField40     *string        `json:"custom_field_40,omitempty"`
	CustomField41     *string        `json:"custom_field_41,omitempty"`
	CustomField42     *string        `json:"custom_field_42,omitempty"`
	CustomField43     *string        `json:"custom_field_43,omitempty"`
	CustomField44     *string        `json:"custom_field_44,omitempty"`
	CustomField45     *string        `json:"custom_field_45,omitempty"`
	CustomField46     *string        `json:"custom_field_46,omitempty"`
	CustomField47     *string        `json:"custom_field_47,omitempty"`
	CustomField48     *string        `json:"custom_field_48,omitempty"`
	CustomField49     *string        `json:"custom_field_49,omitempty"`
	CustomField50     *string        `json:"custom_field_50,omitempty"`
	CustomField51     *string        `json:"custom_field_51,omitempty"`
	CustomField52     *string        `json:"custom_field_52,omitempty"`
	CustomField53     *string        `json:"custom_field_53,omitempty"`
	CustomField54     *string        `json:"custom_field_54,omitempty"`
	CustomField55     *string        `json:"custom_field_55,omitempty"`
	CustomField56     *string        `json:"custom_field_56,omitempty"`
	CustomField57     *string        `json:"custom_field_57,omitempty"`
	CustomField58     *string        `json:"custom_field_58,omitempty"`
	CustomField59     *string        `json:"custom_field_59,omitempty"`
	CustomField60     *string        `json:"custom_field_60,omitempty"`
	CustomField61     *string        `json:"custom_field_61,omitempty"`
	CustomField62     *string        `json:"custom_field_62,omitempty"`
	CustomField63     *string        `json:"custom_field_63,omitempty"`
	CustomField64     *string        `json:"custom_field_64,omitempty"`
	CustomField65     *string        `json:"custom_field_65,omitempty"`
	CustomField66     *string        `json:"custom_field_66,omitempty"`
	CustomField67     *string        `json:"custom_field_67,omitempty"`
	CustomField68     *string        `json:"custom_field_68,omitempty"`
	CustomField69     *string        `json:"custom_field_69,omitempty"`
	CustomField70     *string        `json:"custom_field_70,omitempty"`
	CustomField71     *string        `json:"custom_field_71,omitempty"`
	CustomField72     *string        `json:"custom_field_72,omitempty"`
	CustomField73     *string        `json:"custom_field_73,omitempty"`
	CustomField74     *string        `json:"custom_field_74,omitempty"`
	CustomField75     *string        `json:"custom_field_75,omitempty"`
	CustomField76     *string        `json:"custom_field_76,omitempty"`
	CustomField77     *string        `json:"custom_field_77,omitempty"`
	CustomField78     *string        `json:"custom_field_78,omitempty"`
	CustomField79     *string        `json:"custom_field_79,omitempty"`
	CustomField80     *string        `json:"custom_field_80,omitempty"`
	CustomField81     *string        `json:"custom_field_81,omitempty"`
	CustomField82     *string        `json:"custom_field_82,omitempty"`
	CustomField83     *string        `json:"custom_field_83,omitempty"`
	CustomField84     *string        `json:"custom_field_84,omitempty"`
	CustomField85     *string        `json:"custom_field_85,omitempty"`
	CustomField86     *string        `json:"custom_field_86,omitempty"`
	CustomField87     *string        `json:"custom_field_87,omitempty"`
	CustomField88     *string        `json:"custom_field_88,omitempty"`
	CustomField89     *string        `json:"custom_field_89,omitempty"`
	CustomField90     *string        `json:"custom_field_90,omitempty"`
	CustomField91     *string        `json:"custom_field_91,omitempty"`
	CustomField92     *string        `json:"custom_field_92,omitempty"`
	CustomField93     *string        `json:"custom_field_93,omitempty"`
	CustomField94     *string        `json:"custom_field_94,omitempty"`
	CustomField95     *string        `json:"custom_field_95,omitempty"`
	CustomField96     *string        `json:"custom_field_96,omitempty"`
	CustomField97     *string        `json:"custom_field_97,omitempty"`
	CustomField98     *string        `json:"custom_field_98,omitempty"`
	CustomField99     *string        `json:"custom_field_99,omitempty"`
	CustomField100    *string        `json:"custom_field_100,omitempty"`
	CustomField101    *string        `json:"custom_field_101,omitempty"`
	CustomField102    *string        `json:"custom_field_102,omitempty"`
	CustomField103    *string        `json:"custom_field_103,omitempty"`
	CustomField104    *string        `json:"custom_field_104,omitempty"`
	CustomField105    *string        `json:"custom_field_105,omitempty"`
	CustomField106    *string        `json:"custom_field_106,omitempty"`
	CustomField107    *string        `json:"custom_field_107,omitempty"`
	CustomField108    *string        `json:"custom_field_108,omitempty"`
	CustomField109    *string        `json:"custom_field_109,omitempty"`
	CustomField110    *string        `json:"custom_field_110,omitempty"`
	CustomField111    *string        `json:"custom_field_111,omitempty"`
	CustomField112    *string        `json:"custom_field_112,omitempty"`
	CustomField113    *string        `json:"custom_field_113,omitempty"`
	CustomField114    *string        `json:"custom_field_114,omitempty"`
	CustomField115    *string        `json:"custom_field_115,omitempty"`
	CustomField116    *string        `json:"custom_field_116,omitempty"`
	CustomField117    *string        `json:"custom_field_117,omitempty"`
	CustomField118    *string        `json:"custom_field_118,omitempty"`
	CustomField119    *string        `json:"custom_field_119,omitempty"`
	CustomField120    *string        `json:"custom_field_120,omitempty"`
	CustomField121    *string        `json:"custom_field_121,omitempty"`
	CustomField122    *string        `json:"custom_field_122,omitempty"`
	CustomField123    *string        `json:"custom_field_123,omitempty"`
	CustomField124    *string        `json:"custom_field_124,omitempty"`
	CustomField125    *string        `json:"custom_field_125,omitempty"`
	CustomField126    *string        `json:"custom_field_126,omitempty"`
	CustomField127    *string        `json:"custom_field_127,omitempty"`
	CustomField128    *string        `json:"custom_field_128,omitempty"`
	CustomField129    *string        `json:"custom_field_129,omitempty"`
	CustomField130    *string        `json:"custom_field_130,omitempty"`
	CustomField131    *string        `json:"custom_field_131,omitempty"`
	CustomField132    *string        `json:"custom_field_132,omitempty"`
	CustomField133    *string        `json:"custom_field_133,omitempty"`
	CustomField134    *string        `json:"custom_field_134,omitempty"`
	CustomField135    *string        `json:"custom_field_135,omitempty"`
	CustomField136    *string        `json:"custom_field_136,omitempty"`
	CustomField137    *string        `json:"custom_field_137,omitempty"`
	CustomField138    *string        `json:"custom_field_138,omitempty"`
	CustomField139    *string        `json:"custom_field_139,omitempty"`
	CustomField140    *string        `json:"custom_field_140,omitempty"`
	CustomField141    *string        `json:"custom_field_141,omitempty"`
	CustomField142    *string        `json:"custom_field_142,omitempty"`
	CustomField143    *string        `json:"custom_field_143,omitempty"`
	CustomField144    *string        `json:"custom_field_144,omitempty"`
	CustomField145    *string        `json:"custom_field_145,omitempty"`
	CustomField146    *string        `json:"custom_field_146,omitempty"`
	CustomField147    *string        `json:"custom_field_147,omitempty"`
	CustomField148    *string        `json:"custom_field_148,omitempty"`
	CustomField149    *string        `json:"custom_field_149,omitempty"`
	CustomField150    *string        `json:"custom_field_150,omitempty"`
	CustomField151    *string        `json:"custom_field_151,omitempty"`
	CustomField152    *string        `json:"custom_field_152,omitempty"`
	CustomField153    *string        `json:"custom_field_153,omitempty"`
	CustomField154    *string        `json:"custom_field_154,omitempty"`
	CustomField155    *string        `json:"custom_field_155,omitempty"`
	CustomField156    *string        `json:"custom_field_156,omitempty"`
	CustomField157    *string        `json:"custom_field_157,omitempty"`
	CustomField158    *string        `json:"custom_field_158,omitempty"`
	CustomField159    *string        `json:"custom_field_159,omitempty"`
	CustomField160    *string        `json:"custom_field_160,omitempty"`
	CustomField161    *string        `json:"custom_field_161,omitempty"`
	CustomField162    *string        `json:"custom_field_162,omitempty"`
	CustomField163    *string        `json:"custom_field_163,omitempty"`
	CustomField164    *string        `json:"custom_field_164,omitempty"`
	CustomField165    *string        `json:"custom_field_165,omitempty"`
	CustomField166    *string        `json:"custom_field_166,omitempty"`
	CustomField167    *string        `json:"custom_field_167,omitempty"`
	CustomField168    *string        `json:"custom_field_168,omitempty"`
	CustomField169    *string        `json:"custom_field_169,omitempty"`
	CustomField170    *string        `json:"custom_field_170,omitempty"`
	CustomField171    *string        `json:"custom_field_171,omitempty"`
	CustomField172    *string        `json:"custom_field_172,omitempty"`
	CustomField173    *string        `json:"custom_field_173,omitempty"`
	CustomField174    *string        `json:"custom_field_174,omitempty"`
	CustomField175    *string        `json:"custom_field_175,omitempty"`
	CustomField176    *string        `json:"custom_field_176,omitempty"`
	CustomField177    *string        `json:"custom_field_177,omitempty"`
	CustomField178    *string        `json:"custom_field_178,omitempty"`
	CustomField179    *string        `json:"custom_field_179,omitempty"`
	CustomField180    *string        `json:"custom_field_180,omitempty"`
	CustomField181    *string        `json:"custom_field_181,omitempty"`
	CustomField182    *string        `json:"custom_field_182,omitempty"`
	CustomField183    *string        `json:"custom_field_183,omitempty"`
	CustomField184    *string        `json:"custom_field_184,omitempty"`
	CustomField185    *string        `json:"custom_field_185,omitempty"`
	CustomField186    *string        `json:"custom_field_186,omitempty"`
	CustomField187    *string        `json:"custom_field_187,omitempty"`
	CustomField188    *string        `json:"custom_field_188,omitempty"`
	CustomField189    *string        `json:"custom_field_189,omitempty"`
	CustomField190    *string        `json:"custom_field_190,omitempty"`
	CustomField191    *string        `json:"custom_field_191,omitempty"`
	CustomField192    *string        `json:"custom_field_192,omitempty"`
	CustomField193    *string        `json:"custom_field_193,omitempty"`
	CustomField194    *string        `json:"custom_field_194,omitempty"`
	CustomField195    *string        `json:"custom_field_195,omitempty"`
	CustomField196    *string        `json:"custom_field_196,omitempty"`
	CustomField197    *string        `json:"custom_field_197,omitempty"`
	CustomField198    *string        `json:"custom_field_198,omitempty"`
	CustomField199    *string        `json:"custom_field_199,omitempty"`
	CustomField200    *string        `json:"custom_field_200,omitempty"`
	CustomPlanField1  *string        `json:"custom_plan_field_1,omitempty"`
	CustomPlanField2  *string        `json:"custom_plan_field_2,omitempty"`
	CustomPlanField3  *string        `json:"custom_plan_field_3,omitempty"`
	CustomPlanField4  *string        `json:"custom_plan_field_4,omitempty"`
	CustomPlanField5  *string        `json:"custom_plan_field_5,omitempty"`
	CustomPlanField6  *string        `json:"custom_plan_field_6,omitempty"`
	CustomPlanField7  *string        `json:"custom_plan_field_7,omitempty"`
	CustomPlanField8  *string        `json:"custom_plan_field_8,omitempty"`
	CustomPlanField9  *string        `json:"custom_plan_field_9,omitempty"`
	CustomPlanField10 *string        `json:"custom_plan_field_10,omitempty"`
}

// -----------------------------------------------------------------------------
// 更新需求的需求类别
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求所有字段及候选值
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取需求所有字段的中英文
// -----------------------------------------------------------------------------

// GetStoryTemplates 获取需求模板列表
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_template_list.html
func (s *StoryService) GetStoryTemplates(
	ctx context.Context, request *GetStoryTemplatesRequest, opts ...RequestOption,
) ([]*StoryTemplate, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/template_list", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		WorkitemTemplate *StoryTemplate `json:"WorkitemTemplate,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	templates := make([]*StoryTemplate, 0, len(items))
	for _, item := range items {
		templates = append(templates, item.WorkitemTemplate)
	}

	return templates, resp, nil
}

type GetStoryTemplatesRequest struct {
	WorkspaceID    *int `url:"workspace_id,omitempty"`     // [必须]项目ID
	WorkitemTypeID *int `url:"workitem_type_id,omitempty"` // 需求类别ID
}

type StoryTemplate struct {
	ID          string `json:"id,omitempty"`          // 模板ID
	Name        string `json:"name,omitempty"`        // 标题
	Description string `json:"description,omitempty"` // 详细描述
	Sort        string `json:"sort,omitempty"`        // 排序
	Default     string `json:"default,omitempty"`     // 是否启用
	Creator     string `json:"creator,omitempty"`     // 提交人
	EditorType  string `json:"editor_type,omitempty"` // 编辑器类型
}

// GetStoryTemplateFields 获取需求模板字段
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_default_story_template.html
func (s *StoryService) GetStoryTemplateFields(
	ctx context.Context, request *GetStoryTemplateFieldsRequest, opts ...RequestOption,
) ([]*StoryTemplateField, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_default_story_template", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		WorkitemTemplateField *StoryTemplateField `json:"WorkitemTemplateField,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	templates := make([]*StoryTemplateField, 0, len(items))
	for _, item := range items {
		templates = append(templates, item.WorkitemTemplateField)
	}

	return templates, resp, nil
}

type GetStoryTemplateFieldsRequest struct {
	WorkspaceID *int   `url:"workspace_id,omitempty"` // [必须]项目ID
	TemplateID  *int64 `url:"template_id,omitempty"`  // [必须]模板ID
}

type StoryTemplateField struct {
	ID           string `json:"id,omitempty"`           // 模板字段ID
	WorkspaceID  string `json:"workspace_id,omitempty"` // 项目ID
	Type         string `json:"type,omitempty"`         // 类型
	TemplateID   string `json:"template_id,omitempty"`  // 模板ID
	Field        string `json:"field,omitempty"`        // 字段名称
	Value        string `json:"value,omitempty"`        // 默认值
	Required     string `json:"required,omitempty"`     // 是否必填
	Sort         string `json:"sort,omitempty"`         // 排序
	LinkageRules string `json:"linkage_rules,omitempty"`
}

// -----------------------------------------------------------------------------
// 更新需求分类
// -----------------------------------------------------------------------------

// GetRemovedStories 获取回收站中的需求
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_removed_stories.html
func (s *StoryService) GetRemovedStories(
	ctx context.Context, request *GetRemovedStoriesRequest, opts ...RequestOption,
) ([]*RemovedStory, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_removed_stories", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		RemovedStory *RemovedStory `json:"RemovedStory"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	stories := make([]*RemovedStory, 0, len(items))
	for _, item := range items {
		stories = append(stories, item.RemovedStory)
	}

	return stories, resp, nil
}

type GetRemovedStoriesRequest struct {
	WorkspaceID *int        `url:"workspace_id,omitempty"` // [必须]项目ID
	ID          *Multi[int] `url:"id,omitempty"`           // 需求ID
	Creator     *string     `url:"creator,omitempty"`      // 创建人
	IsArchived  *int        `url:"is_archived,omitempty"`  // 是否为归档。默认取 0，为不返回归档的需求。传 is_archived=1 参数则仅返回归档的需求
	Created     *string     `url:"created,omitempty"`      // 创建时间
	Deleted     *string     `url:"deleted,omitempty"`      // 删除时间
	Limit       *int        `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
	Page        *int        `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1（第一页）
}

type RemovedStory struct {
	ID            string `json:"id,omitempty"`             // 需求ID
	Name          string `json:"name,omitempty"`           // 标题
	Creator       string `json:"creator,omitempty"`        // 创建人
	Created       string `json:"created,omitempty"`        // 创建时间
	OperationUser string `json:"operation_user,omitempty"` // 删除人
	IsArchived    string `json:"is_archived,omitempty"`    // 是否为归档
	Deleted       string `json:"deleted,omitempty"`        // 删除时间
}

// GetStoryRelatedBugs 获取需求关联的缺陷
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/get_story_related_bugs.html
func (s *StoryService) GetStoryRelatedBugs(
	ctx context.Context, request *GetStoryRelatedBugsRequest, opts ...RequestOption,
) ([]*StoryRelatedBug, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_related_bugs", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var bugs []*StoryRelatedBug
	resp, err := s.client.Do(req, &bugs)
	if err != nil {
		return nil, resp, err
	}

	return bugs, resp, nil
}

type GetStoryRelatedBugsRequest struct {
	WorkspaceID *int          `url:"workspace_id,omitempty"`
	StoryID     *Multi[int64] `url:"story_id,omitempty"`
}

type StoryRelatedBug struct {
	WorkspaceID int    `json:"workspace_id,omitempty"`
	StoryID     string `json:"story_id,omitempty"`
	BugID       string `json:"bug_id,omitempty"`
}

// -----------------------------------------------------------------------------
// 解除需求缺陷关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 更新父需求
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 创建需求与缺陷关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 创建需求与测试用例关联关系
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 获取视图对应的需求列表
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// 转换需求ID成列表queryToken
// -----------------------------------------------------------------------------

// GetConvertStoryIDsToQueryToken 转换需求ID成列表queryToken
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/story/story_ids_to_query_token.html
func (s *StoryService) GetConvertStoryIDsToQueryToken(
	ctx context.Context, request *GetConvertStoryIDsToQueryTokenRequest, opts ...RequestOption,
) (*GetConvertStoryIDsToQueryTokenResponse, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "stories/ids_to_query_token", request, opts)
	if err != nil {
		return nil, nil, err
	}

	response := new(GetConvertStoryIDsToQueryTokenResponse)
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

type GetConvertStoryIDsToQueryTokenRequest struct {
	WorkspaceID *int          `json:"workspace_id,omitempty"` // 项目ID
	StoryIDs    *Multi[int64] `json:"ids,omitempty"`          // 需求ID
}

type GetConvertStoryIDsToQueryTokenResponse struct {
	QueryToken string `json:"queryToken,omitempty"` // 列表queryToken
	Href       string `json:"href,omitempty"`       // 对应的TAPD需求列表链接
}

// -----------------------------------------------------------------------------
// 创建需求关联关系
// -----------------------------------------------------------------------------
