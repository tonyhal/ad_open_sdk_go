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

// PromotionUpdateV30NativeSettingIsFeedAndFavSee
type PromotionUpdateV30NativeSettingIsFeedAndFavSee string

// List of promotion_update_v3.0_native_setting_is_feed_and_fav_see
const (
	OFF_PromotionUpdateV30NativeSettingIsFeedAndFavSee PromotionUpdateV30NativeSettingIsFeedAndFavSee = "OFF"
	ON_PromotionUpdateV30NativeSettingIsFeedAndFavSee  PromotionUpdateV30NativeSettingIsFeedAndFavSee = "ON"
)

// All allowed values of PromotionUpdateV30NativeSettingIsFeedAndFavSee enum
var AllowedPromotionUpdateV30NativeSettingIsFeedAndFavSeeEnumValues = []PromotionUpdateV30NativeSettingIsFeedAndFavSee{
	"OFF",
	"ON",
}

// NewPromotionUpdateV30NativeSettingIsFeedAndFavSeeFromValue returns a pointer to a valid PromotionUpdateV30NativeSettingIsFeedAndFavSee
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionUpdateV30NativeSettingIsFeedAndFavSeeFromValue(v string) (*PromotionUpdateV30NativeSettingIsFeedAndFavSee, error) {
	ev := PromotionUpdateV30NativeSettingIsFeedAndFavSee(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionUpdateV30NativeSettingIsFeedAndFavSee: valid values are %v", v, AllowedPromotionUpdateV30NativeSettingIsFeedAndFavSeeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionUpdateV30NativeSettingIsFeedAndFavSee) IsValid() bool {
	for _, existing := range AllowedPromotionUpdateV30NativeSettingIsFeedAndFavSeeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_update_v3.0_native_setting_is_feed_and_fav_see value
func (v PromotionUpdateV30NativeSettingIsFeedAndFavSee) Ptr() *PromotionUpdateV30NativeSettingIsFeedAndFavSee {
	return &v
}
