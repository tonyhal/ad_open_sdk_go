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

// ReportCreativeGetV2FilteringInventoryTypes
type ReportCreativeGetV2FilteringInventoryTypes string

// List of report_creative_get_v2_filtering_inventory_types
const (
	INVENTORY_PIPIXIA_ReportCreativeGetV2FilteringInventoryTypes      ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_PIPIXIA"
	INVENTORY_VIDEO_FEED_ReportCreativeGetV2FilteringInventoryTypes   ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_VIDEO_FEED"
	INVENTORY_AUTOMOBILE_ReportCreativeGetV2FilteringInventoryTypes   ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_AUTOMOBILE"
	INVENTORY_UNIVERSAL_ReportCreativeGetV2FilteringInventoryTypes    ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_UNIVERSAL"
	INVENTORY_FEED_ReportCreativeGetV2FilteringInventoryTypes         ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_FEED"
	INVENTORY_BEAUTY_ReportCreativeGetV2FilteringInventoryTypes       ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_BEAUTY"
	INVENTORY_AWEME_FEED_ReportCreativeGetV2FilteringInventoryTypes   ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_AWEME_FEED"
	UNION_BOUTIQUE_GAME_ReportCreativeGetV2FilteringInventoryTypes    ReportCreativeGetV2FilteringInventoryTypes = "UNION_BOUTIQUE_GAME"
	INVENTORY_STUDY_ReportCreativeGetV2FilteringInventoryTypes        ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_STUDY"
	INVENTORY_HOTSOON_FEED_ReportCreativeGetV2FilteringInventoryTypes ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_HOTSOON_FEED"
	INVENTORY_FURNISH_ReportCreativeGetV2FilteringInventoryTypes      ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_FURNISH"
	INVENTORY_TOMATO_NOVEL_ReportCreativeGetV2FilteringInventoryTypes ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ReportCreativeGetV2FilteringInventoryTypes   ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_UNION_SLOT"
	INVENTORY_FACE_U_ReportCreativeGetV2FilteringInventoryTypes       ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_FACE_U"
	INVENTORY_SEARCH_ReportCreativeGetV2FilteringInventoryTypes       ReportCreativeGetV2FilteringInventoryTypes = "INVENTORY_SEARCH"
)

// All allowed values of ReportCreativeGetV2FilteringInventoryTypes enum
var AllowedReportCreativeGetV2FilteringInventoryTypesEnumValues = []ReportCreativeGetV2FilteringInventoryTypes{
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

// NewReportCreativeGetV2FilteringInventoryTypesFromValue returns a pointer to a valid ReportCreativeGetV2FilteringInventoryTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportCreativeGetV2FilteringInventoryTypesFromValue(v string) (*ReportCreativeGetV2FilteringInventoryTypes, error) {
	ev := ReportCreativeGetV2FilteringInventoryTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportCreativeGetV2FilteringInventoryTypes: valid values are %v", v, AllowedReportCreativeGetV2FilteringInventoryTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportCreativeGetV2FilteringInventoryTypes) IsValid() bool {
	for _, existing := range AllowedReportCreativeGetV2FilteringInventoryTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_creative_get_v2_filtering_inventory_types value
func (v ReportCreativeGetV2FilteringInventoryTypes) Ptr() *ReportCreativeGetV2FilteringInventoryTypes {
	return &v
}
