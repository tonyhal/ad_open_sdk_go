/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInner struct for StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInner
type StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInner struct {
	// 审核状态 0 = 审核不通过 99 = 审核中 1 = 审核通过 3 = 作者删除 4 = 跳过审核
	AuditStatus *int64 `json:"audit_status,omitempty"`
	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`
	// 审核详细信息
	DetailAuditInfo []*StarOrderGetScriptV2ResponseDataOrderScriptListInnerScriptListInnerDetailAuditInfoInner `json:"detail_audit_info,omitempty"`
	// 脚本文件名
	ResourceFilename *string `json:"resource_filename,omitempty"`
	// 资源ID
	ResourceId *int64 `json:"resource_id,omitempty"`
	// 脚本文件地址
	ResourceUrl *string `json:"resource_url,omitempty"`
	// 资源状态 1 = 新上传 2 = 已驳回 3 = 已确认 4 = 非最新无意见 5 = 已删除
	Status *int64 `json:"status,omitempty"`
}
