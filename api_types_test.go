package tapd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes_Multi(t *testing.T) {
	values := &url.Values{}

	// string
	multi1 := NewMulti("a", "b", "c")
	assert.NoError(t, multi1.EncodeValues("key1", values))
	assert.Equal(t, "a,b,c", values.Get("key1"))

	// int
	multi2 := NewMulti(1, 2, 3)
	assert.NoError(t, multi2.EncodeValues("key2", values))
	assert.Equal(t, "1,2,3", values.Get("key2"))

	// float64
	multi3 := NewMulti(1.1, 2.2, 3.3)
	assert.NoError(t, multi3.EncodeValues("key3", values))
	assert.Equal(t, "1.1,2.2,3.3", values.Get("key3"))

	// Multi{}
	multi4 := Multi[string]([]string{"a", "b", "c"})
	assert.NoError(t, multi4.EncodeValues("key4", values))
	assert.Equal(t, "a,b,c", values.Get("key4"))
}

func TestTypes_Enum(t *testing.T) {
	values := &url.Values{}

	// string
	enum1 := NewEnum("a", "b", "c")
	assert.NoError(t, enum1.EncodeValues("key1", values))
	assert.Equal(t, "a|b|c", values.Get("key1"))

	// int
	enum2 := NewEnum(1, 2, 3)
	assert.NoError(t, enum2.EncodeValues("key2", values))
	assert.Equal(t, "1|2|3", values.Get("key2"))

	// float64
	enum3 := NewEnum(1.1, 2.2, 3.3)
	assert.NoError(t, enum3.EncodeValues("key3", values))
	assert.Equal(t, "1.1|2.2|3.3", values.Get("key3"))

	// Enum{}
	enum4 := Enum[string]([]string{"a", "b", "c"})
	assert.NoError(t, enum4.EncodeValues("key4", values))
	assert.Equal(t, "a|b|c", values.Get("key4"))
}

func TestTypes_Order(t *testing.T) {
	tests := []struct {
		name  string
		order *Order
		want  string
	}{
		{"", NewOrder("created"), "created asc"},
		{"", NewOrder("created", OrderByAsc), "created asc"},
		{"", NewOrder("created", OrderByDesc), "created desc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := json.Marshal(tt.order)
			assert.NoError(t, err)
			assert.Equal(t, fmt.Sprintf("\"%s\"", tt.want), string(bytes))

			var o Order
			err = json.Unmarshal(bytes, &o)
			assert.NoError(t, err)
			assert.Equal(t, tt.order, &o)

			values := &url.Values{}
			assert.NoError(t, tt.order.EncodeValues("key", values))
			assert.Equal(t, tt.want, values.Get("key"))
		})
	}
}

func TestTypes_PriorityLabel(t *testing.T) {
	tests := []struct {
		name   string
		label  PriorityLabel
		expect string
	}{
		{"", "", ""},
		{"", PriorityLabelHigh, "High"},
		{"", PriorityLabelMiddle, "Middle"},
		{"", PriorityLabelLow, "Low"},
		{"", PriorityLabelNiceToHave, "Nice To Have"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expect, tt.label.String())
		})
	}
}

func TestTypes_Order_Custom(t *testing.T) {
	type Demo struct {
		Order *Order `json:"order"`
	}

	demo := &Demo{
		Order: NewOrder("id", OrderByAsc),
	}

	bytes, err := json.Marshal(demo)
	assert.NoError(t, err)
	assert.Equal(t, `{"order":"id asc"}`, string(bytes))
}
