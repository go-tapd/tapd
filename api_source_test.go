package tapd

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSourceService_AddCodeCommitInfo(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/code_commit_infos", r.URL.Path)

		var req AddCodeCommitInfoRequest
		assert.NoError(t, json.NewDecoder(r.Body).Decode(&req))
		assert.Equal(t, 20375571, *req.WorkspaceID)
		assert.Equal(t, "zxxxxx", *req.CommitID)
		assert.Equal(t, "terrysxu", *req.Author)
		assert.Equal(t, "--story=854927829 ASA-c2s", *req.Message)
		assert.Equal(t, []string{"U xxx.php", "A xxx.js", "M xxx.html"}, *req.Files)
		assert.Equal(t, "repos/xxx_proj", *req.Repo)
		assert.Equal(t, "abcd1234-avcd-1234-avcd-1234abcdefgh", *req.RepoID)
		assert.Equal(t, "2019-07-22 19:11:11", *req.CommitTime)

		_, _ = w.Write(loadData(t, "internal/testdata/api/source/add_code_commit_info.json"))
	}))

	info, _, err := client.SourceService.AddCodeCommitInfo(ctx, &AddCodeCommitInfoRequest{
		WorkspaceID: Ptr(20375571),
		CommitID:    Ptr("zxxxxx"),
		Author:      Ptr("terrysxu"),
		Message:     Ptr("--story=854927829 ASA-c2s"),
		Files:       &[]string{"U xxx.php", "A xxx.js", "M xxx.html"},
		Repo:        Ptr("repos/xxx_proj"),
		RepoID:      Ptr("abcd1234-avcd-1234-avcd-1234abcdefgh"),
		CommitTime:  Ptr("2019-07-22 19:11:11"),
	})
	require.NoError(t, err)
	assert.Equal(t, "1020375571000321465", info.ID)
	assert.Equal(t, "20375571", info.WorkspaceID.String())
	assert.Equal(t, "zxxxxx", info.CommitID)
	require.Len(t, info.Related, 1)
	assert.Equal(t, EntityTypeStory, info.Related[0].Type)
	assert.Equal(t, "1020375571854927829", info.Related[0].ObjectID)
}

func TestSourceService_GetCodeCommitInfos(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/code_commit_infos", r.URL.Path)
		assert.Equal(t, "20358374", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "story", r.URL.Query().Get("type"))
		assert.Equal(t, "1020358374854843133", r.URL.Query().Get("object_id"))
		assert.Equal(t, "source_code", r.URL.Query().Get("related_type"))
		assert.Equal(t, "50", r.URL.Query().Get("limit"))
		assert.Equal(t, "2", r.URL.Query().Get("page"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/source/get_code_commit_infos.json"))
	}))

	infos, _, err := client.SourceService.GetCodeCommitInfos(ctx, &GetCodeCommitInfosRequest{
		WorkspaceID: Ptr(20358374),
		Type:        Ptr(EntityTypeStory),
		ObjectID:    Ptr[int64](1020358374854843133),
		RelatedType: Ptr(CodeCommitRelatedTypeSourceCode),
		Limit:       Ptr(50),
		Page:        Ptr(2),
	})
	require.NoError(t, err)
	require.Len(t, infos, 1)
	assert.Equal(t, "1020358374000262989", infos[0].ID)
	assert.Equal(t, "20358374", infos[0].WorkspaceID.String())
	assert.Equal(t, "111", infos[0].CommitID)
	assert.Equal(t, 0, infos[0].FileSort["tapd_interface_test/cases/cloud/company_settings.jmx"])
}

func TestSourceService_GetCommitObjects(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/code_commit_objects/workitems", r.URL.Path)
		assert.Equal(t, "20355782", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "7b0645c6a467a502fe1d3b678fea8bdf2890aa8d,047e5764c392bef48fd0e4176c147c7c30a9f32a", r.URL.Query().Get("commit_id"))
		assert.Equal(t, "task", r.URL.Query().Get("entity_type"))
		assert.Equal(t, "gitlab", r.URL.Query().Get("scm_type"))
		assert.Equal(t, "id,name,status", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/source/get_commit_objects.json"))
	}))

	objects, _, err := client.SourceService.GetCommitObjects(ctx, &GetCommitObjectsRequest{
		WorkspaceID: Ptr(20355782),
		CommitID: NewMulti(
			"7b0645c6a467a502fe1d3b678fea8bdf2890aa8d",
			"047e5764c392bef48fd0e4176c147c7c30a9f32a",
		),
		EntityType: Ptr(EntityTypeTask),
		SCMType:    Ptr("gitlab"),
		Fields:     NewMulti("id", "name", "status"),
	})
	require.NoError(t, err)
	require.Len(t, objects, 1)
	require.NotNil(t, objects[0].Task)
	assert.Equal(t, "1020355782500602947", objects[0].Task.ID)
	assert.Equal(t, "666", objects[0].Task.Name)
	assert.Equal(t, TaskStatusOpen, objects[0].Task.Status)
}
