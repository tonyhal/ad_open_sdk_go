/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetVideoV2ResponseDataOrderVideoListInnerVideoListInner struct for StarOrderGetVideoV2ResponseDataOrderVideoListInnerVideoListInner
type StarOrderGetVideoV2ResponseDataOrderVideoListInnerVideoListInner struct {
	// 审核状态（同ScriptInfo中同名字段）
	AuditStatus *int64 `json:"audit_status,omitempty"`
	// 视频封面图
	CoverUrl *string `json:"cover_url,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 审核详细信息
	DetailAuditInfo []*StarOrderDetailV2ResponseDataScriptListInnerDetailAuditInfoInner `json:"detail_audit_info,omitempty"`
	// 视频时长
	Duration *int64 `json:"duration,omitempty"`
	// 视频ID
	ItemId *int64 `json:"item_id,omitempty"`
	// 在线（公开）状态。1: 公开，2: 非公开
	OnlineStatus *int64 `json:"online_status,omitempty"`
	// 资源ID
	ResourceId *int64 `json:"resource_id,omitempty"`
	// 视频播放地址
	ResourceUrl *string `json:"resource_url,omitempty"`
	// 资源状态（同ScriptInfo中同名字段）
	Status *int64 `json:"status,omitempty"`
}
