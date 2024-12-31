package tapd

import (
	"context"
	"net/http"
)

type Bug struct {
	ID                string        `json:"id"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	Priority          string        `json:"priority"`
	Severity          string        `json:"severity"`
	Module            string        `json:"module"`
	Status            string        `json:"status"`
	Reporter          string        `json:"reporter"`
	Created           string        `json:"created"`
	BugType           string        `json:"bugtype"`
	Resolved          string        `json:"resolved"`
	Closed            string        `json:"closed"`
	Modified          string        `json:"modified"`
	LastModify        string        `json:"lastmodify"`
	Auditer           string        `json:"auditer"`
	De                string        `json:"de"`
	Fixer             string        `json:"fixer"`
	VersionTest       string        `json:"version_test"`
	VersionReport     string        `json:"version_report"`
	VersionClose      string        `json:"version_close"`
	VersionFix        string        `json:"version_fix"`
	BaselineFind      string        `json:"baseline_find"`
	BaselineJoin      string        `json:"baseline_join"`
	BaselineClose     string        `json:"baseline_close"`
	BaselineTest      string        `json:"baseline_test"`
	SourcePhase       string        `json:"sourcephase"`
	Te                string        `json:"te"`
	CurrentOwner      string        `json:"current_owner"`
	IterationID       string        `json:"iteration_id"`
	Resolution        string        `json:"resolution"`
	Source            string        `json:"source"`
	OriginPhase       string        `json:"originphase"`
	Confirmer         string        `json:"confirmer"`
	Milestone         string        `json:"milestone"`
	Participator      string        `json:"participator"`
	Closer            string        `json:"closer"`
	Platform          string        `json:"platform"`
	Os                string        `json:"os"`
	TestType          string        `json:"testtype"`
	TestPhase         string        `json:"testphase"`
	Frequency         string        `json:"frequency"`
	CC                string        `json:"cc"`
	RegressionNumber  string        `json:"regression_number"`
	Flows             string        `json:"flows"`
	Feature           string        `json:"feature"`
	TestMode          string        `json:"testmode"`
	Estimate          string        `json:"estimate"`
	IssueID           string        `json:"issue_id"`
	CreatedFrom       string        `json:"created_from"`
	ReleaseID         string        `json:"release_id"`
	VerifyTime        string        `json:"verify_time"`
	RejectTime        string        `json:"reject_time"`
	ReopenTime        string        `json:"reopen_time"`
	AuditTime         string        `json:"audit_time"`
	SuspendTime       string        `json:"suspend_time"`
	Due               string        `json:"due"`
	Begin             string        `json:"begin"`
	Deadline          string        `json:"deadline"`
	InProgressTime    string        `json:"in_progress_time"`
	AssignedTime      string        `json:"assigned_time"`
	TemplateID        string        `json:"template_id"`
	StoryID           string        `json:"story_id"`
	Label             string        `json:"label"`
	Size              string        `json:"size"`
	Effort            string        `json:"effort"`
	EffortCompleted   string        `json:"effort_completed"`
	Exceed            string        `json:"exceed"`
	Remain            string        `json:"remain"`
	CustomFieldOne    string        `json:"custom_field_one"`
	CustomFieldTwo    string        `json:"custom_field_two"`
	CustomFieldThree  string        `json:"custom_field_three"`
	CustomFieldFour   string        `json:"custom_field_four"`
	CustomFieldFive   string        `json:"custom_field_five"`
	CustomField6      string        `json:"custom_field_6"`
	CustomField7      string        `json:"custom_field_7"`
	CustomField8      string        `json:"custom_field_8"`
	CustomField9      string        `json:"custom_field_9"`
	CustomField10     string        `json:"custom_field_10"`
	CustomField11     string        `json:"custom_field_11"`
	CustomField12     string        `json:"custom_field_12"`
	CustomField13     string        `json:"custom_field_13"`
	CustomField14     string        `json:"custom_field_14"`
	CustomField15     string        `json:"custom_field_15"`
	CustomField16     string        `json:"custom_field_16"`
	CustomField17     string        `json:"custom_field_17"`
	CustomField18     string        `json:"custom_field_18"`
	CustomField19     string        `json:"custom_field_19"`
	CustomField20     string        `json:"custom_field_20"`
	CustomField21     string        `json:"custom_field_21"`
	CustomField22     string        `json:"custom_field_22"`
	CustomField23     string        `json:"custom_field_23"`
	CustomField24     string        `json:"custom_field_24"`
	CustomField25     string        `json:"custom_field_25"`
	CustomField26     string        `json:"custom_field_26"`
	CustomField27     string        `json:"custom_field_27"`
	CustomField28     string        `json:"custom_field_28"`
	CustomField29     string        `json:"custom_field_29"`
	CustomField30     string        `json:"custom_field_30"`
	CustomField31     string        `json:"custom_field_31"`
	CustomField32     string        `json:"custom_field_32"`
	CustomField33     string        `json:"custom_field_33"`
	CustomField34     string        `json:"custom_field_34"`
	CustomField35     string        `json:"custom_field_35"`
	CustomField36     string        `json:"custom_field_36"`
	CustomField37     string        `json:"custom_field_37"`
	CustomField38     string        `json:"custom_field_38"`
	CustomField39     string        `json:"custom_field_39"`
	CustomField40     string        `json:"custom_field_40"`
	CustomField41     string        `json:"custom_field_41"`
	CustomField42     string        `json:"custom_field_42"`
	CustomField43     string        `json:"custom_field_43"`
	CustomField44     string        `json:"custom_field_44"`
	CustomField45     string        `json:"custom_field_45"`
	CustomField46     string        `json:"custom_field_46"`
	CustomField47     string        `json:"custom_field_47"`
	CustomField48     string        `json:"custom_field_48"`
	CustomField49     string        `json:"custom_field_49"`
	CustomField50     string        `json:"custom_field_50"`
	CustomField51     string        `json:"custom_field_51"`
	CustomField52     string        `json:"custom_field_52"`
	CustomField53     string        `json:"custom_field_53"`
	CustomField54     string        `json:"custom_field_54"`
	CustomField55     string        `json:"custom_field_55"`
	CustomField56     string        `json:"custom_field_56"`
	CustomField57     string        `json:"custom_field_57"`
	CustomField58     string        `json:"custom_field_58"`
	CustomField59     string        `json:"custom_field_59"`
	CustomField60     string        `json:"custom_field_60"`
	CustomField61     string        `json:"custom_field_61"`
	CustomField62     string        `json:"custom_field_62"`
	CustomField63     string        `json:"custom_field_63"`
	CustomField64     string        `json:"custom_field_64"`
	CustomField65     string        `json:"custom_field_65"`
	CustomField66     string        `json:"custom_field_66"`
	CustomField67     string        `json:"custom_field_67"`
	CustomField68     string        `json:"custom_field_68"`
	CustomField69     string        `json:"custom_field_69"`
	CustomField70     string        `json:"custom_field_70"`
	CustomField71     string        `json:"custom_field_71"`
	CustomField72     string        `json:"custom_field_72"`
	CustomField73     string        `json:"custom_field_73"`
	CustomField74     string        `json:"custom_field_74"`
	CustomField75     string        `json:"custom_field_75"`
	CustomField76     string        `json:"custom_field_76"`
	CustomField77     string        `json:"custom_field_77"`
	CustomField78     string        `json:"custom_field_78"`
	CustomField79     string        `json:"custom_field_79"`
	CustomField80     string        `json:"custom_field_80"`
	CustomField81     string        `json:"custom_field_81"`
	CustomField82     string        `json:"custom_field_82"`
	CustomField83     string        `json:"custom_field_83"`
	CustomField84     string        `json:"custom_field_84"`
	CustomField85     string        `json:"custom_field_85"`
	CustomField86     string        `json:"custom_field_86"`
	CustomField87     string        `json:"custom_field_87"`
	CustomField88     string        `json:"custom_field_88"`
	CustomField89     string        `json:"custom_field_89"`
	CustomField90     string        `json:"custom_field_90"`
	CustomField91     string        `json:"custom_field_91"`
	CustomField92     string        `json:"custom_field_92"`
	CustomField93     string        `json:"custom_field_93"`
	CustomField94     string        `json:"custom_field_94"`
	CustomField95     string        `json:"custom_field_95"`
	CustomField96     string        `json:"custom_field_96"`
	CustomField97     string        `json:"custom_field_97"`
	CustomField98     string        `json:"custom_field_98"`
	CustomField99     string        `json:"custom_field_99"`
	CustomField100    string        `json:"custom_field_100"`
	CustomField101    string        `json:"custom_field_101"`
	CustomField102    string        `json:"custom_field_102"`
	CustomField103    string        `json:"custom_field_103"`
	CustomField104    string        `json:"custom_field_104"`
	CustomField105    string        `json:"custom_field_105"`
	CustomField106    string        `json:"custom_field_106"`
	CustomField107    string        `json:"custom_field_107"`
	CustomField108    string        `json:"custom_field_108"`
	CustomField109    string        `json:"custom_field_109"`
	CustomField110    string        `json:"custom_field_110"`
	CustomField111    string        `json:"custom_field_111"`
	CustomField112    string        `json:"custom_field_112"`
	CustomField113    string        `json:"custom_field_113"`
	CustomField114    string        `json:"custom_field_114"`
	CustomField115    string        `json:"custom_field_115"`
	CustomField116    string        `json:"custom_field_116"`
	CustomField117    string        `json:"custom_field_117"`
	CustomField118    string        `json:"custom_field_118"`
	CustomField119    string        `json:"custom_field_119"`
	CustomField120    string        `json:"custom_field_120"`
	CustomField121    string        `json:"custom_field_121"`
	CustomField122    string        `json:"custom_field_122"`
	CustomField123    string        `json:"custom_field_123"`
	CustomField124    string        `json:"custom_field_124"`
	CustomField125    string        `json:"custom_field_125"`
	CustomField126    string        `json:"custom_field_126"`
	CustomField127    string        `json:"custom_field_127"`
	CustomField128    string        `json:"custom_field_128"`
	CustomField129    string        `json:"custom_field_129"`
	CustomField130    string        `json:"custom_field_130"`
	CustomField131    string        `json:"custom_field_131"`
	CustomField132    string        `json:"custom_field_132"`
	CustomField133    string        `json:"custom_field_133"`
	CustomField134    string        `json:"custom_field_134"`
	CustomField135    string        `json:"custom_field_135"`
	CustomField136    string        `json:"custom_field_136"`
	CustomField137    string        `json:"custom_field_137"`
	CustomField138    string        `json:"custom_field_138"`
	CustomField139    string        `json:"custom_field_139"`
	CustomField140    string        `json:"custom_field_140"`
	CustomField141    string        `json:"custom_field_141"`
	CustomField142    string        `json:"custom_field_142"`
	CustomField143    string        `json:"custom_field_143"`
	CustomField144    string        `json:"custom_field_144"`
	CustomField145    string        `json:"custom_field_145"`
	CustomField146    string        `json:"custom_field_146"`
	CustomField147    string        `json:"custom_field_147"`
	CustomField148    string        `json:"custom_field_148"`
	CustomField149    string        `json:"custom_field_149"`
	CustomField150    string        `json:"custom_field_150"`
	CustomPlanField1  string        `json:"custom_plan_field_1"`
	CustomPlanField2  string        `json:"custom_plan_field_2"`
	CustomPlanField3  string        `json:"custom_plan_field_3"`
	CustomPlanField4  string        `json:"custom_plan_field_4"`
	CustomPlanField5  string        `json:"custom_plan_field_5"`
	CustomPlanField6  string        `json:"custom_plan_field_6"`
	CustomPlanField7  string        `json:"custom_plan_field_7"`
	CustomPlanField8  string        `json:"custom_plan_field_8"`
	CustomPlanField9  string        `json:"custom_plan_field_9"`
	CustomPlanField10 string        `json:"custom_plan_field_10"`
	PriorityLabel     PriorityLabel `json:"priority_label"`
	WorkspaceID       string        `json:"workspace_id"`
}

// BugService todo: add more methods
type BugService struct {
	client *Client
}

// 创建缺陷
// 复制缺陷
// 获取缺陷变更历史
// 获取缺陷变更次数
// 获取缺陷自定义字段配置

// GetBugs 获取缺陷
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/get_bugs.html
func (s *BugService) GetBugs(
	ctx context.Context, request *GetBugsRequest, opts ...RequestOption,
) ([]*Bug, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "bugs", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		Bug *Bug `json:"Bug"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	bugs := make([]*Bug, 0, len(items))
	for _, item := range items {
		bugs = append(bugs, item.Bug)
	}

	return bugs, resp, nil
}

