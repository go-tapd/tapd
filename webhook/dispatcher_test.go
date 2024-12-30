package webhook

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func loadData(t *testing.T, filepath string) []byte {
	content, err := os.ReadFile(filepath)
	require.NoError(t, err)
	return content
}

func loadWebhookData(t *testing.T, filename string) []byte {
	return loadData(t, "../.testdata/webhook/"+filename)
}

func loadAndParseWebhookData(t *testing.T, filename string, v interface{}) {
	require.NoError(t, json.Unmarshal(loadWebhookData(t, filename), v))
}
