/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsWechatGameListV30Filtering
type ToolsWechatGameListV30Filtering struct {
	AssetStatus *ToolsWechatGameListV30FilteringAssetStatus `json:"asset_status,omitempty"`
	AuditStatus *ToolsWechatGameListV30FilteringAuditStatus `json:"audit_status,omitempty"`
	CreateTime  *ToolsWechatGameListV30FilteringCreateTime  `json:"create_time,omitempty"`
	//
	Name       *string                                    `json:"name,omitempty"`
	SearchType *ToolsWechatGameListV30FilteringSearchType `json:"search_type,omitempty"`
}
