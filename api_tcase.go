package tapd

import (
	"context"
	"net/http"
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

type (
	// TestCase 测试用例
	TestCase struct {
		ID                 string         `json:"id,omitempty"`                  // 测试用例ID
		MID                string         `json:"mid,omitempty"`                 // 测试用例MID
		Steps              *string        `json:"steps,omitempty"`               // 用例步骤
		WorkspaceID        string         `json:"workspace_id,omitempty"`        // 项目ID
		CategoryID         string         `json:"category_id,omitempty"`         // 用例目录
		Version            string         `json:"version,omitempty"`             // 版本
		Created            string         `json:"created,omitempty"`             // 创建时间
		Modifier           string         `json:"modifier,omitempty"`            // 最后修改人
		Modified           string         `json:"modified,omitempty"`            // 最后修改时间
		Creator            string         `json:"creator,omitempty"`             // 创建人
		Status             TestCaseStatus `json:"status,omitempty"`              // 用例状态
		Name               string         `json:"name,omitempty"`                // 用例名称
		Precondition       *string        `json:"precondition,omitempty"`        // 前置条件
		Expectation        *string        `json:"expectation,omitempty"`         // 预期结果
		Sort               string         `json:"sort,omitempty"`                // 排序
		IndexCode          string         `json:"indexcode,omitempty"`           // 编号
		Type               string         `json:"type,omitempty"`                // 用例类型
		Priority           string         `json:"priority,omitempty"`            // 用例等级
		IsAutomated        string         `json:"is_automated,omitempty"`        // 是否自动化
		AutomationType     string         `json:"automation_type,omitempty"`     // 自动化类型
		AutomationPlatform string         `json:"automation_platform,omitempty"` // 自动化平台
		IsServing          string         `json:"is_serving,omitempty"`          // 是否服务中
		TemplateID         string         `json:"template_id,omitempty"`         // 模板ID
		CreatedFrom        string         `json:"created_from,omitempty"`        // 创建来源
		CustomField1       *string        `json:"custom_field_1,omitempty"`      // 自定义字段
		CustomField2       *string        `json:"custom_field_2,omitempty"`
		CustomField3       *string        `json:"custom_field_3,omitempty"`
		CustomField4       *string        `json:"custom_field_4,omitempty"`
		CustomField5       *string        `json:"custom_field_5,omitempty"`
		CustomField6       *string        `json:"custom_field_6,omitempty"`
		CustomField7       *string        `json:"custom_field_7,omitempty"`
		CustomField8       *string        `json:"custom_field_8,omitempty"`
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
