/*
Oceanengine Open Api

巨量引擎开放平台 Open Api

API version: 1.1.14
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"fmt"
)

// BrandAwemeListV30DataAwemeUserInfoListAuthType
type BrandAwemeListV30DataAwemeUserInfoListAuthType string

// List of brand_aweme_list_v3.0_data_aweme_user_info_list_auth_type
const (
	ACCOUNT_BrandAwemeListV30DataAwemeUserInfoListAuthType       BrandAwemeListV30DataAwemeUserInfoListAuthType = "ACCOUNT"
	ITEM_BrandAwemeListV30DataAwemeUserInfoListAuthType          BrandAwemeListV30DataAwemeUserInfoListAuthType = "ITEM"
	ITEM_AND_LIVE_BrandAwemeListV30DataAwemeUserInfoListAuthType BrandAwemeListV30DataAwemeUserInfoListAuthType = "ITEM_AND_LIVE"
	LIVE_BrandAwemeListV30DataAwemeUserInfoListAuthType          BrandAwemeListV30DataAwemeUserInfoListAuthType = "LIVE"
)

// All allowed values of BrandAwemeListV30DataAwemeUserInfoListAuthType enum
var AllowedBrandAwemeListV30DataAwemeUserInfoListAuthTypeEnumValues = []BrandAwemeListV30DataAwemeUserInfoListAuthType{
	"ACCOUNT",
	"ITEM",
	"ITEM_AND_LIVE",
	"LIVE",
}

// NewBrandAwemeListV30DataAwemeUserInfoListAuthTypeFromValue returns a pointer to a valid BrandAwemeListV30DataAwemeUserInfoListAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandAwemeListV30DataAwemeUserInfoListAuthTypeFromValue(v string) (*BrandAwemeListV30DataAwemeUserInfoListAuthType, error) {
	ev := BrandAwemeListV30DataAwemeUserInfoListAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandAwemeListV30DataAwemeUserInfoListAuthType: valid values are %v", v, AllowedBrandAwemeListV30DataAwemeUserInfoListAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandAwemeListV30DataAwemeUserInfoListAuthType) IsValid() bool {
	for _, existing := range AllowedBrandAwemeListV30DataAwemeUserInfoListAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to brand_aweme_list_v3.0_data_aweme_user_info_list_auth_type value
func (v BrandAwemeListV30DataAwemeUserInfoListAuthType) Ptr() *BrandAwemeListV30DataAwemeUserInfoListAuthType {
	return &v
}
