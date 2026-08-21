package tapd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelpers_Ptr(t *testing.T) {
	tests := []struct {
		name string
		v    any
	}{
		{"", 1},
		{"", "foo"},
		{"", true},
		{"", struct{}{}},
		{"", []int{1, 2, 3}},
		{"", map[string]int{"a": 1, "b": 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ptr(tt.v)
			assert.NotNil(t, got)
			assert.Equal(t, tt.v, *got)
		})
	}
}
