/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliveryRangeInventoryCatalog
type ProjectCreateV30DeliveryRangeInventoryCatalog string

// List of project_create_v3.0_delivery_range_inventory_catalog
const (
	MANUAL_ProjectCreateV30DeliveryRangeInventoryCatalog          ProjectCreateV30DeliveryRangeInventoryCatalog = "MANUAL"
	UNIVERSAL_SMART_ProjectCreateV30DeliveryRangeInventoryCatalog ProjectCreateV30DeliveryRangeInventoryCatalog = "UNIVERSAL_SMART"
)

// All allowed values of ProjectCreateV30DeliveryRangeInventoryCatalog enum
var AllowedProjectCreateV30DeliveryRangeInventoryCatalogEnumValues = []ProjectCreateV30DeliveryRangeInventoryCatalog{
	"MANUAL",
	"UNIVERSAL_SMART",
}

// NewProjectCreateV30DeliveryRangeInventoryCatalogFromValue returns a pointer to a valid ProjectCreateV30DeliveryRangeInventoryCatalog
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliveryRangeInventoryCatalogFromValue(v string) (*ProjectCreateV30DeliveryRangeInventoryCatalog, error) {
	ev := ProjectCreateV30DeliveryRangeInventoryCatalog(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliveryRangeInventoryCatalog: valid values are %v", v, AllowedProjectCreateV30DeliveryRangeInventoryCatalogEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliveryRangeInventoryCatalog) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliveryRangeInventoryCatalogEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_range_inventory_catalog value
func (v ProjectCreateV30DeliveryRangeInventoryCatalog) Ptr() *ProjectCreateV30DeliveryRangeInventoryCatalog {
	return &v
}
