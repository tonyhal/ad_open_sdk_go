/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// ReportAdGetV2FilteringInventoryTypes
type ReportAdGetV2FilteringInventoryTypes string

// List of report_ad_get_v2_filtering_inventory_types
const (
	INVENTORY_PIPIXIA_ReportAdGetV2FilteringInventoryTypes      ReportAdGetV2FilteringInventoryTypes = "INVENTORY_PIPIXIA"
	INVENTORY_VIDEO_FEED_ReportAdGetV2FilteringInventoryTypes   ReportAdGetV2FilteringInventoryTypes = "INVENTORY_VIDEO_FEED"
	INVENTORY_AUTOMOBILE_ReportAdGetV2FilteringInventoryTypes   ReportAdGetV2FilteringInventoryTypes = "INVENTORY_AUTOMOBILE"
	INVENTORY_UNIVERSAL_ReportAdGetV2FilteringInventoryTypes    ReportAdGetV2FilteringInventoryTypes = "INVENTORY_UNIVERSAL"
	INVENTORY_FEED_ReportAdGetV2FilteringInventoryTypes         ReportAdGetV2FilteringInventoryTypes = "INVENTORY_FEED"
	INVENTORY_BEAUTY_ReportAdGetV2FilteringInventoryTypes       ReportAdGetV2FilteringInventoryTypes = "INVENTORY_BEAUTY"
	INVENTORY_AWEME_FEED_ReportAdGetV2FilteringInventoryTypes   ReportAdGetV2FilteringInventoryTypes = "INVENTORY_AWEME_FEED"
	UNION_BOUTIQUE_GAME_ReportAdGetV2FilteringInventoryTypes    ReportAdGetV2FilteringInventoryTypes = "UNION_BOUTIQUE_GAME"
	INVENTORY_STUDY_ReportAdGetV2FilteringInventoryTypes        ReportAdGetV2FilteringInventoryTypes = "INVENTORY_STUDY"
	INVENTORY_HOTSOON_FEED_ReportAdGetV2FilteringInventoryTypes ReportAdGetV2FilteringInventoryTypes = "INVENTORY_HOTSOON_FEED"
	INVENTORY_FURNISH_ReportAdGetV2FilteringInventoryTypes      ReportAdGetV2FilteringInventoryTypes = "INVENTORY_FURNISH"
	INVENTORY_TOMATO_NOVEL_ReportAdGetV2FilteringInventoryTypes ReportAdGetV2FilteringInventoryTypes = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ReportAdGetV2FilteringInventoryTypes   ReportAdGetV2FilteringInventoryTypes = "INVENTORY_UNION_SLOT"
	INVENTORY_FACE_U_ReportAdGetV2FilteringInventoryTypes       ReportAdGetV2FilteringInventoryTypes = "INVENTORY_FACE_U"
	INVENTORY_SEARCH_ReportAdGetV2FilteringInventoryTypes       ReportAdGetV2FilteringInventoryTypes = "INVENTORY_SEARCH"
)

// All allowed values of ReportAdGetV2FilteringInventoryTypes enum
var AllowedReportAdGetV2FilteringInventoryTypesEnumValues = []ReportAdGetV2FilteringInventoryTypes{
	"INVENTORY_PIPIXIA",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_AUTOMOBILE",
	"INVENTORY_UNIVERSAL",
	"INVENTORY_FEED",
	"INVENTORY_BEAUTY",
	"INVENTORY_AWEME_FEED",
	"UNION_BOUTIQUE_GAME",
	"INVENTORY_STUDY",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_FURNISH",
	"INVENTORY_TOMATO_NOVEL",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_FACE_U",
	"INVENTORY_SEARCH",
}

// NewReportAdGetV2FilteringInventoryTypesFromValue returns a pointer to a valid ReportAdGetV2FilteringInventoryTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2FilteringInventoryTypesFromValue(v string) (*ReportAdGetV2FilteringInventoryTypes, error) {
	ev := ReportAdGetV2FilteringInventoryTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2FilteringInventoryTypes: valid values are %v", v, AllowedReportAdGetV2FilteringInventoryTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2FilteringInventoryTypes) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2FilteringInventoryTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_filtering_inventory_types value
func (v ReportAdGetV2FilteringInventoryTypes) Ptr() *ReportAdGetV2FilteringInventoryTypes {
	return &v
}
