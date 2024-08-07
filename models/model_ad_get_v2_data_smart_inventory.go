/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataSmartInventory
type AdGetV2DataSmartInventory string

// List of ad_get_v2_data_smart_inventory
const (
	UNIVERSAL_ALL_AdGetV2DataSmartInventory AdGetV2DataSmartInventory = "UNIVERSAL_ALL"
	UNIVERSAL_AdGetV2DataSmartInventory     AdGetV2DataSmartInventory = "UNIVERSAL"
	SMART_AdGetV2DataSmartInventory         AdGetV2DataSmartInventory = "SMART"
	NORMAL_AdGetV2DataSmartInventory        AdGetV2DataSmartInventory = "NORMAL"
)

// All allowed values of AdGetV2DataSmartInventory enum
var AllowedAdGetV2DataSmartInventoryEnumValues = []AdGetV2DataSmartInventory{
	"UNIVERSAL_ALL",
	"UNIVERSAL",
	"SMART",
	"NORMAL",
}

// NewAdGetV2DataSmartInventoryFromValue returns a pointer to a valid AdGetV2DataSmartInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataSmartInventoryFromValue(v string) (*AdGetV2DataSmartInventory, error) {
	ev := AdGetV2DataSmartInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataSmartInventory: valid values are %v", v, AllowedAdGetV2DataSmartInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataSmartInventory) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataSmartInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_smart_inventory value
func (v AdGetV2DataSmartInventory) Ptr() *AdGetV2DataSmartInventory {
	return &v
}
