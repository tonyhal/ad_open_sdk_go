/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.15
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ToolsAppManagementExtendPackageListV2V2Filtering
type ToolsAppManagementExtendPackageListV2V2Filtering struct {
	// 渠道号 可通过渠道号筛选应用分包，单次支持传入渠道号个数<=50
	ChannelId  []string                                                    `json:"channel_id,omitempty"`
	Status     *ToolsAppManagementExtendPackageListV2V2FilteringStatus     `json:"status,omitempty"`
	UpdateTime *ToolsAppManagementExtendPackageListV2V2FilteringUpdateTime `json:"update_time,omitempty"`
}
