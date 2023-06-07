/*
API Title

巨量引擎开放平台

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// PromotionListV30DataListNativeSettingIsFeedAndFavSee
type PromotionListV30DataListNativeSettingIsFeedAndFavSee string

// List of promotion_list_v3.0_data_list_native_setting_is_feed_and_fav_see
const (
	OFF_PromotionListV30DataListNativeSettingIsFeedAndFavSee PromotionListV30DataListNativeSettingIsFeedAndFavSee = "OFF"
	ON_PromotionListV30DataListNativeSettingIsFeedAndFavSee  PromotionListV30DataListNativeSettingIsFeedAndFavSee = "ON"
)

// All allowed values of PromotionListV30DataListNativeSettingIsFeedAndFavSee enum
var AllowedPromotionListV30DataListNativeSettingIsFeedAndFavSeeEnumValues = []PromotionListV30DataListNativeSettingIsFeedAndFavSee{
	"OFF",
	"ON",
}

// NewPromotionListV30DataListNativeSettingIsFeedAndFavSeeFromValue returns a pointer to a valid PromotionListV30DataListNativeSettingIsFeedAndFavSee
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPromotionListV30DataListNativeSettingIsFeedAndFavSeeFromValue(v string) (*PromotionListV30DataListNativeSettingIsFeedAndFavSee, error) {
	ev := PromotionListV30DataListNativeSettingIsFeedAndFavSee(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PromotionListV30DataListNativeSettingIsFeedAndFavSee: valid values are %v", v, AllowedPromotionListV30DataListNativeSettingIsFeedAndFavSeeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PromotionListV30DataListNativeSettingIsFeedAndFavSee) IsValid() bool {
	for _, existing := range AllowedPromotionListV30DataListNativeSettingIsFeedAndFavSeeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to promotion_list_v3.0_data_list_native_setting_is_feed_and_fav_see value
func (v PromotionListV30DataListNativeSettingIsFeedAndFavSee) Ptr() *PromotionListV30DataListNativeSettingIsFeedAndFavSee {
	return &v
}
