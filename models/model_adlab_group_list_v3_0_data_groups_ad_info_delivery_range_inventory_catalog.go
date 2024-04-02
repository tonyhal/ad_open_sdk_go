/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog
type AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog string

// List of adlab_group_list_v3.0_data_groups_ad_info_delivery_range_inventory_catalog
const (
	MANUAL_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog          AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog = "MANUAL"
	SCENE_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog           AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog = "SCENE"
	SMART_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog           AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog = "SMART"
	UNIVERSAL_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog       AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog = "UNIVERSAL"
	UNIVERSAL_SMART_AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog = "UNIVERSAL_SMART"
)

// All allowed values of AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog enum
var AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalogEnumValues = []AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog{
	"MANUAL",
	"SCENE",
	"SMART",
	"UNIVERSAL",
	"UNIVERSAL_SMART",
}

// NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalogFromValue returns a pointer to a valid AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalogFromValue(v string) (*AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog, error) {
	ev := AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog: valid values are %v", v, AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalogEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog) IsValid() bool {
	for _, existing := range AllowedAdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalogEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_list_v3.0_data_groups_ad_info_delivery_range_inventory_catalog value
func (v AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog) Ptr() *AdlabGroupListV30DataGroupsAdInfoDeliveryRangeInventoryCatalog {
	return &v
}
