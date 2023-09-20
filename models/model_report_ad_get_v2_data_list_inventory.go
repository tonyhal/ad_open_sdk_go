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

// ReportAdGetV2DataListInventory
type ReportAdGetV2DataListInventory string

// List of report_ad_get_v2_data_list_inventory
const (
	INVENTORY_PIPIXIA_ReportAdGetV2DataListInventory      ReportAdGetV2DataListInventory = "INVENTORY_PIPIXIA"
	INVENTORY_VIDEO_FEED_ReportAdGetV2DataListInventory   ReportAdGetV2DataListInventory = "INVENTORY_VIDEO_FEED"
	INVENTORY_AUTOMOBILE_ReportAdGetV2DataListInventory   ReportAdGetV2DataListInventory = "INVENTORY_AUTOMOBILE"
	INVENTORY_UNIVERSAL_ReportAdGetV2DataListInventory    ReportAdGetV2DataListInventory = "INVENTORY_UNIVERSAL"
	INVENTORY_FEED_ReportAdGetV2DataListInventory         ReportAdGetV2DataListInventory = "INVENTORY_FEED"
	INVENTORY_BEAUTY_ReportAdGetV2DataListInventory       ReportAdGetV2DataListInventory = "INVENTORY_BEAUTY"
	INVENTORY_AWEME_FEED_ReportAdGetV2DataListInventory   ReportAdGetV2DataListInventory = "INVENTORY_AWEME_FEED"
	UNION_BOUTIQUE_GAME_ReportAdGetV2DataListInventory    ReportAdGetV2DataListInventory = "UNION_BOUTIQUE_GAME"
	INVENTORY_STUDY_ReportAdGetV2DataListInventory        ReportAdGetV2DataListInventory = "INVENTORY_STUDY"
	INVENTORY_HOTSOON_FEED_ReportAdGetV2DataListInventory ReportAdGetV2DataListInventory = "INVENTORY_HOTSOON_FEED"
	INVENTORY_FURNISH_ReportAdGetV2DataListInventory      ReportAdGetV2DataListInventory = "INVENTORY_FURNISH"
	INVENTORY_TOMATO_NOVEL_ReportAdGetV2DataListInventory ReportAdGetV2DataListInventory = "INVENTORY_TOMATO_NOVEL"
	INVENTORY_UNION_SLOT_ReportAdGetV2DataListInventory   ReportAdGetV2DataListInventory = "INVENTORY_UNION_SLOT"
	INVENTORY_FACE_U_ReportAdGetV2DataListInventory       ReportAdGetV2DataListInventory = "INVENTORY_FACE_U"
	INVENTORY_SEARCH_ReportAdGetV2DataListInventory       ReportAdGetV2DataListInventory = "INVENTORY_SEARCH"
)

// All allowed values of ReportAdGetV2DataListInventory enum
var AllowedReportAdGetV2DataListInventoryEnumValues = []ReportAdGetV2DataListInventory{
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

// NewReportAdGetV2DataListInventoryFromValue returns a pointer to a valid ReportAdGetV2DataListInventory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAdGetV2DataListInventoryFromValue(v string) (*ReportAdGetV2DataListInventory, error) {
	ev := ReportAdGetV2DataListInventory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAdGetV2DataListInventory: valid values are %v", v, AllowedReportAdGetV2DataListInventoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAdGetV2DataListInventory) IsValid() bool {
	for _, existing := range AllowedReportAdGetV2DataListInventoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to report_ad_get_v2_data_list_inventory value
func (v ReportAdGetV2DataListInventory) Ptr() *ReportAdGetV2DataListInventory {
	return &v
}
