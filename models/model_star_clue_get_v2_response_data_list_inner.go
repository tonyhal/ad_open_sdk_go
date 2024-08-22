/*
Oceanengine Open Api

巨量引擎开放平台 Open Api


*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarClueGetV2ResponseDataListInner struct for StarClueGetV2ResponseDataListInner
type StarClueGetV2ResponseDataListInner struct {
	// 地址
	Address *string `json:"address,omitempty"`
	// 组件名称
	AnchorName string `json:"anchor_name"`
	// 达人昵称
	AuthorName *string `json:"author_name,omitempty"`
	// 达人抖音短号
	AuthorUid *int64 `json:"author_uid,omitempty"`
	// 车系
	CarSeries *string `json:"car_series,omitempty"`
	// 车型
	CarType *string `json:"car_type,omitempty"`
	// 城市
	City          *string                            `json:"city,omitempty"`
	ComponentType StarClueGetV2DataListComponentType `json:"component_type"`
	// 创建时间
	CreateTime string `json:"create_time"`
	// 任务id
	DemandId int64 `json:"demand_id"`
	// 视频id
	ItemId int64 `json:"item_id"`
	// 姓名
	Name *string `json:"name,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id"`
	// 电话
	Phone *string `json:"phone,omitempty"`
	// 省会
	Province *string `json:"province,omitempty"`
	// 客户id
	StarId int64 `json:"star_id"`
}
