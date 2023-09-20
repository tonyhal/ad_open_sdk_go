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

// ProjectListV30FilteringInventoryType
type ProjectListV30FilteringInventoryType string

// List of project_list_v3.0_filtering_inventory_type
const (
	INVENTORY_AUTOMOBILE_ProjectListV30FilteringInventoryType      ProjectListV30FilteringInventoryType = "INVENTORY_AUTOMOBILE"
	INVENTORY_AWEME_FEED_ProjectListV30FilteringInventoryType      ProjectListV30FilteringInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_BEAUTY_ProjectListV30FilteringInventoryType          ProjectListV30FilteringInventoryType = "INVENTORY_BEAUTY"
	INVENTORY_FACE_U_ProjectListV30FilteringInventoryType          ProjectListV30FilteringInventoryType = "INVENTORY_FACE_U"
	INVENTORY_FEED_ProjectListV30FilteringInventoryType            ProjectListV30FilteringInventoryType = "INVENTORY_FEED"
	INVENTORY_HOMED_AGGREGATE_ProjectListV30FilteringInventoryType ProjectListV30FilteringInventoryType = "INVENTORY_HOMED_AGGREGATE"
	INVENTORY_HOTSOON_FEED_ProjectListV30FilteringInventoryType    ProjectListV30FilteringInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_PIPIXIA_ProjectListV30FilteringInventoryType         ProjectListV30FilteringInventoryType = "INVENTORY_PIPIXIA"
	INVENTORY_SEARCH_ProjectListV30FilteringInventoryType          ProjectListV30FilteringInventoryType = "INVENTORY_SEARCH"
	INVENTORY_STUDY_ProjectListV30FilteringInventoryType           ProjectListV30FilteringInventoryType = "INVENTORY_STUDY"
	INVENTORY_TOMATO_NOVEL_ProjectListV30FilteringInventoryType    ProjectListV30FilteringInventoryType = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ProjectListV30FilteringInventoryType      ProjectListV30FilteringInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_UNIVERSAL_ProjectListV30FilteringInventoryType       ProjectListV30FilteringInventoryType = "INVENTORY_UNIVERSAL"
	INVENTORY_VIDEO_FEED_ProjectListV30FilteringInventoryType      ProjectListV30FilteringInventoryType = "INVENTORY_VIDEO_FEED"
	UNION_BOUTIQUE_GAME_ProjectListV30FilteringInventoryType       ProjectListV30FilteringInventoryType = "UNION_BOUTIQUE_GAME"
)

// All allowed values of ProjectListV30FilteringInventoryType enum
var AllowedProjectListV30FilteringInventoryTypeEnumValues = []ProjectListV30FilteringInventoryType{
	"INVENTORY_AUTOMOBILE",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_BEAUTY",
	"INVENTORY_FACE_U",
	"INVENTORY_FEED",
	"INVENTORY_HOMED_AGGREGATE",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_PIPIXIA",
	"INVENTORY_SEARCH",
	"INVENTORY_STUDY",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_VIDEO_FEED",
	"UNION_BOUTIQUE_GAME",
}

// NewProjectListV30FilteringInventoryTypeFromValue returns a pointer to a valid ProjectListV30FilteringInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectListV30FilteringInventoryTypeFromValue(v string) (*ProjectListV30FilteringInventoryType, error) {
	ev := ProjectListV30FilteringInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectListV30FilteringInventoryType: valid values are %v", v, AllowedProjectListV30FilteringInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectListV30FilteringInventoryType) IsValid() bool {
	for _, existing := range AllowedProjectListV30FilteringInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to project_list_v3.0_filtering_inventory_type value
func (v ProjectListV30FilteringInventoryType) Ptr() *ProjectListV30FilteringInventoryType {
	return &v
}
