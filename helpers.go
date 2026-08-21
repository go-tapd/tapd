package tapd

import "encoding/json"

func stringifyJSONRaw(raw json.RawMessage) string {
	if len(raw) == 0 || string(raw) == "null" {
		return ""
	}
	if raw[0] != '"' {
		return string(raw)
	}

	var value string
	_ = json.Unmarshal(raw, &value)
	return value
}
