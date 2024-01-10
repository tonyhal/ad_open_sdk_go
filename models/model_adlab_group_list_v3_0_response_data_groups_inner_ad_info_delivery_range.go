/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.16
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupListV30ResponseDataGroupsInnerAdInfoDeliveryRange 广告版位
type AdlabGroupListV30ResponseDataGroupsInnerAdInfoDeliveryRange struct {
	InventoryCatalog *AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog `json:"inventory_catalog,omitempty"`
	// 广告投放位置（首选媒体）
	InventoryType  []*AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryType `json:"inventory_type,omitempty"`
	UnionVideoType *AdlabGroupListV30DataGroupsAdInfoDeliveryRangeUnionVideoType  `json:"union_video_type,omitempty"`
}
