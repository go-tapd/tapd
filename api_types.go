package tapd

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// -----------------------------------------------------------------------------
// Multi is a type for multi values.
// Multi{Value1, Value2, Value3} => value1,value2,value3
//
// Useful for ID/Fields/...
// -----------------------------------------------------------------------------

type Multi[T any] []T

var (
	_ query.Encoder  = (*Multi[string])(nil)
	_ json.Marshaler = (*Multi[string])(nil)
)

// NewMulti creates a new multi value.
//
// Example:
//
//	NewMulti("a", "b", "c") => "a,b,c"
//	NewMulti(1, 2, 3) => "1,2,3"
func NewMulti[T any](values ...T) *Multi[T] {
	return (*Multi[T])(&values)
}

func (m Multi[T]) EncodeValues(key string, v *url.Values) error {
	if len(m) > 0 {
		v.Add(key, m.String())
	}
	return nil
}

func (m Multi[T]) String() string {
	if len(m) > 0 {
		var values []string
		for _, value := range m {
			values = append(values, fmt.Sprint(value))
		}
		return strings.Join(values, ",")
	}
	return ""
}

func (m Multi[T]) MarshalJSON() ([]byte, error) {
	if len(m) > 0 {
		return json.Marshal(m.String())
	}
	return json.Marshal(nil)
}

// -----------------------------------------------------------------------------
// Enum is a type for enum values.
// Enum{Value1, Value2, Value3} => value1|value2|value3
// -----------------------------------------------------------------------------

type Enum[T any] []T

var (
	_ query.Encoder  = (*Enum[string])(nil)
	_ json.Marshaler = (*Enum[string])(nil)
)

// NewEnum creates a new enum value.
//
// Example:
//
//	NewEnum("a", "b", "c") => "a|b|c"
//	NewEnum(1, 2, 3) => "1|2|3"
func NewEnum[T any](values ...T) *Enum[T] {
	return (*Enum[T])(&values)
}

func (e Enum[T]) EncodeValues(key string, v *url.Values) error {
	if len(e) > 0 {
		values := make([]string, 0, len(e))
		for _, value := range e {
			values = append(values, fmt.Sprint(value))
		}
		v.Add(key, strings.Join(values, "|"))
	}
	return nil
}

func (e Enum[T]) MarshalJSON() ([]byte, error) {
	if len(e) <= 0 {
		return json.Marshal(nil)
	}

	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, fmt.Sprint(value))
	}
	return json.Marshal(strings.Join(values, "|"))
}

// -----------------------------------------------------------------------------
// PriorityLabel is a type for priority labels.
// -----------------------------------------------------------------------------

type PriorityLabel string

const (
	PriorityLabelHigh       PriorityLabel = "High"
	PriorityLabelMiddle     PriorityLabel = "Middle"
	PriorityLabelLow        PriorityLabel = "Low"
	PriorityLabelNiceToHave PriorityLabel = "Nice To Have"
)

func (p PriorityLabel) String() string {
	return string(p)
}

// -----------------------------------------------------------------------------
// Order is a query encoder for order parameters.
// -----------------------------------------------------------------------------

type OrderType string

const (
	OrderTypeAsc  OrderType = "asc"
	OrderTypeDesc OrderType = "desc"
)

// Order is a type for order parameters.
type Order struct {
	field     string
	orderType OrderType
}

var (
	_ json.Marshaler   = (*Order)(nil)
	_ json.Unmarshaler = (*Order)(nil)
	_ query.Encoder    = (*Order)(nil)
)

type OrderOption func(*Order)

func WithOrderType(orderType OrderType) OrderOption {
	return func(o *Order) {
		o.orderType = orderType
	}
}

var (
	OrderByAsc  = WithOrderType(OrderTypeAsc)
	OrderByDesc = WithOrderType(OrderTypeDesc)
)

// NewOrder creates a new order parameter.
//
// Example:
//
//	NewOrder("created") => "created asc"
//	NewOrder("created", OrderByAsc) => "created asc"
//	NewOrder("created", OrderByDesc) => "created desc"
func NewOrder(field string, opts ...OrderOption) *Order {
	o := &Order{
		field:     field,
		orderType: OrderTypeAsc,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func (o *Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%s %s", o.field, o.orderType))
}

func (o *Order) UnmarshalJSON(bytes []byte) error {
	var s string
	if err := json.Unmarshal(bytes, &s); err != nil {
		return err
	}
	_, err := fmt.Sscanf(s, "%s %s", &o.field, &o.orderType)
	return err
}

func (o *Order) EncodeValues(key string, v *url.Values) error {
	v.Add(key, fmt.Sprintf("%s %s", o.field, o.orderType))
	return nil
}

// -----------------------------------------------------------------------------
// EntityType is a type for entity types.
// -----------------------------------------------------------------------------

type EntityType string

const (
	EntityTypeTask  EntityType = "task"
	EntityTypeStory EntityType = "story"
	EntityTypeBug   EntityType = "bug"
)

// -----------------------------------------------------------------------------
// OperateType is a type for operate types.
//
// 操作类型，默认为所有，可以填写add,delete,download,upload中的一个
// -----------------------------------------------------------------------------

type OperateType string

const (
	OperateTypeAdd      OperateType = "add"
	OperateTypeDelete   OperateType = "delete"
	OperateTypeDownload OperateType = "download"
	OperateTypeUpload   OperateType = "upload"
)

// -----------------------------------------------------------------------------
// OperateObject is a type for operate objects.
//
// 操作对象，默认为所有，可以填写attachment,board,bug,document,
// iteration,launch,member_activity_log,
// release,story,task,tcase,testplan,wiki中的一个
// -----------------------------------------------------------------------------

type OperateObject string

const (
	OperateObjectAttachment        OperateObject = "attachment"
	OperateObjectBoard             OperateObject = "board"
	OperateObjectBug               OperateObject = "bug"
	OperateObjectDocument          OperateObject = "document"
	OperateObjectIteration         OperateObject = "iteration"
	OperateObjectLaunch            OperateObject = "launch"
	OperateObjectMemberActivityLog OperateObject = "member_activity_log"
	OperateObjectRelease           OperateObject = "release"
	OperateObjectStory             OperateObject = "story"
	OperateObjectTask              OperateObject = "task"
	OperateObjectTestCase          OperateObject = "tcase"
	OperateObjectTestPlan          OperateObject = "testplan"
	OperateObjectWiki              OperateObject = "wiki"
)
