/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarDemandOrderListV2ResponseDataListInner struct for StarDemandOrderListV2ResponseDataListInner
type StarDemandOrderListV2ResponseDataListInner struct {
	// 达人id
	AuthorId *int64 `json:"author_id,omitempty"`
	// 达人名称
	AuthorName *string `json:"author_name,omitempty"`
	// 达人头像
	AvatarUri *string `json:"avatar_uri,omitempty"`
	// 需求ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 订单创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime *string `json:"create_time,omitempty"`
	// 任务id
	DemandId *int64 `json:"demand_id,omitempty"`
	// 封面图
	HeadImageUri *string `json:"head_image_uri,omitempty"`
	// 视频id，与星图平台前端video_url中展现的视频id一致，每个视频唯一
	ItemId *int64 `json:"item_id,omitempty"`
	// 订单id
	OrderId *int64 `json:"order_id,omitempty"`
	// 指派任务产出物发布时间
	ReleaseTime *string `json:"release_time,omitempty"`
	// 作品名称
	Title                *string                                            `json:"title,omitempty"`
	UniversalOrderStatus *StarDemandOrderListV2DataListUniversalOrderStatus `json:"universal_order_status,omitempty"`
	// 视频id，每个视频唯一（建议使用item_id）
	VideoId *string `json:"video_id,omitempty"`
	// 视频链接
	VideoUrl *string `json:"video_url,omitempty"`
}
