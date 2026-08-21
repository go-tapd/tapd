package tapd

import "encoding/json"

// Ptr returns a pointer to the value.
//
// Deprecated: use the built-in [new] function instead, which accepts
// an expression since Go 1.26.
//
// [new]: https://pkg.go.dev/builtin#new
func Ptr[T any](v T) *T {
	return new(v)
}

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
