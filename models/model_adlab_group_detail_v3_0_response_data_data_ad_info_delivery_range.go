/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// AdlabGroupDetailV30ResponseDataDataAdInfoDeliveryRange 广告版位
type AdlabGroupDetailV30ResponseDataDataAdInfoDeliveryRange struct {
	InventoryCatalog *AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog `json:"inventory_catalog,omitempty"`
	// 广告投放位置（首选媒体）
	InventoryType  []*AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryType `json:"inventory_type,omitempty"`
	UnionVideoType *AdlabGroupDetailV30DataDataAdInfoDeliveryRangeUnionVideoType  `json:"union_video_type,omitempty"`
}