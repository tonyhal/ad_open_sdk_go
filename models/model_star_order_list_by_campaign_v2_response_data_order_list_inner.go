/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderListByCampaignV2ResponseDataOrderListInner struct for StarOrderListByCampaignV2ResponseDataOrderListInner
type StarOrderListByCampaignV2ResponseDataOrderListInner struct {
	// 达人ID
	AuthorId *int64 `json:"author_id,omitempty"`
	// 达人昵称
	AuthorName *string `json:"author_name,omitempty"`
	// 达人头像
	AvatarUri *string `json:"avatar_uri,omitempty"`
	// 需求ID
	CampaignId *int64 `json:"campaign_id,omitempty"`
	// 任务创建时间 unix timestamp
	CreateTime *int64 `json:"create_time,omitempty"`
	// 任务ID
	DemandId *int64 `json:"demand_id,omitempty"`
	// 任务ID
	OrderId *int64 `json:"order_id,omitempty"`
	// 任务状态 待接收 = -1 达人已接单 = 1 已关闭 = 2 已完成 = 3 已取消 = 4 待付尾款 = 10 脚本已上传 = 41 脚本已拒绝 = 42 脚本已确认 = 43 脚本已跳过 = 44 视频已上传 = 51 视频已拒绝 = 52 视频已确认 = 53 视频已发布 = 54
	OrderStatus *int64 `json:"order_status,omitempty"`
}
