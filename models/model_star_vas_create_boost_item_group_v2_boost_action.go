/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// StarVasCreateBoostItemGroupV2BoostAction
type StarVasCreateBoostItemGroupV2BoostAction string

// List of star_vas_create_boost_item_group_v2_boost_action
const (
	LINK_ACTION_StarVasCreateBoostItemGroupV2BoostAction   StarVasCreateBoostItemGroupV2BoostAction = "LINK_ACTION"
	NATIVE_ACTION_StarVasCreateBoostItemGroupV2BoostAction StarVasCreateBoostItemGroupV2BoostAction = "NATIVE_ACTION"
)

// All allowed values of StarVasCreateBoostItemGroupV2BoostAction enum
var AllowedStarVasCreateBoostItemGroupV2BoostActionEnumValues = []StarVasCreateBoostItemGroupV2BoostAction{
	"LINK_ACTION",
	"NATIVE_ACTION",
}

// NewStarVasCreateBoostItemGroupV2BoostActionFromValue returns a pointer to a valid StarVasCreateBoostItemGroupV2BoostAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStarVasCreateBoostItemGroupV2BoostActionFromValue(v string) (*StarVasCreateBoostItemGroupV2BoostAction, error) {
	ev := StarVasCreateBoostItemGroupV2BoostAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StarVasCreateBoostItemGroupV2BoostAction: valid values are %v", v, AllowedStarVasCreateBoostItemGroupV2BoostActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StarVasCreateBoostItemGroupV2BoostAction) IsValid() bool {
	for _, existing := range AllowedStarVasCreateBoostItemGroupV2BoostActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to star_vas_create_boost_item_group_v2_boost_action value
func (v StarVasCreateBoostItemGroupV2BoostAction) Ptr() *StarVasCreateBoostItemGroupV2BoostAction {
	return &v
}
