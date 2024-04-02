/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupCreateV30RequestAdInfoDeliveryRange 广告版位
type AdlabGroupCreateV30RequestAdInfoDeliveryRange struct {
	InventoryCatalog AdlabGroupCreateV30AdInfoDeliveryRangeInventoryCatalog `json:"inventory_catalog"`
	// 广告投放位置（首选媒体）
	InventoryType  []*AdlabGroupCreateV30AdInfoDeliveryRangeInventoryType `json:"inventory_type,omitempty"`
	UnionVideoType *AdlabGroupCreateV30AdInfoDeliveryRangeUnionVideoType  `json:"union_video_type,omitempty"`
}
