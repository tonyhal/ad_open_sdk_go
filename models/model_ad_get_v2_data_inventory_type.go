/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.0.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// AdGetV2DataInventoryType
type AdGetV2DataInventoryType string

// List of ad_get_v2_data_inventory_type
const (
	INVENTORY_TOMATO_NOVEL_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_TOMATO_NOVEL"
	UNION_BOUTIQUE_GAME_AdGetV2DataInventoryType         AdGetV2DataInventoryType = "UNION_BOUTIQUE_GAME"
	INVENTORY_TEXT_LINK_AdGetV2DataInventoryType         AdGetV2DataInventoryType = "INVENTORY_TEXT_LINK"
	INVENTORY_AWEME_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_AWEME_FEED"
	INVENTORY_SEARCH_AdGetV2DataInventoryType            AdGetV2DataInventoryType = "INVENTORY_SEARCH"
	INVENTORY_ESSAY_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_ESSAY_FEED"
	INVENTORY_HOTSOON_FEED_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_HOTSOON_FEED"
	INVENTORY_UNION_SPLASH_SLOT_AdGetV2DataInventoryType AdGetV2DataInventoryType = "INVENTORY_UNION_SPLASH_SLOT"
	INVENTORY_FEED_AdGetV2DataInventoryType              AdGetV2DataInventoryType = "INVENTORY_FEED"
	INVENTORY_VIDEO_FEED_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_VIDEO_FEED"
	INVENTORY_UNION_SLOT_AdGetV2DataInventoryType        AdGetV2DataInventoryType = "INVENTORY_UNION_SLOT"
	INVENTORY_AWEME_SEARCH_AdGetV2DataInventoryType      AdGetV2DataInventoryType = "INVENTORY_AWEME_SEARCH"
)

// All allowed values of AdGetV2DataInventoryType enum
var AllowedAdGetV2DataInventoryTypeEnumValues = []AdGetV2DataInventoryType{
	"INVENTORY_TOMATO_NOVEL",
	"UNION_BOUTIQUE_GAME",
	"INVENTORY_TEXT_LINK",
	"INVENTORY_AWEME_FEED",
	"INVENTORY_SEARCH",
	"INVENTORY_ESSAY_FEED",
	"INVENTORY_HOTSOON_FEED",
	"INVENTORY_UNION_SPLASH_SLOT",
	"INVENTORY_FEED",
	"INVENTORY_VIDEO_FEED",
	"INVENTORY_UNION_SLOT",
	"INVENTORY_AWEME_SEARCH",
}

// NewAdGetV2DataInventoryTypeFromValue returns a pointer to a valid AdGetV2DataInventoryType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdGetV2DataInventoryTypeFromValue(v string) (*AdGetV2DataInventoryType, error) {
	ev := AdGetV2DataInventoryType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdGetV2DataInventoryType: valid values are %v", v, AllowedAdGetV2DataInventoryTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdGetV2DataInventoryType) IsValid() bool {
	for _, existing := range AllowedAdGetV2DataInventoryTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ad_get_v2_data_inventory_type value
func (v AdGetV2DataInventoryType) Ptr() *AdGetV2DataInventoryType {
	return &v
}
