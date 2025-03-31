package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelService_GetLabels(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/label", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "111,222", r.URL.Query().Get("id"))
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "tapd-clientID", r.URL.Query().Get("creator"))
		assert.Equal(t, "2024-08-26", r.URL.Query().Get("created"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id asc", r.URL.Query().Get("order"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/label/get_labels.json"))
	}))

	labels, _, err := client.LabelService.GetLabels(ctx, &GetLabelsRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti(111, 222),
		Name:        Ptr("test"),
		Creator:     Ptr("tapd-clientID"),
		Created:     Ptr("2024-08-26"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Order:       NewOrder("id", OrderByAsc),
	})
	assert.NoError(t, err)
	assert.Len(t, labels, 9)
	assert.Equal(t, "1134190502001000811", labels[0].ID)
	assert.Equal(t, "11112222", labels[0].WorkspaceID)
	assert.Equal(t, "test2", labels[0].Name)
	assert.Equal(t, LabelColor2, labels[0].Color)
	assert.Equal(t, "workspace", labels[0].Category)
	assert.Equal(t, "tapd-clientID", labels[0].Creator)
	assert.Equal(t, "tapd-clientID", labels[0].Modifier)
	assert.Equal(t, "2024-08-26 21:38:29", labels[0].Created)
	assert.Equal(t, "2024-08-26 21:41:59", labels[0].Modified)
	assert.Equal(t, "#FF6770", labels[0].ColorValue)
}

func TestLabelService_GetLabelCount(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/label/count", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "111,222", r.URL.Query().Get("id"))
		assert.Equal(t, "test", r.URL.Query().Get("name"))
		assert.Equal(t, "tapd-clientID", r.URL.Query().Get("creator"))
		assert.Equal(t, "2024-08-26", r.URL.Query().Get("created"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/label/get_label_count.json"))
	}))

	count, _, err := client.LabelService.GetLabelsCount(ctx, &GetLabelCountRequest{
		WorkspaceID: Ptr(11112222),
		ID:          NewMulti(111, 222),
		Name:        Ptr("test"),
		Creator:     Ptr("tapd-clientID"),
		Created:     Ptr("2024-08-26"),
	})
	assert.NoError(t, err)
	assert.Equal(t, 15, count)
}

func TestLabelService_CreateLabel(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/label", r.URL.Path)

		var req struct {
			WorkspaceID int        `json:"workspace_id"`
			Name        string     `json:"name"`
			Color       LabelColor `json:"color"`
			Creator     string     `json:"creator"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))

		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, "test", req.Name)
		assert.Equal(t, LabelColor1, req.Color)
		assert.Equal(t, "tapd-clientID", req.Creator)

		_, _ = w.Write(loadData(t, "internal/testdata/api/label/create_label.json"))
	}))

	label, _, err := client.LabelService.CreateLabel(ctx, &CreateLabelRequest{
		WorkspaceID: Ptr(11112222),
		Name:        Ptr("test"),
		Color:       Ptr(LabelColor1),
		Creator:     Ptr("tapd-clientID"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1134190502001000812", label.ID)
	assert.Equal(t, "1111222", label.WorkspaceID)
	assert.Equal(t, "test3", label.Name)
	assert.Equal(t, LabelColor1, label.Color)
	assert.Equal(t, "workspace", label.Category)
	assert.Equal(t, "tapd-clientID", label.Creator)
	assert.Equal(t, "", label.Modifier)
	assert.Equal(t, "2024-08-26 22:33:14", label.Created)
	assert.Equal(t, "2024-08-26 22:33:14", label.Modified)
	assert.Equal(t, "", label.ColorValue)
}

func TestLabelService_UpdateLabel(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/label", r.URL.Path)

		var req struct {
			ID          int        `json:"id"`
			WorkspaceID int        `json:"workspace_id"`
			Color       LabelColor `json:"color"`
			Modifier    string     `json:"modifier"`
		}
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))

		assert.Equal(t, 1134190502001000812, req.ID)
		assert.Equal(t, 11112222, req.WorkspaceID)
		assert.Equal(t, LabelColor1, req.Color)
		assert.Equal(t, "tapd-clientID", req.Modifier)

		_, _ = w.Write(loadData(t, "internal/testdata/api/label/update_label.json"))
	}))

	label, _, err := client.LabelService.UpdateLabel(ctx, &UpdateLabelRequest{
		ID:          Ptr(1134190502001000812),
		WorkspaceID: Ptr(11112222),
		Color:       Ptr(LabelColor1),
		Modifier:    Ptr("tapd-clientID"),
	})
	assert.NoError(t, err)
	assert.Equal(t, "1134190502001000812", label.ID)
	assert.Equal(t, "1111222", label.WorkspaceID)
	assert.Equal(t, "test3", label.Name)
	assert.Equal(t, LabelColor1, label.Color)
	assert.Equal(t, "workspace", label.Category)
	assert.Equal(t, "tapd-clientID", label.Creator)
	assert.Equal(t, "", label.Modifier)
	assert.Equal(t, "2024-08-26 22:33:14", label.Created)
	assert.Equal(t, "2024-08-26 22:33:14", label.Modified)
	assert.Equal(t, "", label.ColorValue)
}
