package tapd

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachmentService_UploadAttachment(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/files/upload_attachment", r.URL.Path)
		assert.Contains(t, r.Header.Get("Content-Type"), "multipart/form-data; boundary=")

		err := r.ParseMultipartForm(1024)
		assert.NoError(t, err)
		assert.Equal(t, "11112222", r.FormValue("workspace_id"))
		assert.Equal(t, "story_custom_field", r.FormValue("type"))
		assert.Equal(t, "custom_field_one", r.FormValue("custom_field"))
		assert.Equal(t, "1069993260856110917", r.FormValue("entry_id"))
		assert.Equal(t, "go-tapd", r.FormValue("owner"))

		file, header, err := r.FormFile("file")
		assert.NoError(t, err)
		defer file.Close() //nolint:errcheck
		assert.Equal(t, "orangetest.jpg", header.Filename)

		content, err := io.ReadAll(file)
		assert.NoError(t, err)
		assert.Equal(t, "demo image content", string(content))

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/upload_attachment.json"))
	}))

	attachment, _, err := client.AttachmentService.UploadAttachment(ctx, &UploadAttachmentRequest{
		WorkspaceID: Ptr(11112222),
		Type:        Ptr("story_custom_field"),
		CustomField: Ptr("custom_field_one"),
		EntryID:     Ptr[int64](1069993260856110917),
		Owner:       Ptr("go-tapd"),
		Filename:    Ptr("orangetest.jpg"),
		File:        strings.NewReader("demo image content"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1069993260503455439", attachment.ID)
	assert.Equal(t, "story_custom_field", attachment.Type)
	assert.Equal(t, "1069993260856110917", attachment.EntryID)
	assert.Equal(t, "orangetest.jpg", attachment.Filename)
	assert.Equal(t, "image/jpeg", attachment.ContentType)
	assert.Equal(t, "2023-07-07 21:36:08", attachment.Created)
	assert.Equal(t, "69993260", attachment.WorkspaceID)
	assert.Equal(t, "go-tapd", attachment.Owner)
}

func TestAttachmentService_UploadImageBase64(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/files/upload_image_base64", r.URL.Path)
		assert.Contains(t, r.Header.Get("Content-Type"), "multipart/form-data; boundary=")

		err := r.ParseMultipartForm(1024)
		assert.NoError(t, err)
		assert.Equal(t, "69995768", r.FormValue("workspace_id"))
		assert.Equal(t, "story_custom_field", r.FormValue("type"))
		assert.Equal(t, "custom_field_one", r.FormValue("custom_field"))
		assert.Equal(t, "1069995768115415038", r.FormValue("entry_id"))
		assert.Equal(t, "go-tapd", r.FormValue("owner"))
		assert.Equal(t, "base64-image-data", r.FormValue("base64_data"))

		_, _, err = r.FormFile("file")
		assert.Error(t, err)

		_, _ = w.Write(loadData(t, "internal/testdata/api/attachment/upload_image_base64.json"))
	}))

	attachment, _, err := client.AttachmentService.UploadImageBase64(ctx, &UploadImageBase64Request{
		WorkspaceID: Ptr(69995768),
		Type:        Ptr("story_custom_field"),
		CustomField: Ptr("custom_field_one"),
		EntryID:     Ptr[int64](1069995768115415038),
		Owner:       Ptr("go-tapd"),
		Base64Data:  Ptr("base64-image-data"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1069995768523406033", attachment.ID)
	assert.Equal(t, "story_custom_field", attachment.Type)
	assert.Equal(t, "1069995768115415038", attachment.EntryID)
	assert.Equal(t, "tapd_base64_3031935_1701845640.jpg", attachment.Filename)
	assert.Equal(t, "image/png", attachment.ContentType)
	assert.Equal(t, "2023-12-05 14:54:01", attachment.Created)
}

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