type GetBugsRequest struct {
	ID                *Multi[int]    `url:"id,omitempty"`               // ID 支持多ID查询
	Title             *string        `url:"title,omitempty"`            // 标题 支持模糊匹配
	Priority          *string        `url:"priority,omitempty"`         // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
	PriorityLabel     *PriorityLabel `url:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	Severity          *string        `url:"severity,omitempty"`         // 严重程度 支持枚举查询
	Status            *Enum[string]  `url:"status,omitempty"`           // 状态 支持不等于查询、枚举查询
	VStatus           *string        `url:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	Label             *Enum[string]  `url:"label,omitempty"`            // 标签查询 支持枚举查询
	IterationID       *Enum[string]  `url:"iteration_id,omitempty"`     // 迭代 支持枚举查询
	Module            *Enum[string]  `url:"module,omitempty"`           // 模块 支持枚举查询
	ReleaseID         *int           `url:"release_id,omitempty"`       // 发布计划
	VersionReport     *Enum[string]  `url:"version_report,omitempty"`   // 发现版本 枚举查询
	VersionTest       *string        `url:"version_test,omitempty"`     // 验证版本
	VersionFix        *string        `url:"version_fix,omitempty"`      // 合入版本
	VersionClose      *string        `url:"version_close,omitempty"`    // 关闭版本
	BaselineFind      *string        `url:"baseline_find,omitempty"`    // 发现基线
	BaselineJoin      *string        `url:"baseline_join,omitempty"`    // 合入基线
	BaselineTest      *string        `url:"baseline_test,omitempty"`    // 验证基线
	BaselineClose     *string        `url:"baseline_close,omitempty"`   // 关闭基线
	Feature           *string        `url:"feature,omitempty"`          // 特性
	CurrentOwner      *string        `url:"current_owner,omitempty"`    // 处理人 支持模糊匹配
	CC                *string        `url:"cc,omitempty"`               // 抄送人
	Reporter          *Multi[string] `url:"reporter,omitempty"`         // 创建人 支持多人员查询
	Participator      *Multi[string] `url:"participator,omitempty"`     // 参与人 支持多人员查询
	TE                *string        `url:"te,omitempty"`               // 测试人员 支持模糊匹配
	DE                *string        `url:"de,omitempty"`               // 开发人员 支持模糊匹配
	Auditer           *string        `url:"auditer,omitempty"`          // 审核人
	Confirmer         *string        `url:"confirmer,omitempty"`        // 验证人
	Fixer             *string        `url:"fixer,omitempty"`            // 修复人
	Closer            *string        `url:"closer,omitempty"`           // 关闭人
	LastModify        *string        `url:"lastmodify,omitempty"`       // 最后修改人
	Created           *string        `url:"created,omitempty"`          // 创建时间 支持时间查询
	InProgressTime    *string        `url:"in_progress_time,omitempty"` // 接受处理时间 支持时间查询
	Resolved          *string        `url:"resolved,omitempty"`         // 解决时间 支持时间查询
	VerifyTime        *string        `url:"verify_time,omitempty"`      // 验证时间 支持时间查询
	Closed            *string        `url:"closed,omitempty"`           // 关闭时间 支持时间查询
	RejectTime        *string        `url:"reject_time,omitempty"`      // 拒绝时间 支持时间查询
	Modified          *string        `url:"modified,omitempty"`         // 最后修改时间 支持时间查询
	Begin             *string        `url:"begin,omitempty"`            // 预计开始
	Due               *string        `url:"due,omitempty"`              // 预计结束
	Deadline          *string        `url:"deadline,omitempty"`         // 解决期限
	OS                *string        `url:"os,omitempty"`               // 操作系统
	Platform          *string        `url:"platform,omitempty"`         // 软件平台
	TestMode          *string        `url:"testmode,omitempty"`         // 测试方式
	TestPhase         *string        `url:"testphase,omitempty"`        // 测试阶段
	TestType          *string        `url:"testtype,omitempty"`         // 测试类型
	Source            *Enum[string]  `url:"source,omitempty"`           // 缺陷根源 支持枚举查询
	BugType           *string        `url:"bugtype,omitempty"`          // 缺陷类型
	Frequency         *Enum[string]  `url:"frequency,omitempty"`        // 重现规律 支持枚举查询
	OriginPhase       *string        `url:"originphase,omitempty"`      // 发现阶段
	SourcePhase       *string        `url:"sourcephase,omitempty"`      // 引入阶段
	Resolution        *Enum[string]  `url:"resolution,omitempty"`       // 解决方法 支持枚举查询
	Estimate          *int           `url:"estimate,omitempty"`         // 预计解决时间
	Description       *string        `url:"description,omitempty"`      // 详细描述 支持模糊匹配
	WorkspaceID       *int           `url:"workspace_id,omitempty"`     // 项目ID
	CustomFieldOne    *string        `url:"custom_field_one,omitempty"` // 自定义字段参数，具体字段名通过接口 获取缺陷自定义字段配置 获取 支持枚举查询
	CustomFieldTwo    *string        `url:"custom_field_two,omitempty"`
	CustomFieldThree  *string        `url:"custom_field_three,omitempty"`
	CustomFieldFour   *string        `url:"custom_field_four,omitempty"`
	CustomFieldFive   *string        `url:"custom_field_five,omitempty"`
	CustomField6      *string        `url:"custom_field_6,omitempty"`
	CustomField7      *string        `url:"custom_field_7,omitempty"`
	CustomField8      *string        `url:"custom_field_8,omitempty"`
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
	CustomPlanField1  *string        `url:"custom_plan_field_1,omitempty"` // 自定义计划应用参数，具体字段名通过接口 获取自定义计划应用 获取
	CustomPlanField2  *string        `url:"custom_plan_field_2,omitempty"`
	CustomPlanField3  *string        `url:"custom_plan_field_3,omitempty"`
	CustomPlanField4  *string        `url:"custom_plan_field_4,omitempty"`
	CustomPlanField5  *string        `url:"custom_plan_field_5,omitempty"`
	CustomPlanField6  *string        `url:"custom_plan_field_6,omitempty"`
	CustomPlanField7  *string        `url:"custom_plan_field_7,omitempty"`
	CustomPlanField8  *string        `url:"custom_plan_field_8,omitempty"`
	CustomPlanField9  *string        `url:"custom_plan_field_9,omitempty"`
	CustomPlanField10 *string        `url:"custom_plan_field_10,omitempty"`
	Limit             *int           `url:"limit,omitempty"`  // 设置返回数量限制，默认为30
	Page              *int           `url:"page,omitempty"`   // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Order             *Order         `url:"order,omitempty"`  // 排序规则，规则：字段名 ASC或者DESC，然后 urlencode 如按创建时间逆序：order=created%20desc
	Fields            *Multi[string] `url:"fields,omitempty"` // 设置获取的字段，多个字段间以','逗号隔开
}

