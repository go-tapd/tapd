package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachmentService_GetAttachments(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/attachments", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444", r.URL.Query().Get("id"))
		assert.Equal(t, "bug", r.URL.Query().Get("type"))
		assert.Equal(t, "55556666", r.URL.Query().Get("entry_id"))
		assert.Equal(t, "demo.jpg", r.URL.Query().Get("filename"))
		assert.Equal(t, "go-tapd", r.URL.Query().Get("owner"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/get_attachments.json"))
	}))

	attachments, _, err := client.AttachmentService.GetAttachments(ctx, &GetAttachmentsRequest{
		WorkspaceID: Ptr(11112222),
		ID:          Ptr(33334444),
		Type:        Ptr("bug"),
		EntryID:     Ptr(55556666),
		Filename:    Ptr("demo.jpg"),
		Owner:       Ptr("go-tapd"),
	})
	assert.NoError(t, err)
	assert.True(t, len(attachments) > 0)
	assert.Equal(t, "1111112222001002462", attachments[0].ID)
	assert.Equal(t, "bug", attachments[0].Type)
	assert.Equal(t, "1111112222001020342", attachments[0].EntryID)
	assert.Equal(t, "demo.jpg", attachments[0].Filename)
	assert.Equal(t, "this is a demo image", attachments[0].Description)
	assert.Equal(t, "image/jpeg", attachments[0].ContentType)
	assert.Equal(t, "2022-04-20 17:32:37", attachments[0].Created)
	assert.Equal(t, "11112222", attachments[0].WorkspaceID)
	assert.Equal(t, "Go-Tapd", attachments[0].Owner)
	assert.Equal(t, "", attachments[0].DownloadURL)
}

func TestAttachmentService_GetAttachmentDownloadURL(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/attachments/down", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444", r.URL.Query().Get("id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/get_attachment_download_url.json"))
	}))

	attachment, _, err := client.AttachmentService.GetAttachmentDownloadURL(ctx, &GetAttachmentDownloadURLRequest{
		WorkspaceID: Ptr(11112222),
		ID:          Ptr(33334444),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1111112222001002462", attachment.ID)
	assert.Equal(t, "bug", attachment.Type)
	assert.Equal(t, "1111112222001020342", attachment.EntryID)
	assert.Equal(t, "demo.jpg", attachment.Filename)
	assert.Equal(t, "this is a demo image", attachment.Description)
	assert.Equal(t, "image/jpeg", attachment.ContentType)
	assert.Equal(t, "2022-04-20 17:32:37", attachment.Created)
	assert.Equal(t, "11112222", attachment.WorkspaceID)
	assert.Equal(t, "Go-Tapd", attachment.Owner)
	assert.Equal(t, "https://download.com/url/demo.jpg", attachment.DownloadURL)
}

func TestAttachmentService_GetImageDownloadURL(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/files/get_image", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "/demo/demo.jpg", r.URL.Query().Get("image_path"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/get_image_download_url.json"))
	}))

	attachment, _, err := client.AttachmentService.GetImageDownloadURL(ctx, &GetImageDownloadURLRequest{
		WorkspaceID: Ptr(11112222),
		ImagePath:   Ptr("/demo/demo.jpg"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "tfl_image", attachment.Type)
	assert.Equal(t, "/tfl/pictures/202409/demo.jpg", attachment.Value)
	assert.Equal(t, 11112222, attachment.WorkspaceID)
	assert.Equal(t, "demo.jpg", attachment.Filename)
	assert.Contains(t, attachment.DownloadURL, "file.tapd.cn")
}

func TestAttachmentService_GetDocumentDownloadURL(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/documents/down", r.URL.Path)

		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "33334444", r.URL.Query().Get("id"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/get_document_download_url.json"))
	}))

	attachment, _, err := client.AttachmentService.GetDocumentDownloadURL(ctx, &GetDocumentDownloadURLRequest{
		WorkspaceID: Ptr(11112222),
		ID:          Ptr(33334444),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1134190502001000725", attachment.ID)
	assert.Equal(t, "11112222", attachment.WorkspaceID)
	assert.Equal(t, "文档功能使用秘籍", attachment.Name)
	assert.Equal(t, "word", attachment.Type)
	assert.Equal(t, "1134190502001000443", attachment.FolderID)
	assert.Equal(t, "TAPD", attachment.Creator)
	assert.Equal(t, "TAPD", attachment.Modifier)
	assert.Equal(t, "2022-06-10 10:04:13", attachment.Created)
	assert.Equal(t, "2022-06-10 10:04:13", attachment.Modified)
	assert.Contains(t, attachment.DownloadURL, "file.tapd.cn")
}
