/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

// ProjectCreateV30RequestDeliveryRange
type ProjectCreateV30RequestDeliveryRange struct {
	InventoryCatalog ProjectCreateV30DeliveryRangeInventoryCatalog `json:"inventory_catalog"`
	//
	InventoryType  []*ProjectCreateV30DeliveryRangeInventoryType `json:"inventory_type,omitempty"`
	UnionVideoType *ProjectCreateV30DeliveryRangeUnionVideoType  `json:"union_video_type,omitempty"`
}
