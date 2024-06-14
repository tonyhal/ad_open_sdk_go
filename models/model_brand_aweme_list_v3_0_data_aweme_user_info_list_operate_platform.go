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

// BrandAwemeListV30DataAwemeUserInfoListOperatePlatform
type BrandAwemeListV30DataAwemeUserInfoListOperatePlatform string

// List of brand_aweme_list_v3.0_data_aweme_user_info_list_operate_platform
const (
	BIDDING_BrandAwemeListV30DataAwemeUserInfoListOperatePlatform BrandAwemeListV30DataAwemeUserInfoListOperatePlatform = "BIDDING"
	BRAND_BrandAwemeListV30DataAwemeUserInfoListOperatePlatform   BrandAwemeListV30DataAwemeUserInfoListOperatePlatform = "BRAND"
)

// All allowed values of BrandAwemeListV30DataAwemeUserInfoListOperatePlatform enum
var AllowedBrandAwemeListV30DataAwemeUserInfoListOperatePlatformEnumValues = []BrandAwemeListV30DataAwemeUserInfoListOperatePlatform{
	"BIDDING",
	"BRAND",
}

// NewBrandAwemeListV30DataAwemeUserInfoListOperatePlatformFromValue returns a pointer to a valid BrandAwemeListV30DataAwemeUserInfoListOperatePlatform
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAwemeListV30DataAwemeUserInfoListOperatePlatformFromValue(v string) (*BrandAwemeListV30DataAwemeUserInfoListOperatePlatform, error) {
	ev := BrandAwemeListV30DataAwemeUserInfoListOperatePlatform(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAwemeListV30DataAwemeUserInfoListOperatePlatform: valid values are %v", v, AllowedBrandAwemeListV30DataAwemeUserInfoListOperatePlatformEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAwemeListV30DataAwemeUserInfoListOperatePlatform) IsValid() bool {
	for _, existing := range AllowedBrandAwemeListV30DataAwemeUserInfoListOperatePlatformEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_aweme_list_v3.0_data_aweme_user_info_list_operate_platform value
func (v BrandAwemeListV30DataAwemeUserInfoListOperatePlatform) Ptr() *BrandAwemeListV30DataAwemeUserInfoListOperatePlatform {
	return &v
}
