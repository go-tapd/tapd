package tapd

import (
	"context"
	"net/http"
)

type LifeTimesRequest struct {
	EntityID    *int           `url:"entity_id,omitempty"`    // [必须]业务对象ID
	EntityType  *EntityType    `url:"entity_type,omitempty"`  // [必须]业务对象类型 目前type可选值：task,story,bug
	WorkspaceID *int           `url:"workspace_id,omitempty"` // [必须]项目ID
	Created     *string        `url:"created,omitempty"`      // 创建时间
	Limit       *int           `url:"limit,omitempty"`        // 设置返回数量限制，默认为30
	Page        *int           `url:"page,omitempty"`         // 返回当前数量限制下第N页的数据，默认为1（第一页）
	Fields      *Multi[string] `url:"fields,omitempty"`       // 设置获取的字段，多个字段间以','逗号隔开
}

type MeasureService interface {
	// LifeTimes 获取状态流转时间
	// Note: 一次插入一条数据。注意同一 entity_type、entity_id、spentdate、owner ，只能有一条工时记录
	// https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/api_reference/measure/get_life_times.html
	LifeTimes(ctx context.Context, request *LifeTimesRequest, opts ...RequestOption) ([]*LifeTime, *Response, error)
}

type (
	// LifeTime 状态流转时间
	LifeTime struct {
		ID              string     `json:"id"`                // id
		WorkspaceID     string     `json:"workspace_id"`      // 项目ID
		EntityType      EntityType `json:"entity_type"`       // 业务对象类型
		EntityID        string     `json:"entity_id"`         // 业务对象ID
		Status          string     `json:"status"`            // 状态
		Owner           string     `json:"owner"`             //
		IsRepeated      string     `json:"is_repeated"`       // 是否重复
		BeginDate       string     `json:"begin_date"`        // 开始时间
		EndDate         string     `json:"end_date"`          // 结束时间
		TimeCost        string     `json:"time_cost"`         // 停留时长，单位：小时
		TimeCostReduced string     `json:"time_cost_reduced"` // 停留时长，单位：小时
		Created         string     `json:"created"`           // 创建时间（变更时间）
		Operator        string     `json:"operator"`          // 操作人
		IsLatest        string     `json:"is_latest"`         // 是否最新
		IsDelete        string     `json:"is_delete"`         // 是否删除
	}
)

type measureService struct {
	client *Client
}

var _ MeasureService = (*measureService)(nil)

func NewMeasureService(client *Client) MeasureService {
	return &measureService{
		client: client,
	}
}

func (s *measureService) LifeTimes(
	ctx context.Context, request *LifeTimesRequest, opts ...RequestOption,
) ([]*LifeTime, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, "life_times", request, opts)
	if err != nil {
		return nil, nil, err
	}

	var items []struct {
		LifeTime *LifeTime `json:"LifeTime"`
	}
	resp, err := s.client.Do(req, &items)
	if err != nil {
		return nil, resp, err
	}

	lifeTimes := make([]*LifeTime, 0, len(items))
	for _, item := range items {
		lifeTimes = append(lifeTimes, item.LifeTime)
	}

	return lifeTimes, resp, nil
}
