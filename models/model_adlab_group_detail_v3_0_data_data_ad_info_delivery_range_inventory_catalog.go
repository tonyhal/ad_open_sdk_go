/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.12
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog
type AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog string

// List of adlab_group_detail_v3.0_data_data_ad_info_delivery_range_inventory_catalog
const (
	MANUAL_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog          AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog = "MANUAL"
	SCENE_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog           AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog = "SCENE"
	SMART_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog           AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog = "SMART"
	UNIVERSAL_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog       AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog = "UNIVERSAL"
	UNIVERSAL_SMART_AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog = "UNIVERSAL_SMART"
)

// All allowed values of AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog enum
var AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalogEnumValues = []AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog{
	"MANUAL",
	"SCENE",
	"SMART",
	"UNIVERSAL",
	"UNIVERSAL_SMART",
}

// NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalogFromValue returns a pointer to a valid AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalogFromValue(v string) (*AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog, error) {
	ev := AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog: valid values are %v", v, AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalogEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog) IsValid() bool {
	for _, existing := range AllowedAdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalogEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to adlab_group_detail_v3.0_data_data_ad_info_delivery_range_inventory_catalog value
func (v AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog) Ptr() *AdlabGroupDetailV30DataDataAdInfoDeliveryRangeInventoryCatalog {
	return &v
}
