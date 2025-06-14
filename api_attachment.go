package tapd

import (
	"context"
	"net/http"
)

type (
	// Attachment 附件
	Attachment struct {
		ID          string `json:"id,omitempty"`           // 附件ID
		Type        string `json:"type,omitempty"`         // 类型
		EntryID     string `json:"entry_id,omitempty"`     // 依赖对象ID
		Filename    string `json:"filename,omitempty"`     // 附件名称
		Description string `json:"description,omitempty"`  // 描述
		ContentType string `json:"content_type,omitempty"` // 内容类型
		Created     string `json:"created,omitempty"`      // 创建时间
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		Owner       string `json:"owner,omitempty"`        // 上传人
		DownloadURL string `json:"download_url,omitempty"` // 下载链接(仅在获取单个附件时返回)
	}

	GetAttachmentsRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"`  // [必须]项目ID
		ID          *int    `url:"id,omitempty"`            // [可选]ID
		Type        *string `url:"type,omitempty"`          // [可选]类型
		EntryID     *int    `url:"entry_id,omitempty"`      // [可选]依赖对象ID
		Filename    *string `url:"filename,omitempty"`      // [可选]附件名称
		Owner       *string `url:"owner,omitempty"`         // [可选]上传人
		DownloadURL string  `json:"download_url,omitempty"` // 下载链接(仅在获取单个附件时返回)
	}

	GetAttachmentDownloadURLRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
		ID          *int `url:"id,omitempty"`           // [必须]附件ID
	}

	GetImageDownloadURLRequest struct {
		WorkspaceID *int    `url:"workspace_id,omitempty"` // [必须]项目ID
		ImagePath   *string `url:"image_path,omitempty"`   // [必须]图片路径, 支持完整url地址, 图片所属项目必须和传入的项目id一致
	}

	ImageAttachment struct {
		Type        string `json:"type,omitempty"`         // 文件类型
		Value       string `json:"value,omitempty"`        // 图片路径
		WorkspaceID int    `json:"workspace_id,omitempty"` // 项目id
		Filename    string `json:"filename,omitempty"`     // 图片文件名
		DownloadURL string `json:"download_url,omitempty"` // 单个图片下载地址
	}

	GetDocumentDownloadURLRequest struct {
		WorkspaceID *int `url:"workspace_id,omitempty"` // [必须]项目ID
		ID          *int `url:"id,omitempty"`           // [必须]文档ID
	}

	DocumentAttachment struct {
		ID          string `json:"id,omitempty"`           // 文档ID
		WorkspaceID string `json:"workspace_id,omitempty"` // 项目ID
		Name        string `json:"name,omitempty"`         // 标题
		Type        string `json:"type,omitempty"`         // 文档类型
		FolderID    string `json:"folder_id,omitempty"`    // 文件夹ID
		Creator     string `json:"creator,omitempty"`      // 创建人
		Modifier    string `json:"modifier,omitempty"`     // 最后修改人
		Status      string `json:"status,omitempty"`       // 状态
		Created     string `json:"created,omitempty"`      // 创建时间
		Modified    string `json:"modified,omitempty"`     // 最后修改时间
		DownloadURL string `json:"download_url,omitempty"` // 下载链接
	}
)

// AttachmentService is the service to communicate with Attachment API.
//
// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/
type AttachmentService interface {
	// GetAttachments 获取附件
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/get_attachments.html
	GetAttachments(ctx context.Context, request *GetAttachmentsRequest, opts ...RequestOption) ([]*Attachment, *Response, error)

	// GetAttachmentDownloadURL 获取单个附件下载链接
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/get_one_attachment.html
	GetAttachmentDownloadURL(ctx context.Context, request *GetAttachmentDownloadURLRequest, opts ...RequestOption) (*Attachment, *Response, error)

	// GetImageDownloadURL 获取单个图片下载链接
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/get_image.html
	GetImageDownloadURL(ctx context.Context, request *GetImageDownloadURLRequest, opts ...RequestOption) (*ImageAttachment, *Response, error)

	// GetDocumentDownloadURL 获取单个文档下载链接
	//
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/attachment/documents_down.html
	GetDocumentDownloadURL(ctx context.Context, request *GetDocumentDownloadURLRequest, opts ...RequestOption) (*DocumentAttachment, *Response, error)
}

type attachmentService struct {
	client *Client
}

var _ AttachmentService = (*attachmentService)(nil)

func NewAttachmentService(client *Client) AttachmentService {
	return &attachmentService{
		client: client,
	}
}

func (s *attachmentService) GetAttachments(
	ctx context.Context, request *GetAttachmentsRequest, opts ...RequestOption,
) ([]*Attachment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "attachments", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []*struct {
		Attachment *Attachment `json:"attachment,omitempty"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	attachments := make([]*Attachment, 0, len(items))
	for _, item := range items {
		attachments = append(attachments, item.Attachment)
	}

	return attachments, resp, nil
}

func (s *attachmentService) GetAttachmentDownloadURL(
	ctx context.Context, request *GetAttachmentDownloadURLRequest, opts ...RequestOption,
) (*Attachment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "attachments/down", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Attachment *Attachment `json:"attachment,omitempty"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Attachment, resp, nil
}

func (s *attachmentService) GetImageDownloadURL(
	ctx context.Context, request *GetImageDownloadURLRequest, opts ...RequestOption,
) (*ImageAttachment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "files/get_image", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Attachment *ImageAttachment `json:"Attachment,omitempty"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Attachment, resp, nil
}

func (s *attachmentService) GetDocumentDownloadURL(
	ctx context.Context, request *GetDocumentDownloadURLRequest, opts ...RequestOption,
) (*DocumentAttachment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "documents/down", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var response struct {
		Document *DocumentAttachment `json:"Document,omitempty"`
	}
	resp, err := s.client.Do(req, &response)
	if err != nil {
		return nil, resp, err
	}

	return response.Document, resp, nil
}
