package tapd

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasureService_LifeTimes(t *testing.T) {
	_, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/life_times", r.URL.Path)
		assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))
		assert.Equal(t, "11112222", r.URL.Query().Get("entity_id"))
		assert.Equal(t, "story", r.URL.Query().Get("entity_type"))
		assert.Equal(t, "2024-08-26", r.URL.Query().Get("created"))
		assert.Equal(t, "10", r.URL.Query().Get("limit"))
		assert.Equal(t, "1", r.URL.Query().Get("page"))
		assert.Equal(t, "id,workspace_id", r.URL.Query().Get("fields"))

		_, _ = w.Write(loadData(t, "internal/testdata/api/measure/life-times.json"))
	}))

	lifeTimes, _, err := client.MeasureService.LifeTimes(ctx, &LifeTimesRequest{
		EntityID:    Ptr(11112222),
		EntityType:  Ptr(EntityTypeStory),
		WorkspaceID: Ptr(11112222),
		Created:     Ptr("2024-08-26"),
		Limit:       Ptr(10),
		Page:        Ptr(1),
		Fields:      NewMulti("id", "workspace_id"),
	})
	assert.NoError(t, err)
	assert.Len(t, lifeTimes, 1)
	assert.Equal(t, "1134190502001601111", lifeTimes[0].ID)
	assert.Equal(t, "11112222", lifeTimes[0].WorkspaceID)
	assert.Equal(t, EntityTypeStory, lifeTimes[0].EntityType)
	assert.Equal(t, "1134190502001601111", lifeTimes[0].EntityID)
	assert.Equal(t, "planning", lifeTimes[0].Status)
	assert.Equal(t, "", lifeTimes[0].Owner)
	assert.Equal(t, "0", lifeTimes[0].IsRepeated)
	assert.Equal(t, "2024-05-11 18:02:50", lifeTimes[0].BeginDate)
	assert.Equal(t, "2024-05-11 18:02:50", lifeTimes[0].EndDate)
	assert.Equal(t, "0", lifeTimes[0].TimeCost)
	assert.Equal(t, "", lifeTimes[0].TimeCostReduced)
	assert.Equal(t, "2024-05-11 18:02:50", lifeTimes[0].Created)
	assert.Equal(t, "Go-Tapd-Operator", lifeTimes[0].Operator)
	assert.Equal(t, "", lifeTimes[0].IsLatest)
	assert.Equal(t, "", lifeTimes[0].IsDelete)
}