// 获取缺陷数量
// 获取缺陷与其它缺陷的所有关联关系
// 获取缺陷模板列表
// 获取缺陷模板字段
// 获取视图对应的缺陷列表
// 获取缺陷所有字段及候选值
// 获取缺陷所有字段的中英文

// UpdateBug 更新缺陷
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/bug/update_bug.html
func (s *BugService) UpdateBug(
	ctx context.Context, request *UpdateBugRequest, opts ...RequestOption,
) (*Bug, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, "bugs", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var item struct {
		Bug *Bug `json:"Bug"`
	}
	resp, err := s.client.Do(req, &item)
	if err != nil {
		return nil, resp, err
	}

	return item.Bug, resp, nil
}

type UpdateBugRequest struct {
	ID                *int           `json:"id,omitempty"`               // ID 支持多ID查询
	Title             *string        `json:"title,omitempty"`            // 标题 支持模糊匹配
	Priority          *string        `json:"priority,omitempty"`         // 优先级。为了兼容自定义优先级，请使用 priority_label 字段，详情参考：如何兼容自定义优先级
	PriorityLabel     *PriorityLabel `json:"priority_label,omitempty"`   // 优先级。推荐使用这个字段
	Severity          *string        `json:"severity,omitempty"`         // 严重程度 支持枚举查询
	Status            *Enum[string]  `json:"status,omitempty"`           // 状态 支持不等于查询、枚举查询
	VStatus           *string        `json:"v_status,omitempty"`         // 状态(支持传入中文状态名称)
	Label             *Enum[string]  `json:"label,omitempty"`            // 标签查询 支持枚举查询
	IterationID       *Enum[string]  `json:"iteration_id,omitempty"`     // 迭代 支持枚举查询
	Module            *Enum[string]  `json:"module,omitempty"`           // 模块 支持枚举查询
	ReleaseID         *int           `json:"release_id,omitempty"`       // 发布计划
	VersionReport     *Enum[string]  `json:"version_report,omitempty"`   // 发现版本 枚举查询
	VersionTest       *string        `json:"version_test,omitempty"`     // 验证版本
	VersionFix        *string        `json:"version_fix,omitempty"`      // 合入版本
	VersionClose      *string        `json:"version_close,omitempty"`    // 关闭版本
	BaselineFind      *string        `json:"baseline_find,omitempty"`    // 发现基线
	BaselineJoin      *string        `json:"baseline_join,omitempty"`    // 合入基线
	BaselineTest      *string        `json:"baseline_test,omitempty"`    // 验证基线
	BaselineClose     *string        `json:"baseline_close,omitempty"`   // 关闭基线
	Feature           *string        `json:"feature,omitempty"`          // 特性
	CurrentOwner      *string        `json:"current_owner,omitempty"`    // 处理人 支持模糊匹配
	CC                *string        `json:"cc,omitempty"`               // 抄送人
	Reporter          *Multi[string] `json:"reporter,omitempty"`         // 创建人 支持多人员查询
	Participator      *Multi[string] `json:"participator,omitempty"`     // 参与人 支持多人员查询
	TE                *string        `json:"te,omitempty"`               // 测试人员 支持模糊匹配
	DE                *string        `json:"de,omitempty"`               // 开发人员 支持模糊匹配
	Auditer           *string        `json:"auditer,omitempty"`          // 审核人
	Confirmer         *string        `json:"confirmer,omitempty"`        // 验证人
	Fixer             *string        `json:"fixer,omitempty"`            // 修复人
	Closer            *string        `json:"closer,omitempty"`           // 关闭人
	LastModify        *string        `json:"lastmodify,omitempty"`       // 最后修改人
	Created           *string        `json:"created,omitempty"`          // 创建时间 支持时间查询
	InProgressTime    *string        `json:"in_progress_time,omitempty"` // 接受处理时间 支持时间查询
	Resolved          *string        `json:"resolved,omitempty"`         // 解决时间 支持时间查询
	VerifyTime        *string        `json:"verify_time,omitempty"`      // 验证时间 支持时间查询
	Closed            *string        `json:"closed,omitempty"`           // 关闭时间 支持时间查询
	RejectTime        *string        `json:"reject_time,omitempty"`      // 拒绝时间 支持时间查询
	Modified          *string        `json:"modified,omitempty"`         // 最后修改时间 支持时间查询
	Begin             *string        `json:"begin,omitempty"`            // 预计开始
	Due               *string        `json:"due,omitempty"`              // 预计结束
	Deadline          *string        `json:"deadline,omitempty"`         // 解决期限
	OS                *string        `json:"os,omitempty"`               // 操作系统
	Platform          *string        `json:"platform,omitempty"`         // 软件平台
	TestMode          *string        `json:"testmode,omitempty"`         // 测试方式
	TestPhase         *string        `json:"testphase,omitempty"`        // 测试阶段
	TestType          *string        `json:"testtype,omitempty"`         // 测试类型
	Source            *Enum[string]  `json:"source,omitempty"`           // 缺陷根源 支持枚举查询
	BugType           *string        `json:"bugtype,omitempty"`          // 缺陷类型
	Frequency         *Enum[string]  `json:"frequency,omitempty"`        // 重现规律 支持枚举查询
	OriginPhase       *string        `json:"originphase,omitempty"`      // 发现阶段
	SourcePhase       *string        `json:"sourcephase,omitempty"`      // 引入阶段
	Resolution        *Enum[string]  `json:"resolution,omitempty"`       // 解决方法 支持枚举查询
	Estimate          *int           `json:"estimate,omitempty"`         // 预计解决时间
	Description       *string        `json:"description,omitempty"`      // 详细描述 支持模糊匹配
	WorkspaceID       *int           `json:"workspace_id,omitempty"`     // 项目ID
	CustomFieldOne    *string        `json:"custom_field_one,omitempty"` // 自定义字段参数，具体字段名通过接口 获取缺陷自定义字段配置 获取 支持枚举查询
	CustomFieldTwo    *string        `json:"custom_field_two,omitempty"`
	CustomFieldThree  *string        `json:"custom_field_three,omitempty"`
	CustomFieldFour   *string        `json:"custom_field_four,omitempty"`
	CustomFieldFive   *string        `json:"custom_field_five,omitempty"`
	CustomField6      *string        `json:"custom_field_6,omitempty"`
	CustomField7      *string        `json:"custom_field_7,omitempty"`
	CustomField8      *string        `json:"custom_field_8,omitempty"`
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
	CustomPlanField1  *string        `json:"custom_plan_field_1,omitempty"` // 自定义计划应用参数，具体字段名通过接口 获取自定义计划应用 获取
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

// 获取回收站下的缺陷
// 获取缺陷关联的需求ID
// 转换缺陷ID成列表queryToken
// 缺陷说明
