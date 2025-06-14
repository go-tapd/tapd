package tapd

import (
	"context"
	"net/http"
)

// ReportStatus 报告状态(sent:已发送, draft:草稿, abandon:已删除)
type ReportStatus string

const (
	ReportStatusSent    ReportStatus = "sent"
	ReportStatusDraft   ReportStatus = "draft"
	ReportStatusAbandon ReportStatus = "abandon"
)

// ReportType 报告类型(normal:项目进度报告, totest:项目转测试, test:测试报告)
type ReportType string

const (
	ReportTypeNormal ReportType = "normal"
	ReportTypeToTest ReportType = "totest"
	ReportTypeTest   ReportType = "test"
)

type (
	// Report 项目报告
	Report struct {
		ID                      string       `json:"id,omitempty"`
		Title                   string       `json:"title,omitempty"`
		WorkspaceID             string       `json:"workspace_id,omitempty"`
		ReportType              ReportType   `json:"report_type,omitempty"`
		Receiver                string       `json:"receiver,omitempty"`
		Cc                      string       `json:"cc,omitempty"`
		ReceiverOrganizationIDs string       `json:"receiver_organization_ids,omitempty"`
		CcOrganizationIDs       string       `json:"cc_organization_ids,omitempty"`
		Sender                  string       `json:"sender,omitempty"`
		SendTime                string       `json:"send_time,omitempty"`
		Author                  string       `json:"author,omitempty"`
		Created                 string       `json:"created,omitempty"`
		Status                  ReportStatus `json:"status,omitempty"`
		Modified                string       `json:"modified,omitempty"`
		LastModify              string       `json:"last_modify,omitempty"`
	}

	GetReportsRequest struct {
		WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目 ID
		ID          *int           `url:"id,omitempty"`           // ID
		Title       *string        `url:"title,omitempty"`        // 标题
		Author      *string        `url:"author,omitempty"`       // 创建人
		Created     *string        `url:"created,omitempty"`      // 创建时间
		Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
		Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1（第一页）
		Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
	}
)

// ReportService is a service to interact with report related API
type ReportService interface {
	// GetReports 获取项目报告
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/report/get_workspace_reports.html
	GetReports(ctx context.Context, request *GetReportsRequest, opts ...RequestOption) ([]*Report, *Response, error)
}

type reportService struct {
	client *Client
}

var _ ReportService = (*reportService)(nil)

func NewReportService(client *Client) ReportService {
	return &reportService{
		client: client,
	}
}

func (s *reportService) GetReports(
	ctx context.Context, request *GetReportsRequest, opts ...RequestOption,
) ([]*Report, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "workspace_reports", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		WorkspaceReport *Report `json:"WorkspaceReport"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	reports := make([]*Report, 0, len(items))
	for _, item := range items {
		reports = append(reports, item.WorkspaceReport)
	}

	return reports, resp, nil
}
