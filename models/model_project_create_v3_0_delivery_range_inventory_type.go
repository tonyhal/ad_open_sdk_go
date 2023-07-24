/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 0.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ProjectCreateV30DeliveryRangeInventoryType
type ProjectCreateV30DeliveryRangeInventoryType string

// List of project_create_v3.0_delivery_range_inventory_type
const (
	INVENTORY_AWEME_FEED_ProjectCreateV30DeliveryRangeInventoryType      ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_FEED_ProjectCreateV30DeliveryRangeInventoryType            ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_FEED"
	INVENTORY_HOMED_AGGREGATE_ProjectCreateV30DeliveryRangeInventoryType ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_HOMED_AGGREGATE"
	INVENTORY_HOTSOON_FEED_ProjectCreateV30DeliveryRangeInventoryType    ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_SEARCH_ProjectCreateV30DeliveryRangeInventoryType          ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_SEARCH"
	INVENTORY_TOMATO_NOVEL_ProjectCreateV30DeliveryRangeInventoryType    ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ProjectCreateV30DeliveryRangeInventoryType      ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_VIDEO_FEED_ProjectCreateV30DeliveryRangeInventoryType      ProjectCreateV30DeliveryRangeInventoryType = "INVENTORY_VIDEO_FEED"
	UNION_BOUTIQUE_GAME_ProjectCreateV30DeliveryRangeInventoryType       ProjectCreateV30DeliveryRangeInventoryType = "UNION_BOUTIQUE_GAME"
)

// All allowed values of ProjectCreateV30DeliveryRangeInventoryType enum
var AllowedProjectCreateV30DeliveryRangeInventoryTypeEnumValues = []ProjectCreateV30DeliveryRangeInventoryType{
	"INVENTORY_AWEME_FEED",
	"INVENTORY_FEED",
	"INVENTORY_HOMED_AGGREGATE",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_SEARCH",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_VIDEO_FEED",
	"UNION_BOUTIQUE_GAME",
}

// NewProjectCreateV30DeliveryRangeInventoryTypeFromValue returns a pointer to a valid ProjectCreateV30DeliveryRangeInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectCreateV30DeliveryRangeInventoryTypeFromValue(v string) (*ProjectCreateV30DeliveryRangeInventoryType, error) {
	ev := ProjectCreateV30DeliveryRangeInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectCreateV30DeliveryRangeInventoryType: valid values are %v", v, AllowedProjectCreateV30DeliveryRangeInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectCreateV30DeliveryRangeInventoryType) IsValid() bool {
	for _, existing := range AllowedProjectCreateV30DeliveryRangeInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_create_v3.0_delivery_range_inventory_type value
func (v ProjectCreateV30DeliveryRangeInventoryType) Ptr() *ProjectCreateV30DeliveryRangeInventoryType {
	return &v
}
